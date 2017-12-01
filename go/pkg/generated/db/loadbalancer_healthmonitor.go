package db
// loadbalancer_healthmonitor

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`fq_name`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`delay`,`expected_codes`,`max_retries`,`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerHealthmonitorQuery = "update `loadbalancer_healthmonitor` set `fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?,`delay` = ?,`expected_codes` = ?,`max_retries` = ?,`http_method` = ?,`admin_state` = ?,`timeout` = ?,`url_path` = ?,`monitor_type` = ?;"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor`"
const selectLoadbalancerHealthmonitorQuery = "select `fq_name`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`delay`,`expected_codes`,`max_retries`,`http_method`,`admin_state`,`timeout`,`url_path`,`monitor_type` from `loadbalancer_healthmonitor`"

func CreateLoadbalancerHealthmonitor(tx *sql.Tx, model *models.LoadbalancerHealthmonitor) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID,
    model.LoadbalancerHealthmonitorProperties.Delay,
    model.LoadbalancerHealthmonitorProperties.ExpectedCodes,
    model.LoadbalancerHealthmonitorProperties.MaxRetries,
    model.LoadbalancerHealthmonitorProperties.HTTPMethod,
    model.LoadbalancerHealthmonitorProperties.AdminState,
    model.LoadbalancerHealthmonitorProperties.Timeout,
    model.LoadbalancerHealthmonitorProperties.URLPath,
    model.LoadbalancerHealthmonitorProperties.MonitorType)
    return err
}

func ListLoadbalancerHealthmonitor(tx *sql.Tx) ([]*models.LoadbalancerHealthmonitor, error) {
    result := models.MakeLoadbalancerHealthmonitorSlice()
    rows, err := tx.Query(selectLoadbalancerHealthmonitorQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLoadbalancerHealthmonitor()
            if err := rows.Scan(&m.FQName,
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
                &m.IDPerms.Created,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.UUID,
                &m.LoadbalancerHealthmonitorProperties.Delay,
                &m.LoadbalancerHealthmonitorProperties.ExpectedCodes,
                &m.LoadbalancerHealthmonitorProperties.MaxRetries,
                &m.LoadbalancerHealthmonitorProperties.HTTPMethod,
                &m.LoadbalancerHealthmonitorProperties.AdminState,
                &m.LoadbalancerHealthmonitorProperties.Timeout,
                &m.LoadbalancerHealthmonitorProperties.URLPath,
                &m.LoadbalancerHealthmonitorProperties.MonitorType); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowLoadbalancerHealthmonitor(db *sql.DB, id string, model *models.LoadbalancerHealthmonitor) error {
    return nil
}

func UpdateLoadbalancerHealthmonitor(db *sql.DB, id string, model *models.LoadbalancerHealthmonitor) error {
    return nil
}

func DeleteLoadbalancerHealthmonitor(db *sql.DB, id string) error {
    return nil
}