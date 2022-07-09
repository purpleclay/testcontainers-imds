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

func TestContainerWithDefaults(t *testing.T) {
	containerWithDefaults(t)

	out := get(t, "http://localhost:1338/latest/meta-data")
	assert.Contains(t, string(out), "local-ipv4")
}

func containerWithDefaults(t *testing.T) {
	t.Helper()

	container, err := aemm.Container(context.Background())
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(context.Background())
	})
}

func get(t *testing.T, url string) string {
	t.Helper()

	resp, err := http.Get(url)
	require.NoError(t, err)

	t.Cleanup(func() {
		resp.Body.Close()
	})

	out, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	return string(out)
}
