package db
// firewall_rule

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertFirewallRuleQuery = "insert into `firewall_rule` (`display_name`,`ip_prefix`,`ip_prefix_len`,`tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`end_port`,`start_port`,`src_ports_end_port`,`src_ports_start_port`,`protocol_id`,`protocol`,`tag_type`,`match_tags`,`key_value_pair`,`uuid`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`endpoint_2_address_group`,`qos_action`,`assign_routing_instance`,`encapsulation`,`udp_port`,`analyzer_mac_address`,`nic_assisted_mirroring`,`nh_mode`,`nic_assisted_mirroring_vlan`,`juniper_header`,`routing_instance`,`analyzer_name`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_ip_address`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`direction`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallRuleQuery = "update `firewall_rule` set `display_name` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`tags` = ?,`tag_ids` = ?,`virtual_network` = ?,`any` = ?,`address_group` = ?,`end_port` = ?,`start_port` = ?,`src_ports_end_port` = ?,`src_ports_start_port` = ?,`protocol_id` = ?,`protocol` = ?,`tag_type` = ?,`match_tags` = ?,`key_value_pair` = ?,`uuid` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`endpoint_2_subnet_ip_prefix` = ?,`endpoint_2_subnet_ip_prefix_len` = ?,`endpoint_2_tags` = ?,`endpoint_2_tag_ids` = ?,`endpoint_2_virtual_network` = ?,`endpoint_2_any` = ?,`endpoint_2_address_group` = ?,`qos_action` = ?,`assign_routing_instance` = ?,`encapsulation` = ?,`udp_port` = ?,`analyzer_mac_address` = ?,`nic_assisted_mirroring` = ?,`nh_mode` = ?,`nic_assisted_mirroring_vlan` = ?,`juniper_header` = ?,`routing_instance` = ?,`analyzer_name` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`vni` = ?,`analyzer_ip_address` = ?,`simple_action` = ?,`apply_service` = ?,`gateway_name` = ?,`log` = ?,`alert` = ?,`direction` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?;"
const deleteFirewallRuleQuery = "delete from `firewall_rule`"
const selectFirewallRuleQuery = "select `display_name`,`ip_prefix`,`ip_prefix_len`,`tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`end_port`,`start_port`,`src_ports_end_port`,`src_ports_start_port`,`protocol_id`,`protocol`,`tag_type`,`match_tags`,`key_value_pair`,`uuid`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`endpoint_2_address_group`,`qos_action`,`assign_routing_instance`,`encapsulation`,`udp_port`,`analyzer_mac_address`,`nic_assisted_mirroring`,`nh_mode`,`nic_assisted_mirroring_vlan`,`juniper_header`,`routing_instance`,`analyzer_name`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`vni`,`analyzer_ip_address`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`direction`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name` from `firewall_rule`"

func CreateFirewallRule(tx *sql.Tx, model *models.FirewallRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallRuleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Endpoint1.Subnet.IPPrefix,
    model.Endpoint1.Subnet.IPPrefixLen,
    model.Endpoint1.Tags,
    model.Endpoint1.TagIds,
    model.Endpoint1.VirtualNetwork,
    model.Endpoint1.Any,
    model.Endpoint1.AddressGroup,
    model.Service.DSTPorts.EndPort,
    model.Service.DSTPorts.StartPort,
    model.Service.SRCPorts.EndPort,
    model.Service.SRCPorts.StartPort,
    model.Service.ProtocolID,
    model.Service.Protocol,
    model.MatchTagTypes.TagType,
    model.MatchTags,
    model.Annotations.KeyValuePair,
    model.UUID,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.Endpoint2.Subnet.IPPrefix,
    model.Endpoint2.Subnet.IPPrefixLen,
    model.Endpoint2.Tags,
    model.Endpoint2.TagIds,
    model.Endpoint2.VirtualNetwork,
    model.Endpoint2.Any,
    model.Endpoint2.AddressGroup,
    model.ActionList.QosAction,
    model.ActionList.AssignRoutingInstance,
    model.ActionList.MirrorTo.Encapsulation,
    model.ActionList.MirrorTo.UDPPort,
    model.ActionList.MirrorTo.AnalyzerMacAddress,
    model.ActionList.MirrorTo.NicAssistedMirroring,
    model.ActionList.MirrorTo.NHMode,
    model.ActionList.MirrorTo.NicAssistedMirroringVlan,
    model.ActionList.MirrorTo.JuniperHeader,
    model.ActionList.MirrorTo.RoutingInstance,
    model.ActionList.MirrorTo.AnalyzerName,
    model.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
    model.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
    model.ActionList.MirrorTo.StaticNHHeader.Vni,
    model.ActionList.MirrorTo.AnalyzerIPAddress,
    model.ActionList.SimpleAction,
    model.ActionList.ApplyService,
    model.ActionList.GatewayName,
    model.ActionList.Log,
    model.ActionList.Alert,
    model.Direction,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.FQName)
    return err
}

func ListFirewallRule(tx *sql.Tx) ([]*models.FirewallRule, error) {
    result := models.MakeFirewallRuleSlice()
    rows, err := tx.Query(selectFirewallRuleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeFirewallRule()
            if err := rows.Scan(&m.DisplayName,
                &m.Endpoint1.Subnet.IPPrefix,
                &m.Endpoint1.Subnet.IPPrefixLen,
                &m.Endpoint1.Tags,
                &m.Endpoint1.TagIds,
                &m.Endpoint1.VirtualNetwork,
                &m.Endpoint1.Any,
                &m.Endpoint1.AddressGroup,
                &m.Service.DSTPorts.EndPort,
                &m.Service.DSTPorts.StartPort,
                &m.Service.SRCPorts.EndPort,
                &m.Service.SRCPorts.StartPort,
                &m.Service.ProtocolID,
                &m.Service.Protocol,
                &m.MatchTagTypes.TagType,
                &m.MatchTags,
                &m.Annotations.KeyValuePair,
                &m.UUID,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.Endpoint2.Subnet.IPPrefix,
                &m.Endpoint2.Subnet.IPPrefixLen,
                &m.Endpoint2.Tags,
                &m.Endpoint2.TagIds,
                &m.Endpoint2.VirtualNetwork,
                &m.Endpoint2.Any,
                &m.Endpoint2.AddressGroup,
                &m.ActionList.QosAction,
                &m.ActionList.AssignRoutingInstance,
                &m.ActionList.MirrorTo.Encapsulation,
                &m.ActionList.MirrorTo.UDPPort,
                &m.ActionList.MirrorTo.AnalyzerMacAddress,
                &m.ActionList.MirrorTo.NicAssistedMirroring,
                &m.ActionList.MirrorTo.NHMode,
                &m.ActionList.MirrorTo.NicAssistedMirroringVlan,
                &m.ActionList.MirrorTo.JuniperHeader,
                &m.ActionList.MirrorTo.RoutingInstance,
                &m.ActionList.MirrorTo.AnalyzerName,
                &m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
                &m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
                &m.ActionList.MirrorTo.StaticNHHeader.Vni,
                &m.ActionList.MirrorTo.AnalyzerIPAddress,
                &m.ActionList.SimpleAction,
                &m.ActionList.ApplyService,
                &m.ActionList.GatewayName,
                &m.ActionList.Log,
                &m.ActionList.Alert,
                &m.Direction,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
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

func ShowFirewallRule(db *sql.DB, id string, model *models.FirewallRule) error {
    return nil
}

func UpdateFirewallRule(db *sql.DB, id string, model *models.FirewallRule) error {
    return nil
}

func DeleteFirewallRule(db *sql.DB, id string) error {
    return nil
}