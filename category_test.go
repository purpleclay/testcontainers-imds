/*
Copyright (c) 2022 - 2023 Purple Clay

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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategory(t *testing.T) {
	t.Parallel()
	container := startWithDefaults(t)

	t.Run("AMIID", checkIMDSCategory(container, imds.PathAMIID, imds.ValueAMIID))
	t.Run("AMILaunchIndex", checkIMDSCategory(container, imds.PathAMILaunchIndex, imds.ValueAMILaunchIndex))
	t.Run("AMIManifestPath", checkIMDSCategory(container, imds.PathAMIManifestPath, imds.ValueAMIManifestPath))
	t.Run("BlockDeviceMappingAMI", checkIMDSCategory(container, imds.PathBlockDeviceMappingAMI, imds.ValueBlockDeviceMappingAMI))
	t.Run("BlockDeviceMappingEBS0", checkIMDSCategory(container, imds.PathBlockDeviceMappingEBS2, imds.ValueBlockDeviceMappingEBS2))
	t.Run("BlockDeviceMappingRoot", checkIMDSCategory(container, imds.PathBlockDeviceMappingRoot, imds.ValueBlockDeviceMappingRoot))
	t.Run("Hostname", checkIMDSCategory(container, imds.PathHostname, imds.ValueHostname))
	t.Run("IAMInfo", checkIMDSCategory(container, imds.PathIAMInfo, imds.ValueIAMInfo))
	t.Run("IAMSecurityCredentials", checkIMDSCategory(container, imds.PathIAMSecurityCredentials, imds.ValueIAMSecurityCredentials))
	t.Run("InstanceAction", checkIMDSCategory(container, imds.PathInstanceAction, imds.ValueInstanceAction))
	t.Run("InstanceID", checkIMDSCategory(container, imds.PathInstanceID, imds.ValueInstanceID))
	t.Run("InstanceLifecycle", checkIMDSCategory(container, imds.PathInstanceLifecycle, imds.ValueInstanceLifecycle))
	t.Run("InstanceType", checkIMDSCategory(container, imds.PathInstanceType, imds.ValueInstanceType))
	t.Run("LocalHostname", checkIMDSCategory(container, imds.PathLocalHostname, imds.ValueLocalHostname))
	t.Run("LocalIPv4", checkIMDSCategory(container, imds.PathLocalIPv4, imds.ValueLocalIPv4))
	t.Run("MAC", checkIMDSCategory(container, imds.PathMAC, imds.ValueMAC))
	t.Run("NetworkInterfaces0DeviceNumber", checkIMDSCategory(container, imds.PathNetworkInterfaces0DeviceNumber, imds.ValueNetworkInterfaces0DeviceNumber))
	t.Run("NetworkInterfaces0InterfaceID", checkIMDSCategory(container, imds.PathNetworkInterfaces0InterfaceID, imds.ValueNetworkInterfaces0InterfaceID))
	t.Run("NetworkInterfaces0LocalHostname", checkIMDSCategory(container, imds.PathNetworkInterfaces0LocalHostname, imds.ValueNetworkInterfaces0LocalHostname))
	t.Run("NetworkInterfaces0LocalIPv4s", checkIMDSCategory(container, imds.PathNetworkInterfaces0LocalIPv4s, imds.ValueNetworkInterfaces0LocalIPv4s))
	t.Run("NetworkInterfaces0MAC", checkIMDSCategory(container, imds.PathNetworkInterfaces0MAC, imds.ValueNetworkInterfaces0MAC))
	t.Run("NetworkInterfaces0OwnerID ", checkIMDSCategory(container, imds.PathNetworkInterfaces0OwnerID, imds.ValueNetworkInterfaces0OwnerID))
	t.Run("NetworkInterfaces0SecurityGroups", checkIMDSCategory(container, imds.PathNetworkInterfaces0SecurityGroups, imds.ValueNetworkInterfaces0SecurityGroups))
	t.Run("NetworkInterfaces0SecurityGroupIDs", checkIMDSCategory(container, imds.PathNetworkInterfaces0SecurityGroupIDs, imds.ValueNetworkInterfaces0SecurityGroupIDs))
	t.Run("NetworkInterfaces0SubnetID", checkIMDSCategory(container, imds.PathNetworkInterfaces0SubnetID, imds.ValueNetworkInterfaces0SubnetID))
	t.Run("NetworkInterfaces0SubnetIPv4CIDRBlock", checkIMDSCategory(container, imds.PathNetworkInterfaces0SubnetIPv4CIDRBlock, imds.ValueNetworkInterfaces0SubnetIPv4CIDRBlock))
	t.Run("NetworkInterfaces0VPCID", checkIMDSCategory(container, imds.PathNetworkInterfaces0VPCID, imds.ValueNetworkInterfaces0VPCID))
	t.Run("NetworkInterfaces0VPCIDPv4CIDRBlock", checkIMDSCategory(container, imds.PathNetworkInterfaces0VPCIDPv4CIDRBlock, imds.ValueNetworkInterfaces0VPCIDPv4CIDRBlock))
	t.Run("NetworkInterfaces0VPCIDPv4CIDRBlocks", checkIMDSCategory(container, imds.PathNetworkInterfaces0VPCIDPv4CIDRBlocks, imds.ValueNetworkInterfaces0VPCIDPv4CIDRBlocks))
	t.Run("NetworkInterfaces0VPCIDPv6CIDRBlocks", checkIMDSCategory(container, imds.PathNetworkInterfaces0VPCIDPv6CIDRBlocks, imds.ValueNetworkInterfaces0VPCIDPv6CIDRBlocks))
	t.Run("PlacementAvailabilityZone", checkIMDSCategory(container, imds.PathPlacementAvailabilityZone, imds.ValuePlacementAvailabilityZone))
	t.Run("PlacementAvailabilityZoneID", checkIMDSCategory(container, imds.PathPlacementAvailabilityZoneID, imds.ValuePlacementAvailabilityZoneID))
	t.Run("PlacementRegion", checkIMDSCategory(container, imds.PathPlacementRegion, imds.ValuePlacementRegion))
	t.Run("PublicKeys0OpenSSHKey", checkIMDSCategory(container, imds.PathPublicKeys0OpenSSHKey, imds.ValuePublicKeys0OpenSSHKey))
	t.Run("ReservationID", checkIMDSCategory(container, imds.PathReservationID, imds.ValueReservationID))
	t.Run("SecurityGroups", checkIMDSCategory(container, imds.PathSecurityGroups, imds.ValueSecurityGroups))
	t.Run("ServicesDomain", checkIMDSCategory(container, imds.PathServicesDomain, imds.ValueServicesDomain))
	t.Run("ServicesPartition", checkIMDSCategory(container, imds.PathServicesPartition, imds.ValueServicesPartition))
	t.Run("TagsInstance", checkIMDSCategory(container, imds.PathTagsInstance, imds.ValueTagsInstance))
}

// Utility wrapper, tidying up inline test functions
func checkIMDSCategory(container *imds.Container, category, value string) func(t *testing.T) {
	return func(t *testing.T) {
		v, _, _ := container.Get(category)
		require.Equal(t, value, v)
	}
}

func TestInstanceTagPath(t *testing.T) {
	assert.Equal(t, "tags/instance/Name", imds.InstanceTagPath("Name"))
}
