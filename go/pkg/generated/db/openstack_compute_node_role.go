package db

// openstack_compute_node_role

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertOpenstackComputeNodeRoleQuery = "insert into `openstack_compute_node_role` (`fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`default_gateway`,`vrouter_bond_interface_members`,`vrouter_type`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`key_value_pair`,`provisioning_log`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`provisioning_start_time`,`provisioning_state`,`vrouter_bond_interface`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackComputeNodeRoleQuery = "update `openstack_compute_node_role` set `fq_name` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`default_gateway` = ?,`vrouter_bond_interface_members` = ?,`vrouter_type` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`key_value_pair` = ?,`provisioning_log` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`vrouter_bond_interface` = ?,`display_name` = ?;"
const deleteOpenstackComputeNodeRoleQuery = "delete from `openstack_compute_node_role` where uuid = ?"
const listOpenstackComputeNodeRoleQuery = "select `fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`default_gateway`,`vrouter_bond_interface_members`,`vrouter_type`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`key_value_pair`,`provisioning_log`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`provisioning_start_time`,`provisioning_state`,`vrouter_bond_interface`,`display_name` from `openstack_compute_node_role`"
const showOpenstackComputeNodeRoleQuery = "select `fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`default_gateway`,`vrouter_bond_interface_members`,`vrouter_type`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`key_value_pair`,`provisioning_log`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`provisioning_start_time`,`provisioning_state`,`vrouter_bond_interface`,`display_name` from `openstack_compute_node_role` where uuid = ?"

func CreateOpenstackComputeNodeRole(tx *sql.Tx, model *models.OpenstackComputeNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackComputeNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.DefaultGateway),
		string(model.VrouterBondInterfaceMembers),
		string(model.VrouterType),
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
		string(model.IDPerms.Description),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.VrouterBondInterface),
		string(model.DisplayName))
	return err
}

func scanOpenstackComputeNodeRole(rows *sql.Rows) (*models.OpenstackComputeNodeRole, error) {
	m := models.MakeOpenstackComputeNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.DefaultGateway,
		&m.VrouterBondInterfaceMembers,
		&m.VrouterType,
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
		&m.IDPerms.Description,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningLog,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.VrouterBondInterface,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createOpenstackComputeNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["default_gateway"]; ok {
		results = append(results, "default_gateway = ?")
		values = append(values, value)
	}

	if value, ok := where["vrouter_bond_interface_members"]; ok {
		results = append(results, "vrouter_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["vrouter_type"]; ok {
		results = append(results, "vrouter_type = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["vrouter_bond_interface"]; ok {
		results = append(results, "vrouter_bond_interface = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListOpenstackComputeNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.OpenstackComputeNodeRole, error) {
	result := models.MakeOpenstackComputeNodeRoleSlice()
	whereQuery, values := createOpenstackComputeNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listOpenstackComputeNodeRoleQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
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
