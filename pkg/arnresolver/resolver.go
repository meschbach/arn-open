// Copyright 2020 Mark Eschbach. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package arnresolver

import (
	"errors"
	"strings"
	awsARN "github.com/aws/aws-sdk-go/aws/arn"
)

type ARNTypeToURL func( arn *awsARN.ARN, resourceType string, resourceID string )(string,error)

type ARNResolver struct {
	ResourceTypes map[string]ARNTypeToURL
}

func (m *ARNResolver) Register( resourceType string, mapFunc ARNTypeToURL ) error  {
	if m.ResourceTypes[resourceType] != nil {
		return errors.New("resource type already registered")
	}
	m.ResourceTypes[resourceType] = mapFunc
	return nil
}

func (m *ARNResolver) URLFrom(arn awsARN.ARN) (string,error)  {
	firstColon := strings.Index(arn.Resource, ":")
	firstForwardSlash := strings.Index(arn.Resource,"/")

	var separateAt int
	if firstColon == -1 {
		if firstForwardSlash == -1 {
			return "", errors.New("no resource type separator")
		} else {
			separateAt = firstForwardSlash
		}
	} else if firstForwardSlash == -1 {
		separateAt = firstForwardSlash
	} else if firstForwardSlash < firstColon {
		separateAt = firstForwardSlash
	} else {
		separateAt = firstColon
	}

	resourceType := arn.Resource[0:separateAt]
	resourceID := arn.Resource[separateAt+1:]
	if fn, ok := m.ResourceTypes[resourceType]; ok {
		return fn(&arn,resourceType, resourceID)
	} else {
		return "", errors.New("Unknown resource type "+ resourceType)
	}
}

func NewDefaultAWSResolver() (*ARNResolver) {
	resolver := &ARNResolver{ ResourceTypes: map[string]ARNTypeToURL{} }
	resolver.Register("instance", EC2Instance)
	resolver.Register("role", IAMResolver)
	resolver.Register("route-table", VPCRoutingTable)
	resolver.Register("image", EC2Image )
	resolver.Register("security-group", EC2SecurityGroup )
	return resolver
}
