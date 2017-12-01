package db

// virtual_machine_interface

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertVirtualMachineInterfaceQuery = "insert into `virtual_machine_interface` (`key_value_pair`,`mac_address`,`dhcp_option`,`virtual_machine_interface_device_owner`,`port_security_enabled`,`uuid`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`virtual_machine_interface_disable_policy`,`share`,`owner`,`owner_access`,`global_access`,`display_name`,`route`,`virtual_machine_interface_bindings`,`allowed_address_pair`,`local_preference`,`analyzer_ip_address`,`udp_port`,`nic_assisted_mirroring_vlan`,`juniper_header`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_name`,`nh_mode`,`nic_assisted_mirroring`,`encapsulation`,`analyzer_mac_address`,`routing_instance`,`traffic_direction`,`service_interface_type`,`sub_interface_vlan_tag`,`fq_name`,`virtual_machine_interface_fat_flow_protocols`,`vlan_tag_based_bridge_domain`,`vrf_assign_rule`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineInterfaceQuery = "update `virtual_machine_interface` set `key_value_pair` = ?,`mac_address` = ?,`dhcp_option` = ?,`virtual_machine_interface_device_owner` = ?,`port_security_enabled` = ?,`uuid` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`virtual_machine_interface_disable_policy` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`display_name` = ?,`route` = ?,`virtual_machine_interface_bindings` = ?,`allowed_address_pair` = ?,`local_preference` = ?,`analyzer_ip_address` = ?,`udp_port` = ?,`nic_assisted_mirroring_vlan` = ?,`juniper_header` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`vni` = ?,`analyzer_name` = ?,`nh_mode` = ?,`nic_assisted_mirroring` = ?,`encapsulation` = ?,`analyzer_mac_address` = ?,`routing_instance` = ?,`traffic_direction` = ?,`service_interface_type` = ?,`sub_interface_vlan_tag` = ?,`fq_name` = ?,`virtual_machine_interface_fat_flow_protocols` = ?,`vlan_tag_based_bridge_domain` = ?,`vrf_assign_rule` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?;"
const deleteVirtualMachineInterfaceQuery = "delete from `virtual_machine_interface` where uuid = ?"
const listVirtualMachineInterfaceQuery = "select `key_value_pair`,`mac_address`,`dhcp_option`,`virtual_machine_interface_device_owner`,`port_security_enabled`,`uuid`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`virtual_machine_interface_disable_policy`,`share`,`owner`,`owner_access`,`global_access`,`display_name`,`route`,`virtual_machine_interface_bindings`,`allowed_address_pair`,`local_preference`,`analyzer_ip_address`,`udp_port`,`nic_assisted_mirroring_vlan`,`juniper_header`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_name`,`nh_mode`,`nic_assisted_mirroring`,`encapsulation`,`analyzer_mac_address`,`routing_instance`,`traffic_direction`,`service_interface_type`,`sub_interface_vlan_tag`,`fq_name`,`virtual_machine_interface_fat_flow_protocols`,`vlan_tag_based_bridge_domain`,`vrf_assign_rule`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created` from `virtual_machine_interface`"
const showVirtualMachineInterfaceQuery = "select `key_value_pair`,`mac_address`,`dhcp_option`,`virtual_machine_interface_device_owner`,`port_security_enabled`,`uuid`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`virtual_machine_interface_disable_policy`,`share`,`owner`,`owner_access`,`global_access`,`display_name`,`route`,`virtual_machine_interface_bindings`,`allowed_address_pair`,`local_preference`,`analyzer_ip_address`,`udp_port`,`nic_assisted_mirroring_vlan`,`juniper_header`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_name`,`nh_mode`,`nic_assisted_mirroring`,`encapsulation`,`analyzer_mac_address`,`routing_instance`,`traffic_direction`,`service_interface_type`,`sub_interface_vlan_tag`,`fq_name`,`virtual_machine_interface_fat_flow_protocols`,`vlan_tag_based_bridge_domain`,`vrf_assign_rule`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created` from `virtual_machine_interface` where uuid = ?"

