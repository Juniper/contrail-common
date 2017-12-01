package db
// virtual_machine_interface

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`fq_name`,`virtual_machine_interface_bindings`,`virtual_machine_interface_disable_policy`,`vrf_assign_rule`,`port_security_enabled`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`mac_address`,`vlan_tag_based_bridge_domain`,`virtual_machine_interface_device_owner`,`dhcp_option`,`route`,`allowed_address_pair`,`virtual_machine_interface_fat_flow_protocols`,`local_preference`,`encapsulation`,`juniper_header`,`routing_instance`,`udp_port`,`nic_assisted_mirroring`,`vtep_dst_mac_address`,`vni`,`vtep_dst_ip_address`,`nh_mode`,`nic_assisted_mirroring_vlan`,`analyzer_ip_address`,`analyzer_mac_address`,`analyzer_name`,`traffic_direction`,`service_interface_type`,`sub_interface_vlan_tag`,`uuid`,`display_name`,`key_value_pair`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set `fq_name` = ?,`virtual_machine_interface_bindings` = ?,`virtual_machine_interface_disable_policy` = ?,`vrf_assign_rule` = ?,`port_security_enabled` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`mac_address` = ?,`vlan_tag_based_bridge_domain` = ?,`virtual_machine_interface_device_owner` = ?,`dhcp_option` = ?,`route` = ?,`allowed_address_pair` = ?,`virtual_machine_interface_fat_flow_protocols` = ?,`local_preference` = ?,`encapsulation` = ?,`juniper_header` = ?,`routing_instance` = ?,`udp_port` = ?,`nic_assisted_mirroring` = ?,`vtep_dst_mac_address` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`nh_mode` = ?,`nic_assisted_mirroring_vlan` = ?,`analyzer_ip_address` = ?,`analyzer_mac_address` = ?,`analyzer_name` = ?,`traffic_direction` = ?,`service_interface_type` = ?,`sub_interface_vlan_tag` = ?,`uuid` = ?,`display_name` = ?,`key_value_pair` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface`"
const selectVirtualMachineInterfaceQuery = "select `fq_name`,`virtual_machine_interface_bindings`,`virtual_machine_interface_disable_policy`,`vrf_assign_rule`,`port_security_enabled`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`mac_address`,`vlan_tag_based_bridge_domain`,`virtual_machine_interface_device_owner`,`dhcp_option`,`route`,`allowed_address_pair`,`virtual_machine_interface_fat_flow_protocols`,`local_preference`,`encapsulation`,`juniper_header`,`routing_instance`,`udp_port`,`nic_assisted_mirroring`,`vtep_dst_mac_address`,`vni`,`vtep_dst_ip_address`,`nh_mode`,`nic_assisted_mirroring_vlan`,`analyzer_ip_address`,`analyzer_mac_address`,`analyzer_name`,`traffic_direction`,`service_interface_type`,`sub_interface_vlan_tag`,`uuid`,`display_name`,`key_value_pair`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `virtual_machine_interface`"

func CreateVirtualMachineInterface(tx *sql.Tx, model *models.VirtualMachineInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineInterfaceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.VirtualMachineInterfaceBindings,
    model.VirtualMachineInterfaceDisablePolicy,
    model.VRFAssignTable.VRFAssignRule,
    model.PortSecurityEnabled,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.VirtualMachineInterfaceMacAddresses.MacAddress,
    model.VlanTagBasedBridgeDomain,
    model.VirtualMachineInterfaceDeviceOwner,
    model.VirtualMachineInterfaceDHCPOptionList.DHCPOption,
    model.VirtualMachineInterfaceHostRoutes.Route,
    model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair,
    model.VirtualMachineInterfaceFatFlowProtocols,
    model.VirtualMachineInterfaceProperties.LocalPreference,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName,
    model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection,
    model.VirtualMachineInterfaceProperties.ServiceInterfaceType,
    model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag,
    model.UUID,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.EcmpHashingIncludeFields.IPProtocol,
    model.EcmpHashingIncludeFields.SourceIP,
    model.EcmpHashingIncludeFields.HashingConfigured,
    model.EcmpHashingIncludeFields.SourcePort,
    model.EcmpHashingIncludeFields.DestinationPort,
    model.EcmpHashingIncludeFields.DestinationIP,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess)
    return err
}

func ListVirtualMachineInterface(tx *sql.Tx) ([]*models.VirtualMachineInterface, error) {
    result := models.MakeVirtualMachineInterfaceSlice()
    rows, err := tx.Query(selectVirtualMachineInterfaceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualMachineInterface()
            if err := rows.Scan(&m.FQName,
                &m.VirtualMachineInterfaceBindings,
                &m.VirtualMachineInterfaceDisablePolicy,
                &m.VRFAssignTable.VRFAssignRule,
                &m.PortSecurityEnabled,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.VirtualMachineInterfaceMacAddresses.MacAddress,
                &m.VlanTagBasedBridgeDomain,
                &m.VirtualMachineInterfaceDeviceOwner,
                &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption,
                &m.VirtualMachineInterfaceHostRoutes.Route,
                &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair,
                &m.VirtualMachineInterfaceFatFlowProtocols,
                &m.VirtualMachineInterfaceProperties.LocalPreference,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName,
                &m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection,
                &m.VirtualMachineInterfaceProperties.ServiceInterfaceType,
                &m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag,
                &m.UUID,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.EcmpHashingIncludeFields.IPProtocol,
                &m.EcmpHashingIncludeFields.SourceIP,
                &m.EcmpHashingIncludeFields.HashingConfigured,
                &m.EcmpHashingIncludeFields.SourcePort,
                &m.EcmpHashingIncludeFields.DestinationPort,
                &m.EcmpHashingIncludeFields.DestinationIP,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowVirtualMachineInterface(db *sql.DB, id string, model *models.VirtualMachineInterface) error {
    return nil
}

func UpdateVirtualMachineInterface(db *sql.DB, id string, model *models.VirtualMachineInterface) error {
    return nil
}

func DeleteVirtualMachineInterface(db *sql.DB, id string) error {
    return nil
}