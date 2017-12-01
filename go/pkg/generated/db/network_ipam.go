package db
// network_ipam

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertNetworkIpamQuery = "insert into `network_ipam` (`ipam_subnets`,`ipam_subnet_method`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ip_prefix_len`,`ip_prefix`,`fq_name`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkIpamQuery = "update `network_ipam` set `ipam_subnets` = ?,`ipam_subnet_method` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`ipam_method` = ?,`ipam_dns_method` = ?,`ip_address` = ?,`virtual_dns_server_name` = ?,`dhcp_option` = ?,`route` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`fq_name` = ?,`uuid` = ?;"
const deleteNetworkIpamQuery = "delete from `network_ipam`"
const selectNetworkIpamQuery = "select `ipam_subnets`,`ipam_subnet_method`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name`,`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ip_prefix_len`,`ip_prefix`,`fq_name`,`uuid` from `network_ipam`"

func CreateNetworkIpam(tx *sql.Tx, model *models.NetworkIpam) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkIpamQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IpamSubnets,
    model.IpamSubnetMethod,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.DisplayName,
    model.NetworkIpamMGMT.IpamMethod,
    model.NetworkIpamMGMT.IpamDNSMethod,
    model.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress,
    model.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName,
    model.NetworkIpamMGMT.DHCPOptionList.DHCPOption,
    model.NetworkIpamMGMT.HostRoutes.Route,
    model.NetworkIpamMGMT.CidrBlock.IPPrefixLen,
    model.NetworkIpamMGMT.CidrBlock.IPPrefix,
    model.FQName,
    model.UUID)
    return err
}

func ListNetworkIpam(tx *sql.Tx) ([]*models.NetworkIpam, error) {
    result := models.MakeNetworkIpamSlice()
    rows, err := tx.Query(selectNetworkIpamQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeNetworkIpam()
            if err := rows.Scan(&m.IpamSubnets,
                &m.IpamSubnetMethod,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.IDPerms.Created,
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
                &m.DisplayName,
                &m.NetworkIpamMGMT.IpamMethod,
                &m.NetworkIpamMGMT.IpamDNSMethod,
                &m.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress,
                &m.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName,
                &m.NetworkIpamMGMT.DHCPOptionList.DHCPOption,
                &m.NetworkIpamMGMT.HostRoutes.Route,
                &m.NetworkIpamMGMT.CidrBlock.IPPrefixLen,
                &m.NetworkIpamMGMT.CidrBlock.IPPrefix,
                &m.FQName,
                &m.UUID); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowNetworkIpam(db *sql.DB, id string, model *models.NetworkIpam) error {
    return nil
}

func UpdateNetworkIpam(db *sql.DB, id string, model *models.NetworkIpam) error {
    return nil
}

func DeleteNetworkIpam(db *sql.DB, id string) error {
    return nil
}