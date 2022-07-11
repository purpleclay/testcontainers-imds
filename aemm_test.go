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

package aemm_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	aemm "github.com/purpleclay/testcontainers-aemm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	startWithDefaults(t)

	out, _ := get(t, "http://localhost:1338/latest/meta-data")
	assert.Contains(t, string(out), "local-ipv4")
}

func TestMustStart_Panics(t *testing.T) {
	// Deliberately spin up a container that blocks the default port of 1338
	startWithDefaults(t)

	require.Panics(t, func() {
		aemm.MustStart(context.Background())
	})
}

func TestMustStart(t *testing.T) {
	require.NotPanics(t, func() {
		ctx := context.Background()

		container := aemm.MustStart(ctx)
		container.Terminate(ctx)
	})
}

func TestStartWith_StrictIMDSv2Unauthorised(t *testing.T) {
	startWithOptions(t, aemm.Options{StrictIMDSv2: true})

	out, status := get(t, "http://localhost:1338/latest/meta-data")

	assert.Contains(t, string(out), "<h1>401 - Unauthorized</h1>")
	assert.Equal(t, http.StatusUnauthorized, status)
}

func TestStartWith_StrictIMDSv2(t *testing.T) {
	startWithOptions(t, aemm.Options{StrictIMDSv2: true})

	out, _ := getAuthorised(t, "http://localhost:1338/latest/meta-data")

	assert.Contains(t, string(out), "local-ipv4")
}

func TestMustStartWith_Panics(t *testing.T) {
	require.Panics(t, func() {
		aemm.MustStartWith(context.Background(), aemm.Options{Image: "image-pull-failure"})
	})
}

func TestMustStartWith(t *testing.T) {
	require.NotPanics(t, func() {
		ctx := context.Background()

		container := aemm.MustStartWith(ctx, aemm.Options{})
		container.Terminate(ctx)
	})
}

func startWithDefaults(t *testing.T) {
	t.Helper()

	container, err := aemm.Start(context.Background())
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})
}

func startWithOptions(t *testing.T, opts aemm.Options) {
	t.Helper()

	container, err := aemm.StartWith(context.Background(), opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})
}

func get(t *testing.T, url string) (string, int) {
	t.Helper()

	resp, err := http.Get(url)
	require.NoError(t, err)

	t.Cleanup(func() {
		resp.Body.Close()
	})

	out, err := ioutil.ReadAll(resp.Body)
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

	data, err := ioutil.ReadAll(authResp.Body)
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

	out, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	return string(out), resp.StatusCode
}
