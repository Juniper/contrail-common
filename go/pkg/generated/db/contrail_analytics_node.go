package db
// contrail_analytics_node

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertContrailAnalyticsNodeQuery = "insert into `contrail_analytics_node` (`provisioning_progress_stage`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`provisioning_log`,`fq_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsNodeQuery = "update `contrail_analytics_node` set `provisioning_progress_stage` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`provisioning_log` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?;"
const deleteContrailAnalyticsNodeQuery = "delete from `contrail_analytics_node`"
const selectContrailAnalyticsNodeQuery = "select `provisioning_progress_stage`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`provisioning_log`,`fq_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state` from `contrail_analytics_node`"

func CreateContrailAnalyticsNode(tx *sql.Tx, model *models.ContrailAnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsNodeQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningProgressStage,
    model.IDPerms.Enable,
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
    model.UUID,
    model.ProvisioningLog,
    model.FQName,
    model.ProvisioningProgress,
    model.ProvisioningStartTime,
    model.ProvisioningState)
    return err
}

func ListContrailAnalyticsNode(tx *sql.Tx) ([]*models.ContrailAnalyticsNode, error) {
    result := models.MakeContrailAnalyticsNodeSlice()
    rows, err := tx.Query(selectContrailAnalyticsNodeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeContrailAnalyticsNode()
            if err := rows.Scan(&m.ProvisioningProgressStage,
                &m.IDPerms.Enable,
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
                &m.UUID,
                &m.ProvisioningLog,
                &m.FQName,
                &m.ProvisioningProgress,
                &m.ProvisioningStartTime,
                &m.ProvisioningState); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowContrailAnalyticsNode(db *sql.DB, id string, model *models.ContrailAnalyticsNode) error {
    return nil
}

func UpdateContrailAnalyticsNode(db *sql.DB, id string, model *models.ContrailAnalyticsNode) error {
    return nil
}

func DeleteContrailAnalyticsNode(db *sql.DB, id string) error {
    return nil
}