func CreateVirtualMachineInterface(tx *sql.Tx, model *models.VirtualMachineInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.VirtualMachineInterfaceMacAddresses.MacAddress),
		util.MustJSON(model.VirtualMachineInterfaceDHCPOptionList.DHCPOption),
		string(model.VirtualMachineInterfaceDeviceOwner),
		bool(model.PortSecurityEnabled),
		string(model.UUID),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.VirtualMachineInterfaceDisablePolicy),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.DisplayName),
		util.MustJSON(model.VirtualMachineInterfaceHostRoutes.Route),
		util.MustJSON(model.VirtualMachineInterfaceBindings),
		util.MustJSON(model.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair),
		int(model.VirtualMachineInterfaceProperties.LocalPreference),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		int(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode),
		bool(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance),
		string(model.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection),
		string(model.VirtualMachineInterfaceProperties.ServiceInterfaceType),
		int(model.VirtualMachineInterfaceProperties.SubInterfaceVlanTag),
		util.MustJSON(model.FQName),
		util.MustJSON(model.VirtualMachineInterfaceFatFlowProtocols),
		bool(model.VlanTagBasedBridgeDomain),
		util.MustJSON(model.VRFAssignTable.VRFAssignRule),
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
		string(model.IDPerms.Created))
	return err
}

func scanVirtualMachineInterface(rows *sql.Rows) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	var jsonAnnotationsKeyValuePair string

	var jsonVirtualMachineInterfaceMacAddressesMacAddress string

	var jsonVirtualMachineInterfaceDHCPOptionListDHCPOption string

	var jsonPerms2Share string

	var jsonVirtualMachineInterfaceHostRoutesRoute string

	var jsonVirtualMachineInterfaceBindings string

	var jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair string

	var jsonFQName string

	var jsonVirtualMachineInterfaceFatFlowProtocols string

	var jsonVRFAssignTableVRFAssignRule string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&jsonVirtualMachineInterfaceMacAddressesMacAddress,
		&jsonVirtualMachineInterfaceDHCPOptionListDHCPOption,
		&m.VirtualMachineInterfaceDeviceOwner,
		&m.PortSecurityEnabled,
		&m.UUID,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.VirtualMachineInterfaceDisablePolicy,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.DisplayName,
		&jsonVirtualMachineInterfaceHostRoutesRoute,
		&jsonVirtualMachineInterfaceBindings,
		&jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair,
		&m.VirtualMachineInterfaceProperties.LocalPreference,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance,
		&m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection,
		&m.VirtualMachineInterfaceProperties.ServiceInterfaceType,
		&m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag,
		&jsonFQName,
		&jsonVirtualMachineInterfaceFatFlowProtocols,
		&m.VlanTagBasedBridgeDomain,
		&jsonVRFAssignTableVRFAssignRule,
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
		&m.IDPerms.Created); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceMacAddressesMacAddress), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceDHCPOptionListDHCPOption), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceHostRoutesRoute), &m.VirtualMachineInterfaceHostRoutes.Route)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceBindings), &m.VirtualMachineInterfaceBindings)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceAllowedAddressPairsAllowedAddressPair), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonVirtualMachineInterfaceFatFlowProtocols), &m.VirtualMachineInterfaceFatFlowProtocols)

	json.Unmarshal([]byte(jsonVRFAssignTableVRFAssignRule), &m.VRFAssignTable.VRFAssignRule)

	return m, nil
}

func createVirtualMachineInterfaceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["virtual_machine_interface_device_owner"]; ok {
		results = append(results, "virtual_machine_interface_device_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_ip_address"]; ok {
		results = append(results, "analyzer_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["vtep_dst_ip_address"]; ok {
		results = append(results, "vtep_dst_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["vtep_dst_mac_address"]; ok {
		results = append(results, "vtep_dst_mac_address = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_name"]; ok {
		results = append(results, "analyzer_name = ?")
		values = append(values, value)
	}

	if value, ok := where["nh_mode"]; ok {
		results = append(results, "nh_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["encapsulation"]; ok {
		results = append(results, "encapsulation = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_mac_address"]; ok {
		results = append(results, "analyzer_mac_address = ?")
		values = append(values, value)
	}

	if value, ok := where["routing_instance"]; ok {
		results = append(results, "routing_instance = ?")
		values = append(values, value)
	}

	if value, ok := where["traffic_direction"]; ok {
		results = append(results, "traffic_direction = ?")
		values = append(values, value)
	}

	if value, ok := where["service_interface_type"]; ok {
		results = append(results, "service_interface_type = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualMachineInterface(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualMachineInterface, error) {
	result := models.MakeVirtualMachineInterfaceSlice()
	whereQuery, values := createVirtualMachineInterfaceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualMachineInterfaceQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
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
