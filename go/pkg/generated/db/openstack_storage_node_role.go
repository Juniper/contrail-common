package db

// openstack_storage_node_role

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertOpenstackStorageNodeRoleQuery = "insert into `openstack_storage_node_role` (`fq_name`,`provisioning_progress`,`provisioning_start_time`,`journal_drives`,`osd_drives`,`storage_backend_bond_interface_members`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`provisioning_log`,`provisioning_progress_stage`,`uuid`,`provisioning_state`,`storage_access_bond_interface_members`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackStorageNodeRoleQuery = "update `openstack_storage_node_role` set `fq_name` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`journal_drives` = ?,`osd_drives` = ?,`storage_backend_bond_interface_members` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?,`uuid` = ?,`provisioning_state` = ?,`storage_access_bond_interface_members` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteOpenstackStorageNodeRoleQuery = "delete from `openstack_storage_node_role` where uuid = ?"
const listOpenstackStorageNodeRoleQuery = "select `fq_name`,`provisioning_progress`,`provisioning_start_time`,`journal_drives`,`osd_drives`,`storage_backend_bond_interface_members`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`provisioning_log`,`provisioning_progress_stage`,`uuid`,`provisioning_state`,`storage_access_bond_interface_members`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`display_name`,`key_value_pair` from `openstack_storage_node_role`"
const showOpenstackStorageNodeRoleQuery = "select `fq_name`,`provisioning_progress`,`provisioning_start_time`,`journal_drives`,`osd_drives`,`storage_backend_bond_interface_members`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`provisioning_log`,`provisioning_progress_stage`,`uuid`,`provisioning_state`,`storage_access_bond_interface_members`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`display_name`,`key_value_pair` from `openstack_storage_node_role` where uuid = ?"

func CreateOpenstackStorageNodeRole(tx *sql.Tx, model *models.OpenstackStorageNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackStorageNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.JournalDrives),
		string(model.OsdDrives),
		string(model.StorageBackendBondInterfaceMembers),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.ProvisioningLog),
		string(model.ProvisioningProgressStage),
		string(model.UUID),
		string(model.ProvisioningState),
		string(model.StorageAccessBondInterfaceMembers),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanOpenstackStorageNodeRole(rows *sql.Rows) (*models.OpenstackStorageNodeRole, error) {
	m := models.MakeOpenstackStorageNodeRole()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.JournalDrives,
		&m.OsdDrives,
		&m.StorageBackendBondInterfaceMembers,
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
		&m.IDPerms.LastModified,
		&m.ProvisioningLog,
		&m.ProvisioningProgressStage,
		&m.UUID,
		&m.ProvisioningState,
		&m.StorageAccessBondInterfaceMembers,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListOpenstackStorageNodeRole(tx *sql.Tx) ([]*models.OpenstackStorageNodeRole, error) {
	result := models.MakeOpenstackStorageNodeRoleSlice()
	rows, err := tx.Query(listOpenstackStorageNodeRoleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanOpenstackStorageNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackStorageNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackStorageNodeRole, error) {
	rows, err := tx.Query(showOpenstackStorageNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackStorageNodeRole(rows)
	}
	return nil, nil
}

func UpdateOpenstackStorageNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackStorageNodeRole) error {
	return nil
}

func DeleteOpenstackStorageNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackStorageNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
