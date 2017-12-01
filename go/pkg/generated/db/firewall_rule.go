package db

// firewall_rule

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertFirewallRuleQuery = "insert into `firewall_rule` (`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`address_group`,`ip_prefix`,`ip_prefix_len`,`tags`,`tag_ids`,`virtual_network`,`any`,`nh_mode`,`analyzer_ip_address`,`encapsulation`,`analyzer_name`,`udp_port`,`nic_assisted_mirroring_vlan`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`juniper_header`,`routing_instance`,`analyzer_mac_address`,`nic_assisted_mirroring`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`qos_action`,`assign_routing_instance`,`tag_type`,`match_tags`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`endpoint_1_virtual_network`,`endpoint_1_any`,`endpoint_1_address_group`,`endpoint_1_subnet_ip_prefix`,`endpoint_1_subnet_ip_prefix_len`,`endpoint_1_tags`,`endpoint_1_tag_ids`,`end_port`,`start_port`,`src_ports_start_port`,`src_ports_end_port`,`protocol_id`,`protocol`,`direction`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallRuleQuery = "update `firewall_rule` set `last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`key_value_pair` = ?,`address_group` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`tags` = ?,`tag_ids` = ?,`virtual_network` = ?,`any` = ?,`nh_mode` = ?,`analyzer_ip_address` = ?,`encapsulation` = ?,`analyzer_name` = ?,`udp_port` = ?,`nic_assisted_mirroring_vlan` = ?,`vni` = ?,`vtep_dst_ip_address` = ?,`vtep_dst_mac_address` = ?,`juniper_header` = ?,`routing_instance` = ?,`analyzer_mac_address` = ?,`nic_assisted_mirroring` = ?,`simple_action` = ?,`apply_service` = ?,`gateway_name` = ?,`log` = ?,`alert` = ?,`qos_action` = ?,`assign_routing_instance` = ?,`tag_type` = ?,`match_tags` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`endpoint_1_virtual_network` = ?,`endpoint_1_any` = ?,`endpoint_1_address_group` = ?,`endpoint_1_subnet_ip_prefix` = ?,`endpoint_1_subnet_ip_prefix_len` = ?,`endpoint_1_tags` = ?,`endpoint_1_tag_ids` = ?,`end_port` = ?,`start_port` = ?,`src_ports_start_port` = ?,`src_ports_end_port` = ?,`protocol_id` = ?,`protocol` = ?,`direction` = ?,`display_name` = ?;"
const deleteFirewallRuleQuery = "delete from `firewall_rule` where uuid = ?"
const listFirewallRuleQuery = "select `last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`address_group`,`ip_prefix`,`ip_prefix_len`,`tags`,`tag_ids`,`virtual_network`,`any`,`nh_mode`,`analyzer_ip_address`,`encapsulation`,`analyzer_name`,`udp_port`,`nic_assisted_mirroring_vlan`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`juniper_header`,`routing_instance`,`analyzer_mac_address`,`nic_assisted_mirroring`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`qos_action`,`assign_routing_instance`,`tag_type`,`match_tags`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`endpoint_1_virtual_network`,`endpoint_1_any`,`endpoint_1_address_group`,`endpoint_1_subnet_ip_prefix`,`endpoint_1_subnet_ip_prefix_len`,`endpoint_1_tags`,`endpoint_1_tag_ids`,`end_port`,`start_port`,`src_ports_start_port`,`src_ports_end_port`,`protocol_id`,`protocol`,`direction`,`display_name` from `firewall_rule`"
const showFirewallRuleQuery = "select `last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`address_group`,`ip_prefix`,`ip_prefix_len`,`tags`,`tag_ids`,`virtual_network`,`any`,`nh_mode`,`analyzer_ip_address`,`encapsulation`,`analyzer_name`,`udp_port`,`nic_assisted_mirroring_vlan`,`vni`,`vtep_dst_ip_address`,`vtep_dst_mac_address`,`juniper_header`,`routing_instance`,`analyzer_mac_address`,`nic_assisted_mirroring`,`simple_action`,`apply_service`,`gateway_name`,`log`,`alert`,`qos_action`,`assign_routing_instance`,`tag_type`,`match_tags`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`endpoint_1_virtual_network`,`endpoint_1_any`,`endpoint_1_address_group`,`endpoint_1_subnet_ip_prefix`,`endpoint_1_subnet_ip_prefix_len`,`endpoint_1_tags`,`endpoint_1_tag_ids`,`end_port`,`start_port`,`src_ports_start_port`,`src_ports_end_port`,`protocol_id`,`protocol`,`direction`,`display_name` from `firewall_rule` where uuid = ?"

