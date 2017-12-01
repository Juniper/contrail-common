package db

// appformix_node_role

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertAppformixNodeRoleQuery = "insert into `appformix_node_role` (`provisioning_log`,`uuid`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAppformixNodeRoleQuery = "update `appformix_node_role` set `provisioning_log` = ?,`uuid` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`provisioning_state` = ?;"
const deleteAppformixNodeRoleQuery = "delete from `appformix_node_role` where uuid = ?"
const listAppformixNodeRoleQuery = "select `provisioning_log`,`uuid`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`provisioning_state` from `appformix_node_role`"
const showAppformixNodeRoleQuery = "select `provisioning_log`,`uuid`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`provisioning_state` from `appformix_node_role` where uuid = ?"

func CreateAppformixNodeRole(tx *sql.Tx, model *models.AppformixNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAppformixNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningLog),
		string(model.UUID),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.ProvisioningState))
	return err
}

func scanAppformixNodeRole(rows *sql.Rows) (*models.AppformixNodeRole, error) {
	m := models.MakeAppformixNodeRole()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.ProvisioningLog,
		&m.UUID,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.ProvisioningStartTime,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.ProvisioningState); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createAppformixNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
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

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListAppformixNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AppformixNodeRole, error) {
	result := models.MakeAppformixNodeRoleSlice()
	whereQuery, values := createAppformixNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAppformixNodeRoleQuery)
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
