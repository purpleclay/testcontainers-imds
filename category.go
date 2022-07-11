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

package aemm

// Instance Metadata is divided into categories. To retrieve instance metadata, a category
// is provided within the request. To find a comprehensive description of each category,
// view the official AWS documentation at:
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html
const (
	PathAMIID                                  = "ami-id"
	PathAMILaunchIndex                         = "ami-launch-index"
	PathAMIManifestPath                        = "ami-manifest-path"
	PathBlockDeviceMappingAMI                  = "block-device-mapping/ami"
	PathBlockDeviceMappingEBS0                 = "block-device-mapping/ebs0"
	PathBlockDeviceMappingEphemeral0           = "block-device-mapping/ephemeral0"
	PathBlockDeviceMappingRoot                 = "block-device-mapping/root"
	PathBlockDeviceMappingSwap                 = "block-device-mapping/swap"
	PathElasticInferenceAssociation            = "elastic-inference/associations/eia-bfa21c7904f64a82a21b9f4540169ce1"
	PathEventsMaintenanceScheduled             = "events/maintenance/scheduled"
	PathEventsRecommendationsRebalance         = "events/recommendations/rebalance"
	PathHostname                               = "hostname"
	PathIAMInfo                                = "iam/info"
	PathIAMSecurityCredentials                 = "iam/security-credentials/baskinc-role"
	PathInstanceAction                         = "instance-action"
	PathInstanceID                             = "instance-id"
	PathInstanceLifecycle                      = "instance-life-cycle"
	PathInstanceType                           = "instance-type"
	PathKernelID                               = "kernel-id"
	PathLocalHostname                          = "local-hostname"
	PathLocalIPv4                              = "local-ipv4"
	PathMAC                                    = "mac"
	PathNetworkInterfaces0DeviceNumber         = "network/interfaces/macs/0e:49:61:0f:c3:11/device-number"
	PathNetworkInterfaces0InterfaceID          = "network/interfaces/macs/0e:49:61:0f:c3:11/interface-id"
	PathNetworkInterfaces0IPv4Associations     = "network/interfaces/macs/0e:49:61:0f:c3:11/ipv4-associations/192.0.2.54"
	PathNetworkInterfaces0IPv6s                = "network/interfaces/macs/0e:49:61:0f:c3:11/ipv6s"
	PathNetworkInterfaces0LocalHostname        = "network/interfaces/macs/0e:49:61:0f:c3:11/local-hostname"
	PathNetworkInterfaces0LocalIPv4s           = "network/interfaces/macs/0e:49:61:0f:c3:11/local-ipv4s"
	PathNetworkInterfaces0MAC                  = "network/interfaces/macs/0e:49:61:0f:c3:11/mac"
	PathNetworkInterfaces0NetworkCardIndex     = "network/interfaces/macs/0e:49:61:0f:c3:11/network-card-index"
	PathNetworkInterfaces0OwnerID              = "network/interfaces/macs/0e:49:61:0f:c3:11/owner-id"
	PathNetworkInterfaces0PublicHostname       = "network/interfaces/macs/0e:49:61:0f:c3:11/public-hostname"
	PathNetworkInterfaces0PublicIPv4s          = "network/interfaces/macs/0e:49:61:0f:c3:11/public-ipv4s"
	PathNetworkInterfaces0SecurityGroups       = "network/interfaces/macs/0e:49:61:0f:c3:11/security-groups"
	PathNetworkInterfaces0SecurityGroupIDs     = "network/interfaces/macs/0e:49:61:0f:c3:11/security-group-ids"
	PathNetworkInterfaces0SubnetID             = "network/interfaces/macs/0e:49:61:0f:c3:11/subnet-id"
	PathNetworkInterfaces0SubnetIPv4CIDRBlock  = "network/interfaces/macs/0e:49:61:0f:c3:11/subnet-ipv4-cidr-block"
	PathNetworkInterfaces0SubnetIPv6CIDRBlocks = "network/interfaces/macs/0e:49:61:0f:c3:11/subnet-ipv6-cidr-blocks"
	PathNetworkInterfaces0VPCID                = "network/interfaces/macs/0e:49:61:0f:c3:11/vpc-id"
	PathNetworkInterfaces0VPCIDPv4CIDRBlock    = "network/interfaces/macs/0e:49:61:0f:c3:11/vpc-ipv4-cidr-block"
	PathNetworkInterfaces0VPCIDPv4CIDRBlocks   = "network/interfaces/macs/0e:49:61:0f:c3:11/vpc-ipv4-cidr-blocks"
	PathNetworkInterfaces0VPCIDPv6CIDRBlocks   = "network/interfaces/macs/0e:49:61:0f:c3:11/vpc-ipv6-cidr-blocks"
	PathPlacementAvailabilityZone              = "placement/availability-zone"
	PathPlacementAvailabilityZoneID            = "placement/availability-zone-id"
	PathPlacementGroupName                     = "placement/group-name"
	PathPlacementHostID                        = "placement/host-id"
	PathPlacementPartitionNumber               = "placement/partition-number"
	PathPlacementRegion                        = "placement/region"
	PathProductCodes                           = "product-codes"
	PathPublicHostname                         = "public-hostname"
	PathPublicIPv4                             = "public-ipv4"
	PathPublicKeys0OpenSSHKey                  = "public-keys/0/openssh-key"
	PathRAMDiskID                              = "ramdisk-id"
	PathReservationID                          = "reservation-id"
	PathSecurityGroups                         = "security-groups"
	PathServicesDomain                         = "services/domain"
	PathServicesPartition                      = "services/partition"
	PathSpotInstanceAction                     = "spot/instance-action"
	PathSpotTerminationTime                    = "spot/termination-time"
	PathTagsInstance                           = "tags/instance"
)

