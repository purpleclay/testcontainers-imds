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

package imds_test

import (
	"testing"

	imds "github.com/purpleclay/testcontainers-imds"
	"github.com/stretchr/testify/require"
)

func TestCategory(t *testing.T) {
	t.Parallel()
	url := startWithDefaults(t)

	t.Run("AMIID", checkIMDSCategory(url, imds.PathAMIID, imds.ValueAMIID))
	t.Run("AMILaunchIndex", checkIMDSCategory(url, imds.PathAMILaunchIndex, imds.ValueAMILaunchIndex))
	t.Run("AMIManifestPath", checkIMDSCategory(url, imds.PathAMIManifestPath, imds.ValueAMIManifestPath))
	t.Run("BlockDeviceMappingAMI", checkIMDSCategory(url, imds.PathBlockDeviceMappingAMI, imds.ValueBlockDeviceMappingAMI))
	t.Run("BlockDeviceMappingEBS0", checkIMDSCategory(url, imds.PathBlockDeviceMappingEBS2, imds.ValueBlockDeviceMappingEBS2))
	t.Run("BlockDeviceMappingRoot", checkIMDSCategory(url, imds.PathBlockDeviceMappingRoot, imds.ValueBlockDeviceMappingRoot))
	t.Run("Hostname", checkIMDSCategory(url, imds.PathHostname, imds.ValueHostname))
	t.Run("IAMInfo", checkIMDSCategory(url, imds.PathIAMInfo, imds.ValueIAMInfo))
	t.Run("IAMSecurityCredentials", checkIMDSCategory(url, imds.PathIAMSecurityCredentials, imds.ValueIAMSecurityCredentials))
	t.Run("InstanceAction", checkIMDSCategory(url, imds.PathInstanceAction, imds.ValueInstanceAction))
	t.Run("InstanceID", checkIMDSCategory(url, imds.PathInstanceID, imds.ValueInstanceID))
	t.Run("InstanceLifecycle", checkIMDSCategory(url, imds.PathInstanceLifecycle, imds.ValueInstanceLifecycle))
	t.Run("InstanceType", checkIMDSCategory(url, imds.PathInstanceType, imds.ValueInstanceType))
	t.Run("LocalHostname", checkIMDSCategory(url, imds.PathLocalHostname, imds.ValueLocalHostname))
	t.Run("LocalIPv4", checkIMDSCategory(url, imds.PathLocalIPv4, imds.ValueLocalIPv4))
	t.Run("MAC", checkIMDSCategory(url, imds.PathMAC, imds.ValueMAC))
	t.Run("NetworkInterfaces0DeviceNumber", checkIMDSCategory(url, imds.PathNetworkInterfaces0DeviceNumber, imds.ValueNetworkInterfaces0DeviceNumber))
	t.Run("NetworkInterfaces0InterfaceID", checkIMDSCategory(url, imds.PathNetworkInterfaces0InterfaceID, imds.ValueNetworkInterfaces0InterfaceID))
	t.Run("NetworkInterfaces0LocalHostname", checkIMDSCategory(url, imds.PathNetworkInterfaces0LocalHostname, imds.ValueNetworkInterfaces0LocalHostname))
	t.Run("NetworkInterfaces0LocalIPv4s", checkIMDSCategory(url, imds.PathNetworkInterfaces0LocalIPv4s, imds.ValueNetworkInterfaces0LocalIPv4s))
	t.Run("NetworkInterfaces0MAC", checkIMDSCategory(url, imds.PathNetworkInterfaces0MAC, imds.ValueNetworkInterfaces0MAC))
	t.Run("NetworkInterfaces0OwnerID ", checkIMDSCategory(url, imds.PathNetworkInterfaces0OwnerID, imds.ValueNetworkInterfaces0OwnerID))
	t.Run("NetworkInterfaces0SecurityGroups", checkIMDSCategory(url, imds.PathNetworkInterfaces0SecurityGroups, imds.ValueNetworkInterfaces0SecurityGroups))
	t.Run("NetworkInterfaces0SecurityGroupIDs", checkIMDSCategory(url, imds.PathNetworkInterfaces0SecurityGroupIDs, imds.ValueNetworkInterfaces0SecurityGroupIDs))
	t.Run("NetworkInterfaces0SubnetID", checkIMDSCategory(url, imds.PathNetworkInterfaces0SubnetID, imds.ValueNetworkInterfaces0SubnetID))
	t.Run("NetworkInterfaces0SubnetIPv4CIDRBlock", checkIMDSCategory(url, imds.PathNetworkInterfaces0SubnetIPv4CIDRBlock, imds.ValueNetworkInterfaces0SubnetIPv4CIDRBlock))
	t.Run("NetworkInterfaces0VPCID", checkIMDSCategory(url, imds.PathNetworkInterfaces0VPCID, imds.ValueNetworkInterfaces0VPCID))
	t.Run("NetworkInterfaces0VPCIDPv4CIDRBlock", checkIMDSCategory(url, imds.PathNetworkInterfaces0VPCIDPv4CIDRBlock, imds.ValueNetworkInterfaces0VPCIDPv4CIDRBlock))
	t.Run("NetworkInterfaces0VPCIDPv4CIDRBlocks", checkIMDSCategory(url, imds.PathNetworkInterfaces0VPCIDPv4CIDRBlocks, imds.ValueNetworkInterfaces0VPCIDPv4CIDRBlocks))
	t.Run("NetworkInterfaces0VPCIDPv6CIDRBlocks", checkIMDSCategory(url, imds.PathNetworkInterfaces0VPCIDPv6CIDRBlocks, imds.ValueNetworkInterfaces0VPCIDPv6CIDRBlocks))
	t.Run("PlacementAvailabilityZone", checkIMDSCategory(url, imds.PathPlacementAvailabilityZone, imds.ValuePlacementAvailabilityZone))
	t.Run("PlacementAvailabilityZoneID", checkIMDSCategory(url, imds.PathPlacementAvailabilityZoneID, imds.ValuePlacementAvailabilityZoneID))
	t.Run("PlacementRegion", checkIMDSCategory(url, imds.PathPlacementRegion, imds.ValuePlacementRegion))
	t.Run("PublicKeys0OpenSSHKey", checkIMDSCategory(url, imds.PathPublicKeys0OpenSSHKey, imds.ValuePublicKeys0OpenSSHKey))
	t.Run("ReservationID", checkIMDSCategory(url, imds.PathReservationID, imds.ValueReservationID))
	t.Run("SecurityGroups", checkIMDSCategory(url, imds.PathSecurityGroups, imds.ValueSecurityGroups))
	t.Run("ServicesDomain", checkIMDSCategory(url, imds.PathServicesDomain, imds.ValueServicesDomain))
	t.Run("ServicesPartition", checkIMDSCategory(url, imds.PathServicesPartition, imds.ValueServicesPartition))
	t.Run("TagsInstance", checkIMDSCategory(url, imds.PathTagsInstance, imds.ValueTagsInstance))
}

// Utility wrapper, tidying up inline test functions
func checkIMDSCategory(url, cat, value string) func(t *testing.T) {
	return func(t *testing.T) {
		v, _ := get(t, url+cat)
		require.Equal(t, value, v)
	}
}
