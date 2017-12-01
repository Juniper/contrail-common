package db

// analytics_node

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAnalyticsNodeQuery = "insert into `analytics_node` (`key_value_pair`,`analytics_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAnalyticsNodeQuery = "update `analytics_node` set `key_value_pair` = ?,`analytics_node_ip_address` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteAnalyticsNodeQuery = "delete from `analytics_node` where uuid = ?"
const listAnalyticsNodeQuery = "select `key_value_pair`,`analytics_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name` from `analytics_node`"
const showAnalyticsNodeQuery = "select `key_value_pair`,`analytics_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`display_name` from `analytics_node` where uuid = ?"

func CreateAnalyticsNode(tx *sql.Tx, model *models.AnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		string(model.AnalyticsNodeIPAddress),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
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

func scanAnalyticsNode(rows *sql.Rows) (*models.AnalyticsNode, error) {
	m := models.MakeAnalyticsNode()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.AnalyticsNodeIPAddress,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
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

func ListAnalyticsNode(tx *sql.Tx) ([]*models.AnalyticsNode, error) {
	result := models.MakeAnalyticsNodeSlice()
	rows, err := tx.Query(listAnalyticsNodeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAnalyticsNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAnalyticsNode(tx *sql.Tx, uuid string) (*models.AnalyticsNode, error) {
	rows, err := tx.Query(showAnalyticsNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAnalyticsNode(rows)
	}
	return nil, nil
}

func UpdateAnalyticsNode(tx *sql.Tx, uuid string, model *models.AnalyticsNode) error {
	return nil
}

func DeleteAnalyticsNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
