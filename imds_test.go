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
	"strings"
	"testing"

	imds "github.com/purpleclay/testcontainers-imds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStart_CheckURL(t *testing.T) {
	url := startWithDefaults(t)

	assert.Equal(t, "http://localhost:1338/latest/meta-data/", url)
}

func TestStart(t *testing.T) {
	url := startWithDefaults(t)

	out, _ := get(t, url)
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

func TestStartWith_StrictIMDSv2Unauthorised(t *testing.T) {
	url := startWithOptions(t, imds.Options{StrictIMDSv2: true})

	out, status := get(t, url)

	require.Equal(t, http.StatusUnauthorized, status)
	assert.Contains(t, out, "<h1>401 - Unauthorized</h1>")
}

func TestStartWith_StrictIMDSv2(t *testing.T) {
	url := startWithOptions(t, imds.Options{StrictIMDSv2: true})

	out, _ := getAuthorised(t, url)

	assert.Contains(t, string(out), "local-ipv4")
}

func TestMustStartWith_Panics(t *testing.T) {
	require.Panics(t, func() {
		imds.MustStartWith(context.Background(), imds.Options{Image: "image-pull-failure"})
	})
}

func TestStartWith_CheckURL(t *testing.T) {
	url := startWithOptions(t, imds.Options{ExposedPort: "2233"})

	assert.Equal(t, "http://localhost:2233/latest/meta-data/", url)
}

func TestMustStartWith(t *testing.T) {
	require.NotPanics(t, func() {
		ctx := context.Background()

		container := imds.MustStartWith(ctx, imds.Options{})
		container.Terminate(ctx)
	})
}

func TestStartWith_Pretty(t *testing.T) {
	url := startWithOptions(t, imds.Options{Pretty: true})

	out, status := get(t, url+"iam/info")

	require.Equal(t, http.StatusOK, status)
	assert.True(t, strings.HasPrefix(out, "{\n  \"Code\": \"Success\""))
}

func TestStartWith_ExcludeInstanceTags(t *testing.T) {
	url := startWithOptions(t, imds.Options{ExcludeInstanceTags: true})

	_, status := get(t, url+"tags/instance")

	require.Equal(t, http.StatusNotFound, status)
}

func TestStartWith_InstanceTags(t *testing.T) {
	url := startWithOptions(t, imds.Options{InstanceTags: map[string]string{
		"Name":        "testing",
		"Environment": "dev",
	}})

	out, status := get(t, url+"tags/instance")

	require.Equal(t, http.StatusOK, status)

	tags := strings.Split(out, "\n")
	require.Len(t, tags, 2)
	assert.Contains(t, tags, "Name")
	assert.Contains(t, tags, "Environment")
}

func startWithDefaults(t *testing.T) string {
	t.Helper()

	container, err := imds.Start(context.Background())
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})

	return container.URL
}

func startWithOptions(t *testing.T, opts imds.Options) string {
	t.Helper()

	container, err := imds.StartWith(context.Background(), opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})

	return container.URL
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

func getAuthorised(t *testing.T, url string) (string, int) {
	t.Helper()

	// Request an authorisation token using supported maximum duration
	authReq, err := http.NewRequest(http.MethodPut, "http://localhost:1338/latest/api/token", http.NoBody)
	require.NoError(t, err)

	authReq.Header.Add("X-aws-ec2-metadata-token-ttl-seconds", "21600")

	authResp, err := http.DefaultClient.Do(authReq)
	require.NoError(t, err)

	t.Cleanup(func() {
		authResp.Body.Close()
	})

	data, err := io.ReadAll(authResp.Body)
	require.NoError(t, err)
	token := string(data)

	// Perform IMDS request using authorisation token
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	require.NoError(t, err)

	req.Header.Add("X-aws-ec2-metadata-token", token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	t.Cleanup(func() {
		resp.Body.Close()
	})

	out, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return string(out), resp.StatusCode
}
