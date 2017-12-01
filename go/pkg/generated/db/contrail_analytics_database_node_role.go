package db
// contrail_analytics_database_node_role

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertContrailAnalyticsDatabaseNodeRoleQuery = "insert into `contrail_analytics_database_node_role` (`provisioning_log`,`provisioning_progress_stage`,`provisioning_start_time`,`uuid`,`fq_name`,`provisioning_progress`,`provisioning_state`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsDatabaseNodeRoleQuery = "update `contrail_analytics_database_node_role` set `provisioning_log` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`uuid` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_state` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteContrailAnalyticsDatabaseNodeRoleQuery = "delete from `contrail_analytics_database_node_role`"
const selectContrailAnalyticsDatabaseNodeRoleQuery = "select `provisioning_log`,`provisioning_progress_stage`,`provisioning_start_time`,`uuid`,`fq_name`,`provisioning_progress`,`provisioning_state`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `contrail_analytics_database_node_role`"

func CreateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningLog,
    model.ProvisioningProgressStage,
    model.ProvisioningStartTime,
    model.UUID,
    model.FQName,
    model.ProvisioningProgress,
    model.ProvisioningState,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share)
    return err
}

func ListContrailAnalyticsDatabaseNodeRole(tx *sql.Tx) ([]*models.ContrailAnalyticsDatabaseNodeRole, error) {
    result := models.MakeContrailAnalyticsDatabaseNodeRoleSlice()
    rows, err := tx.Query(selectContrailAnalyticsDatabaseNodeRoleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeContrailAnalyticsDatabaseNodeRole()
            if err := rows.Scan(&m.ProvisioningLog,
                &m.ProvisioningProgressStage,
                &m.ProvisioningStartTime,
                &m.UUID,
                &m.FQName,
                &m.ProvisioningProgress,
                &m.ProvisioningState,
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
                &m.IDPerms.LastModified,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowContrailAnalyticsDatabaseNodeRole(db *sql.DB, id string, model *models.ContrailAnalyticsDatabaseNodeRole) error {
    return nil
}

func UpdateContrailAnalyticsDatabaseNodeRole(db *sql.DB, id string, model *models.ContrailAnalyticsDatabaseNodeRole) error {
    return nil
}

func DeleteContrailAnalyticsDatabaseNodeRole(db *sql.DB, id string) error {
    return nil
}