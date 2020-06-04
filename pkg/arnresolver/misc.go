// Copyright 2020 Mark Eschbach. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package arnresolver

import (
	awsARN "github.com/aws/aws-sdk-go/aws/arn"
	"strings"
)

func VPCRoutingTable( arn *awsARN.ARN, resourceType string, resourceID string ) (string,error)  {
	parts := []string{"https://", arn.Region, ".console.", arn.Partition, ".amazon.com",
		"/vpc/home?&region=", arn.Region, "#RouteTables:routeTableId=" , resourceID, ";sort=routeTableId"}
	return strings.Join(parts, ""), nil
}

func IAMResolver( arn *awsARN.ARN, resourceType string, resourceID string ) (string,error)  {
	resourcePoints := strings.Split(resourceID, "/")
	roleName := resourcePoints[len(resourcePoints) - 1]
	parts := []string{"https://console.", arn.Partition, ".amazon.com", "/iam/home", "#/roles/" , roleName}
	return strings.Join(parts, ""), nil
}

