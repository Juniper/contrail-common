package db

// contrail_controller_node_role

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertContrailControllerNodeRoleQuery = "insert into `contrail_controller_node_role` (`uuid`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_log`,`provisioning_progress_stage`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailControllerNodeRoleQuery = "update `contrail_controller_node_role` set `uuid` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?;"
const deleteContrailControllerNodeRoleQuery = "delete from `contrail_controller_node_role` where uuid = ?"
const listContrailControllerNodeRoleQuery = "select `uuid`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_log`,`provisioning_progress_stage` from `contrail_controller_node_role`"
const showContrailControllerNodeRoleQuery = "select `uuid`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_log`,`provisioning_progress_stage` from `contrail_controller_node_role` where uuid = ?"

func CreateContrailControllerNodeRole(tx *sql.Tx, model *models.ContrailControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		util.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.ProvisioningLog),
		string(model.ProvisioningProgressStage))
	return err
}

func scanContrailControllerNodeRole(rows *sql.Rows) (*models.ContrailControllerNodeRole, error) {
	m := models.MakeContrailControllerNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&jsonFQName,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.ProvisioningLog,
		&m.ProvisioningProgressStage); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListContrailControllerNodeRole(tx *sql.Tx) ([]*models.ContrailControllerNodeRole, error) {
	result := models.MakeContrailControllerNodeRoleSlice()
	rows, err := tx.Query(listContrailControllerNodeRoleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanContrailControllerNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailControllerNodeRole(tx *sql.Tx, uuid string) (*models.ContrailControllerNodeRole, error) {
	rows, err := tx.Query(showContrailControllerNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailControllerNodeRole(rows)
	}
	return nil, nil
}

func UpdateContrailControllerNodeRole(tx *sql.Tx, uuid string, model *models.ContrailControllerNodeRole) error {
	return nil
}

func DeleteContrailControllerNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
