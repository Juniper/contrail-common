package db

// port_tuple

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertPortTupleQuery = "insert into `port_tuple` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePortTupleQuery = "update `port_tuple` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?,`fq_name` = ?;"
const deletePortTupleQuery = "delete from `port_tuple` where uuid = ?"
const listPortTupleQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`fq_name` from `port_tuple`"
const showPortTupleQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`fq_name` from `port_tuple` where uuid = ?"

func CreatePortTuple(tx *sql.Tx, model *models.PortTuple) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPortTupleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanPortTuple(rows *sql.Rows) (*models.PortTuple, error) {
	m := models.MakePortTuple()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Description,
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
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListPortTuple(tx *sql.Tx) ([]*models.PortTuple, error) {
	result := models.MakePortTupleSlice()
	rows, err := tx.Query(listPortTupleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanPortTuple(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPortTuple(tx *sql.Tx, uuid string) (*models.PortTuple, error) {
	rows, err := tx.Query(showPortTupleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPortTuple(rows)
	}
	return nil, nil
}

func UpdatePortTuple(tx *sql.Tx, uuid string, model *models.PortTuple) error {
	return nil
}

func DeletePortTuple(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePortTupleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
