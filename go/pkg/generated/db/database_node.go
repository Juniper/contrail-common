package db

// database_node

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertDatabaseNodeQuery = "insert into `database_node` (`display_name`,`key_value_pair`,`database_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDatabaseNodeQuery = "update `database_node` set `display_name` = ?,`key_value_pair` = ?,`database_node_ip_address` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?;"
const deleteDatabaseNodeQuery = "delete from `database_node` where uuid = ?"
const listDatabaseNodeQuery = "select `display_name`,`key_value_pair`,`database_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `database_node`"
const showDatabaseNodeQuery = "select `display_name`,`key_value_pair`,`database_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `database_node` where uuid = ?"

func CreateDatabaseNode(tx *sql.Tx, model *models.DatabaseNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDatabaseNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.DatabaseNodeIPAddress),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified))
	return err
}

func scanDatabaseNode(rows *sql.Rows) (*models.DatabaseNode, error) {
	m := models.MakeDatabaseNode()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.DatabaseNodeIPAddress,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListDatabaseNode(tx *sql.Tx) ([]*models.DatabaseNode, error) {
	result := models.MakeDatabaseNodeSlice()
	rows, err := tx.Query(listDatabaseNodeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanDatabaseNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDatabaseNode(tx *sql.Tx, uuid string) (*models.DatabaseNode, error) {
	rows, err := tx.Query(showDatabaseNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDatabaseNode(rows)
	}
	return nil, nil
}

func UpdateDatabaseNode(tx *sql.Tx, uuid string, model *models.DatabaseNode) error {
	return nil
}

func DeleteDatabaseNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDatabaseNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