// Instance Metadata values as returned by the AEMM mock for each supported category
const (
	ValueAMIID                                  = "ami-0a887e401f7654935"
	ValueAMILaunchIndex                         = "0"
	ValueAMIManifestPath                        = "(unknown)"
	ValueBlockDeviceMappingAMI                  = "/dev/xvda"
	ValueBlockDeviceMappingEBS0                 = "sdb"
	ValueBlockDeviceMappingEphemeral0           = "sdb"
	ValueBlockDeviceMappingRoot                 = "/dev/xvda"
	ValueBlockDeviceMappingSwap                 = "sdcs"
	ValueElasticInferenceAssociation            = `{"version_2018_04_12":{"elastic-inference-accelerator-id":"eia-bfa21c7904f64a82a21b9f4540169ce1","elastic-inference-accelerator-type":"eia1.medium"}}`
	ValueEventsMaintenanceScheduled             = valueEventsMaintenanceScheduled
	ValueEventsRecommendationsRebalance         = valueEventsRecommendationsRebalance
	ValueHostname                               = "ip-172-16-34-43.ec2.internal"
	ValueIAMInfo                                = valueIAMInfo
	ValueIAMSecurityCredentials                 = valueIAMSecurityCredentials
	ValueInstanceAction                         = "none"
	ValueInstanceID                             = "i-1234567890abcdef0"
	ValueInstanceLifecycle                      = "on-demand"
	ValueInstanceType                           = "m4.xlarge"
	ValueKernelID                               = "aki-5c21674b"
	ValueLocalHostname                          = "ip-172-16-34-43.ec2.internal"
	ValueLocalIPv4                              = "172.16.34.43"
	ValueMAC                                    = "0e:49:61:0f:c3:11"
	ValueNetworkInterfaces0DeviceNumber         = "0"
	ValueNetworkInterfaces0InterfaceID          = "eni-0f95d3625f5c521cc"
	ValueNetworkInterfaces0IPv4Associations     = "192.0.2.54"
	ValueNetworkInterfaces0IPv6s                = "2001:db8:8:4::2"
	ValueNetworkInterfaces0LocalHostname        = "ip-172-16-34-43.ec2.internal"
	ValueNetworkInterfaces0LocalIPv4s           = "172.16.34.43"
	ValueNetworkInterfaces0MAC                  = "0e:49:61:0f:c3:11"
	ValueNetworkInterfaces0NetworkCardIndex     = "0"
	ValueNetworkInterfaces0OwnerID              = "515336597381"
	ValueNetworkInterfaces0PublicHostname       = "ec2-192-0-2-54.compute-1.amazonaws.com"
	ValueNetworkInterfaces0PublicIPv4s          = "192.0.2.54"
	ValueNetworkInterfaces0SecurityGroups       = "ura-launch-wizard-harry-1"
	ValueNetworkInterfaces0SecurityGroupIDs     = "sg-0b07f8f6cb485d4df"
	ValueNetworkInterfaces0SubnetID             = "subnet-0ac62554"
	ValueNetworkInterfaces0SubnetIPv4CIDRBlock  = "192.0.2.0/24"
	ValueNetworkInterfaces0SubnetIPv6CIDRBlocks = "2001:db8::/32"
	ValueNetworkInterfaces0VPCID                = "vpc-d295a6a7"
	ValueNetworkInterfaces0VPCIDPv4CIDRBlock    = "192.0.2.0/24"
	ValueNetworkInterfaces0VPCIDPv4CIDRBlocks   = "192.0.2.0/24"
	ValueNetworkInterfaces0VPCIDPv6CIDRBlocks   = "2001:db8::/32"
	ValuePlacementAvailabilityZone              = "us-east-1a"
	ValuePlacementAvailabilityZoneID            = "use1-az4"
	ValuePlacementGroupName                     = "a-placement-group"
	ValuePlacementHostID                        = "h-0da999999f9999fb9"
	ValuePlacementPartitionNumber               = "1"
	ValuePlacementRegion                        = "us-east-1"
	ValueProductCodes                           = "3iplms73etrdhxdepv72l6ywj"
	ValuePublicHostname                         = "ec2-192-0-2-54.compute-1.amazonaws.com"
	ValuePublicIPv4                             = "192.0.2.54"
	ValuePublicKeys0OpenSSHKey                  = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC/JxGByvHDHgQAU+0nRFWdvMPi22OgNUn9ansrI8QN1ZJGxD1ML8DRnJ3Q3zFKqqjGucfNWW0xpVib+ttkIBp8G9P/EOcX9C3FF63O3SnnIUHJsp5faRAZsTJPx0G5HUbvhBvnAcCtSqQgmr02c1l582vAWx48pOmeXXMkl9qe9V/s7K3utmeZkRLo9DqnbsDlg5GWxLC/rWKYaZR66CnMEyZ7yBy3v3abKaGGRovLkHNAgWjSSgmUTI1nT5/S2OLxxuDnsC7+BiABLPaqlIE70SzcWZ0swx68Bo2AY9T9ymGqeAM/1T4yRtg0sPB98TpT7WrY5A3iia2UVtLO/xcTt test"
	ValueRAMDiskID                              = "ari-01bb5768"
	ValueReservationID                          = "r-046cb3eca3e201d2f"
	ValueSecurityGroups                         = "ura-launch-wizard-harry-1"
	ValueServicesDomain                         = "amazonaws.com"
	ValueServicesPartition                      = "aws"
	ValueSpotInstanceAction                     = valueSpotInstanceActionJSON
	ValueSpotTerminationTime                    = "2022-07-11T09:58:52Z"
	ValueTagsInstance                           = `Name
Test`
)

