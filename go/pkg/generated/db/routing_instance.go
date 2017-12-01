package db

// routing_instance

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertRoutingInstanceQuery = "insert into `routing_instance` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRoutingInstanceQuery = "update `routing_instance` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteRoutingInstanceQuery = "delete from `routing_instance` where uuid = ?"
const listRoutingInstanceQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair` from `routing_instance`"
const showRoutingInstanceQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair` from `routing_instance` where uuid = ?"

func CreateRoutingInstance(tx *sql.Tx, model *models.RoutingInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanRoutingInstance(rows *sql.Rows) (*models.RoutingInstance, error) {
	m := models.MakeRoutingInstance()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Creator,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListRoutingInstance(tx *sql.Tx) ([]*models.RoutingInstance, error) {
	result := models.MakeRoutingInstanceSlice()
	rows, err := tx.Query(listRoutingInstanceQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanRoutingInstance(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRoutingInstance(tx *sql.Tx, uuid string) (*models.RoutingInstance, error) {
	rows, err := tx.Query(showRoutingInstanceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRoutingInstance(rows)
	}
	return nil, nil
}

func UpdateRoutingInstance(tx *sql.Tx, uuid string, model *models.RoutingInstance) error {
	return nil
}

func DeleteRoutingInstance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRoutingInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
