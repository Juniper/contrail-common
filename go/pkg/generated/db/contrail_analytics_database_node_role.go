package db

// contrail_analytics_database_node_role

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertContrailAnalyticsDatabaseNodeRoleQuery = "insert into `contrail_analytics_database_node_role` (`display_name`,`key_value_pair`,`provisioning_log`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`uuid`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsDatabaseNodeRoleQuery = "update `contrail_analytics_database_node_role` set `display_name` = ?,`key_value_pair` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`uuid` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_progress_stage` = ?,`fq_name` = ?;"
const deleteContrailAnalyticsDatabaseNodeRoleQuery = "delete from `contrail_analytics_database_node_role` where uuid = ?"
const listContrailAnalyticsDatabaseNodeRoleQuery = "select `display_name`,`key_value_pair`,`provisioning_log`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`uuid`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`fq_name` from `contrail_analytics_database_node_role`"
const showContrailAnalyticsDatabaseNodeRoleQuery = "select `display_name`,`key_value_pair`,`provisioning_log`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`uuid`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`fq_name` from `contrail_analytics_database_node_role` where uuid = ?"

func CreateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.UUID),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.ProvisioningProgressStage),
		util.MustJSON(model.FQName))
	return err
}

func scanContrailAnalyticsDatabaseNodeRole(rows *sql.Rows) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
	m := models.MakeContrailAnalyticsDatabaseNodeRole()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningLog,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.UUID,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ProvisioningProgressStage,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListContrailAnalyticsDatabaseNodeRole(tx *sql.Tx) ([]*models.ContrailAnalyticsDatabaseNodeRole, error) {
	result := models.MakeContrailAnalyticsDatabaseNodeRoleSlice()
	rows, err := tx.Query(listContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanContrailAnalyticsDatabaseNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string) (*models.ContrailAnalyticsDatabaseNodeRole, error) {
	rows, err := tx.Query(showContrailAnalyticsDatabaseNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailAnalyticsDatabaseNodeRole(rows)
	}
	return nil, nil
}

func UpdateContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string, model *models.ContrailAnalyticsDatabaseNodeRole) error {
	return nil
}

func DeleteContrailAnalyticsDatabaseNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailAnalyticsDatabaseNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
