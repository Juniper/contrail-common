package db

// service_health_check

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceHealthCheckQuery = "insert into `service_health_check` (`expected_codes`,`max_retries`,`delayUsecs`,`timeoutUsecs`,`delay`,`health_check_type`,`enabled`,`url_path`,`http_method`,`timeout`,`monitor_type`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceHealthCheckQuery = "update `service_health_check` set `expected_codes` = ?,`max_retries` = ?,`delayUsecs` = ?,`timeoutUsecs` = ?,`delay` = ?,`health_check_type` = ?,`enabled` = ?,`url_path` = ?,`http_method` = ?,`timeout` = ?,`monitor_type` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?;"
const deleteServiceHealthCheckQuery = "delete from `service_health_check` where uuid = ?"
const listServiceHealthCheckQuery = "select `expected_codes`,`max_retries`,`delayUsecs`,`timeoutUsecs`,`delay`,`health_check_type`,`enabled`,`url_path`,`http_method`,`timeout`,`monitor_type`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`display_name` from `service_health_check`"
const showServiceHealthCheckQuery = "select `expected_codes`,`max_retries`,`delayUsecs`,`timeoutUsecs`,`delay`,`health_check_type`,`enabled`,`url_path`,`http_method`,`timeout`,`monitor_type`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`display_name` from `service_health_check` where uuid = ?"

func CreateServiceHealthCheck(tx *sql.Tx, model *models.ServiceHealthCheck) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceHealthCheckQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ServiceHealthCheckProperties.ExpectedCodes),
		int(model.ServiceHealthCheckProperties.MaxRetries),
		int(model.ServiceHealthCheckProperties.DelayUsecs),
		int(model.ServiceHealthCheckProperties.TimeoutUsecs),
		int(model.ServiceHealthCheckProperties.Delay),
		string(model.ServiceHealthCheckProperties.HealthCheckType),
		bool(model.ServiceHealthCheckProperties.Enabled),
		string(model.ServiceHealthCheckProperties.URLPath),
		string(model.ServiceHealthCheckProperties.HTTPMethod),
		int(model.ServiceHealthCheckProperties.Timeout),
		string(model.ServiceHealthCheckProperties.MonitorType),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName))
	return err
}

func scanServiceHealthCheck(rows *sql.Rows) (*models.ServiceHealthCheck, error) {
	m := models.MakeServiceHealthCheck()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.ServiceHealthCheckProperties.ExpectedCodes,
		&m.ServiceHealthCheckProperties.MaxRetries,
		&m.ServiceHealthCheckProperties.DelayUsecs,
		&m.ServiceHealthCheckProperties.TimeoutUsecs,
		&m.ServiceHealthCheckProperties.Delay,
		&m.ServiceHealthCheckProperties.HealthCheckType,
		&m.ServiceHealthCheckProperties.Enabled,
		&m.ServiceHealthCheckProperties.URLPath,
		&m.ServiceHealthCheckProperties.HTTPMethod,
		&m.ServiceHealthCheckProperties.Timeout,
		&m.ServiceHealthCheckProperties.MonitorType,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListServiceHealthCheck(tx *sql.Tx) ([]*models.ServiceHealthCheck, error) {
	result := models.MakeServiceHealthCheckSlice()
	rows, err := tx.Query(listServiceHealthCheckQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceHealthCheck(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceHealthCheck(tx *sql.Tx, uuid string) (*models.ServiceHealthCheck, error) {
	rows, err := tx.Query(showServiceHealthCheckQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceHealthCheck(rows)
	}
	return nil, nil
}

func UpdateServiceHealthCheck(tx *sql.Tx, uuid string, model *models.ServiceHealthCheck) error {
	return nil
}

func DeleteServiceHealthCheck(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceHealthCheckQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
