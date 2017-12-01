package db

// network_ipam

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertNetworkIpamQuery = "insert into `network_ipam` (`dhcp_option`,`route`,`ip_prefix_len`,`ip_prefix`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`ipam_subnet_method`,`uuid`,`display_name`,`owner`,`owner_access`,`global_access`,`share`,`ipam_subnets`,`fq_name`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkIpamQuery = "update `network_ipam` set `dhcp_option` = ?,`route` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`ipam_method` = ?,`ipam_dns_method` = ?,`ip_address` = ?,`virtual_dns_server_name` = ?,`ipam_subnet_method` = ?,`uuid` = ?,`display_name` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`ipam_subnets` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`key_value_pair` = ?;"
const deleteNetworkIpamQuery = "delete from `network_ipam` where uuid = ?"
const listNetworkIpamQuery = "select `dhcp_option`,`route`,`ip_prefix_len`,`ip_prefix`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`ipam_subnet_method`,`uuid`,`display_name`,`owner`,`owner_access`,`global_access`,`share`,`ipam_subnets`,`fq_name`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`key_value_pair` from `network_ipam`"
const showNetworkIpamQuery = "select `dhcp_option`,`route`,`ip_prefix_len`,`ip_prefix`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`ipam_subnet_method`,`uuid`,`display_name`,`owner`,`owner_access`,`global_access`,`share`,`ipam_subnets`,`fq_name`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`key_value_pair` from `network_ipam` where uuid = ?"

func CreateNetworkIpam(tx *sql.Tx, model *models.NetworkIpam) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.NetworkIpamMGMT.DHCPOptionList.DHCPOption),
		util.MustJSON(model.NetworkIpamMGMT.HostRoutes.Route),
		int(model.NetworkIpamMGMT.CidrBlock.IPPrefixLen),
		string(model.NetworkIpamMGMT.CidrBlock.IPPrefix),
		string(model.NetworkIpamMGMT.IpamMethod),
		string(model.NetworkIpamMGMT.IpamDNSMethod),
		string(model.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress),
		string(model.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName),
		string(model.IpamSubnetMethod),
		string(model.UUID),
		string(model.DisplayName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		util.MustJSON(model.IpamSubnets),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.Creator),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanNetworkIpam(rows *sql.Rows) (*models.NetworkIpam, error) {
	m := models.MakeNetworkIpam()

	var jsonNetworkIpamMGMTDHCPOptionListDHCPOption string

	var jsonNetworkIpamMGMTHostRoutesRoute string

	var jsonPerms2Share string

	var jsonIpamSubnets string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonNetworkIpamMGMTDHCPOptionListDHCPOption,
		&jsonNetworkIpamMGMTHostRoutesRoute,
		&m.NetworkIpamMGMT.CidrBlock.IPPrefixLen,
		&m.NetworkIpamMGMT.CidrBlock.IPPrefix,
		&m.NetworkIpamMGMT.IpamMethod,
		&m.NetworkIpamMGMT.IpamDNSMethod,
		&m.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress,
		&m.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName,
		&m.IpamSubnetMethod,
		&m.UUID,
		&m.DisplayName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonIpamSubnets,
		&jsonFQName,
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
		&m.IDPerms.Creator,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonNetworkIpamMGMTDHCPOptionListDHCPOption), &m.NetworkIpamMGMT.DHCPOptionList.DHCPOption)

	json.Unmarshal([]byte(jsonNetworkIpamMGMTHostRoutesRoute), &m.NetworkIpamMGMT.HostRoutes.Route)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonIpamSubnets), &m.IpamSubnets)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListNetworkIpam(tx *sql.Tx) ([]*models.NetworkIpam, error) {
	result := models.MakeNetworkIpamSlice()
	rows, err := tx.Query(listNetworkIpamQuery)
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
