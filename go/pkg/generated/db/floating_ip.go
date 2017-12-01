package db
// floating_ip

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertFloatingIPQuery = "insert into `floating_ip` (`share`,`owner`,`owner_access`,`global_access`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_address`,`floating_ip_fixed_ip_address`,`floating_ip_traffic_direction`,`display_name`,`key_value_pair`,`floating_ip_is_virtual_ip`,`floating_ip_port_mappings_enable`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPQuery = "update `floating_ip` set `share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`floating_ip_address_family` = ?,`floating_ip_port_mappings` = ?,`floating_ip_address` = ?,`floating_ip_fixed_ip_address` = ?,`floating_ip_traffic_direction` = ?,`display_name` = ?,`key_value_pair` = ?,`floating_ip_is_virtual_ip` = ?,`floating_ip_port_mappings_enable` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`uuid` = ?;"
const deleteFloatingIPQuery = "delete from `floating_ip`"
const selectFloatingIPQuery = "select `share`,`owner`,`owner_access`,`global_access`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_address`,`floating_ip_fixed_ip_address`,`floating_ip_traffic_direction`,`display_name`,`key_value_pair`,`floating_ip_is_virtual_ip`,`floating_ip_port_mappings_enable`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`uuid` from `floating_ip`"

func CreateFloatingIP(tx *sql.Tx, model *models.FloatingIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.FloatingIPAddressFamily,
    model.FloatingIPPortMappings,
    model.FloatingIPAddress,
    model.FloatingIPFixedIPAddress,
    model.FloatingIPTrafficDirection,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.FloatingIPIsVirtualIP,
    model.FloatingIPPortMappingsEnable,
    model.FQName,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.UUID)
    return err
}

func ListFloatingIP(tx *sql.Tx) ([]*models.FloatingIP, error) {
    result := models.MakeFloatingIPSlice()
    rows, err := tx.Query(selectFloatingIPQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeFloatingIP()
            if err := rows.Scan(&m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.FloatingIPAddressFamily,
                &m.FloatingIPPortMappings,
                &m.FloatingIPAddress,
                &m.FloatingIPFixedIPAddress,
                &m.FloatingIPTrafficDirection,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.FloatingIPIsVirtualIP,
                &m.FloatingIPPortMappingsEnable,
                &m.FQName,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
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

func ShowFloatingIP(db *sql.DB, id string, model *models.FloatingIP) error {
    return nil
}

func UpdateFloatingIP(db *sql.DB, id string, model *models.FloatingIP) error {
    return nil
}

func DeleteFloatingIP(db *sql.DB, id string) error {
    return nil
}