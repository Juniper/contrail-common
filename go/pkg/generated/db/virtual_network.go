package db

// virtual_network

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertVirtualNetworkQuery = "insert into `virtual_network` (`pbb_etree_enable`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`virtual_network_network_id`,`pbb_evpn_enable`,`segmentation_id`,`physical_network`,`fq_name`,`share`,`owner`,`owner_access`,`global_access`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`vxlan_network_identifier`,`mac_limit_action`,`mac_limit`,`multi_policy_service_chains_enabled`,`address_allocation_mode`,`external_ipam`,`is_shared`,`mac_learning_enabled`,`port_security_enabled`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`layer2_control_word`,`uuid`,`route_target`,`mac_aging_time`,`export_route_target_list_route_target`,`flood_unknown_unicast`,`key_value_pair`,`router_external`,`route_target_list_route_target`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualNetworkQuery = "update `virtual_network` set `pbb_etree_enable` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`ip_protocol` = ?,`virtual_network_network_id` = ?,`pbb_evpn_enable` = ?,`segmentation_id` = ?,`physical_network` = ?,`fq_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`rpf` = ?,`forwarding_mode` = ?,`allow_transit` = ?,`network_id` = ?,`mirror_destination` = ?,`vxlan_network_identifier` = ?,`mac_limit_action` = ?,`mac_limit` = ?,`multi_policy_service_chains_enabled` = ?,`address_allocation_mode` = ?,`external_ipam` = ?,`is_shared` = ?,`mac_learning_enabled` = ?,`port_security_enabled` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`layer2_control_word` = ?,`uuid` = ?,`route_target` = ?,`mac_aging_time` = ?,`export_route_target_list_route_target` = ?,`flood_unknown_unicast` = ?,`key_value_pair` = ?,`router_external` = ?,`route_target_list_route_target` = ?;"
const deleteVirtualNetworkQuery = "delete from `virtual_network` where uuid = ?"
const listVirtualNetworkQuery = "select `pbb_etree_enable`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`virtual_network_network_id`,`pbb_evpn_enable`,`segmentation_id`,`physical_network`,`fq_name`,`share`,`owner`,`owner_access`,`global_access`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`vxlan_network_identifier`,`mac_limit_action`,`mac_limit`,`multi_policy_service_chains_enabled`,`address_allocation_mode`,`external_ipam`,`is_shared`,`mac_learning_enabled`,`port_security_enabled`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`layer2_control_word`,`uuid`,`route_target`,`mac_aging_time`,`export_route_target_list_route_target`,`flood_unknown_unicast`,`key_value_pair`,`router_external`,`route_target_list_route_target` from `virtual_network`"
const showVirtualNetworkQuery = "select `pbb_etree_enable`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`virtual_network_network_id`,`pbb_evpn_enable`,`segmentation_id`,`physical_network`,`fq_name`,`share`,`owner`,`owner_access`,`global_access`,`rpf`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`vxlan_network_identifier`,`mac_limit_action`,`mac_limit`,`multi_policy_service_chains_enabled`,`address_allocation_mode`,`external_ipam`,`is_shared`,`mac_learning_enabled`,`port_security_enabled`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`layer2_control_word`,`uuid`,`route_target`,`mac_aging_time`,`export_route_target_list_route_target`,`flood_unknown_unicast`,`key_value_pair`,`router_external`,`route_target_list_route_target` from `virtual_network` where uuid = ?"

func CreateVirtualNetwork(tx *sql.Tx, model *models.VirtualNetwork) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.PBBEtreeEnable),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		int(model.VirtualNetworkNetworkID),
		bool(model.PBBEvpnEnable),
		int(model.ProviderProperties.SegmentationID),
		string(model.ProviderProperties.PhysicalNetwork),
		util.MustJSON(model.FQName),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.VirtualNetworkProperties.RPF),
		string(model.VirtualNetworkProperties.ForwardingMode),
		bool(model.VirtualNetworkProperties.AllowTransit),
		int(model.VirtualNetworkProperties.NetworkID),
		bool(model.VirtualNetworkProperties.MirrorDestination),
		int(model.VirtualNetworkProperties.VxlanNetworkIdentifier),
		string(model.MacLimitControl.MacLimitAction),
		int(model.MacLimitControl.MacLimit),
		bool(model.MultiPolicyServiceChainsEnabled),
		string(model.AddressAllocationMode),
		bool(model.ExternalIpam),
		bool(model.IsShared),
		bool(model.MacLearningEnabled),
		bool(model.PortSecurityEnabled),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		bool(model.Layer2ControlWord),
		string(model.UUID),
		util.MustJSON(model.ImportRouteTargetList.RouteTarget),
		int(model.MacAgingTime),
		util.MustJSON(model.ExportRouteTargetList.RouteTarget),
		bool(model.FloodUnknownUnicast),
		util.MustJSON(model.Annotations.KeyValuePair),
		bool(model.RouterExternal),
		util.MustJSON(model.RouteTargetList.RouteTarget))
	return err
}

func scanVirtualNetwork(rows *sql.Rows) (*models.VirtualNetwork, error) {
	m := models.MakeVirtualNetwork()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonImportRouteTargetListRouteTarget string

	var jsonExportRouteTargetListRouteTarget string

	var jsonAnnotationsKeyValuePair string

	var jsonRouteTargetListRouteTarget string

	if err := rows.Scan(&m.PBBEtreeEnable,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.VirtualNetworkNetworkID,
		&m.PBBEvpnEnable,
		&m.ProviderProperties.SegmentationID,
		&m.ProviderProperties.PhysicalNetwork,
		&jsonFQName,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.VirtualNetworkProperties.RPF,
		&m.VirtualNetworkProperties.ForwardingMode,
		&m.VirtualNetworkProperties.AllowTransit,
		&m.VirtualNetworkProperties.NetworkID,
		&m.VirtualNetworkProperties.MirrorDestination,
		&m.VirtualNetworkProperties.VxlanNetworkIdentifier,
		&m.MacLimitControl.MacLimitAction,
		&m.MacLimitControl.MacLimit,
		&m.MultiPolicyServiceChainsEnabled,
		&m.AddressAllocationMode,
		&m.ExternalIpam,
		&m.IsShared,
		&m.MacLearningEnabled,
		&m.PortSecurityEnabled,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&m.Layer2ControlWord,
		&m.UUID,
		&jsonImportRouteTargetListRouteTarget,
		&m.MacAgingTime,
		&jsonExportRouteTargetListRouteTarget,
		&m.FloodUnknownUnicast,
		&jsonAnnotationsKeyValuePair,
		&m.RouterExternal,
		&jsonRouteTargetListRouteTarget); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonImportRouteTargetListRouteTarget), &m.ImportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonExportRouteTargetListRouteTarget), &m.ExportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonRouteTargetListRouteTarget), &m.RouteTargetList.RouteTarget)

	return m, nil
}

func createVirtualNetworkWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["physical_network"]; ok {
		results = append(results, "physical_network = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["rpf"]; ok {
		results = append(results, "rpf = ?")
		values = append(values, value)
	}

	if value, ok := where["forwarding_mode"]; ok {
		results = append(results, "forwarding_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_limit_action"]; ok {
		results = append(results, "mac_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["address_allocation_mode"]; ok {
		results = append(results, "address_allocation_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_move_limit_action"]; ok {
		results = append(results, "mac_move_limit_action = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualNetwork(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualNetwork, error) {
	result := models.MakeVirtualNetworkSlice()
	whereQuery, values := createVirtualNetworkWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualNetworkQuery)
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
