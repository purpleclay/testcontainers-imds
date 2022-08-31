/*
Copyright (c) 2022 Purple Clay

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package imds

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/creasty/defaults"
	imdsmock "github.com/purpleclay/imds-mock/pkg/imds"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	// MinTokenTTLInSeconds defines the minimum duration of a session token in seconds
	MinTokenTTLInSeconds = 1

	// MaxTokenTTLInSeconds defines the maximum duration of a session token in seconds
	MaxTokenTTLInSeconds = 21600

	// AllCategories triggers the retrieval of all instance metadata categories when
	// using either the Get() or GetWithToken() operations
	AllCategories = ""
)

// Container represents an instance of an AEMM container
type Container struct {
	testcontainers.Container

	metadataURL string
	tokenURL    string
	client      *http.Client
}

// Start will create and start an instance of the Instance Metadata Mock (imds-mock),
// simulating the Amazon EC2 Metadata Service (IMDS). Once started, IMDS will
// be accessible through the expected endpoint. As the caller it is your responsibility
// to terminate the container by invoking the Terminate() method on the container.
//
// http://localhost:1338/latest/meta-data/
//
// By using the default settings, both IMDSv1 and IMDSv2 are supported. Metadata about the
// mocked EC2 instance can then be retrieved using any of the documented categories,
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
//
// For example:
//
//	curl http://localhost:1338/latest/meta-data/local-ipv4
//	curl http://localhost:1338/latest/meta-data/block-device-mapping/root
//
// To ensure your AWS config is configured to call this mock, the required option needs to
// be set:
//
//	package main
//
//	import "github.com/aws/aws-sdk-go-v2/config"
//
//	func main() {
//		config.LoadDefaultConfig(context.TODO(), config.WithEC2IMDSEndpoint("http://localhost:1338/latest/meta-data/"))
//	}
func Start(ctx context.Context) (*Container, error) {
	return StartWith(ctx, Options{})
}

// MustStart behaves in the same way as Start but panics if the container cannot
// be started for any reason. This removes the need to handle any returned errors,
// simplifying initialisation.
//
// As the caller it is your responsibility to terminate the container by invoking
// the Terminate() method on the container.
func MustStart(ctx context.Context) *Container {
	container, err := StartWith(ctx, Options{})
	if err != nil {
		panic(`aemm: MustStart(): ` + err.Error())
	}

	return container
}

// Options defines all configurable options when starting the AEMM container
type Options struct {
	// ExcludeInstanceTags will ensure any tags associated with the instance
	// are not exposed through the tags/instance category. Enable this to
	// simulate the default behaviour of an EC2
	// 	@Default false
	ExcludeInstanceTags bool

	// ExposedPort defines which port on the host will be mapped to the default port
	// of the container
	//	@Default 1338
	ExposedPort string `default:"1338"`

	// Image is the name of the Instance Metadata Mock image to pull when
	// launching the container
	// 	@Default ghcr.io/purpleclay/imds-mock
	Image string `default:"ghcr.io/purpleclay/imds-mock"`

	// ImageTag is the version of the Instance Metadata Mock image to pull
	// from the source docker registry
	//	@Default latest
	ImageTag string `default:"latest"`

	// InstanceTags defines a list of instance tags that should be exposed through
	// the instance/tags metadata category, overwriting any existing defaults
	//	@Default existing instance tags will not be overwritten
	InstanceTags map[string]string

	// Pretty print any JSON response
	//	@Default false
	Pretty bool

	// Spot is a flag that controls the simulation of a spot instance and interruption
	// notice
	//	@Default false
	Spot bool

	// SpotAction is used in conjunction with the spot flag to control both the type
	// and initial delay of the spot interruption notice.
	//   @Default
	SpotAction imdsmock.SpotActionEvent `default:"{\"Action\":\"terminate\", \"Duration\": \"0s\"}"`

	// IMDSv2 will enforce IMDSv2 and require a session token when making metadata
	// requests. A token is requested by issuing a PUT request to the token endpoint, and
	// supplying a TTL of between 1 and 2600 seconds.
	//
	//	PUT localhost:1338/latest/api/token -H "X-aws-ec2-metadata-token-ttl-seconds: 21600"
	//
	// Any subsequent request must provide the token as a header:
	//
	// 	GET localhost:1338/latest/meta-data/local-ipv4 -H "X-aws-ec2-metadata-token: $TOKEN"
	//
	//	@Default false
	IMDSv2 bool
}

// StartWith will create and start an instance of the Instance Metadata Mock (imds-mock),
// simulating the Amazon EC2 Metadata Service (IMDS). The launch behaviour of the mock
// can be configured through the provided LaunchOptions. Once started, IMDS will
// be accessible through the endpoint. As the caller it is your responsibility to
// terminate the container by invoking the Terminate() method on the container.
//
// http://localhost:1338/latest/meta-data/
//
// By using the default settings, both IMDSv1 and IMDSv2 are supported. Metadata about the
// mocked EC2 instance can then be retrieved using any of the documented categories,
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
//
// For example:
//
//	curl http://localhost:1338/latest/meta-data/local-ipv4
//	curl http://localhost:1338/latest/meta-data/block-device-mapping/root
//
// To ensure your AWS config is configured to call this mock, the required option needs to
// be set:
//
//	package main
//
//	import "github.com/aws/aws-sdk-go-v2/config"
//
//	func main() {
//		config.LoadDefaultConfig(context.TODO(), config.WithEC2IMDSEndpoint("http://localhost:1338/latest/meta-data/"))
//	}
func StartWith(ctx context.Context, opts Options) (*Container, error) {
	// Adjust the wait strategy based on the options
	waitStrategy := wait.ForHTTP("/latest/meta-data/").WithPort("1338")

	// Ensure all defaults are set before launching the container
	defaults.Set(&opts)
	fmt.Printf("%#v\n", opts)

	flags := []string{}
	if opts.ExcludeInstanceTags {
		flags = append(flags, "--exclude-instance-tags")
	}

	if len(opts.InstanceTags) > 0 {
		flags = append(flags, "--instance-tags")
		flags = append(flags, keyValueListFlag(opts.InstanceTags))
	}

	if opts.Pretty {
		flags = append(flags, "--pretty")
	}

	if opts.Spot {
		flags = append(flags, "--spot")
		flags = append(flags, "--spot-action")
		flags = append(flags, spotActionFlag(opts.SpotAction))
	}

	if opts.IMDSv2 {
		flags = append(flags, "--imdsv2")

		// 401 should be issued without a token
		waitStrategy = wait.ForHTTP("/latest/meta-data/").
			WithPort("1338").
			WithStatusCodeMatcher(func(status int) bool { return status == http.StatusUnauthorized })
	}

	req := testcontainers.ContainerRequest{
		Image:        fmt.Sprintf("%s:%s", opts.Image, opts.ImageTag),
		Cmd:          flags,
		ExposedPorts: []string{opts.ExposedPort + ":1338/tcp"},
		WaitingFor:   waitStrategy,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &Container{
		Container:   container,
		metadataURL: fmt.Sprintf("http://localhost:%s/latest/meta-data/", opts.ExposedPort),
		tokenURL:    fmt.Sprintf("http://localhost:%s/latest/api/token", opts.ExposedPort),
		client:      &http.Client{Timeout: 1 * time.Second},
	}, nil
}

func keyValueListFlag(in map[string]string) string {
	kv := make([]string, 0, len(in))
	for key, value := range in {
		kv = append(kv, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(kv, ",")
}

func spotActionFlag(event imdsmock.SpotActionEvent) string {
	return fmt.Sprintf("%s=%s", string(event.Action), event.Duration.String())
}

// MustStartWith behaves in the same way as StartWith but panics if the container cannot
// be started for any reason. This removes the need to handle any returned errors,
// simplifying initialisation.
//
// As the caller it is your responsibility to terminate the container by invoking
// the Terminate() method on the container.
func MustStartWith(ctx context.Context, opts Options) *Container {
	container, err := StartWith(ctx, opts)
	if err != nil {
		panic(`aemm: MustStartWith(` + fmt.Sprintf("%#v", opts) + `): ` + err.Error())
	}

	return container
}

// Get will attempt to retrieve an instance category from the running container. The raw
// value of the category will be returned from the container upon success. If any HTTP
// failure occurs while trying to retrieve a category, the raw error is returned
//
// Status Codes:
//
//	200: category was retrieved
//	404: category does not exist
func (c *Container) Get(category string) (string, int, error) {
	return c.GetV2(category, "")
}

// GetV2 will attempt to retrieve an instance category from the running container
// using an authenticated session token based request. If the container was not started
// in IMDSv2 mode, the token will have no effect. The raw value of the category will
// be returned from the container upon success. If any HTTP failure occurs while trying
// to retrieve a category, the raw error is returned
//
// Status Codes:
//
//	200: category was retrieved
//	404: category does not exist
//	401: session token is either invalid or expired
func (c *Container) GetV2(category, token string) (string, int, error) {
	req, _ := http.NewRequest(http.MethodGet, c.metadataURL+category, http.NoBody)
	if token != "" {
		req.Header.Add("X-aws-ec2-metadata-token", token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	return string(data), resp.StatusCode, nil
}

// TokenWithTTL will attempt to generate a session token with the provided TTL in seconds.
// If any HTTP failure occurs while trying to retrieve a category, the raw error is returned
//
// Status Codes:
//
//	200: token was created
//	400: TTL was outside the expected bounds (min: 1, max: 21600)
func (c *Container) TokenWithTTL(ttl int) (string, int, error) {
	req, _ := http.NewRequest(http.MethodPut, c.tokenURL, http.NoBody)
	req.Header.Add("X-aws-ec2-metadata-token-ttl-seconds", strconv.Itoa(ttl))

	resp, err := c.client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	return string(data), resp.StatusCode, nil
}
