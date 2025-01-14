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

package dbcluster

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/docdb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/docdb/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeDBClustersInput returns input for read
// operation.
func GenerateDescribeDBClustersInput(cr *svcapitypes.DBCluster) *svcsdk.DescribeDBClustersInput {
	res := &svcsdk.DescribeDBClustersInput{}

	return res
}

// GenerateDBCluster returns the current state in the form of *svcapitypes.DBCluster.
func GenerateDBCluster(resp *svcsdk.DescribeDBClustersOutput) *svcapitypes.DBCluster {
	cr := &svcapitypes.DBCluster{}

	found := false
	for _, elem := range resp.DBClusters {
		if elem.AssociatedRoles != nil {
			f0 := []*svcapitypes.DBClusterRole{}
			for _, f0iter := range elem.AssociatedRoles {
				f0elem := &svcapitypes.DBClusterRole{}
				if f0iter.RoleArn != nil {
					f0elem.RoleARN = f0iter.RoleArn
				}
				if f0iter.Status != nil {
					f0elem.Status = f0iter.Status
				}
				f0 = append(f0, f0elem)
			}
			cr.Status.AtProvider.AssociatedRoles = f0
		}
		if elem.AvailabilityZones != nil {
			f1 := []*string{}
			for _, f1iter := range elem.AvailabilityZones {
				var f1elem string
				f1elem = *f1iter
				f1 = append(f1, &f1elem)
			}
			cr.Spec.ForProvider.AvailabilityZones = f1
		}
		if elem.BackupRetentionPeriod != nil {
			cr.Spec.ForProvider.BackupRetentionPeriod = elem.BackupRetentionPeriod
		}
		if elem.ClusterCreateTime != nil {
			cr.Status.AtProvider.ClusterCreateTime = &metav1.Time{*elem.ClusterCreateTime}
		}
		if elem.DBClusterArn != nil {
			cr.Status.AtProvider.DBClusterARN = elem.DBClusterArn
		}
		if elem.DBClusterIdentifier != nil {
			cr.Status.AtProvider.DBClusterIdentifier = elem.DBClusterIdentifier
		}
		if elem.DBClusterMembers != nil {
			f6 := []*svcapitypes.DBClusterMember{}
			for _, f6iter := range elem.DBClusterMembers {
				f6elem := &svcapitypes.DBClusterMember{}
				if f6iter.DBClusterParameterGroupStatus != nil {
					f6elem.DBClusterParameterGroupStatus = f6iter.DBClusterParameterGroupStatus
				}
				if f6iter.DBInstanceIdentifier != nil {
					f6elem.DBInstanceIdentifier = f6iter.DBInstanceIdentifier
				}
				if f6iter.IsClusterWriter != nil {
					f6elem.IsClusterWriter = f6iter.IsClusterWriter
				}
				if f6iter.PromotionTier != nil {
					f6elem.PromotionTier = f6iter.PromotionTier
				}
				f6 = append(f6, f6elem)
			}
			cr.Status.AtProvider.DBClusterMembers = f6
		}
		if elem.DBClusterParameterGroup != nil {
			cr.Status.AtProvider.DBClusterParameterGroup = elem.DBClusterParameterGroup
		}
		if elem.DBSubnetGroup != nil {
			cr.Status.AtProvider.DBSubnetGroup = elem.DBSubnetGroup
		}
		if elem.DbClusterResourceId != nil {
			cr.Status.AtProvider.DBClusterResourceID = elem.DbClusterResourceId
		}
		if elem.DeletionProtection != nil {
			cr.Spec.ForProvider.DeletionProtection = elem.DeletionProtection
		}
		if elem.EarliestRestorableTime != nil {
			cr.Status.AtProvider.EarliestRestorableTime = &metav1.Time{*elem.EarliestRestorableTime}
		}
		if elem.EnabledCloudwatchLogsExports != nil {
			f12 := []*string{}
			for _, f12iter := range elem.EnabledCloudwatchLogsExports {
				var f12elem string
				f12elem = *f12iter
				f12 = append(f12, &f12elem)
			}
			cr.Status.AtProvider.EnabledCloudwatchLogsExports = f12
		}
		if elem.Endpoint != nil {
			cr.Status.AtProvider.Endpoint = elem.Endpoint
		}
		if elem.Engine != nil {
			cr.Spec.ForProvider.Engine = elem.Engine
		}
		if elem.EngineVersion != nil {
			cr.Spec.ForProvider.EngineVersion = elem.EngineVersion
		}
		if elem.HostedZoneId != nil {
			cr.Status.AtProvider.HostedZoneID = elem.HostedZoneId
		}
		if elem.KmsKeyId != nil {
			cr.Spec.ForProvider.KMSKeyID = elem.KmsKeyId
		}
		if elem.LatestRestorableTime != nil {
			cr.Status.AtProvider.LatestRestorableTime = &metav1.Time{*elem.LatestRestorableTime}
		}
		if elem.MasterUsername != nil {
			cr.Spec.ForProvider.MasterUsername = elem.MasterUsername
		}
		if elem.MultiAZ != nil {
			cr.Status.AtProvider.MultiAZ = elem.MultiAZ
		}
		if elem.PercentProgress != nil {
			cr.Status.AtProvider.PercentProgress = elem.PercentProgress
		}
		if elem.Port != nil {
			cr.Spec.ForProvider.Port = elem.Port
		}
		if elem.PreferredBackupWindow != nil {
			cr.Spec.ForProvider.PreferredBackupWindow = elem.PreferredBackupWindow
		}
		if elem.PreferredMaintenanceWindow != nil {
			cr.Spec.ForProvider.PreferredMaintenanceWindow = elem.PreferredMaintenanceWindow
		}
		if elem.ReaderEndpoint != nil {
			cr.Status.AtProvider.ReaderEndpoint = elem.ReaderEndpoint
		}
		if elem.Status != nil {
			cr.Status.AtProvider.Status = elem.Status
		}
		if elem.StorageEncrypted != nil {
			cr.Spec.ForProvider.StorageEncrypted = elem.StorageEncrypted
		}
		if elem.VpcSecurityGroups != nil {
			f28 := []*svcapitypes.VPCSecurityGroupMembership{}
			for _, f28iter := range elem.VpcSecurityGroups {
				f28elem := &svcapitypes.VPCSecurityGroupMembership{}
				if f28iter.Status != nil {
					f28elem.Status = f28iter.Status
				}
				if f28iter.VpcSecurityGroupId != nil {
					f28elem.VPCSecurityGroupID = f28iter.VpcSecurityGroupId
				}
				f28 = append(f28, f28elem)
			}
			cr.Status.AtProvider.VPCSecurityGroups = f28
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDBClusterInput returns a create input.
func GenerateCreateDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.CreateDBClusterInput {
	res := &svcsdk.CreateDBClusterInput{}

	if cr.Spec.ForProvider.AvailabilityZones != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.AvailabilityZones {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetAvailabilityZones(f0)
	}
	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.DBClusterParameterGroupName != nil {
		res.SetDBClusterParameterGroupName(*cr.Spec.ForProvider.DBClusterParameterGroupName)
	}
	if cr.Spec.ForProvider.DBSubnetGroupName != nil {
		res.SetDBSubnetGroupName(*cr.Spec.ForProvider.DBSubnetGroupName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.EnableCloudwatchLogsExports != nil {
		f5 := []*string{}
		for _, f5iter := range cr.Spec.ForProvider.EnableCloudwatchLogsExports {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		res.SetEnableCloudwatchLogsExports(f5)
	}
	if cr.Spec.ForProvider.Engine != nil {
		res.SetEngine(*cr.Spec.ForProvider.Engine)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.MasterUsername != nil {
		res.SetMasterUsername(*cr.Spec.ForProvider.MasterUsername)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreSignedURL != nil {
		res.SetPreSignedUrl(*cr.Spec.ForProvider.PreSignedURL)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.StorageEncrypted != nil {
		res.SetStorageEncrypted(*cr.Spec.ForProvider.StorageEncrypted)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f15 := []*svcsdk.Tag{}
		for _, f15iter := range cr.Spec.ForProvider.Tags {
			f15elem := &svcsdk.Tag{}
			if f15iter.Key != nil {
				f15elem.SetKey(*f15iter.Key)
			}
			if f15iter.Value != nil {
				f15elem.SetValue(*f15iter.Value)
			}
			f15 = append(f15, f15elem)
		}
		res.SetTags(f15)
	}
	if cr.Spec.ForProvider.VPCSecurityGroupIDs != nil {
		f16 := []*string{}
		for _, f16iter := range cr.Spec.ForProvider.VPCSecurityGroupIDs {
			var f16elem string
			f16elem = *f16iter
			f16 = append(f16, &f16elem)
		}
		res.SetVpcSecurityGroupIds(f16)
	}

	return res
}

// GenerateModifyDBClusterInput returns an update input.
func GenerateModifyDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.ModifyDBClusterInput {
	res := &svcsdk.ModifyDBClusterInput{}

	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.DBClusterParameterGroupName != nil {
		res.SetDBClusterParameterGroupName(*cr.Spec.ForProvider.DBClusterParameterGroupName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.VPCSecurityGroupIDs != nil {
		f10 := []*string{}
		for _, f10iter := range cr.Spec.ForProvider.VPCSecurityGroupIDs {
			var f10elem string
			f10elem = *f10iter
			f10 = append(f10, &f10elem)
		}
		res.SetVpcSecurityGroupIds(f10)
	}

	return res
}

// GenerateDeleteDBClusterInput returns a deletion input.
func GenerateDeleteDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.DeleteDBClusterInput {
	res := &svcsdk.DeleteDBClusterInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "DBClusterNotFoundFault"
}
