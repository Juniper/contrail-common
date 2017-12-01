package db
// global_qos_config

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertGlobalQosConfigQuery = "insert into `global_qos_config` (`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`control`,`analytics`,`dns`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalQosConfigQuery = "update `global_qos_config` set `user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`control` = ?,`analytics` = ?,`dns` = ?,`fq_name` = ?;"
const deleteGlobalQosConfigQuery = "delete from `global_qos_config`"
const selectGlobalQosConfigQuery = "select `user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`control`,`analytics`,`dns`,`fq_name` from `global_qos_config`"

func CreateGlobalQosConfig(tx *sql.Tx, model *models.GlobalQosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalQosConfigQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.ControlTrafficDSCP.Control,
    model.ControlTrafficDSCP.Analytics,
    model.ControlTrafficDSCP.DNS,
    model.FQName)
    return err
}

func ListGlobalQosConfig(tx *sql.Tx) ([]*models.GlobalQosConfig, error) {
    result := models.MakeGlobalQosConfigSlice()
    rows, err := tx.Query(selectGlobalQosConfigQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeGlobalQosConfig()
            if err := rows.Scan(&m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.ControlTrafficDSCP.Control,
                &m.ControlTrafficDSCP.Analytics,
                &m.ControlTrafficDSCP.DNS,
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

func ShowGlobalQosConfig(db *sql.DB, id string, model *models.GlobalQosConfig) error {
    return nil
}

func UpdateGlobalQosConfig(db *sql.DB, id string, model *models.GlobalQosConfig) error {
    return nil
}

func DeleteGlobalQosConfig(db *sql.DB, id string) error {
    return nil
}