package db

// appformix_node_role

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAppformixNodeRoleQuery = "insert into `appformix_node_role` (`provisioning_state`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`key_value_pair`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAppformixNodeRoleQuery = "update `appformix_node_role` set `provisioning_state` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`key_value_pair` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`display_name` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`provisioning_log` = ?;"
const deleteAppformixNodeRoleQuery = "delete from `appformix_node_role` where uuid = ?"
const listAppformixNodeRoleQuery = "select `provisioning_state`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`key_value_pair`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`provisioning_log` from `appformix_node_role`"
const showAppformixNodeRoleQuery = "select `provisioning_state`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`key_value_pair`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`provisioning_log` from `appformix_node_role` where uuid = ?"

func CreateAppformixNodeRole(tx *sql.Tx, model *models.AppformixNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAppformixNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningState),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		string(model.DisplayName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		string(model.ProvisioningLog))
	return err
}

func scanAppformixNodeRole(rows *sql.Rows) (*models.AppformixNodeRole, error) {
	m := models.MakeAppformixNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ProvisioningState,
		&jsonFQName,
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
		&m.IDPerms.Enable,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.ProvisioningStartTime,
		&m.DisplayName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&m.ProvisioningLog); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListAppformixNodeRole(tx *sql.Tx) ([]*models.AppformixNodeRole, error) {
	result := models.MakeAppformixNodeRoleSlice()
	rows, err := tx.Query(listAppformixNodeRoleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAppformixNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAppformixNodeRole(tx *sql.Tx, uuid string) (*models.AppformixNodeRole, error) {
	rows, err := tx.Query(showAppformixNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAppformixNodeRole(rows)
	}
	return nil, nil
}

func UpdateAppformixNodeRole(tx *sql.Tx, uuid string, model *models.AppformixNodeRole) error {
	return nil
}

func DeleteAppformixNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAppformixNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
