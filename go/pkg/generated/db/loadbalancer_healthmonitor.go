package db

// loadbalancer_healthmonitor

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`max_retries`,`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerHealthmonitorQuery = "update `loadbalancer_healthmonitor` set `max_retries` = ?,`http_method` = ?,`admin_state` = ?,`timeout` = ?,`url_path` = ?,`monitor_type` = ?,`delay` = ?,`expected_codes` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor` where uuid = ?"
const listLoadbalancerHealthmonitorQuery = "select `max_retries`,`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `loadbalancer_healthmonitor`"
const showLoadbalancerHealthmonitorQuery = "select `max_retries`,`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type`,`delay`,`expected_codes`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `loadbalancer_healthmonitor` where uuid = ?"

func CreateLoadbalancerHealthmonitor(tx *sql.Tx, model *models.LoadbalancerHealthmonitor) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.LoadbalancerHealthmonitorProperties.MaxRetries),
		string(model.LoadbalancerHealthmonitorProperties.HTTPMethod),
		bool(model.LoadbalancerHealthmonitorProperties.AdminState),
		int(model.LoadbalancerHealthmonitorProperties.Timeout),
		string(model.LoadbalancerHealthmonitorProperties.URLPath),
		string(model.LoadbalancerHealthmonitorProperties.MonitorType),
		int(model.LoadbalancerHealthmonitorProperties.Delay),
		string(model.LoadbalancerHealthmonitorProperties.ExpectedCodes),
		string(model.UUID),
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
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share))
	return err
}

func scanLoadbalancerHealthmonitor(rows *sql.Rows) (*models.LoadbalancerHealthmonitor, error) {
	m := models.MakeLoadbalancerHealthmonitor()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.LoadbalancerHealthmonitorProperties.MaxRetries,
		&m.LoadbalancerHealthmonitorProperties.HTTPMethod,
		&m.LoadbalancerHealthmonitorProperties.AdminState,
		&m.LoadbalancerHealthmonitorProperties.Timeout,
		&m.LoadbalancerHealthmonitorProperties.URLPath,
		&m.LoadbalancerHealthmonitorProperties.MonitorType,
		&m.LoadbalancerHealthmonitorProperties.Delay,
		&m.LoadbalancerHealthmonitorProperties.ExpectedCodes,
		&m.UUID,
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

func ListLoadbalancerHealthmonitor(tx *sql.Tx) ([]*models.LoadbalancerHealthmonitor, error) {
	result := models.MakeLoadbalancerHealthmonitorSlice()
	rows, err := tx.Query(listLoadbalancerHealthmonitorQuery)
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
