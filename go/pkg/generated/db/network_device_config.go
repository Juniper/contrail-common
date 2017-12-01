package db
// network_device_config

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertNetworkDeviceConfigQuery = "insert into `network_device_config` (`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkDeviceConfigQuery = "update `network_device_config` set `last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteNetworkDeviceConfigQuery = "delete from `network_device_config`"
const selectNetworkDeviceConfigQuery = "select `last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name` from `network_device_config`"

func CreateNetworkDeviceConfig(tx *sql.Tx, model *models.NetworkDeviceConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkDeviceConfigQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.FQName)
    return err
}

func ListNetworkDeviceConfig(tx *sql.Tx) ([]*models.NetworkDeviceConfig, error) {
    result := models.MakeNetworkDeviceConfigSlice()
    rows, err := tx.Query(selectNetworkDeviceConfigQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeNetworkDeviceConfig()
            if err := rows.Scan(&m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
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

func ShowNetworkDeviceConfig(db *sql.DB, id string, model *models.NetworkDeviceConfig) error {
    return nil
}

func UpdateNetworkDeviceConfig(db *sql.DB, id string, model *models.NetworkDeviceConfig) error {
    return nil
}

func DeleteNetworkDeviceConfig(db *sql.DB, id string) error {
    return nil
}