func CreateFirewallRule(tx *sql.Tx, model *models.FirewallRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Endpoint2.AddressGroup),
		string(model.Endpoint2.Subnet.IPPrefix),
		int(model.Endpoint2.Subnet.IPPrefixLen),
		util.MustJSON(model.Endpoint2.Tags),
		util.MustJSON(model.Endpoint2.TagIds),
		string(model.Endpoint2.VirtualNetwork),
		bool(model.Endpoint2.Any),
		string(model.ActionList.MirrorTo.NHMode),
		string(model.ActionList.MirrorTo.AnalyzerIPAddress),
		string(model.ActionList.MirrorTo.Encapsulation),
		string(model.ActionList.MirrorTo.AnalyzerName),
		int(model.ActionList.MirrorTo.UDPPort),
		int(model.ActionList.MirrorTo.NicAssistedMirroringVlan),
		int(model.ActionList.MirrorTo.StaticNHHeader.Vni),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress),
		string(model.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress),
		bool(model.ActionList.MirrorTo.JuniperHeader),
		string(model.ActionList.MirrorTo.RoutingInstance),
		string(model.ActionList.MirrorTo.AnalyzerMacAddress),
		bool(model.ActionList.MirrorTo.NicAssistedMirroring),
		string(model.ActionList.SimpleAction),
		util.MustJSON(model.ActionList.ApplyService),
		string(model.ActionList.GatewayName),
		bool(model.ActionList.Log),
		bool(model.ActionList.Alert),
		string(model.ActionList.QosAction),
		string(model.ActionList.AssignRoutingInstance),
		util.MustJSON(model.MatchTagTypes.TagType),
		util.MustJSON(model.MatchTags),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.Endpoint1.VirtualNetwork),
		bool(model.Endpoint1.Any),
		string(model.Endpoint1.AddressGroup),
		string(model.Endpoint1.Subnet.IPPrefix),
		int(model.Endpoint1.Subnet.IPPrefixLen),
		util.MustJSON(model.Endpoint1.Tags),
		util.MustJSON(model.Endpoint1.TagIds),
		int(model.Service.DSTPorts.EndPort),
		int(model.Service.DSTPorts.StartPort),
		int(model.Service.SRCPorts.StartPort),
		int(model.Service.SRCPorts.EndPort),
		int(model.Service.ProtocolID),
		string(model.Service.Protocol),
		string(model.Direction),
		string(model.DisplayName))
	return err
}

