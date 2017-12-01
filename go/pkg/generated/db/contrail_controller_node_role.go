package db
// contrail_controller_node_role

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertContrailControllerNodeRoleQuery = "insert into `contrail_controller_node_role` (`provisioning_start_time`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_state`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailControllerNodeRoleQuery = "update `contrail_controller_node_role` set `provisioning_start_time` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`provisioning_state` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteContrailControllerNodeRoleQuery = "delete from `contrail_controller_node_role`"
const selectContrailControllerNodeRoleQuery = "select `provisioning_start_time`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_state`,`display_name`,`key_value_pair` from `contrail_controller_node_role`"

func CreateContrailControllerNodeRole(tx *sql.Tx, model *models.ContrailControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailControllerNodeRoleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningStartTime,
    model.UUID,
    model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.ProvisioningLog,
    model.ProvisioningProgress,
    model.ProvisioningProgressStage,
    model.ProvisioningState,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListContrailControllerNodeRole(tx *sql.Tx) ([]*models.ContrailControllerNodeRole, error) {
    result := models.MakeContrailControllerNodeRoleSlice()
    rows, err := tx.Query(selectContrailControllerNodeRoleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeContrailControllerNodeRole()
            if err := rows.Scan(&m.ProvisioningStartTime,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.ProvisioningLog,
                &m.ProvisioningProgress,
                &m.ProvisioningProgressStage,
                &m.ProvisioningState,
                &m.DisplayName,
                &m.Annotations.KeyValuePair); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowContrailControllerNodeRole(db *sql.DB, id string, model *models.ContrailControllerNodeRole) error {
    return nil
}

func UpdateContrailControllerNodeRole(db *sql.DB, id string, model *models.ContrailControllerNodeRole) error {
    return nil
}

func DeleteContrailControllerNodeRole(db *sql.DB, id string) error {
    return nil
}