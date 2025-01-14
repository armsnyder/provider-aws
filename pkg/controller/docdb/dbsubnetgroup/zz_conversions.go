/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package dbsubnetgroup

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/docdb"

	svcapitypes "github.com/crossplane/provider-aws/apis/docdb/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeDBSubnetGroupsInput returns input for read
// operation.
func GenerateDescribeDBSubnetGroupsInput(cr *svcapitypes.DBSubnetGroup) *svcsdk.DescribeDBSubnetGroupsInput {
	res := &svcsdk.DescribeDBSubnetGroupsInput{}

	return res
}

// GenerateDBSubnetGroup returns the current state in the form of *svcapitypes.DBSubnetGroup.
func GenerateDBSubnetGroup(resp *svcsdk.DescribeDBSubnetGroupsOutput) *svcapitypes.DBSubnetGroup {
	cr := &svcapitypes.DBSubnetGroup{}

	found := false
	for _, elem := range resp.DBSubnetGroups {
		if elem.DBSubnetGroupArn != nil {
			cr.Status.AtProvider.DBSubnetGroupARN = elem.DBSubnetGroupArn
		}
		if elem.DBSubnetGroupDescription != nil {
			cr.Spec.ForProvider.DBSubnetGroupDescription = elem.DBSubnetGroupDescription
		}
		if elem.DBSubnetGroupName != nil {
			cr.Status.AtProvider.DBSubnetGroupName = elem.DBSubnetGroupName
		}
		if elem.SubnetGroupStatus != nil {
			cr.Status.AtProvider.SubnetGroupStatus = elem.SubnetGroupStatus
		}
		if elem.Subnets != nil {
			f4 := []*svcapitypes.Subnet{}
			for _, f4iter := range elem.Subnets {
				f4elem := &svcapitypes.Subnet{}
				if f4iter.SubnetAvailabilityZone != nil {
					f4elemf0 := &svcapitypes.AvailabilityZone{}
					if f4iter.SubnetAvailabilityZone.Name != nil {
						f4elemf0.Name = f4iter.SubnetAvailabilityZone.Name
					}
					f4elem.SubnetAvailabilityZone = f4elemf0
				}
				if f4iter.SubnetIdentifier != nil {
					f4elem.SubnetIdentifier = f4iter.SubnetIdentifier
				}
				if f4iter.SubnetStatus != nil {
					f4elem.SubnetStatus = f4iter.SubnetStatus
				}
				f4 = append(f4, f4elem)
			}
			cr.Status.AtProvider.Subnets = f4
		}
		if elem.VpcId != nil {
			cr.Status.AtProvider.VPCID = elem.VpcId
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDBSubnetGroupInput returns a create input.
func GenerateCreateDBSubnetGroupInput(cr *svcapitypes.DBSubnetGroup) *svcsdk.CreateDBSubnetGroupInput {
	res := &svcsdk.CreateDBSubnetGroupInput{}

	if cr.Spec.ForProvider.DBSubnetGroupDescription != nil {
		res.SetDBSubnetGroupDescription(*cr.Spec.ForProvider.DBSubnetGroupDescription)
	}
	if cr.Spec.ForProvider.SubnetIDs != nil {
		f1 := []*string{}
		for _, f1iter := range cr.Spec.ForProvider.SubnetIDs {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetSubnetIds(f1)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f2 := []*svcsdk.Tag{}
		for _, f2iter := range cr.Spec.ForProvider.Tags {
			f2elem := &svcsdk.Tag{}
			if f2iter.Key != nil {
				f2elem.SetKey(*f2iter.Key)
			}
			if f2iter.Value != nil {
				f2elem.SetValue(*f2iter.Value)
			}
			f2 = append(f2, f2elem)
		}
		res.SetTags(f2)
	}

	return res
}

// GenerateModifyDBSubnetGroupInput returns an update input.
func GenerateModifyDBSubnetGroupInput(cr *svcapitypes.DBSubnetGroup) *svcsdk.ModifyDBSubnetGroupInput {
	res := &svcsdk.ModifyDBSubnetGroupInput{}

	if cr.Spec.ForProvider.DBSubnetGroupDescription != nil {
		res.SetDBSubnetGroupDescription(*cr.Spec.ForProvider.DBSubnetGroupDescription)
	}
	if cr.Spec.ForProvider.SubnetIDs != nil {
		f1 := []*string{}
		for _, f1iter := range cr.Spec.ForProvider.SubnetIDs {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetSubnetIds(f1)
	}

	return res
}

// GenerateDeleteDBSubnetGroupInput returns a deletion input.
func GenerateDeleteDBSubnetGroupInput(cr *svcapitypes.DBSubnetGroup) *svcsdk.DeleteDBSubnetGroupInput {
	res := &svcsdk.DeleteDBSubnetGroupInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "DBSubnetGroupNotFoundFault"
}
