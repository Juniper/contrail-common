package db

// contrail_analytics_node

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertContrailAnalyticsNodeQuery = "insert into `contrail_analytics_node` (`uuid`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`,`provisioning_log`,`provisioning_progress_stage`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsNodeQuery = "update `contrail_analytics_node` set `uuid` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`fq_name` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteContrailAnalyticsNodeQuery = "delete from `contrail_analytics_node` where uuid = ?"
const listContrailAnalyticsNodeQuery = "select `uuid`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`,`provisioning_log`,`provisioning_progress_stage`,`display_name`,`key_value_pair` from `contrail_analytics_node`"
const showContrailAnalyticsNodeQuery = "select `uuid`,`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`,`provisioning_log`,`provisioning_progress_stage`,`display_name`,`key_value_pair` from `contrail_analytics_node` where uuid = ?"

func CreateContrailAnalyticsNode(tx *sql.Tx, model *models.ContrailAnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		util.MustJSON(model.FQName),
		string(model.ProvisioningLog),
		string(model.ProvisioningProgressStage),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanContrailAnalyticsNode(rows *sql.Rows) (*models.ContrailAnalyticsNode, error) {
	m := models.MakeContrailAnalyticsNode()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.UUID,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonFQName,
		&m.ProvisioningLog,
		&m.ProvisioningProgressStage,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createContrailAnalyticsNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListContrailAnalyticsNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ContrailAnalyticsNode, error) {
	result := models.MakeContrailAnalyticsNodeSlice()
	whereQuery, values := createContrailAnalyticsNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listContrailAnalyticsNodeQuery)
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
		m, _ := scanContrailAnalyticsNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailAnalyticsNode(tx *sql.Tx, uuid string) (*models.ContrailAnalyticsNode, error) {
	rows, err := tx.Query(showContrailAnalyticsNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailAnalyticsNode(rows)
	}
	return nil, nil
}

func UpdateContrailAnalyticsNode(tx *sql.Tx, uuid string, model *models.ContrailAnalyticsNode) error {
	return nil
}

func DeleteContrailAnalyticsNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
