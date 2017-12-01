package db

// virtual_network

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVirtualNetworkQuery = "insert into `virtual_network` (`router_external`,`mac_aging_time`,`route_target`,`is_shared`,`fq_name`,`address_allocation_mode`,`segmentation_id`,`physical_network`,`multi_policy_service_chains_enabled`,`virtual_network_network_id`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`pbb_evpn_enable`,`pbb_etree_enable`,`external_ipam`,`port_security_enabled`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`mac_limit`,`mac_limit_action`,`vxlan_network_identifier`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`import_route_target_list_route_target`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`route_target_list_route_target`,`mac_learning_enabled`,`layer2_control_word`,`flood_unknown_unicast`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualNetworkQuery = "update `virtual_network` set `router_external` = ?,`mac_aging_time` = ?,`route_target` = ?,`is_shared` = ?,`fq_name` = ?,`address_allocation_mode` = ?,`segmentation_id` = ?,`physical_network` = ?,`multi_policy_service_chains_enabled` = ?,`virtual_network_network_id` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`pbb_evpn_enable` = ?,`pbb_etree_enable` = ?,`external_ipam` = ?,`port_security_enabled` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`vxlan_network_identifier` = ?,`rpf` = ?,`forwarding_mode` = ?,`allow_transit` = ?,`network_id` = ?,`mirror_destination` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`import_route_target_list_route_target` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`route_target_list_route_target` = ?,`mac_learning_enabled` = ?,`layer2_control_word` = ?,`flood_unknown_unicast` = ?,`key_value_pair` = ?;"
const deleteVirtualNetworkQuery = "delete from `virtual_network` where uuid = ?"
const listVirtualNetworkQuery = "select `router_external`,`mac_aging_time`,`route_target`,`is_shared`,`fq_name`,`address_allocation_mode`,`segmentation_id`,`physical_network`,`multi_policy_service_chains_enabled`,`virtual_network_network_id`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`pbb_evpn_enable`,`pbb_etree_enable`,`external_ipam`,`port_security_enabled`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`mac_limit`,`mac_limit_action`,`vxlan_network_identifier`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`import_route_target_list_route_target`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`route_target_list_route_target`,`mac_learning_enabled`,`layer2_control_word`,`flood_unknown_unicast`,`key_value_pair` from `virtual_network`"
const showVirtualNetworkQuery = "select `router_external`,`mac_aging_time`,`route_target`,`is_shared`,`fq_name`,`address_allocation_mode`,`segmentation_id`,`physical_network`,`multi_policy_service_chains_enabled`,`virtual_network_network_id`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`pbb_evpn_enable`,`pbb_etree_enable`,`external_ipam`,`port_security_enabled`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`mac_limit`,`mac_limit_action`,`vxlan_network_identifier`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`import_route_target_list_route_target`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`route_target_list_route_target`,`mac_learning_enabled`,`layer2_control_word`,`flood_unknown_unicast`,`key_value_pair` from `virtual_network` where uuid = ?"

func CreateVirtualNetwork(tx *sql.Tx, model *models.VirtualNetwork) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.RouterExternal),
		int(model.MacAgingTime),
		util.MustJSON(model.ExportRouteTargetList.RouteTarget),
		bool(model.IsShared),
		util.MustJSON(model.FQName),
		string(model.AddressAllocationMode),
		int(model.ProviderProperties.SegmentationID),
		string(model.ProviderProperties.PhysicalNetwork),
		bool(model.MultiPolicyServiceChainsEnabled),
		int(model.VirtualNetworkNetworkID),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		bool(model.PBBEvpnEnable),
		bool(model.PBBEtreeEnable),
		bool(model.ExternalIpam),
		bool(model.PortSecurityEnabled),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		int(model.VirtualNetworkProperties.VxlanNetworkIdentifier),
		string(model.VirtualNetworkProperties.RPF),
		string(model.VirtualNetworkProperties.ForwardingMode),
		bool(model.VirtualNetworkProperties.AllowTransit),
		int(model.VirtualNetworkProperties.NetworkID),
		bool(model.VirtualNetworkProperties.MirrorDestination),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		util.MustJSON(model.ImportRouteTargetList.RouteTarget),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		util.MustJSON(model.RouteTargetList.RouteTarget),
		bool(model.MacLearningEnabled),
		bool(model.Layer2ControlWord),
		bool(model.FloodUnknownUnicast),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanVirtualNetwork(rows *sql.Rows) (*models.VirtualNetwork, error) {
	m := models.MakeVirtualNetwork()

	var jsonExportRouteTargetListRouteTarget string

	var jsonFQName string

	var jsonPerms2Share string

	var jsonImportRouteTargetListRouteTarget string

	var jsonRouteTargetListRouteTarget string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.RouterExternal,
		&m.MacAgingTime,
		&jsonExportRouteTargetListRouteTarget,
		&m.IsShared,
		&jsonFQName,
		&m.AddressAllocationMode,
		&m.ProviderProperties.SegmentationID,
		&m.ProviderProperties.PhysicalNetwork,
		&m.MultiPolicyServiceChainsEnabled,
		&m.VirtualNetworkNetworkID,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&m.PBBEvpnEnable,
		&m.PBBEtreeEnable,
		&m.ExternalIpam,
		&m.PortSecurityEnabled,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.VirtualNetworkProperties.VxlanNetworkIdentifier,
		&m.VirtualNetworkProperties.RPF,
		&m.VirtualNetworkProperties.ForwardingMode,
		&m.VirtualNetworkProperties.AllowTransit,
		&m.VirtualNetworkProperties.NetworkID,
		&m.VirtualNetworkProperties.MirrorDestination,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&jsonImportRouteTargetListRouteTarget,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonRouteTargetListRouteTarget,
		&m.MacLearningEnabled,
		&m.Layer2ControlWord,
		&m.FloodUnknownUnicast,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonExportRouteTargetListRouteTarget), &m.ExportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonImportRouteTargetListRouteTarget), &m.ImportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonRouteTargetListRouteTarget), &m.RouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListVirtualNetwork(tx *sql.Tx) ([]*models.VirtualNetwork, error) {
	result := models.MakeVirtualNetworkSlice()
	rows, err := tx.Query(listVirtualNetworkQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualNetwork(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualNetwork(tx *sql.Tx, uuid string) (*models.VirtualNetwork, error) {
	rows, err := tx.Query(showVirtualNetworkQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualNetwork(rows)
	}
	return nil, nil
}

func UpdateVirtualNetwork(tx *sql.Tx, uuid string, model *models.VirtualNetwork) error {
	return nil
}

func DeleteVirtualNetwork(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
