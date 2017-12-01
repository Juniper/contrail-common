package db

// routing_policy

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertRoutingPolicyQuery = "insert into `routing_policy` (`fq_name`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRoutingPolicyQuery = "update `routing_policy` set `fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?;"
const deleteRoutingPolicyQuery = "delete from `routing_policy` where uuid = ?"
const listRoutingPolicyQuery = "select `fq_name`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid` from `routing_policy`"
const showRoutingPolicyQuery = "select `fq_name`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid` from `routing_policy` where uuid = ?"

func CreateRoutingPolicy(tx *sql.Tx, model *models.RoutingPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
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
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID))
	return err
}

func scanRoutingPolicy(rows *sql.Rows) (*models.RoutingPolicy, error) {
	m := models.MakeRoutingPolicy()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
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
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListRoutingPolicy(tx *sql.Tx) ([]*models.RoutingPolicy, error) {
	result := models.MakeRoutingPolicySlice()
	rows, err := tx.Query(listRoutingPolicyQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanRoutingPolicy(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRoutingPolicy(tx *sql.Tx, uuid string) (*models.RoutingPolicy, error) {
	rows, err := tx.Query(showRoutingPolicyQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRoutingPolicy(rows)
	}
	return nil, nil
}

func UpdateRoutingPolicy(tx *sql.Tx, uuid string, model *models.RoutingPolicy) error {
	return nil
}

func DeleteRoutingPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRoutingPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
