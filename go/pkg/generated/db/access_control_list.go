package db

// access_control_list

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAccessControlListQuery = "insert into `access_control_list` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`access_control_list_hash`,`dynamic`,`acl_rule`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAccessControlListQuery = "update `access_control_list` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`access_control_list_hash` = ?,`dynamic` = ?,`acl_rule` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteAccessControlListQuery = "delete from `access_control_list` where uuid = ?"
const listAccessControlListQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`access_control_list_hash`,`dynamic`,`acl_rule`,`uuid`,`fq_name` from `access_control_list`"
const showAccessControlListQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`access_control_list_hash`,`dynamic`,`acl_rule`,`uuid`,`fq_name` from `access_control_list` where uuid = ?"

func CreateAccessControlList(tx *sql.Tx, model *models.AccessControlList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAccessControlListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Description),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		util.MustJSON(model.AccessControlListHash),
		bool(model.AccessControlListEntries.Dynamic),
		util.MustJSON(model.AccessControlListEntries.ACLRule),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanAccessControlList(rows *sql.Rows) (*models.AccessControlList, error) {
	m := models.MakeAccessControlList()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonAccessControlListHash string

	var jsonAccessControlListEntriesACLRule string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Description,
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
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonAccessControlListHash,
		&m.AccessControlListEntries.Dynamic,
		&jsonAccessControlListEntriesACLRule,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAccessControlListHash), &m.AccessControlListHash)

	json.Unmarshal([]byte(jsonAccessControlListEntriesACLRule), &m.AccessControlListEntries.ACLRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListAccessControlList(tx *sql.Tx) ([]*models.AccessControlList, error) {
	result := models.MakeAccessControlListSlice()
	rows, err := tx.Query(listAccessControlListQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAccessControlList(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAccessControlList(tx *sql.Tx, uuid string) (*models.AccessControlList, error) {
	rows, err := tx.Query(showAccessControlListQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAccessControlList(rows)
	}
	return nil, nil
}

func UpdateAccessControlList(tx *sql.Tx, uuid string, model *models.AccessControlList) error {
	return nil
}

func DeleteAccessControlList(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAccessControlListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
