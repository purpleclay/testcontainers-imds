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
	"io"
	"log"
	"net/http"
	"time"

	imds "github.com/purpleclay/testcontainers-imds"
)

// TODO: replace these
const (
	a = "X-aws-ec2-metadata-token-ttl-seconds"
	b = "X-aws-ec2-metadata-token"
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
	instanceID(container.URL, token)

	token = generateToken(container.TokenURL)
	log.Printf("Generated session token with 1 second expiry. %s\n", token)

	// Request the token and then sleep to allow the token to expire
	instanceID(container.URL, token)

	log.Println("Sleeping for 1 second to allow token to expire...")
	time.Sleep(1 * time.Second)
	instanceID(container.URL, token)
}

// TODO: remove the need to write all of this boilerplate code

func instanceID(url, token string) {
	req, err := http.NewRequest(http.MethodGet, url+imds.PathInstanceID, http.NoBody)
	if err != nil {
		log.Fatalf("Failed to create instance ID request. %s\n", err.Error())
	}
	req.Header.Add(b, token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to query instance metadata mock. %s\n", err.Error())
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read instance id. %s\n", err.Error())
	}

	instanceID := string(data)
	if resp.StatusCode != http.StatusOK {
		instanceID = ""
	}

	log.Printf("Retrieved instance ID: %s Status: %d\n", instanceID, resp.StatusCode)
}

func generateToken(url string) string {
	req, err := http.NewRequest(http.MethodPut, url, http.NoBody)
	if err != nil {
		log.Fatalf("Failed to create session token request. %s\n", err.Error())
	}
	req.Header.Add(a, "1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to generate session token. %s\n", err.Error())
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read session token. %s\n", err.Error())
	}

	return string(data)
}
