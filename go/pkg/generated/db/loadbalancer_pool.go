package db
// loadbalancer_pool

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`key_value_pair`,`loadbalancer_pool_provider`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`annotations_key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`session_persistence`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerPoolQuery = "update `loadbalancer_pool` set `key_value_pair` = ?,`loadbalancer_pool_provider` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`annotations_key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`status_description` = ?,`loadbalancer_method` = ?,`status` = ?,`protocol` = ?,`subnet_id` = ?,`session_persistence` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool`"
const selectLoadbalancerPoolQuery = "select `key_value_pair`,`loadbalancer_pool_provider`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`annotations_key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`session_persistence`,`uuid`,`fq_name` from `loadbalancer_pool`"

func CreateLoadbalancerPool(tx *sql.Tx, model *models.LoadbalancerPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerPoolQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.LoadbalancerPoolCustomAttributes.KeyValuePair,
    model.LoadbalancerPoolProvider,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.LoadbalancerPoolProperties.AdminState,
    model.LoadbalancerPoolProperties.PersistenceCookieName,
    model.LoadbalancerPoolProperties.StatusDescription,
    model.LoadbalancerPoolProperties.LoadbalancerMethod,
    model.LoadbalancerPoolProperties.Status,
    model.LoadbalancerPoolProperties.Protocol,
    model.LoadbalancerPoolProperties.SubnetID,
    model.LoadbalancerPoolProperties.SessionPersistence,
    model.UUID,
    model.FQName)
    return err
}

func ListLoadbalancerPool(tx *sql.Tx) ([]*models.LoadbalancerPool, error) {
    result := models.MakeLoadbalancerPoolSlice()
    rows, err := tx.Query(selectLoadbalancerPoolQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLoadbalancerPool()
            if err := rows.Scan(&m.LoadbalancerPoolCustomAttributes.KeyValuePair,
                &m.LoadbalancerPoolProvider,
                &m.IDPerms.Description,
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
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.LoadbalancerPoolProperties.AdminState,
                &m.LoadbalancerPoolProperties.PersistenceCookieName,
                &m.LoadbalancerPoolProperties.StatusDescription,
                &m.LoadbalancerPoolProperties.LoadbalancerMethod,
                &m.LoadbalancerPoolProperties.Status,
                &m.LoadbalancerPoolProperties.Protocol,
                &m.LoadbalancerPoolProperties.SubnetID,
                &m.LoadbalancerPoolProperties.SessionPersistence,
                &m.UUID,
                &m.FQName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowLoadbalancerPool(db *sql.DB, id string, model *models.LoadbalancerPool) error {
    return nil
}

func UpdateLoadbalancerPool(db *sql.DB, id string, model *models.LoadbalancerPool) error {
    return nil
}

func DeleteLoadbalancerPool(db *sql.DB, id string) error {
    return nil
}