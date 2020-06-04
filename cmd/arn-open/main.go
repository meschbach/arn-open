// Copyright 2020 Mark Eschbach. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	awsARN "github.com/aws/aws-sdk-go/aws/arn"
	"fmt"
	"os"
	"os/exec"
	arnresolver "github.com/meschbach/arn-open/pkg/arnresolver"
)

func main() {
	/*
	 * Interpret ARN
	 */
	if( len(os.Args) != 2 ){
		fmt.Printf("Usage %s <arn>\n", os.Args[0])
		return
	}

	arnString := os.Args[1]
	arn, err := awsARN.Parse(arnString)
	if err != nil {
		fmt.Printf("Failed to parse ARN %s\n", arnString, err)
		return
	}

	/*
	 * ARN to URL
	 */
	resolver := arnresolver.NewDefaultAWSResolver()
	url, err := resolver.URLFrom(arn)
	if err != nil {
		fmt.Printf("Failed to convert ARN to URL %s\n", err)
		return
	}

	/*
	 * Invoke Open
	 */
	open, err := exec.LookPath("open")
	if err != nil {
		panic(err)
	}

	out, err := exec.Command(open, url).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	if len(output) > 0 {
		fmt.Println(output)
	}
}
