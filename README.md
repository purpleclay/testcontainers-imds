# testcontainers-imds

Testcontainers wrapper for the [Instance Metadata Mock](https://github.com/purpleclay/imds-mock) (imds-mock) tool. Quickly and easily simulate the [Amazon Instance Metadata Service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) (IMDS) for localised testing.

[![Build status](https://img.shields.io/github/workflow/status/purpleclay/testcontainers-imds/ci?style=flat-square&logo=go)](https://github.com/purpleclay/testcontainers-imds/actions?workflow=ci)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/purpleclay/testcontainers-imds?style=flat-square)](https://goreportcard.com/report/github.com/purpleclay/testcontainers-imds)
[![Go Version](https://img.shields.io/github/go-mod/go-version/purpleclay/testcontainers-imds.svg?style=flat-square)](go.mod)
[![codecov](https://codecov.io/gh/purpleclay/testcontainers-imds/branch/main/graph/badge.svg)](https://codecov.io/gh/purpleclay/testcontainers-imds)

## Quick Start

Import the library into your project:

```sh
go get github.com/purpleclay/testcontainers-imds
```

Then write your first test:

```go
package imds_test

import (
    "context"
    "io/ioutil"
    "net/http"
    "testing"

    imds "github.com/purpleclay/testcontainers-imds"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestInstanceMetadata(t *testing.T) {
    ctx := context.Background()

    container, err := imds.Start(ctx)
    require.NoError(t, err)
    defer container.Terminate(ctx)

    resp, _ := http.Get(container.URL + imds.PathLocalIPv4)
    defer resp.Body.Close()

    out, _ := ioutil.ReadAll(resp.Body)
    assert.Equal(t, imds.ValueLocalIPv4, string(out))
}
```

## Examples

TODO: add links to examples
