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

	imds "github.com/purpleclay/testcontainers-imds"
)

// This example demonstrates how to access EC2 instance tags through the Instance Metadata
// Service (IMDS). When using a real EC2, you must explicitly enable access to instance
// tags using the following guide:
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#allow-access-to-tags-in-IMDS
//
// Details on how the Instance Metadata Mock (imds-mock) handles instance tags can be found
// here: https://docs.purpleclay.dev/imds-mock/configure/instance-tags/
func main() {
	ctx := context.Background()

	container := imds.MustStartWith(ctx, imds.Options{
		InstanceTags: map[string]string{
			"Name":        "mock-ec2",
			"Environment": "development",
			"Role":        "devops",
		},
	})
	defer container.Terminate(ctx)

	log.Println("IMDS mock started with three instance tags...")

	var tag string
	var err error

	if tag, _, err = container.Get(imds.PathTagsInstance + "/" + "Name"); err != nil {
		log.Fatalf("Failed to read instance tag %s. %s\n", tag, err.Error())
	}
	log.Printf("Retrieved tag: Name=%s\n", tag)

	if tag, _, err = container.Get(imds.PathTagsInstance + "/" + "Environment"); err != nil {
		log.Fatalf("Failed to read instance tag %s. %s\n", tag, err.Error())
	}
	log.Printf("Retrieved tag: Environment=%s\n", tag)

	if tag, _, err = container.Get(imds.PathTagsInstance + "/" + "Role"); err != nil {
		log.Fatalf("Failed to read instance tag %s. %s\n", tag, err.Error())
	}
	log.Printf("Retrieved tag: Role=%s\n", tag)
}
