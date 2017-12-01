package db

// controller_node_role

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertControllerNodeRoleQuery = "insert into `controller_node_role` (`performance_drives`,`storage_management_bond_interface_members`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_state`,`provisioning_progress_stage`,`capacity_drives`,`internalapi_bond_interface_members`,`provisioning_progress`,`fq_name`,`provisioning_log`,`uuid`,`provisioning_start_time`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateControllerNodeRoleQuery = "update `controller_node_role` set `performance_drives` = ?,`storage_management_bond_interface_members` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_state` = ?,`provisioning_progress_stage` = ?,`capacity_drives` = ?,`internalapi_bond_interface_members` = ?,`provisioning_progress` = ?,`fq_name` = ?,`provisioning_log` = ?,`uuid` = ?,`provisioning_start_time` = ?;"
const deleteControllerNodeRoleQuery = "delete from `controller_node_role` where uuid = ?"
const listControllerNodeRoleQuery = "select `performance_drives`,`storage_management_bond_interface_members`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_state`,`provisioning_progress_stage`,`capacity_drives`,`internalapi_bond_interface_members`,`provisioning_progress`,`fq_name`,`provisioning_log`,`uuid`,`provisioning_start_time` from `controller_node_role`"
const showControllerNodeRoleQuery = "select `performance_drives`,`storage_management_bond_interface_members`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_state`,`provisioning_progress_stage`,`capacity_drives`,`internalapi_bond_interface_members`,`provisioning_progress`,`fq_name`,`provisioning_log`,`uuid`,`provisioning_start_time` from `controller_node_role` where uuid = ?"

func CreateControllerNodeRole(tx *sql.Tx, model *models.ControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.PerformanceDrives),
		string(model.StorageManagementBondInterfaceMembers),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.ProvisioningState),
		string(model.ProvisioningProgressStage),
		string(model.CapacityDrives),
		string(model.InternalapiBondInterfaceMembers),
		int(model.ProvisioningProgress),
		util.MustJSON(model.FQName),
		string(model.ProvisioningLog),
		string(model.UUID),
		string(model.ProvisioningStartTime))
	return err
}

func scanControllerNodeRole(rows *sql.Rows) (*models.ControllerNodeRole, error) {
	m := models.MakeControllerNodeRole()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.PerformanceDrives,
		&m.StorageManagementBondInterfaceMembers,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ProvisioningState,
		&m.ProvisioningProgressStage,
		&m.CapacityDrives,
		&m.InternalapiBondInterfaceMembers,
		&m.ProvisioningProgress,
		&jsonFQName,
		&m.ProvisioningLog,
		&m.UUID,
		&m.ProvisioningStartTime); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListControllerNodeRole(tx *sql.Tx) ([]*models.ControllerNodeRole, error) {
	result := models.MakeControllerNodeRoleSlice()
	rows, err := tx.Query(listControllerNodeRoleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanControllerNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowControllerNodeRole(tx *sql.Tx, uuid string) (*models.ControllerNodeRole, error) {
	rows, err := tx.Query(showControllerNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanControllerNodeRole(rows)
	}
	return nil, nil
}

func UpdateControllerNodeRole(tx *sql.Tx, uuid string, model *models.ControllerNodeRole) error {
	return nil
}

func DeleteControllerNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
