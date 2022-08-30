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

package imds

// Instance Metadata is divided into categories. To retrieve instance metadata, a category
// is provided within the request. To find a comprehensive description of each category,
// view the official AWS documentation at:
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
const (
	PathAMIID                                 = "ami-id"
	PathAMILaunchIndex                        = "ami-launch-index"
	PathAMIManifestPath                       = "ami-manifest-path"
	PathBlockDeviceMappingAMI                 = "block-device-mapping/ami"
	PathBlockDeviceMappingEBS2                = "block-device-mapping/ebs2"
	PathBlockDeviceMappingRoot                = "block-device-mapping/root"
	PathEventsMaintenanceHistory              = "events/maintenance/history"
	PathEventsMaintenanceScheduled            = "events/maintenance/scheduled"
	PathEventsRecommendationsRebalance        = "events/recommendations/rebalance"
	PathHostname                              = "hostname"
	PathIAMInfo                               = "iam/info"
	PathIAMSecurityCredentials                = "iam/security-credentials/ssm-access"
	PathInstanceAction                        = "instance-action"
	PathInstanceID                            = "instance-id"
	PathInstanceLifecycle                     = "instance-life-cycle"
	PathInstanceType                          = "instance-type"
	PathLocalHostname                         = "local-hostname"
	PathLocalIPv4                             = "local-ipv4"
	PathMAC                                   = "mac"
	PathNetworkInterfaces0DeviceNumber        = "network/interfaces/macs/06:e5:43:29:8f:08/device-number"
	PathNetworkInterfaces0InterfaceID         = "network/interfaces/macs/06:e5:43:29:8f:08/interface-id"
	PathNetworkInterfaces0LocalHostname       = "network/interfaces/macs/06:e5:43:29:8f:08/local-hostname"
	PathNetworkInterfaces0LocalIPv4s          = "network/interfaces/macs/06:e5:43:29:8f:08/local-ipv4s"
	PathNetworkInterfaces0MAC                 = "network/interfaces/macs/06:e5:43:29:8f:08/mac"
	PathNetworkInterfaces0OwnerID             = "network/interfaces/macs/06:e5:43:29:8f:08/owner-id"
	PathNetworkInterfaces0SecurityGroups      = "network/interfaces/macs/06:e5:43:29:8f:08/security-groups"
	PathNetworkInterfaces0SecurityGroupIDs    = "network/interfaces/macs/06:e5:43:29:8f:08/security-group-ids"
	PathNetworkInterfaces0SubnetID            = "network/interfaces/macs/06:e5:43:29:8f:08/subnet-id"
	PathNetworkInterfaces0SubnetIPv4CIDRBlock = "network/interfaces/macs/06:e5:43:29:8f:08/subnet-ipv4-cidr-block"
	PathNetworkInterfaces0VPCID               = "network/interfaces/macs/06:e5:43:29:8f:08/vpc-id"
	PathNetworkInterfaces0VPCIDPv4CIDRBlock   = "network/interfaces/macs/06:e5:43:29:8f:08/vpc-ipv4-cidr-block"
	PathNetworkInterfaces0VPCIDPv4CIDRBlocks  = "network/interfaces/macs/06:e5:43:29:8f:08/vpc-ipv4-cidr-blocks"
	PathNetworkInterfaces0VPCIDPv6CIDRBlocks  = "network/interfaces/macs/06:e5:43:29:8f:08/vpc-ipv6-cidr-blocks"
	PathPlacementAvailabilityZone             = "placement/availability-zone"
	PathPlacementAvailabilityZoneID           = "placement/availability-zone-id"
	PathPlacementRegion                       = "placement/region"
	PathProfile                               = "profile"
	PathPublicKeys0OpenSSHKey                 = "public-keys/0/openssh-key"
	PathReservationID                         = "reservation-id"
	PathSecurityGroups                        = "security-groups"
	PathServicesDomain                        = "services/domain"
	PathServicesPartition                     = "services/partition"
	PathSpotInstanceAction                    = "spot/instance-action"
	PathSpotTerminationTime                   = "spot/termination-time"
	PathTagsInstance                          = "tags/instance"
)

