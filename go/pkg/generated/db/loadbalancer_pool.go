package db

// loadbalancer_pool

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`loadbalancer_pool_provider`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`key_value_pair`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerPoolQuery = "update `loadbalancer_pool` set `protocol` = ?,`subnet_id` = ?,`session_persistence` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`status_description` = ?,`loadbalancer_method` = ?,`status` = ?,`loadbalancer_pool_provider` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`key_value_pair` = ?,`display_name` = ?,`annotations_key_value_pair` = ?;"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool` where uuid = ?"
const listLoadbalancerPoolQuery = "select `protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`loadbalancer_pool_provider`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`key_value_pair`,`display_name`,`annotations_key_value_pair` from `loadbalancer_pool`"
const showLoadbalancerPoolQuery = "select `protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`loadbalancer_pool_provider`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`key_value_pair`,`display_name`,`annotations_key_value_pair` from `loadbalancer_pool` where uuid = ?"

func CreateLoadbalancerPool(tx *sql.Tx, model *models.LoadbalancerPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.LoadbalancerPoolProperties.Protocol),
		string(model.LoadbalancerPoolProperties.SubnetID),
		string(model.LoadbalancerPoolProperties.SessionPersistence),
		bool(model.LoadbalancerPoolProperties.AdminState),
		string(model.LoadbalancerPoolProperties.PersistenceCookieName),
		string(model.LoadbalancerPoolProperties.StatusDescription),
		string(model.LoadbalancerPoolProperties.LoadbalancerMethod),
		string(model.LoadbalancerPoolProperties.Status),
		string(model.LoadbalancerPoolProvider),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		util.MustJSON(model.LoadbalancerPoolCustomAttributes.KeyValuePair),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanLoadbalancerPool(rows *sql.Rows) (*models.LoadbalancerPool, error) {
	m := models.MakeLoadbalancerPool()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonLoadbalancerPoolCustomAttributesKeyValuePair string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.LoadbalancerPoolProperties.Protocol,
		&m.LoadbalancerPoolProperties.SubnetID,
		&m.LoadbalancerPoolProperties.SessionPersistence,
		&m.LoadbalancerPoolProperties.AdminState,
		&m.LoadbalancerPoolProperties.PersistenceCookieName,
		&m.LoadbalancerPoolProperties.StatusDescription,
		&m.LoadbalancerPoolProperties.LoadbalancerMethod,
		&m.LoadbalancerPoolProperties.Status,
		&m.LoadbalancerPoolProvider,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonLoadbalancerPoolCustomAttributesKeyValuePair,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonLoadbalancerPoolCustomAttributesKeyValuePair), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createLoadbalancerPoolWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["protocol"]; ok {
		results = append(results, "protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["subnet_id"]; ok {
		results = append(results, "subnet_id = ?")
		values = append(values, value)
	}

	if value, ok := where["session_persistence"]; ok {
		results = append(results, "session_persistence = ?")
		values = append(values, value)
	}

	if value, ok := where["persistence_cookie_name"]; ok {
		results = append(results, "persistence_cookie_name = ?")
		values = append(values, value)
	}

	if value, ok := where["status_description"]; ok {
		results = append(results, "status_description = ?")
		values = append(values, value)
	}

	if value, ok := where["loadbalancer_method"]; ok {
		results = append(results, "loadbalancer_method = ?")
		values = append(values, value)
	}

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
		values = append(values, value)
	}

	if value, ok := where["loadbalancer_pool_provider"]; ok {
		results = append(results, "loadbalancer_pool_provider = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancerPool(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LoadbalancerPool, error) {
	result := models.MakeLoadbalancerPoolSlice()
	whereQuery, values := createLoadbalancerPoolWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerPoolQuery)
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
