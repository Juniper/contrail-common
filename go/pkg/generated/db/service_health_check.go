package db
// service_health_check

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceHealthCheckQuery = "insert into `service_health_check` (`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`timeoutUsecs`,`expected_codes`,`monitor_type`,`url_path`,`health_check_type`,`http_method`,`delayUsecs`,`enabled`,`delay`,`max_retries`,`timeout`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceHealthCheckQuery = "update `service_health_check` set `uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`timeoutUsecs` = ?,`expected_codes` = ?,`monitor_type` = ?,`url_path` = ?,`health_check_type` = ?,`http_method` = ?,`delayUsecs` = ?,`enabled` = ?,`delay` = ?,`max_retries` = ?,`timeout` = ?;"
const deleteServiceHealthCheckQuery = "delete from `service_health_check`"
const selectServiceHealthCheckQuery = "select `uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`timeoutUsecs`,`expected_codes`,`monitor_type`,`url_path`,`health_check_type`,`http_method`,`delayUsecs`,`enabled`,`delay`,`max_retries`,`timeout` from `service_health_check`"

func CreateServiceHealthCheck(tx *sql.Tx, model *models.ServiceHealthCheck) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceHealthCheckQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.UUID,
    model.FQName,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.ServiceHealthCheckProperties.TimeoutUsecs,
    model.ServiceHealthCheckProperties.ExpectedCodes,
    model.ServiceHealthCheckProperties.MonitorType,
    model.ServiceHealthCheckProperties.URLPath,
    model.ServiceHealthCheckProperties.HealthCheckType,
    model.ServiceHealthCheckProperties.HTTPMethod,
    model.ServiceHealthCheckProperties.DelayUsecs,
    model.ServiceHealthCheckProperties.Enabled,
    model.ServiceHealthCheckProperties.Delay,
    model.ServiceHealthCheckProperties.MaxRetries,
    model.ServiceHealthCheckProperties.Timeout)
    return err
}

func ListServiceHealthCheck(tx *sql.Tx) ([]*models.ServiceHealthCheck, error) {
    result := models.MakeServiceHealthCheckSlice()
    rows, err := tx.Query(selectServiceHealthCheckQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceHealthCheck()
            if err := rows.Scan(&m.UUID,
                &m.FQName,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Enable,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.ServiceHealthCheckProperties.TimeoutUsecs,
                &m.ServiceHealthCheckProperties.ExpectedCodes,
                &m.ServiceHealthCheckProperties.MonitorType,
                &m.ServiceHealthCheckProperties.URLPath,
                &m.ServiceHealthCheckProperties.HealthCheckType,
                &m.ServiceHealthCheckProperties.HTTPMethod,
                &m.ServiceHealthCheckProperties.DelayUsecs,
                &m.ServiceHealthCheckProperties.Enabled,
                &m.ServiceHealthCheckProperties.Delay,
                &m.ServiceHealthCheckProperties.MaxRetries,
                &m.ServiceHealthCheckProperties.Timeout); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceHealthCheck(db *sql.DB, id string, model *models.ServiceHealthCheck) error {
    return nil
}

func UpdateServiceHealthCheck(db *sql.DB, id string, model *models.ServiceHealthCheck) error {
    return nil
}

func DeleteServiceHealthCheck(db *sql.DB, id string) error {
    return nil
}