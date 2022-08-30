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
package main

import (
	"context"
	"log"
	"strings"
	"time"

	imds "github.com/purpleclay/testcontainers-imds"
)

// This examples demonstrates how to enable strict IMDSv2 and require a session token
// to access any instance categories within the Instance Metadata Service (IMDS).
// Enabling IMDSv2 on all EC2 instances is an Amazon best practices, as documented here:
// https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/
//
// Details on how the Instance Metadata Mock (imds-mock) handles IMDSv2 can be found
// here: https://docs.purpleclay.dev/imds-mock/configure/imdsv2/
func main() {
	ctx := context.Background()

	container := imds.MustStartWith(ctx, imds.Options{
		IMDSv2: true,
	})
	defer container.Terminate(ctx)

	log.Println("IMDS mock started with IMDSv2...")

	token := ""
	instanceID(container, token)

	var err error
	if token, _, err = container.GenerateToken(1); err != nil {
		log.Fatalf("Failed to generate session token. %s\n", err.Error())
	}
	log.Printf("Generated session token with 1 second expiry. %s\n", token)

	// Request the token and then sleep to allow the token to expire
	instanceID(container, token)

	log.Println("Sleeping for 1 second to allow token to expire...")
	time.Sleep(1 * time.Second)
	instanceID(container, token)
}

func instanceID(container *imds.Container, token string) {
	instanceID, status, err := container.GetWithToken(imds.PathInstanceID, token)
	if err != nil {
		log.Fatalf("Failed to query instance metadata mock. %s\n", err.Error())
	}

	if strings.Contains(instanceID, "<title>401 -") {
		instanceID = ""
	}

	log.Printf("Retrieved instance ID: %s Status: %d\n", instanceID, status)
}
