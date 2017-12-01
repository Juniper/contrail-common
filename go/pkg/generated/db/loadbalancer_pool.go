package db

// loadbalancer_pool

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`loadbalancer_pool_provider`,`display_name`,`key_value_pair`,`uuid`,`fq_name`,`loadbalancer_pool_custom_attributes_key_value_pair`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerPoolQuery = "update `loadbalancer_pool` set `status_description` = ?,`loadbalancer_method` = ?,`status` = ?,`protocol` = ?,`subnet_id` = ?,`session_persistence` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`loadbalancer_pool_provider` = ?,`display_name` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`loadbalancer_pool_custom_attributes_key_value_pair` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool` where uuid = ?"
const listLoadbalancerPoolQuery = "select `status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`loadbalancer_pool_provider`,`display_name`,`key_value_pair`,`uuid`,`fq_name`,`loadbalancer_pool_custom_attributes_key_value_pair`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `loadbalancer_pool`"
const showLoadbalancerPoolQuery = "select `status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`loadbalancer_pool_provider`,`display_name`,`key_value_pair`,`uuid`,`fq_name`,`loadbalancer_pool_custom_attributes_key_value_pair`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `loadbalancer_pool` where uuid = ?"

func CreateLoadbalancerPool(tx *sql.Tx, model *models.LoadbalancerPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.LoadbalancerPoolProperties.StatusDescription),
		string(model.LoadbalancerPoolProperties.LoadbalancerMethod),
		string(model.LoadbalancerPoolProperties.Status),
		string(model.LoadbalancerPoolProperties.Protocol),
		string(model.LoadbalancerPoolProperties.SubnetID),
		string(model.LoadbalancerPoolProperties.SessionPersistence),
		bool(model.LoadbalancerPoolProperties.AdminState),
		string(model.LoadbalancerPoolProperties.PersistenceCookieName),
		string(model.LoadbalancerPoolProvider),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		util.MustJSON(model.FQName),
		util.MustJSON(model.LoadbalancerPoolCustomAttributes.KeyValuePair),
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
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))
	return err
}

func scanLoadbalancerPool(rows *sql.Rows) (*models.LoadbalancerPool, error) {
	m := models.MakeLoadbalancerPool()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonLoadbalancerPoolCustomAttributesKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.LoadbalancerPoolProperties.StatusDescription,
		&m.LoadbalancerPoolProperties.LoadbalancerMethod,
		&m.LoadbalancerPoolProperties.Status,
		&m.LoadbalancerPoolProperties.Protocol,
		&m.LoadbalancerPoolProperties.SubnetID,
		&m.LoadbalancerPoolProperties.SessionPersistence,
		&m.LoadbalancerPoolProperties.AdminState,
		&m.LoadbalancerPoolProperties.PersistenceCookieName,
		&m.LoadbalancerPoolProvider,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName,
		&jsonLoadbalancerPoolCustomAttributesKeyValuePair,
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
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonLoadbalancerPoolCustomAttributesKeyValuePair), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListLoadbalancerPool(tx *sql.Tx) ([]*models.LoadbalancerPool, error) {
	result := models.MakeLoadbalancerPoolSlice()
	rows, err := tx.Query(listLoadbalancerPoolQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanLoadbalancerPool(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerPool(tx *sql.Tx, uuid string) (*models.LoadbalancerPool, error) {
	rows, err := tx.Query(showLoadbalancerPoolQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerPool(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerPool(tx *sql.Tx, uuid string, model *models.LoadbalancerPool) error {
	return nil
}

func DeleteLoadbalancerPool(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
