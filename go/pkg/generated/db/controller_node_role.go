package db

// controller_node_role

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertControllerNodeRoleQuery = "insert into `controller_node_role` (`performance_drives`,`storage_management_bond_interface_members`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`capacity_drives`,`provisioning_progress_stage`,`provisioning_log`,`internalapi_bond_interface_members`,`uuid`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateControllerNodeRoleQuery = "update `controller_node_role` set `performance_drives` = ?,`storage_management_bond_interface_members` = ?,`fq_name` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`capacity_drives` = ?,`provisioning_progress_stage` = ?,`provisioning_log` = ?,`internalapi_bond_interface_members` = ?,`uuid` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteControllerNodeRoleQuery = "delete from `controller_node_role` where uuid = ?"
const listControllerNodeRoleQuery = "select `performance_drives`,`storage_management_bond_interface_members`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`capacity_drives`,`provisioning_progress_stage`,`provisioning_log`,`internalapi_bond_interface_members`,`uuid`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `controller_node_role`"
const showControllerNodeRoleQuery = "select `performance_drives`,`storage_management_bond_interface_members`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`capacity_drives`,`provisioning_progress_stage`,`provisioning_log`,`internalapi_bond_interface_members`,`uuid`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `controller_node_role` where uuid = ?"

func CreateControllerNodeRole(tx *sql.Tx, model *models.ControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.PerformanceDrives),
		string(model.StorageManagementBondInterfaceMembers),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.CapacityDrives),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningLog),
		string(model.InternalapiBondInterfaceMembers),
		string(model.UUID),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share))
	return err
}

func scanControllerNodeRole(rows *sql.Rows) (*models.ControllerNodeRole, error) {
	m := models.MakeControllerNodeRole()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.PerformanceDrives,
		&m.StorageManagementBondInterfaceMembers,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.CapacityDrives,
		&m.ProvisioningProgressStage,
		&m.ProvisioningLog,
		&m.InternalapiBondInterfaceMembers,
		&m.UUID,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createControllerNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["performance_drives"]; ok {
		results = append(results, "performance_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["storage_management_bond_interface_members"]; ok {
		results = append(results, "storage_management_bond_interface_members = ?")
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

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	if value, ok := where["capacity_drives"]; ok {
		results = append(results, "capacity_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["internalapi_bond_interface_members"]; ok {
		results = append(results, "internalapi_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListControllerNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ControllerNodeRole, error) {
	result := models.MakeControllerNodeRoleSlice()
	whereQuery, values := createControllerNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listControllerNodeRoleQuery)
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
