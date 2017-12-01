package db

// network_ipam

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertNetworkIpamQuery = "insert into `network_ipam` (`ipam_subnet_method`,`key_value_pair`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`ip_prefix`,`ip_prefix_len`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ipam_subnets`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkIpamQuery = "update `network_ipam` set `ipam_subnet_method` = ?,`key_value_pair` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`ipam_method` = ?,`ipam_dns_method` = ?,`ip_address` = ?,`virtual_dns_server_name` = ?,`dhcp_option` = ?,`route` = ?,`ipam_subnets` = ?,`fq_name` = ?;"
const deleteNetworkIpamQuery = "delete from `network_ipam` where uuid = ?"
const listNetworkIpamQuery = "select `ipam_subnet_method`,`key_value_pair`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`ip_prefix`,`ip_prefix_len`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ipam_subnets`,`fq_name` from `network_ipam`"
const showNetworkIpamQuery = "select `ipam_subnet_method`,`key_value_pair`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`ip_prefix`,`ip_prefix_len`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ipam_subnets`,`fq_name` from `network_ipam` where uuid = ?"

func CreateNetworkIpam(tx *sql.Tx, model *models.NetworkIpam) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IpamSubnetMethod),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		string(model.NetworkIpamMGMT.CidrBlock.IPPrefix),
		int(model.NetworkIpamMGMT.CidrBlock.IPPrefixLen),
		string(model.NetworkIpamMGMT.IpamMethod),
		string(model.NetworkIpamMGMT.IpamDNSMethod),
		string(model.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress),
		string(model.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName),
		util.MustJSON(model.NetworkIpamMGMT.DHCPOptionList.DHCPOption),
		util.MustJSON(model.NetworkIpamMGMT.HostRoutes.Route),
		util.MustJSON(model.IpamSubnets),
		util.MustJSON(model.FQName))
	return err
}

func scanNetworkIpam(rows *sql.Rows) (*models.NetworkIpam, error) {
	m := models.MakeNetworkIpam()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonNetworkIpamMGMTDHCPOptionListDHCPOption string

	var jsonNetworkIpamMGMTHostRoutesRoute string

	var jsonIpamSubnets string

	var jsonFQName string

	if err := rows.Scan(&m.IpamSubnetMethod,
		&jsonAnnotationsKeyValuePair,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&m.NetworkIpamMGMT.CidrBlock.IPPrefix,
		&m.NetworkIpamMGMT.CidrBlock.IPPrefixLen,
		&m.NetworkIpamMGMT.IpamMethod,
		&m.NetworkIpamMGMT.IpamDNSMethod,
		&m.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress,
		&m.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName,
		&jsonNetworkIpamMGMTDHCPOptionListDHCPOption,
		&jsonNetworkIpamMGMTHostRoutesRoute,
		&jsonIpamSubnets,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonNetworkIpamMGMTDHCPOptionListDHCPOption), &m.NetworkIpamMGMT.DHCPOptionList.DHCPOption)

	json.Unmarshal([]byte(jsonNetworkIpamMGMTHostRoutesRoute), &m.NetworkIpamMGMT.HostRoutes.Route)

	json.Unmarshal([]byte(jsonIpamSubnets), &m.IpamSubnets)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createNetworkIpamWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["ipam_subnet_method"]; ok {
		results = append(results, "ipam_subnet_method = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["ipam_method"]; ok {
		results = append(results, "ipam_method = ?")
		values = append(values, value)
	}

	if value, ok := where["ipam_dns_method"]; ok {
		results = append(results, "ipam_dns_method = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_address"]; ok {
		results = append(results, "ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_dns_server_name"]; ok {
		results = append(results, "virtual_dns_server_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListNetworkIpam(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.NetworkIpam, error) {
	result := models.MakeNetworkIpamSlice()
	whereQuery, values := createNetworkIpamWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listNetworkIpamQuery)
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
		m, _ := scanNetworkIpam(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNetworkIpam(tx *sql.Tx, uuid string) (*models.NetworkIpam, error) {
	rows, err := tx.Query(showNetworkIpamQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNetworkIpam(rows)
	}
	return nil, nil
}

func UpdateNetworkIpam(tx *sql.Tx, uuid string, model *models.NetworkIpam) error {
	return nil
}

func DeleteNetworkIpam(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