// Instance Metadata values as returned by the Instance Metadata mock for each supported category.
// Values are not provided for the following categories, as the Instance Metadata mock returns
// dated values:
//
//   - events/maintenance/history
//   - events/maintenance/scheduled
//   - events/recommendations/rebalance
//   - spot/instance-action
//   - spot/termination-time
const (
	ValueAMIID                                 = "ami-0e34bbddc66def5ac"
	ValueAMILaunchIndex                        = "0"
	ValueAMIManifestPath                       = "(unknown)"
	ValueBlockDeviceMappingAMI                 = "/dev/xvda"
	ValueBlockDeviceMappingEBS2                = "sdb"
	ValueBlockDeviceMappingRoot                = "/dev/xvda"
	ValueHostname                              = "ip-10-0-1-100.us-east-1.compute.internal"
	ValueIAMInfo                               = valueIAMInfo
	ValueIAMSecurityCredentials                = valueIAMSecurityCredentials
	ValueInstanceAction                        = "none"
	ValueInstanceID                            = "i-0decb1524582da041"
	ValueInstanceLifecycle                     = "on-demand"
	ValueInstanceType                          = "m4.xlarge"
	ValueLocalHostname                         = "ip-10-0-1-100.us-east-1.compute.internal"
	ValueLocalIPv4                             = "10.0.1.100"
	ValueMAC                                   = "06:e5:43:29:8f:08"
	ValueNetworkInterfaces0DeviceNumber        = "0"
	ValueNetworkInterfaces0InterfaceID         = "eni-01180ca4a78168553"
	ValueNetworkInterfaces0LocalHostname       = "ip-10-0-1-100.us-east-1.compute.internal"
	ValueNetworkInterfaces0LocalIPv4s          = "10.0.1.100"
	ValueNetworkInterfaces0MAC                 = "06:e5:43:29:8f:08"
	ValueNetworkInterfaces0OwnerID             = "112233445566"
	ValueNetworkInterfaces0SecurityGroups      = "ssm-sg"
	ValueNetworkInterfaces0SecurityGroupIDs    = "sg-083739656b4679c06"
	ValueNetworkInterfaces0SubnetID            = "subnet-0d908159d6c3e2e54"
	ValueNetworkInterfaces0SubnetIPv4CIDRBlock = "10.0.1.0/24"
	ValueNetworkInterfaces0VPCID               = "vpc-016d173db537793d1"
	ValueNetworkInterfaces0VPCIDPv4CIDRBlock   = "10.0.0.0/16"
	ValueNetworkInterfaces0VPCIDPv4CIDRBlocks  = "10.0.0.0/16"
	ValueNetworkInterfaces0VPCIDPv6CIDRBlocks  = "2a05:d01c:f2d:3200::/56"
	ValuePlacementAvailabilityZone             = "us-east-1a"
	ValuePlacementAvailabilityZoneID           = "use1-az4"
	ValuePlacementRegion                       = "us-east-1"
	ValueProfile                               = "default-hvm"
	ValuePublicKeys0OpenSSHKey                 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCCJB1NiqaDeAmFoLIN7ZKAmoyVEJLP0lY88d7olECKx6yK3bf4S87Yq+yTJ4bpieZBuY2wubNtkH4Qz2/UBVIf4FDCqkWyrrJxbz3keK/X9ZT2lKuc60hK+f6ly77YP+4j8E2E59X6oFrGEMKPdgZHBeljviYh/UU0VcRl9UDd/xXbLj/VIR8UDWrSOsYDfI4mTIun1OFoixW55WvTFHl+Zdm4juldsSnjyZbnTbGG085ixQXAfQm46Z7KBvRF71RRC34McI+a1zNRiULiLS/da4EhIUIX9Po7ZezW+wkls+S7sTketAuRj0H9MDMXTE+f8YfkNlrAsYv1le96RWIL test"
	ValueReservationID                         = "r-0c4dee716c0dbe3c9"
	ValueSecurityGroups                        = `["ssm-sg"]`
	ValueServicesDomain                        = "amazonaws.com"
	ValueServicesPartition                     = "aws"
	ValueTagsInstance                          = "Name"
)

// JSON values returned by the Instance Metadata mock
const (
	valueIAMInfo = `{"Code":"Success","LastUpdated":"2022-08-08T04:25:36Z","InstanceProfileArn":"arn:aws:iam::112233445566:instance-profile/ssm-access","InstanceProfileId":"AIPAYUKXDENX4ZNCZWHF6"}`

	valueIAMSecurityCredentials = `{"Code":"Success","LastUpdated":"2022-08-08T04:26:10Z","Type":"AWS-HMAC","AccessKeyId":"ASIABCDEFGHIJKL","SecretAccessKey":"AAAAAA/abcdefghijnklmnopqrstuvwxyz","Token":"ABCDEFGHIJKLMNOP//////////testing12345/YfenfTTuhJuF3bWoRpkiko7x8NKUMRg==","Expiration":"2022-08-08T11:00:36Z"}`
)

// InstanceTagPath ...
func InstanceTagPath(tag string) string {
	return PathTagsInstance + "/" + tag
}
