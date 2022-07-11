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

package aemm_test

import (
	"testing"

	aemm "github.com/purpleclay/testcontainers-aemm"
	"github.com/stretchr/testify/require"
)

func TestCategory(t *testing.T) {
	t.Parallel()
	url := startWithDefaults(t)

	t.Run("AMIID", checkIMDSCategory(url, aemm.PathAMIID, aemm.ValueAMIID))
	t.Run("AMILaunchIndex", checkIMDSCategory(url, aemm.PathAMILaunchIndex, aemm.ValueAMILaunchIndex))
	t.Run("AMIManifestPath", checkIMDSCategory(url, aemm.PathAMIManifestPath, aemm.ValueAMIManifestPath))
	t.Run("BlockDeviceMappingAMI", checkIMDSCategory(url, aemm.PathBlockDeviceMappingAMI, aemm.ValueBlockDeviceMappingAMI))
	t.Run("BlockDeviceMappingEBS0", checkIMDSCategory(url, aemm.PathBlockDeviceMappingEBS0, aemm.ValueBlockDeviceMappingEBS0))
	t.Run("BlockDeviceMappingEphemeral0", checkIMDSCategory(url, aemm.PathBlockDeviceMappingEphemeral0, aemm.ValueBlockDeviceMappingEphemeral0))
	t.Run("BlockDeviceMappingRoot", checkIMDSCategory(url, aemm.PathBlockDeviceMappingRoot, aemm.ValueBlockDeviceMappingRoot))
	t.Run("BlockDeviceMappingSwap", checkIMDSCategory(url, aemm.PathBlockDeviceMappingSwap, aemm.ValueBlockDeviceMappingSwap))
	t.Run("ElasticInferenceAssociation", checkIMDSCategory(url, aemm.PathElasticInferenceAssociation, aemm.ValueElasticInferenceAssociation))
	t.Run("Hostname", checkIMDSCategory(url, aemm.PathHostname, aemm.ValueHostname))
	t.Run("IAMInfo", checkIMDSCategory(url, aemm.PathIAMInfo, aemm.ValueIAMInfo))
	t.Run("IAMSecurityCredentials", checkIMDSCategory(url, aemm.PathIAMSecurityCredentials, aemm.ValueIAMSecurityCredentials))
	t.Run("InstanceAction", checkIMDSCategory(url, aemm.PathInstanceAction, aemm.ValueInstanceAction))
	t.Run("InstanceID", checkIMDSCategory(url, aemm.PathInstanceID, aemm.ValueInstanceID))
	t.Run("InstanceLifecycle", checkIMDSCategory(url, aemm.PathInstanceLifecycle, aemm.ValueInstanceLifecycle))
	t.Run("InstanceType", checkIMDSCategory(url, aemm.PathInstanceType, aemm.ValueInstanceType))
	t.Run("KernelID", checkIMDSCategory(url, aemm.PathKernelID, aemm.ValueKernelID))
	t.Run("LocalHostname", checkIMDSCategory(url, aemm.PathLocalHostname, aemm.ValueLocalHostname))
	t.Run("LocalIPv4", checkIMDSCategory(url, aemm.PathLocalIPv4, aemm.ValueLocalIPv4))
	t.Run("MAC", checkIMDSCategory(url, aemm.PathMAC, aemm.ValueMAC))
	t.Run("NetworkInterfaces0DeviceNumber", checkIMDSCategory(url, aemm.PathNetworkInterfaces0DeviceNumber, aemm.ValueNetworkInterfaces0DeviceNumber))
	t.Run("NetworkInterfaces0InterfaceID", checkIMDSCategory(url, aemm.PathNetworkInterfaces0InterfaceID, aemm.ValueNetworkInterfaces0InterfaceID))
	t.Run("NetworkInterfaces0IPv4Associations", checkIMDSCategory(url, aemm.PathNetworkInterfaces0IPv4Associations, aemm.ValueNetworkInterfaces0IPv4Associations))
	t.Run("NetworkInterfaces0IPv6s", checkIMDSCategory(url, aemm.PathNetworkInterfaces0IPv6s, aemm.ValueNetworkInterfaces0IPv6s))
	t.Run("NetworkInterfaces0LocalHostname", checkIMDSCategory(url, aemm.PathNetworkInterfaces0LocalHostname, aemm.ValueNetworkInterfaces0LocalHostname))
	t.Run("NetworkInterfaces0LocalIPv4s", checkIMDSCategory(url, aemm.PathNetworkInterfaces0LocalIPv4s, aemm.ValueNetworkInterfaces0LocalIPv4s))
	t.Run("NetworkInterfaces0MAC", checkIMDSCategory(url, aemm.PathNetworkInterfaces0MAC, aemm.ValueNetworkInterfaces0MAC))
	t.Run("NetworkInterfaces0NetworkCardIndex", checkIMDSCategory(url, aemm.PathNetworkInterfaces0NetworkCardIndex, aemm.ValueNetworkInterfaces0NetworkCardIndex))
	t.Run("NetworkInterfaces0OwnerID ", checkIMDSCategory(url, aemm.PathNetworkInterfaces0OwnerID, aemm.ValueNetworkInterfaces0OwnerID))
	t.Run("NetworkInterfaces0PublicHostname", checkIMDSCategory(url, aemm.PathNetworkInterfaces0PublicHostname, aemm.ValueNetworkInterfaces0PublicHostname))
	t.Run("NetworkInterfaces0PublicIPv4s ", checkIMDSCategory(url, aemm.PathNetworkInterfaces0PublicIPv4s, aemm.ValueNetworkInterfaces0PublicIPv4s))
	t.Run("NetworkInterfaces0SecurityGroups", checkIMDSCategory(url, aemm.PathNetworkInterfaces0SecurityGroups, aemm.ValueNetworkInterfaces0SecurityGroups))
	t.Run("NetworkInterfaces0SecurityGroupIDs", checkIMDSCategory(url, aemm.PathNetworkInterfaces0SecurityGroupIDs, aemm.ValueNetworkInterfaces0SecurityGroupIDs))
	t.Run("NetworkInterfaces0SubnetID", checkIMDSCategory(url, aemm.PathNetworkInterfaces0SubnetID, aemm.ValueNetworkInterfaces0SubnetID))
	t.Run("NetworkInterfaces0SubnetIPv4CIDRBlock", checkIMDSCategory(url, aemm.PathNetworkInterfaces0SubnetIPv4CIDRBlock, aemm.ValueNetworkInterfaces0SubnetIPv4CIDRBlock))
	t.Run("NetworkInterfaces0SubnetIPv6CIDRBlocks", checkIMDSCategory(url, aemm.PathNetworkInterfaces0SubnetIPv6CIDRBlocks, aemm.ValueNetworkInterfaces0SubnetIPv6CIDRBlocks))
	t.Run("NetworkInterfaces0VPCID", checkIMDSCategory(url, aemm.PathNetworkInterfaces0VPCID, aemm.ValueNetworkInterfaces0VPCID))
	t.Run("NetworkInterfaces0VPCIDPv4CIDRBlock", checkIMDSCategory(url, aemm.PathNetworkInterfaces0VPCIDPv4CIDRBlock, aemm.ValueNetworkInterfaces0VPCIDPv4CIDRBlock))
	t.Run("NetworkInterfaces0VPCIDPv4CIDRBlocks", checkIMDSCategory(url, aemm.PathNetworkInterfaces0VPCIDPv4CIDRBlocks, aemm.ValueNetworkInterfaces0VPCIDPv4CIDRBlocks))
	t.Run("NetworkInterfaces0VPCIDPv6CIDRBlocks", checkIMDSCategory(url, aemm.PathNetworkInterfaces0VPCIDPv6CIDRBlocks, aemm.ValueNetworkInterfaces0VPCIDPv6CIDRBlocks))
	t.Run("PlacementAvailabilityZone", checkIMDSCategory(url, aemm.PathPlacementAvailabilityZone, aemm.ValuePlacementAvailabilityZone))
	t.Run("PlacementAvailabilityZoneID", checkIMDSCategory(url, aemm.PathPlacementAvailabilityZoneID, aemm.ValuePlacementAvailabilityZoneID))
	t.Run("PlacementGroupName", checkIMDSCategory(url, aemm.PathPlacementGroupName, aemm.ValuePlacementGroupName))
	t.Run("PlacementHostID", checkIMDSCategory(url, aemm.PathPlacementHostID, aemm.ValuePlacementHostID))
	t.Run("PlacementPartitionNumber", checkIMDSCategory(url, aemm.PathPlacementPartitionNumber, aemm.ValuePlacementPartitionNumber))
	t.Run("PlacementRegion", checkIMDSCategory(url, aemm.PathPlacementRegion, aemm.ValuePlacementRegion))
	t.Run("ProductCodes", checkIMDSCategory(url, aemm.PathProductCodes, aemm.ValueProductCodes))
	t.Run("PublicHostname", checkIMDSCategory(url, aemm.PathPublicHostname, aemm.ValuePublicHostname))
	t.Run("PublicIPv4", checkIMDSCategory(url, aemm.PathPublicIPv4, aemm.ValuePublicIPv4))
	t.Run("PublicKeys0OpenSSHKey", checkIMDSCategory(url, aemm.PathPublicKeys0OpenSSHKey, aemm.ValuePublicKeys0OpenSSHKey))
	t.Run("RAMDiskID", checkIMDSCategory(url, aemm.PathRAMDiskID, aemm.ValueRAMDiskID))
	t.Run("ReservationID", checkIMDSCategory(url, aemm.PathReservationID, aemm.ValueReservationID))
	t.Run("SecurityGroups", checkIMDSCategory(url, aemm.PathSecurityGroups, aemm.ValueSecurityGroups))
	t.Run("ServicesDomain", checkIMDSCategory(url, aemm.PathServicesDomain, aemm.ValueServicesDomain))
	t.Run("ServicesPartition", checkIMDSCategory(url, aemm.PathServicesPartition, aemm.ValueServicesPartition))
	t.Run("TagsInstance", checkIMDSCategory(url, aemm.PathTagsInstance, aemm.ValueTagsInstance))
}

// Utility wrapper, tidying up inline test functions
func checkIMDSCategory(url, cat, value string) func(t *testing.T) {
	return func(t *testing.T) {
		v, _ := get(t, url+cat)
		require.Equal(t, value, v)
	}
}
