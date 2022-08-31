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

package imds_test

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	imdsmock "github.com/purpleclay/imds-mock/pkg/imds"
	"github.com/purpleclay/imds-mock/pkg/imds/patch"
	imds "github.com/purpleclay/testcontainers-imds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	startWithDefaults(t)

	out, status := get(t, "http://localhost:1338/latest/meta-data/")
	assert.Equal(t, http.StatusOK, status)
	assert.Contains(t, string(out), "local-ipv4")
}

func TestMustStart_Panics(t *testing.T) {
	// Deliberately spin up a container that blocks the default port of 1338
	startWithDefaults(t)

	require.Panics(t, func() {
		imds.MustStart(context.Background())
	})
}

func TestMustStart(t *testing.T) {
	require.NotPanics(t, func() {
		ctx := context.Background()

		container := imds.MustStart(ctx)
		container.Terminate(ctx)
	})
}

func TestStartWith_IMDSv2(t *testing.T) {
	startWithOptions(t, imds.Options{IMDSv2: true})

	out, status := get(t, "http://localhost:1338/latest/meta-data/")

	require.Equal(t, http.StatusUnauthorized, status)
	assert.Contains(t, out, "<h1>401 - Unauthorized</h1>")
}

func TestMustStartWith_Panics(t *testing.T) {
	require.Panics(t, func() {
		imds.MustStartWith(context.Background(), imds.Options{Image: "image-pull-failure"})
	})
}

func TestMustStartWith(t *testing.T) {
	require.NotPanics(t, func() {
		ctx := context.Background()

		container := imds.MustStartWith(ctx, imds.Options{})
		container.Terminate(ctx)
	})
}

func TestStartWith_ExposedPort(t *testing.T) {
	startWithOptions(t, imds.Options{ExposedPort: "2233"})

	out, _ := get(t, "http://localhost:2233/latest/meta-data/")
	assert.Contains(t, string(out), "local-ipv4")
}

func TestStartWith_Pretty(t *testing.T) {
	startWithOptions(t, imds.Options{Pretty: true})

	out, _ := get(t, "http://localhost:1338/latest/meta-data/iam/info")

	assert.True(t, strings.HasPrefix(out, "{\n  \"Code\": \"Success\""))
}

func TestStartWith_ExcludeInstanceTags(t *testing.T) {
	startWithOptions(t, imds.Options{ExcludeInstanceTags: true})

	_, status := get(t, "http://localhost:1338/latest/meta-data/tags/instance")

	require.Equal(t, http.StatusNotFound, status)
}

func TestStartWith_InstanceTags(t *testing.T) {
	startWithOptions(t, imds.Options{InstanceTags: map[string]string{
		"Name":        "testing",
		"Environment": "dev",
	}})

	out, _ := get(t, "http://localhost:1338/latest/meta-data/tags/instance")

	tags := strings.Split(out, "\n")
	require.Len(t, tags, 2)
	assert.Contains(t, tags, "Name")
	assert.Contains(t, tags, "Environment")
}

func TestStartWith_Spot(t *testing.T) {
	startWithOptions(t, imds.Options{Spot: true})

	out, _ := get(t, "http://localhost:1338/latest/meta-data/spot/instance-action")

	assert.Contains(t, out, `"action":"terminate"`)
}

func TestStartWith_SpotAction(t *testing.T) {
	startWithOptions(t, imds.Options{
		Spot: true,
		SpotAction: imdsmock.SpotActionEvent{
			Action:   patch.StopSpotInstanceAction,
			Duration: 200 * time.Millisecond,
		},
	})

	// Replicate polling the instance for a spot interruption notice
	var out string
	var status int
	for {
		out, status = get(t, "http://localhost:1338/latest/meta-data/spot/instance-action")
		if status == http.StatusOK {
			break
		}
		t.Log("spot interruption hasn't been raised yet. sleep and try again")
		time.Sleep(20 * time.Millisecond)
	}

	assert.Contains(t, out, `"action":"stop"`)
}

func startWithDefaults(t *testing.T) *imds.Container {
	t.Helper()

	container, err := imds.Start(context.Background())
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})

	return container
}

func startWithOptions(t *testing.T, opts imds.Options) *imds.Container {
	t.Helper()

	container, err := imds.StartWith(context.Background(), opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})

	return container
}

func get(t *testing.T, url string) (string, int) {
	t.Helper()

	resp, err := http.Get(url)
	require.NoError(t, err)

	t.Cleanup(func() {
		resp.Body.Close()
	})

	out, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return string(out), resp.StatusCode
}

func getToken(t *testing.T, url string) (string, int) {
	t.Helper()

	req, err := http.NewRequest(http.MethodPut, url, http.NoBody)
	require.NoError(t, err)

	req.Header.Add("X-aws-ec2-metadata-token-ttl-seconds", strconv.Itoa(imds.MaxTokenTTLInSeconds))

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	t.Cleanup(func() {
		resp.Body.Close()
	})

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	token := string(data)

	return token, resp.StatusCode
}

func TestGetAll(t *testing.T) {
	container := startWithDefaults(t)

	out, _, _ := container.Get(imds.AllCategories)

	categories := strings.Split(out, "\n")

	// Just verify a subset
	assert.Contains(t, categories, imds.PathAMIID)
	assert.Contains(t, categories, imds.PathAMILaunchIndex)
	assert.Contains(t, categories, imds.PathAMIManifestPath)
}

func TestGet(t *testing.T) {
	container := startWithDefaults(t)

	ipv4, status, err := container.Get(imds.PathLocalIPv4)

	assert.Equal(t, imds.ValueLocalIPv4, ipv4)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestGetTimeout(t *testing.T) {
	container := startWithDefaults(t)
	// Ensure a connection error is simulated by immediately stopping the container
	container.Stop(context.Background(), nil)

	out, status, err := container.Get(imds.PathLocalIPv4)

	assert.Empty(t, out)
	assert.Equal(t, 0, status)
	assert.Error(t, err)
}

func TestGetV2(t *testing.T) {
	container := startWithOptions(t, imds.Options{IMDSv2: true})

	token, _ := getToken(t, "http://localhost:1338/latest/api/token")
	require.NotEmpty(t, token)

	amiID, status, err := container.GetV2(imds.PathAMIID, token)

	assert.Equal(t, imds.ValueAMIID, amiID)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestTokenWithTTL(t *testing.T) {
	container := startWithOptions(t, imds.Options{IMDSv2: true})

	token, status, err := container.TokenWithTTL(10)

	assert.NotEmpty(t, token)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, err)
}

func TestTokenWithTTLTimeout(t *testing.T) {
	container := startWithDefaults(t)
	// Ensure a connection error is simulated by immediately stopping the container
	container.Stop(context.Background(), nil)

	out, status, err := container.TokenWithTTL(1)

	assert.Empty(t, out)
	assert.Equal(t, 0, status)
	assert.Error(t, err)
}
