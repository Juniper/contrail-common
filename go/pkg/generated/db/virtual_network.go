package db
// virtual_network

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualNetworkQuery = "insert into `virtual_network` (`port_security_enabled`,`is_shared`,`address_allocation_mode`,`mac_limit`,`mac_limit_action`,`display_name`,`virtual_network_network_id`,`segmentation_id`,`physical_network`,`mac_learning_enabled`,`external_ipam`,`router_external`,`route_target`,`pbb_etree_enable`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`multi_policy_service_chains_enabled`,`share`,`owner`,`owner_access`,`global_access`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`vxlan_network_identifier`,`rpf`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`mac_aging_time`,`flood_unknown_unicast`,`uuid`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`import_route_target_list_route_target`,`key_value_pair`,`pbb_evpn_enable`,`export_route_target_list_route_target`,`layer2_control_word`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualNetworkQuery = "update `virtual_network` set `port_security_enabled` = ?,`is_shared` = ?,`address_allocation_mode` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`display_name` = ?,`virtual_network_network_id` = ?,`segmentation_id` = ?,`physical_network` = ?,`mac_learning_enabled` = ?,`external_ipam` = ?,`router_external` = ?,`route_target` = ?,`pbb_etree_enable` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`multi_policy_service_chains_enabled` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`forwarding_mode` = ?,`allow_transit` = ?,`network_id` = ?,`mirror_destination` = ?,`vxlan_network_identifier` = ?,`rpf` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`ip_protocol` = ?,`mac_aging_time` = ?,`flood_unknown_unicast` = ?,`uuid` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`import_route_target_list_route_target` = ?,`key_value_pair` = ?,`pbb_evpn_enable` = ?,`export_route_target_list_route_target` = ?,`layer2_control_word` = ?,`fq_name` = ?;"
const deleteVirtualNetworkQuery = "delete from `virtual_network`"
const selectVirtualNetworkQuery = "select `port_security_enabled`,`is_shared`,`address_allocation_mode`,`mac_limit`,`mac_limit_action`,`display_name`,`virtual_network_network_id`,`segmentation_id`,`physical_network`,`mac_learning_enabled`,`external_ipam`,`router_external`,`route_target`,`pbb_etree_enable`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`multi_policy_service_chains_enabled`,`share`,`owner`,`owner_access`,`global_access`,`forwarding_mode`,`allow_transit`,`network_id`,`mirror_destination`,`vxlan_network_identifier`,`rpf`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`mac_aging_time`,`flood_unknown_unicast`,`uuid`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`import_route_target_list_route_target`,`key_value_pair`,`pbb_evpn_enable`,`export_route_target_list_route_target`,`layer2_control_word`,`fq_name` from `virtual_network`"

func CreateVirtualNetwork(tx *sql.Tx, model *models.VirtualNetwork) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualNetworkQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.PortSecurityEnabled,
    model.IsShared,
    model.AddressAllocationMode,
    model.MacLimitControl.MacLimit,
    model.MacLimitControl.MacLimitAction,
    model.DisplayName,
    model.VirtualNetworkNetworkID,
    model.ProviderProperties.SegmentationID,
    model.ProviderProperties.PhysicalNetwork,
    model.MacLearningEnabled,
    model.ExternalIpam,
    model.RouterExternal,
    model.RouteTargetList.RouteTarget,
    model.PBBEtreeEnable,
    model.MacMoveControl.MacMoveTimeWindow,
    model.MacMoveControl.MacMoveLimit,
    model.MacMoveControl.MacMoveLimitAction,
    model.MultiPolicyServiceChainsEnabled,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.VirtualNetworkProperties.ForwardingMode,
    model.VirtualNetworkProperties.AllowTransit,
    model.VirtualNetworkProperties.NetworkID,
    model.VirtualNetworkProperties.MirrorDestination,
    model.VirtualNetworkProperties.VxlanNetworkIdentifier,
    model.VirtualNetworkProperties.RPF,
    model.EcmpHashingIncludeFields.SourceIP,
    model.EcmpHashingIncludeFields.HashingConfigured,
    model.EcmpHashingIncludeFields.SourcePort,
    model.EcmpHashingIncludeFields.DestinationPort,
    model.EcmpHashingIncludeFields.DestinationIP,
    model.EcmpHashingIncludeFields.IPProtocol,
    model.MacAgingTime,
    model.FloodUnknownUnicast,
    model.UUID,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.ImportRouteTargetList.RouteTarget,
    model.Annotations.KeyValuePair,
    model.PBBEvpnEnable,
    model.ExportRouteTargetList.RouteTarget,
    model.Layer2ControlWord,
    model.FQName)
    return err
}

func ListVirtualNetwork(tx *sql.Tx) ([]*models.VirtualNetwork, error) {
    result := models.MakeVirtualNetworkSlice()
    rows, err := tx.Query(selectVirtualNetworkQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualNetwork()
            if err := rows.Scan(&m.PortSecurityEnabled,
                &m.IsShared,
                &m.AddressAllocationMode,
                &m.MacLimitControl.MacLimit,
                &m.MacLimitControl.MacLimitAction,
                &m.DisplayName,
                &m.VirtualNetworkNetworkID,
                &m.ProviderProperties.SegmentationID,
                &m.ProviderProperties.PhysicalNetwork,
                &m.MacLearningEnabled,
                &m.ExternalIpam,
                &m.RouterExternal,
                &m.RouteTargetList.RouteTarget,
                &m.PBBEtreeEnable,
                &m.MacMoveControl.MacMoveTimeWindow,
                &m.MacMoveControl.MacMoveLimit,
                &m.MacMoveControl.MacMoveLimitAction,
                &m.MultiPolicyServiceChainsEnabled,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.VirtualNetworkProperties.ForwardingMode,
                &m.VirtualNetworkProperties.AllowTransit,
                &m.VirtualNetworkProperties.NetworkID,
                &m.VirtualNetworkProperties.MirrorDestination,
                &m.VirtualNetworkProperties.VxlanNetworkIdentifier,
                &m.VirtualNetworkProperties.RPF,
                &m.EcmpHashingIncludeFields.SourceIP,
                &m.EcmpHashingIncludeFields.HashingConfigured,
                &m.EcmpHashingIncludeFields.SourcePort,
                &m.EcmpHashingIncludeFields.DestinationPort,
                &m.EcmpHashingIncludeFields.DestinationIP,
                &m.EcmpHashingIncludeFields.IPProtocol,
                &m.MacAgingTime,
                &m.FloodUnknownUnicast,
                &m.UUID,
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
                &m.ImportRouteTargetList.RouteTarget,
                &m.Annotations.KeyValuePair,
                &m.PBBEvpnEnable,
                &m.ExportRouteTargetList.RouteTarget,
                &m.Layer2ControlWord,
                &m.FQName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowVirtualNetwork(db *sql.DB, id string, model *models.VirtualNetwork) error {
    return nil
}

func UpdateVirtualNetwork(db *sql.DB, id string, model *models.VirtualNetwork) error {
    return nil
}

func DeleteVirtualNetwork(db *sql.DB, id string) error {
    return nil
}