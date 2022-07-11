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

package aemm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/creasty/defaults"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Container represents an instance of an AEMM container
type Container struct {
	testcontainers.Container

	// URL for querying the default instance metadata endpoint of the container
	//	@Default http://localhost:<EXPOSED_PORT>/latest/meta-data
	URL string
}

// Start will create and start an instance of the Amazon EC2 Metadata Mock (AEMM),
// simulating the Amazon EC2 Metadata Service (IMDS). Once started, IMDS will
// be accessible through the endpoint. As the caller it is your responsibility to
// terminate the container by invoking the Terminate() method on the container.
//
// http://localhost:1338/latest/meta-data
//
// By using the default settings, both IMDSv1 and IMDSv2 are supported. Metadata about the
// mocked EC2 instance can then be retrieved using any of the documented categories,
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
//
// For example:
// 	curl http://localhost:1338/latest/meta-data/local-ipv4
// 	curl http://localhost:1338/latest/meta-data/block-device-mapping/root
//
// To ensure your AWS config is configured to call this mock, the required option needs to
// be set:
//
// 	package main
//
//	import "github.com/aws/aws-sdk-go-v2/config"
//
//	func main() {
//		config.LoadDefaultConfig(context.TODO(), config.WithEC2IMDSEndpoint("http://localhost:1338/latest/meta-data"))
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
	// Image is the name of the AEMM image to pull when launching the container
	// 	@Default public.ecr.aws/aws-ec2/amazon-ec2-metadata-mock
	Image string `default:"public.ecr.aws/aws-ec2/amazon-ec2-metadata-mock"`

	// ImageTag is the version of the AEMM image to pull from the source docker registry
	//	@Default v1.11.1
	ImageTag string `default:"v1.11.1"`

	// ExposedPort defines which port on the host will be mapped to the default port
	// of the container
	//	@Default 1338
	ExposedPort string `default:"1338"`

	// StrictIMDSv2 will enforce IMDSv2 and require a session token when making metadata
	// requests. A token is requested by issuing a PUT request to the token endpoint, and
	// supplying a TTL of between 1 and 2600 seconds.
	//
	//	PUT localhost:1338/latest/api/token -H "X-aws-ec2-metadata-token-ttl-seconds: 21600"
	//
	// Any subsequent request must provide the token as a header:
	//
	// 	GET localhost:1338/latest/meta-data/local-ipv4 -H "X-aws-ec2-metadata-token: $TOKEN"
	StrictIMDSv2 bool
}

// StartWith will create and start an instance of the Amazon EC2 Metadata Mock (AEMM),
// simulating the Amazon EC2 Metadata Service (IMDS). The launch behaviour of the AEMM
// can be configured through the provided LaunchOptions. Once started, IMDS will
// be accessible through the endpoint. As the caller it is your responsibility to
// terminate the container by invoking the Terminate() method on the container.
//
// http://localhost:1338/latest/meta-data
//
// By using the default settings, both IMDSv1 and IMDSv2 are supported. Metadata about the
// mocked EC2 instance can then be retrieved using any of the documented categories,
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
//
// For example:
// 	curl http://localhost:1338/latest/meta-data/local-ipv4
// 	curl http://localhost:1338/latest/meta-data/block-device-mapping/root
//
// To ensure your AWS config is configured to call this mock, the required option needs to
// be set:
//
// 	package main
//
//	import "github.com/aws/aws-sdk-go-v2/config"
//
//	func main() {
//		config.LoadDefaultConfig(context.TODO(), config.WithEC2IMDSEndpoint("http://localhost:1338/latest/meta-data"))
//	}
func StartWith(ctx context.Context, opts Options) (*Container, error) {
	// Adjust the wait strategy based on the options
	waitStrategy := wait.ForHTTP("/latest/meta-data").WithPort("1338")

	flags := []string{}
	if opts.StrictIMDSv2 {
		flags = append(flags, "--imdsv2")

		// 401 should be issued without a token
		waitStrategy = wait.ForHTTP("/latest/meta-data").
			WithPort("1338").
			WithStatusCodeMatcher(func(status int) bool { return status == http.StatusUnauthorized })
	}

	// Ensure all defaults are set before launching the container
	defaults.Set(&opts)

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
		Container: container,
		URL:       fmt.Sprintf("http://localhost:%s/latest/meta-data/", opts.ExposedPort),
	}, nil
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