func scanFirewallRule(rows *sql.Rows) (*models.FirewallRule, error) {
	m := models.MakeFirewallRule()

	var jsonAnnotationsKeyValuePair string

	var jsonEndpoint2Tags string

	var jsonEndpoint2TagIds string

	var jsonActionListApplyService string

	var jsonMatchTagTypesTagType string

	var jsonMatchTags string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonEndpoint1Tags string

	var jsonEndpoint1TagIds string

	if err := rows.Scan(&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&jsonAnnotationsKeyValuePair,
		&m.Endpoint2.AddressGroup,
		&m.Endpoint2.Subnet.IPPrefix,
		&m.Endpoint2.Subnet.IPPrefixLen,
		&jsonEndpoint2Tags,
		&jsonEndpoint2TagIds,
		&m.Endpoint2.VirtualNetwork,
		&m.Endpoint2.Any,
		&m.ActionList.MirrorTo.NHMode,
		&m.ActionList.MirrorTo.AnalyzerIPAddress,
		&m.ActionList.MirrorTo.Encapsulation,
		&m.ActionList.MirrorTo.AnalyzerName,
		&m.ActionList.MirrorTo.UDPPort,
		&m.ActionList.MirrorTo.NicAssistedMirroringVlan,
		&m.ActionList.MirrorTo.StaticNHHeader.Vni,
		&m.ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress,
		&m.ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress,
		&m.ActionList.MirrorTo.JuniperHeader,
		&m.ActionList.MirrorTo.RoutingInstance,
		&m.ActionList.MirrorTo.AnalyzerMacAddress,
		&m.ActionList.MirrorTo.NicAssistedMirroring,
		&m.ActionList.SimpleAction,
		&jsonActionListApplyService,
		&m.ActionList.GatewayName,
		&m.ActionList.Log,
		&m.ActionList.Alert,
		&m.ActionList.QosAction,
		&m.ActionList.AssignRoutingInstance,
		&jsonMatchTagTypesTagType,
		&jsonMatchTags,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
		&m.Endpoint1.VirtualNetwork,
		&m.Endpoint1.Any,
		&m.Endpoint1.AddressGroup,
		&m.Endpoint1.Subnet.IPPrefix,
		&m.Endpoint1.Subnet.IPPrefixLen,
		&jsonEndpoint1Tags,
		&jsonEndpoint1TagIds,
		&m.Service.DSTPorts.EndPort,
		&m.Service.DSTPorts.StartPort,
		&m.Service.SRCPorts.StartPort,
		&m.Service.SRCPorts.EndPort,
		&m.Service.ProtocolID,
		&m.Service.Protocol,
		&m.Direction,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonEndpoint2Tags), &m.Endpoint2.Tags)

	json.Unmarshal([]byte(jsonEndpoint2TagIds), &m.Endpoint2.TagIds)

	json.Unmarshal([]byte(jsonActionListApplyService), &m.ActionList.ApplyService)

	json.Unmarshal([]byte(jsonMatchTagTypesTagType), &m.MatchTagTypes.TagType)

	json.Unmarshal([]byte(jsonMatchTags), &m.MatchTags)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonEndpoint1Tags), &m.Endpoint1.Tags)

	json.Unmarshal([]byte(jsonEndpoint1TagIds), &m.Endpoint1.TagIds)

	return m, nil
}

func createFirewallRuleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["address_group"]; ok {
		results = append(results, "address_group = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_network"]; ok {
		results = append(results, "virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["nh_mode"]; ok {
		results = append(results, "nh_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_ip_address"]; ok {
		results = append(results, "analyzer_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["encapsulation"]; ok {
		results = append(results, "encapsulation = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_name"]; ok {
		results = append(results, "analyzer_name = ?")
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

	if value, ok := where["routing_instance"]; ok {
		results = append(results, "routing_instance = ?")
		values = append(values, value)
	}

	if value, ok := where["analyzer_mac_address"]; ok {
		results = append(results, "analyzer_mac_address = ?")
		values = append(values, value)
	}

	if value, ok := where["simple_action"]; ok {
		results = append(results, "simple_action = ?")
		values = append(values, value)
	}

	if value, ok := where["gateway_name"]; ok {
		results = append(results, "gateway_name = ?")
		values = append(values, value)
	}

	if value, ok := where["qos_action"]; ok {
		results = append(results, "qos_action = ?")
		values = append(values, value)
	}

	if value, ok := where["assign_routing_instance"]; ok {
		results = append(results, "assign_routing_instance = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["endpoint_1_virtual_network"]; ok {
		results = append(results, "endpoint_1_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["endpoint_1_address_group"]; ok {
		results = append(results, "endpoint_1_address_group = ?")
		values = append(values, value)
	}

	if value, ok := where["endpoint_1_subnet_ip_prefix"]; ok {
		results = append(results, "endpoint_1_subnet_ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["protocol"]; ok {
		results = append(results, "protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["direction"]; ok {
		results = append(results, "direction = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListFirewallRule(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.FirewallRule, error) {
	result := models.MakeFirewallRuleSlice()
	whereQuery, values := createFirewallRuleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listFirewallRuleQuery)
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
