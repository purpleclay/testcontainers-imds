# testcontainers-aemm

Testcontainers wrapper for the [Amazon EC2 Metadata Mock](https://github.com/aws/amazon-ec2-metadata-mock) (AEMM) tool. Quickly and easily simulate the [Amazon Instance Metadata Service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) (IMDS) for localised testing.

[![Build status](https://img.shields.io/github/workflow/status/purpleclay/testcontainers-aemm/ci?style=flat-square&logo=go)](https://github.com/purpleclay/testcontainers-aemm/actions?workflow=ci)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/purpleclay/testcontainers-aemm?style=flat-square)](https://goreportcard.com/report/github.com/purpleclay/testcontainers-aemm)
[![Go Version](https://img.shields.io/github/go-mod/go-version/purpleclay/testcontainers-aemm.svg?style=flat-square)](go.mod)
[![codecov](https://codecov.io/gh/purpleclay/testcontainers-aemm/branch/main/graph/badge.svg)](https://codecov.io/gh/purpleclay/testcontainers-aemm)

## Quick Start

Import the library into your project:

```sh
go get github.com/purpleclay/testcontainers-aemm
```

Then write your first test:

```go
package imds_test

import (
    "context"
    "io/ioutil"
    "net/http"
    "testing"

    aemm "github.com/purpleclay/testcontainers-aemm"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestInstanceMetadata(t *testing.T) {
    ctx := context.Background()

    container, err := aemm.Start(ctx)
    require.NoError(t, err)
    defer container.Terminate(ctx)

    resp, _ := http.Get(container.URL + aemm.PathLocalIPv4)
    defer resp.Body.Close()

    out, _ := ioutil.ReadAll(resp.Body)
    assert.Equal(t, aemm.ValueLocalIPv4, string(out))
}
```
