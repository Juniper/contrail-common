package db

// virtual_machine_interface

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`key_value_pair`,`route`,`dhcp_option`,`vlan_tag_based_bridge_domain`,`fq_name`,`global_access`,`share`,`owner`,`owner_access`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`virtual_machine_interface_bindings`,`virtual_machine_interface_device_owner`,`vrf_assign_rule`,`uuid`,`virtual_machine_interface_disable_policy`,`allowed_address_pair`,`virtual_machine_interface_fat_flow_protocols`,`local_preference`,`traffic_direction`,`encapsulation`,`nh_mode`,`udp_port`,`juniper_header`,`analyzer_ip_address`,`analyzer_mac_address`,`nic_assisted_mirroring_vlan`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_name`,`routing_instance`,`nic_assisted_mirroring`,`service_interface_type`,`sub_interface_vlan_tag`,`mac_address`,`port_security_enabled`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set `key_value_pair` = ?,`route` = ?,`dhcp_option` = ?,`vlan_tag_based_bridge_domain` = ?,`fq_name` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`virtual_machine_interface_bindings` = ?,`virtual_machine_interface_device_owner` = ?,`vrf_assign_rule` = ?,`uuid` = ?,`virtual_machine_interface_disable_policy` = ?,`allowed_address_pair` = ?,`virtual_machine_interface_fat_flow_protocols` = ?,`local_preference` = ?,`traffic_direction` = ?,`encapsulation` = ?,`nh_mode` = ?,`udp_port` = ?,`juniper_header` = ?,`analyzer_ip_address` = ?,`analyzer_mac_address` = ?,`nic_assisted_mirroring_vlan` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`vni` = ?,`analyzer_name` = ?,`routing_instance` = ?,`nic_assisted_mirroring` = ?,`service_interface_type` = ?,`sub_interface_vlan_tag` = ?,`mac_address` = ?,`port_security_enabled` = ?;"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface` where uuid = ?"
const listVirtualMachineInterfaceQuery = "select `key_value_pair`,`route`,`dhcp_option`,`vlan_tag_based_bridge_domain`,`fq_name`,`global_access`,`share`,`owner`,`owner_access`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`virtual_machine_interface_bindings`,`virtual_machine_interface_device_owner`,`vrf_assign_rule`,`uuid`,`virtual_machine_interface_disable_policy`,`allowed_address_pair`,`virtual_machine_interface_fat_flow_protocols`,`local_preference`,`traffic_direction`,`encapsulation`,`nh_mode`,`udp_port`,`juniper_header`,`analyzer_ip_address`,`analyzer_mac_address`,`nic_assisted_mirroring_vlan`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_name`,`routing_instance`,`nic_assisted_mirroring`,`service_interface_type`,`sub_interface_vlan_tag`,`mac_address`,`port_security_enabled` from `virtual_machine_interface`"
const showVirtualMachineInterfaceQuery = "select `key_value_pair`,`route`,`dhcp_option`,`vlan_tag_based_bridge_domain`,`fq_name`,`global_access`,`share`,`owner`,`owner_access`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`virtual_machine_interface_bindings`,`virtual_machine_interface_device_owner`,`vrf_assign_rule`,`uuid`,`virtual_machine_interface_disable_policy`,`allowed_address_pair`,`virtual_machine_interface_fat_flow_protocols`,`local_preference`,`traffic_direction`,`encapsulation`,`nh_mode`,`udp_port`,`juniper_header`,`analyzer_ip_address`,`analyzer_mac_address`,`nic_assisted_mirroring_vlan`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_name`,`routing_instance`,`nic_assisted_mirroring`,`service_interface_type`,`sub_interface_vlan_tag`,`mac_address`,`port_security_enabled` from `virtual_machine_interface` where uuid = ?"

func CreateVirtualMachineInterface(tx *sql.Tx, model *models.VirtualMachineInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.VirtualMachineInterfaceHostRoutes.Route),
		util.MustJSON(model.VirtualMachineInterfaceDHCPOptionList.DHCPOption),
		bool(model.VlanTagBasedBridgeDomain),
		util.MustJSON(model.FQName),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		util.MustJSON(model.VirtualMachineInterfaceBindings),
		string(model.VirtualMachineInterfaceDeviceOwner),
		util.MustJSON(model.VRFAssignTable.VRFAssignRule),
		string(model.UUID),
		bool(model.VirtualMachineInterfaceDisablePolicy),
		util.MustJSON(model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair),
		util.MustJSON(model.VirtualMachineInterfaceFatFlowProtocols),
		int(model.VirtualMachineInterfaceProperties.LocalPreference),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring),
		string(model.VirtualMachineInterfaceProperties.ServiceInterfaceType),
		int(model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag),
		util.MustJSON(model.VirtualMachineInterfaceMacAddresses.MacAddress),
		bool(model.PortSecurityEnabled))
	return err
}

func scanVirtualMachineInterface(rows *sql.Rows) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	var jsonAnnotationsKeyValuePair string

	var jsonVirtualMachineInterfaceHostRoutesRoute string

	var jsonVirtualMachineInterfaceDHCPOptionListDHCPOption string

	var jsonFQName string

	var jsonPerms2Share string

	var jsonVirtualMachineInterfaceBindings string

	var jsonVRFAssignTableVRFAssignRule string

	var jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair string

	var jsonVirtualMachineInterfaceFatFlowProtocols string

	var jsonVirtualMachineInterfaceMacAddressesMacAddress string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&jsonVirtualMachineInterfaceHostRoutesRoute,
		&jsonVirtualMachineInterfaceDHCPOptionListDHCPOption,
		&m.VlanTagBasedBridgeDomain,
		&jsonFQName,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&jsonVirtualMachineInterfaceBindings,
		&m.VirtualMachineInterfaceDeviceOwner,
		&jsonVRFAssignTableVRFAssignRule,
		&m.UUID,
		&m.VirtualMachineInterfaceDisablePolicy,
		&jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair,
		&jsonVirtualMachineInterfaceFatFlowProtocols,
		&m.VirtualMachineInterfaceProperties.LocalPreference,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring,
		&m.VirtualMachineInterfaceProperties.ServiceInterfaceType,
		&m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag,
		&jsonVirtualMachineInterfaceMacAddressesMacAddress,
		&m.PortSecurityEnabled); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceHostRoutesRoute), &m.VirtualMachineInterfaceHostRoutes.Route)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceDHCPOptionListDHCPOption), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceBindings), &m.VirtualMachineInterfaceBindings)

	json.Unmarshal([]byte(jsonVRFAssignTableVRFAssignRule), &m.VRFAssignTable.VRFAssignRule)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceFatFlowProtocols), &m.VirtualMachineInterfaceFatFlowProtocols)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceMacAddressesMacAddress), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	return m, nil
}

func ListVirtualMachineInterface(tx *sql.Tx) ([]*models.VirtualMachineInterface, error) {
	result := models.MakeVirtualMachineInterfaceSlice()
	rows, err := tx.Query(listVirtualMachineInterfaceQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualMachineInterface(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualMachineInterface(tx *sql.Tx, uuid string) (*models.VirtualMachineInterface, error) {
	rows, err := tx.Query(showVirtualMachineInterfaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualMachineInterface(rows)
	}
	return nil, nil
}

func UpdateVirtualMachineInterface(tx *sql.Tx, uuid string, model *models.VirtualMachineInterface) error {
	return nil
}

func DeleteVirtualMachineInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
