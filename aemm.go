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

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Container will create and start an instance of the Amazon EC2 Metadata Mock (AEMM),
// simulating the Amazon EC2 Metadata Service (IMDS). Once started, IMDS will
// be accessible through the following endpoint:
//
// http://localhost:1338/latest/metadata
//
// By using the default settings, both IMDSv1 and IMDSv2 are supported. Metadata about the
// mocked EC2 instance can then be retrieved using any of the documented categories,
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
//
// For example:
// 	curl http://localhost:1338/latest/metadata/local-ipv4
// 	curl http://localhost:1338/latest/metadata/block-device-mapping/root
//
// To ensure your AWS config is configured to call this mock, the required option needs to
// be set:
//
// 	package main
//
//	import "github.com/aws/aws-sdk-go-v2/config"
//
//	func main() {
//		config.LoadDefaultConfig(context.TODO(), config.WithEC2IMDSEndpoint("http://localhost:1338/latest/metadata"))
//	}
func Container(ctx context.Context) (testcontainers.Container, error) {
	return ContainerWith(ctx, DefaultOptions)
}

// LaunchOptions ...
type LaunchOptions struct {
	// StrictIMDSv2 ...
	StrictIMDSv2 bool
}

// DefaultOptions ...
var DefaultOptions = LaunchOptions{}

// ContainerWith ...
func ContainerWith(ctx context.Context, opts LaunchOptions) (testcontainers.Container, error) {
	flags := []string{}
	if opts.StrictIMDSv2 {
		flags = append(flags, "--imdsv2")
	}

	req := testcontainers.ContainerRequest{
		Image:        "public.ecr.aws/aws-ec2/amazon-ec2-metadata-mock:v1.10.1",
		Cmd:          flags,
		ExposedPorts: []string{"1338:1338/tcp"},
		WaitingFor:   wait.ForLog("Initiating ec2-metadata-mock for all mocks on port 1338"),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
