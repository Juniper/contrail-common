package db

// contrail_controller_node_role

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertContrailControllerNodeRoleQuery = "insert into `contrail_controller_node_role` (`provisioning_state`,`provisioning_log`,`fq_name`,`display_name`,`provisioning_progress_stage`,`owner_access`,`global_access`,`share`,`owner`,`provisioning_start_time`,`provisioning_progress`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailControllerNodeRoleQuery = "update `contrail_controller_node_role` set `provisioning_state` = ?,`provisioning_log` = ?,`fq_name` = ?,`display_name` = ?,`provisioning_progress_stage` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`provisioning_start_time` = ?,`provisioning_progress` = ?,`uuid` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`key_value_pair` = ?;"
const deleteContrailControllerNodeRoleQuery = "delete from `contrail_controller_node_role` where uuid = ?"
const listContrailControllerNodeRoleQuery = "select `provisioning_state`,`provisioning_log`,`fq_name`,`display_name`,`provisioning_progress_stage`,`owner_access`,`global_access`,`share`,`owner`,`provisioning_start_time`,`provisioning_progress`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair` from `contrail_controller_node_role`"
const showContrailControllerNodeRoleQuery = "select `provisioning_state`,`provisioning_log`,`fq_name`,`display_name`,`provisioning_progress_stage`,`owner_access`,`global_access`,`share`,`owner`,`provisioning_start_time`,`provisioning_progress`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair` from `contrail_controller_node_role` where uuid = ?"

func CreateContrailControllerNodeRole(tx *sql.Tx, model *models.ContrailControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailControllerNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningState),
		string(model.ProvisioningLog),
		util.MustJSON(model.FQName),
		string(model.DisplayName),
		string(model.ProvisioningProgressStage),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.ProvisioningStartTime),
		int(model.ProvisioningProgress),
		string(model.UUID),
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
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanContrailControllerNodeRole(rows *sql.Rows) (*models.ContrailControllerNodeRole, error) {
	m := models.MakeContrailControllerNodeRole()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ProvisioningState,
		&m.ProvisioningLog,
		&jsonFQName,
		&m.DisplayName,
		&m.ProvisioningProgressStage,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.ProvisioningStartTime,
		&m.ProvisioningProgress,
		&m.UUID,
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
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createContrailControllerNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListContrailControllerNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ContrailControllerNodeRole, error) {
	result := models.MakeContrailControllerNodeRoleSlice()
	whereQuery, values := createContrailControllerNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listContrailControllerNodeRoleQuery)
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
