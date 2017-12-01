package db

// firewall_rule

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertFirewallRuleQuery = "insert into `firewall_rule` (`tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`ip_prefix_len`,`ip_prefix`,`qos_action`,`assign_routing_instance`,`nic_assisted_mirroring`,`analyzer_name`,`juniper_header`,`routing_instance`,`vtep_dst_mac_address`,`vni`,`vtep_dst_ip_address`,`analyzer_mac_address`,`udp_port`,`encapsulation`,`nic_assisted_mirroring_vlan`,`nh_mode`,`analyzer_ip_address`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`direction`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`endpoint_2_address_group`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`end_port`,`start_port`,`protocol_id`,`protocol`,`dst_ports_end_port`,`dst_ports_start_port`,`tag_type`,`match_tags`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallRuleQuery = "update `firewall_rule` set `tags` = ?,`tag_ids` = ?,`virtual_network` = ?,`any` = ?,`address_group` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`qos_action` = ?,`assign_routing_instance` = ?,`nic_assisted_mirroring` = ?,`analyzer_name` = ?,`juniper_header` = ?,`routing_instance` = ?,`vtep_dst_mac_address` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`analyzer_mac_address` = ?,`udp_port` = ?,`encapsulation` = ?,`nic_assisted_mirroring_vlan` = ?,`nh_mode` = ?,`analyzer_ip_address` = ?,`simple_action` = ?,`apply_service` = ?,`gateway_name` = ?,`log` = ?,`alert` = ?,`direction` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`endpoint_2_address_group` = ?,`endpoint_2_subnet_ip_prefix` = ?,`endpoint_2_subnet_ip_prefix_len` = ?,`endpoint_2_tags` = ?,`endpoint_2_tag_ids` = ?,`endpoint_2_virtual_network` = ?,`endpoint_2_any` = ?,`end_port` = ?,`start_port` = ?,`protocol_id` = ?,`protocol` = ?,`dst_ports_end_port` = ?,`dst_ports_start_port` = ?,`tag_type` = ?,`match_tags` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteFirewallRuleQuery = "delete from `firewall_rule` where uuid = ?"
const listFirewallRuleQuery = "select `tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`ip_prefix_len`,`ip_prefix`,`qos_action`,`assign_routing_instance`,`nic_assisted_mirroring`,`analyzer_name`,`juniper_header`,`routing_instance`,`vtep_dst_mac_address`,`vni`,`vtep_dst_ip_address`,`analyzer_mac_address`,`udp_port`,`encapsulation`,`nic_assisted_mirroring_vlan`,`nh_mode`,`analyzer_ip_address`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`direction`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`endpoint_2_address_group`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`end_port`,`start_port`,`protocol_id`,`protocol`,`dst_ports_end_port`,`dst_ports_start_port`,`tag_type`,`match_tags`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `firewall_rule`"
const showFirewallRuleQuery = "select `tags`,`tag_ids`,`virtual_network`,`any`,`address_group`,`ip_prefix_len`,`ip_prefix`,`qos_action`,`assign_routing_instance`,`nic_assisted_mirroring`,`analyzer_name`,`juniper_header`,`routing_instance`,`vtep_dst_mac_address`,`vni`,`vtep_dst_ip_address`,`analyzer_mac_address`,`udp_port`,`encapsulation`,`nic_assisted_mirroring_vlan`,`nh_mode`,`analyzer_ip_address`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`direction`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`endpoint_2_address_group`,`endpoint_2_subnet_ip_prefix`,`endpoint_2_subnet_ip_prefix_len`,`endpoint_2_tags`,`endpoint_2_tag_ids`,`endpoint_2_virtual_network`,`endpoint_2_any`,`end_port`,`start_port`,`protocol_id`,`protocol`,`dst_ports_end_port`,`dst_ports_start_port`,`tag_type`,`match_tags`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `firewall_rule` where uuid = ?"

