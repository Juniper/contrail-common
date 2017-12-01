package db
// appformix_node_role

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAppformixNodeRoleQuery = "insert into `appformix_node_role` (`provisioning_start_time`,`provisioning_state`,`key_value_pair`,`uuid`,`fq_name`,`display_name`,`provisioning_log`,`provisioning_progress_stage`,`share`,`owner`,`owner_access`,`global_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`provisioning_progress`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAppformixNodeRoleQuery = "update `appformix_node_role` set `provisioning_start_time` = ?,`provisioning_state` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`provisioning_progress` = ?;"
const deleteAppformixNodeRoleQuery = "delete from `appformix_node_role`"
const selectAppformixNodeRoleQuery = "select `provisioning_start_time`,`provisioning_state`,`key_value_pair`,`uuid`,`fq_name`,`display_name`,`provisioning_log`,`provisioning_progress_stage`,`share`,`owner`,`owner_access`,`global_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`provisioning_progress` from `appformix_node_role`"

func CreateAppformixNodeRole(tx *sql.Tx, model *models.AppformixNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAppformixNodeRoleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningStartTime,
    model.ProvisioningState,
    model.Annotations.KeyValuePair,
    model.UUID,
    model.FQName,
    model.DisplayName,
    model.ProvisioningLog,
    model.ProvisioningProgressStage,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.ProvisioningProgress)
    return err
}

func ListAppformixNodeRole(tx *sql.Tx) ([]*models.AppformixNodeRole, error) {
    result := models.MakeAppformixNodeRoleSlice()
    rows, err := tx.Query(selectAppformixNodeRoleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAppformixNodeRole()
            if err := rows.Scan(&m.ProvisioningStartTime,
                &m.ProvisioningState,
                &m.Annotations.KeyValuePair,
                &m.UUID,
                &m.FQName,
                &m.DisplayName,
                &m.ProvisioningLog,
                &m.ProvisioningProgressStage,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.ProvisioningProgress); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowAppformixNodeRole(db *sql.DB, id string, model *models.AppformixNodeRole) error {
    return nil
}

func UpdateAppformixNodeRole(db *sql.DB, id string, model *models.AppformixNodeRole) error {
    return nil
}

func DeleteAppformixNodeRole(db *sql.DB, id string) error {
    return nil
}