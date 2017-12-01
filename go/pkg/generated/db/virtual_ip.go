package db
// virtual_ip

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualIPQuery = "insert into `virtual_ip` (`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`status_description`,`status`,`persistence_type`,`address`,`protocol_port`,`admin_state`,`protocol`,`subnet_id`,`persistence_cookie_name`,`connection_limit`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualIPQuery = "update `virtual_ip` set `enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`status_description` = ?,`status` = ?,`persistence_type` = ?,`address` = ?,`protocol_port` = ?,`admin_state` = ?,`protocol` = ?,`subnet_id` = ?,`persistence_cookie_name` = ?,`connection_limit` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteVirtualIPQuery = "delete from `virtual_ip`"
const selectVirtualIPQuery = "select `enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`status_description`,`status`,`persistence_type`,`address`,`protocol_port`,`admin_state`,`protocol`,`subnet_id`,`persistence_cookie_name`,`connection_limit`,`uuid`,`fq_name` from `virtual_ip`"

func CreateVirtualIP(tx *sql.Tx, model *models.VirtualIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualIPQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.VirtualIPProperties.StatusDescription,
    model.VirtualIPProperties.Status,
    model.VirtualIPProperties.PersistenceType,
    model.VirtualIPProperties.Address,
    model.VirtualIPProperties.ProtocolPort,
    model.VirtualIPProperties.AdminState,
    model.VirtualIPProperties.Protocol,
    model.VirtualIPProperties.SubnetID,
    model.VirtualIPProperties.PersistenceCookieName,
    model.VirtualIPProperties.ConnectionLimit,
    model.UUID,
    model.FQName)
    return err
}

func ListVirtualIP(tx *sql.Tx) ([]*models.VirtualIP, error) {
    result := models.MakeVirtualIPSlice()
    rows, err := tx.Query(selectVirtualIPQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualIP()
            if err := rows.Scan(&m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.VirtualIPProperties.StatusDescription,
                &m.VirtualIPProperties.Status,
                &m.VirtualIPProperties.PersistenceType,
                &m.VirtualIPProperties.Address,
                &m.VirtualIPProperties.ProtocolPort,
                &m.VirtualIPProperties.AdminState,
                &m.VirtualIPProperties.Protocol,
                &m.VirtualIPProperties.SubnetID,
                &m.VirtualIPProperties.PersistenceCookieName,
                &m.VirtualIPProperties.ConnectionLimit,
                &m.UUID,
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

func ShowVirtualIP(db *sql.DB, id string, model *models.VirtualIP) error {
    return nil
}

func UpdateVirtualIP(db *sql.DB, id string, model *models.VirtualIP) error {
    return nil
}

func DeleteVirtualIP(db *sql.DB, id string) error {
    return nil
}