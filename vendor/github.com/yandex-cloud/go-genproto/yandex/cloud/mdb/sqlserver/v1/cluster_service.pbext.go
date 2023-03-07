// Code generated by protoc-gen-goext. DO NOT EDIT.

package sqlserver

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/sqlserver/v1/config"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *GetClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClustersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListClustersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClustersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClustersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClustersResponse) SetClusters(v []*Cluster) {
	m.Clusters = v
}

func (m *ListClustersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *CreateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *CreateClusterRequest) SetDatabaseSpecs(v []*DatabaseSpec) {
	m.DatabaseSpecs = v
}

func (m *CreateClusterRequest) SetUserSpecs(v []*UserSpec) {
	m.UserSpecs = v
}

func (m *CreateClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateClusterRequest) SetSqlcollation(v string) {
	m.Sqlcollation = v
}

func (m *CreateClusterRequest) SetHostGroupIds(v []string) {
	m.HostGroupIds = v
}

func (m *CreateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterRequest) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RestoreClusterRequest) SetTime(v *timestamppb.Timestamp) {
	m.Time = v
}

func (m *RestoreClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *RestoreClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *RestoreClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *RestoreClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *RestoreClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *RestoreClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *RestoreClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *RestoreClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *RestoreClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *RestoreClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *RestoreClusterRequest) SetHostGroupIds(v []string) {
	m.HostGroupIds = v
}

func (m *RestoreClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *RestoreClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterMetadata) SetBackupId(v string) {
	m.BackupId = v
}

func (m *StartClusterFailoverRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterFailoverRequest) SetHostName(v string) {
	m.HostName = v
}

func (m *StartClusterFailoverMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *LogRecord) SetTimestamp(v *timestamppb.Timestamp) {
	m.Timestamp = v
}

func (m *LogRecord) SetMessage(v map[string]string) {
	m.Message = v
}

func (m *ListClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *ListClusterLogsRequest) SetServiceType(v ListClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *ListClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *ListClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *ListClusterLogsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterLogsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterLogsRequest) SetAlwaysNextPageToken(v bool) {
	m.AlwaysNextPageToken = v
}

func (m *ListClusterLogsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterLogsResponse) SetLogs(v []*LogRecord) {
	m.Logs = v
}

func (m *ListClusterLogsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterOperationsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListClusterOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterBackupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterBackupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterBackupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterBackupsResponse) SetBackups(v []*Backup) {
	m.Backups = v
}

func (m *ListClusterBackupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *StartClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterMetadata) SetSourceFolderId(v string) {
	m.SourceFolderId = v
}

func (m *MoveClusterMetadata) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *UpdateClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *HostSpec) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *HostSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *HostSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *UpdateHostSpec) SetHostName(v string) {
	m.HostName = v
}

func (m *UpdateHostSpec) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateHostSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *UpdateClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterHostsRequest) SetUpdateHostSpecs(v []*UpdateHostSpec) {
	m.UpdateHostSpecs = v
}

type ConfigSpec_SqlserverConfig = isConfigSpec_SqlserverConfig

func (m *ConfigSpec) SetSqlserverConfig(v ConfigSpec_SqlserverConfig) {
	m.SqlserverConfig = v
}

func (m *ConfigSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigSpec) SetSqlserverConfig_2016Sp2Std(v *config.SQLServerConfig2016Sp2Std) {
	m.SqlserverConfig = &ConfigSpec_SqlserverConfig_2016Sp2Std{
		SqlserverConfig_2016Sp2Std: v,
	}
}

func (m *ConfigSpec) SetSqlserverConfig_2016Sp2Ent(v *config.SQLServerConfig2016Sp2Ent) {
	m.SqlserverConfig = &ConfigSpec_SqlserverConfig_2016Sp2Ent{
		SqlserverConfig_2016Sp2Ent: v,
	}
}

func (m *ConfigSpec) SetSqlserverConfig_2017Std(v *config.SQLServerConfig2017Std) {
	m.SqlserverConfig = &ConfigSpec_SqlserverConfig_2017Std{
		SqlserverConfig_2017Std: v,
	}
}

func (m *ConfigSpec) SetSqlserverConfig_2017Ent(v *config.SQLServerConfig2017Ent) {
	m.SqlserverConfig = &ConfigSpec_SqlserverConfig_2017Ent{
		SqlserverConfig_2017Ent: v,
	}
}

func (m *ConfigSpec) SetSqlserverConfig_2019Std(v *config.SQLServerConfig2019Std) {
	m.SqlserverConfig = &ConfigSpec_SqlserverConfig_2019Std{
		SqlserverConfig_2019Std: v,
	}
}

func (m *ConfigSpec) SetSqlserverConfig_2019Ent(v *config.SQLServerConfig2019Ent) {
	m.SqlserverConfig = &ConfigSpec_SqlserverConfig_2019Ent{
		SqlserverConfig_2019Ent: v,
	}
}

func (m *ConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ConfigSpec) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ConfigSpec) SetAccess(v *Access) {
	m.Access = v
}

func (m *ConfigSpec) SetSecondaryConnections(v ClusterConfig_SecondaryConnections) {
	m.SecondaryConnections = v
}