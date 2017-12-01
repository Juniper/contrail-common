package db

// openstack_compute_node_role

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertOpenstackComputeNodeRoleQuery = "insert into `openstack_compute_node_role` (`provisioning_state`,`vrouter_bond_interface`,`vrouter_bond_interface_members`,`uuid`,`display_name`,`provisioning_progress`,`default_gateway`,`fq_name`,`provisioning_progress_stage`,`provisioning_start_time`,`key_value_pair`,`vrouter_type`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackComputeNodeRoleQuery = "update `openstack_compute_node_role` set `provisioning_state` = ?,`vrouter_bond_interface` = ?,`vrouter_bond_interface_members` = ?,`uuid` = ?,`display_name` = ?,`provisioning_progress` = ?,`default_gateway` = ?,`fq_name` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`key_value_pair` = ?,`vrouter_type` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_log` = ?;"
const deleteOpenstackComputeNodeRoleQuery = "delete from `openstack_compute_node_role` where uuid = ?"
const listOpenstackComputeNodeRoleQuery = "select `provisioning_state`,`vrouter_bond_interface`,`vrouter_bond_interface_members`,`uuid`,`display_name`,`provisioning_progress`,`default_gateway`,`fq_name`,`provisioning_progress_stage`,`provisioning_start_time`,`key_value_pair`,`vrouter_type`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log` from `openstack_compute_node_role`"
const showOpenstackComputeNodeRoleQuery = "select `provisioning_state`,`vrouter_bond_interface`,`vrouter_bond_interface_members`,`uuid`,`display_name`,`provisioning_progress`,`default_gateway`,`fq_name`,`provisioning_progress_stage`,`provisioning_start_time`,`key_value_pair`,`vrouter_type`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log` from `openstack_compute_node_role` where uuid = ?"

func CreateOpenstackComputeNodeRole(tx *sql.Tx, model *models.OpenstackComputeNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackComputeNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningState),
		string(model.VrouterBondInterface),
		string(model.VrouterBondInterfaceMembers),
		string(model.UUID),
		string(model.DisplayName),
		int(model.ProvisioningProgress),
		string(model.DefaultGateway),
		util.MustJSON(model.FQName),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.VrouterType),
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
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.ProvisioningLog))
	return err
}

func scanOpenstackComputeNodeRole(rows *sql.Rows) (*models.OpenstackComputeNodeRole, error) {
	m := models.MakeOpenstackComputeNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ProvisioningState,
		&m.VrouterBondInterface,
		&m.VrouterBondInterfaceMembers,
		&m.UUID,
		&m.DisplayName,
		&m.ProvisioningProgress,
		&m.DefaultGateway,
		&jsonFQName,
		&m.ProvisioningProgressStage,
		&m.ProvisioningStartTime,
		&jsonAnnotationsKeyValuePair,
		&m.VrouterType,
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
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ProvisioningLog); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListOpenstackComputeNodeRole(tx *sql.Tx) ([]*models.OpenstackComputeNodeRole, error) {
	result := models.MakeOpenstackComputeNodeRoleSlice()
	rows, err := tx.Query(listOpenstackComputeNodeRoleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanOpenstackComputeNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackComputeNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackComputeNodeRole, error) {
	rows, err := tx.Query(showOpenstackComputeNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackComputeNodeRole(rows)
	}
	return nil, nil
}

func UpdateOpenstackComputeNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackComputeNodeRole) error {
	return nil
}

func DeleteOpenstackComputeNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackComputeNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
