package db

// loadbalancer_healthmonitor

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`max_retries`,`http_method`,`admin_state`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerHealthmonitorQuery = "update `loadbalancer_healthmonitor` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`timeout` = ?,`url_path` = ?,`monitor_type` = ?,`delay` = ?,`expected_codes` = ?,`max_retries` = ?,`http_method` = ?,`admin_state` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor` where uuid = ?"
const listLoadbalancerHealthmonitorQuery = "select `key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`max_retries`,`http_method`,`admin_state`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name` from `loadbalancer_healthmonitor`"
const showLoadbalancerHealthmonitorQuery = "select `key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`max_retries`,`http_method`,`admin_state`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name` from `loadbalancer_healthmonitor` where uuid = ?"

func CreateLoadbalancerHealthmonitor(tx *sql.Tx, model *models.LoadbalancerHealthmonitor) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		int(model.LoadbalancerHealthmonitorProperties.Timeout),
		string(model.LoadbalancerHealthmonitorProperties.URLPath),
		string(model.LoadbalancerHealthmonitorProperties.MonitorType),
		int(model.LoadbalancerHealthmonitorProperties.Delay),
		string(model.LoadbalancerHealthmonitorProperties.ExpectedCodes),
		int(model.LoadbalancerHealthmonitorProperties.MaxRetries),
		string(model.LoadbalancerHealthmonitorProperties.HTTPMethod),
		bool(model.LoadbalancerHealthmonitorProperties.AdminState),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.DisplayName))
	return err
}

func scanLoadbalancerHealthmonitor(rows *sql.Rows) (*models.LoadbalancerHealthmonitor, error) {
	m := models.MakeLoadbalancerHealthmonitor()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.LoadbalancerHealthmonitorProperties.Timeout,
		&m.LoadbalancerHealthmonitorProperties.URLPath,
		&m.LoadbalancerHealthmonitorProperties.MonitorType,
		&m.LoadbalancerHealthmonitorProperties.Delay,
		&m.LoadbalancerHealthmonitorProperties.ExpectedCodes,
		&m.LoadbalancerHealthmonitorProperties.MaxRetries,
		&m.LoadbalancerHealthmonitorProperties.HTTPMethod,
		&m.LoadbalancerHealthmonitorProperties.AdminState,
		&m.UUID,
		&jsonFQName,
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
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createLoadbalancerHealthmonitorWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["url_path"]; ok {
		results = append(results, "url_path = ?")
		values = append(values, value)
	}

	if value, ok := where["monitor_type"]; ok {
		results = append(results, "monitor_type = ?")
		values = append(values, value)
	}

	if value, ok := where["expected_codes"]; ok {
		results = append(results, "expected_codes = ?")
		values = append(values, value)
	}

	if value, ok := where["http_method"]; ok {
		results = append(results, "http_method = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancerHealthmonitor(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LoadbalancerHealthmonitor, error) {
	result := models.MakeLoadbalancerHealthmonitorSlice()
	whereQuery, values := createLoadbalancerHealthmonitorWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerHealthmonitorQuery)
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
		m, _ := scanLoadbalancerHealthmonitor(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerHealthmonitor(tx *sql.Tx, uuid string) (*models.LoadbalancerHealthmonitor, error) {
	rows, err := tx.Query(showLoadbalancerHealthmonitorQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerHealthmonitor(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerHealthmonitor(tx *sql.Tx, uuid string, model *models.LoadbalancerHealthmonitor) error {
	return nil
}

func DeleteLoadbalancerHealthmonitor(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerHealthmonitorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
