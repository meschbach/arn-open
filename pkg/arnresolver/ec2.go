// Copyright 2020 Mark Eschbach. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package arnresolver

import (
	awsARN "github.com/aws/aws-sdk-go/aws/arn"
	"strings"
)

func EC2Instance( arn *awsARN.ARN, resourceType string, resourceID string ) (string,error)  {
	parts := []string{"https://", arn.Region, ".console.", arn.Partition, ".amazon.com",
		"/ec2/v2/home?&region=", arn.Region, "#Instances:search=" , resourceID, ";sort=tag:Name"}
	return strings.Join(parts, ""), nil
}

func EC2Image( arn *awsARN.ARN, resourceType string, resourceID string ) (string,error)  {
	//TODO: UI Forces either public or private image searches, defaulting to public
	//TODO: Might be a nice option to search through the API first
	parts := []string{"https://", arn.Region, ".console.", arn.Partition, ".amazon.com",
		"/ec2/v2/home?&region=", arn.Region, "#Images:visibility=public-images;search=" , resourceID, ";sort=name"}
	return strings.Join(parts, ""), nil
}


func EC2SecurityGroup( arn *awsARN.ARN, resourceType string, resourceID string ) (string,error)  {
	//https://us-west-2.console.aws.amazon.com/ec2/v2/home?&region=us-west-2#SecurityGroups:group-id=sg-65089901;sort=group-id
	parts := []string{"https://", arn.Region, ".console.", arn.Partition, ".amazon.com",
		"/ec2/v2/home?&region=", arn.Region, "#SecurityGroups:group-id=" , resourceID, ";sort=group-id"}
	return strings.Join(parts, ""), nil
}