// JSON values returned by the AEMM mock
const (
	valueEventsMaintenanceScheduled = `[
	{
		"Code": "system-reboot",
		"Description": "The instance is scheduled for system-reboot",
		"State": "active",
		"EventId": "instance-event-1234567890abcdef0",
		"NotBefore": "11 Jul 2022 09:11:54 GMT",
		"NotAfter": "18 Jul 2022 09:11:54 GMT",
		"NotBeforeDeadline": "20 Jul 2022 09:11:54 GMT"
	}
]`

	valueEventsRecommendationsRebalance = `{
	"noticeTime": "2022-07-11T10:20:22Z"
}`

	valueIAMInfo = `{
	"Code": "Success",
	"LastUpdated": "2020-04-02T18:50:40Z",
	"InstanceProfileArn": "arn:aws:iam::896453262835:instance-profile/baskinc-role",
	"InstanceProfileId": "AIPA5BOGHHXZELSK34VU4"
}`

	valueIAMSecurityCredentials = `{
	"Code": "Success",
	"LastUpdated": "2020-04-02T18:50:40Z",
	"Type": "AWS-HMAC",
	"AccessKeyId": "12345678901",
	"SecretAccessKey": "v/12345678901",
	"Token": "TEST92test48TEST+y6RpoTEST92test48TEST/8oWVAiBqTEsT5Ky7ty2tEStxC1T==",
	"Expiration": "2020-04-02T00:49:51Z"
}`

	valueSpotInstanceActionJSON = `{
	"action": "terminate",
	"time": "2022-07-11T10:25:54Z"
}`
)