func CreateFirewallRule(tx *sql.Tx, model *models.FirewallRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Endpoint1.Tags),
		util.MustJSON(model.Endpoint1.TagIds),
		string(model.Endpoint1.VirtualNetwork),
		bool(model.Endpoint1.Any),
		string(model.Endpoint1.AddressGroup),
		int(model.Endpoint1.Subnet.IPPrefixLen),
		string(model.Endpoint1.Subnet.IPPrefix),
		string(model.ActionList.QosAction),
		string(model.ActionList.AssignRoutingInstance),
		bool(model.ActionList.MirrorTo.NicAssistedMirroring),
		string(model.ActionList.MirrorTo.AnalyzerName),
		bool(model.ActionList.MirrorTo.JuniperHeader),
		string(model.ActionList.MirrorTo.RoutingInstance),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		int(model.ActionList.MirrorTo.StaticNHHeader.Vni),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.ActionList.MirrorTo.AnalyzerMacAddress),
		int(model.ActionList.MirrorTo.UDPPort),
		string(model.ActionList.MirrorTo.Encapsulation),
		int(model.ActionList.MirrorTo.NicAssistedMirroringVlan),
		string(model.ActionList.MirrorTo.NHMode),
		string(model.ActionList.MirrorTo.AnalyzerIPAddress),
		string(model.ActionList.SimpleAction),
		util.MustJSON(model.ActionList.ApplyService),
		string(model.ActionList.GatewayName),
		bool(model.ActionList.Log),
		bool(model.ActionList.Alert),
		string(model.Direction),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Endpoint2.AddressGroup),
		string(model.Endpoint2.Subnet.IPPrefix),
		int(model.Endpoint2.Subnet.IPPrefixLen),
		util.MustJSON(model.Endpoint2.Tags),
		util.MustJSON(model.Endpoint2.TagIds),
		string(model.Endpoint2.VirtualNetwork),
		bool(model.Endpoint2.Any),
		int(model.Service.SRCPorts.EndPort),
		int(model.Service.SRCPorts.StartPort),
		int(model.Service.ProtocolID),
		string(model.Service.Protocol),
		int(model.Service.DSTPorts.EndPort),
		int(model.Service.DSTPorts.StartPort),
		util.MustJSON(model.MatchTagTypes.TagType),
		util.MustJSON(model.MatchTags),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	return err
}

func scanFirewallRule(rows *sql.Rows) (*models.FirewallRule, error) {
	m := models.MakeFirewallRule()

	var jsonEndpoint1Tags string

	var jsonEndpoint1TagIds string

	var jsonActionListApplyService string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonEndpoint2Tags string

	var jsonEndpoint2TagIds string

	var jsonMatchTagTypesTagType string

	var jsonMatchTags string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonEndpoint1Tags,
		&jsonEndpoint1TagIds,
		&m.Endpoint1.VirtualNetwork,
		&m.Endpoint1.Any,
		&m.Endpoint1.AddressGroup,
		&m.Endpoint1.Subnet.IPPrefixLen,
		&m.Endpoint1.Subnet.IPPrefix,
		&m.ActionList.QosAction,
		&m.ActionList.AssignRoutingInstance,
		&m.ActionList.MirrorTo.NicAssistedMirroring,
		&m.ActionList.MirrorTo.AnalyzerName,
		&m.ActionList.MirrorTo.JuniperHeader,
		&m.ActionList.MirrorTo.RoutingInstance,
		&m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
		&m.ActionList.MirrorTo.StaticNHHeader.Vni,
		&m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
		&m.ActionList.MirrorTo.AnalyzerMacAddress,
		&m.ActionList.MirrorTo.UDPPort,
		&m.ActionList.MirrorTo.Encapsulation,
		&m.ActionList.MirrorTo.NicAssistedMirroringVlan,
		&m.ActionList.MirrorTo.NHMode,
		&m.ActionList.MirrorTo.AnalyzerIPAddress,
		&m.ActionList.SimpleAction,
		&jsonActionListApplyService,
		&m.ActionList.GatewayName,
		&m.ActionList.Log,
		&m.ActionList.Alert,
		&m.Direction,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Endpoint2.AddressGroup,
		&m.Endpoint2.Subnet.IPPrefix,
		&m.Endpoint2.Subnet.IPPrefixLen,
		&jsonEndpoint2Tags,
		&jsonEndpoint2TagIds,
		&m.Endpoint2.VirtualNetwork,
		&m.Endpoint2.Any,
		&m.Service.SRCPorts.EndPort,
		&m.Service.SRCPorts.StartPort,
		&m.Service.ProtocolID,
		&m.Service.Protocol,
		&m.Service.DSTPorts.EndPort,
		&m.Service.DSTPorts.StartPort,
		&jsonMatchTagTypesTagType,
		&jsonMatchTags,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonEndpoint1Tags), &m.Endpoint1.Tags)

	json.Unmarshal([]byte(jsonEndpoint1TagIds), &m.Endpoint1.TagIds)

	json.Unmarshal([]byte(jsonActionListApplyService), &m.ActionList.ApplyService)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonEndpoint2Tags), &m.Endpoint2.Tags)

	json.Unmarshal([]byte(jsonEndpoint2TagIds), &m.Endpoint2.TagIds)

	json.Unmarshal([]byte(jsonMatchTagTypesTagType), &m.MatchTagTypes.TagType)

	json.Unmarshal([]byte(jsonMatchTags), &m.MatchTags)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListFirewallRule(tx *sql.Tx) ([]*models.FirewallRule, error) {
	result := models.MakeFirewallRuleSlice()
	rows, err := tx.Query(listFirewallRuleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanFirewallRule(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFirewallRule(tx *sql.Tx, uuid string) (*models.FirewallRule, error) {
	rows, err := tx.Query(showFirewallRuleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFirewallRule(rows)
	}
	return nil, nil
}

func UpdateFirewallRule(tx *sql.Tx, uuid string, model *models.FirewallRule) error {
	return nil
}

func DeleteFirewallRule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFirewallRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
