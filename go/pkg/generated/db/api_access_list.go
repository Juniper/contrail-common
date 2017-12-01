package db

// api_access_list

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAPIAccessListQuery = "insert into `api_access_list` (`global_access`,`share`,`owner`,`owner_access`,`rbac_rule`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAPIAccessListQuery = "update `api_access_list` set `global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`rbac_rule` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteAPIAccessListQuery = "delete from `api_access_list` where uuid = ?"
const listAPIAccessListQuery = "select `global_access`,`share`,`owner`,`owner_access`,`rbac_rule`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`key_value_pair` from `api_access_list`"
const showAPIAccessListQuery = "select `global_access`,`share`,`owner`,`owner_access`,`rbac_rule`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`key_value_pair` from `api_access_list` where uuid = ?"

func CreateAPIAccessList(tx *sql.Tx, model *models.APIAccessList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAPIAccessListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		util.MustJSON(model.APIAccessListEntries.RbacRule),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanAPIAccessList(rows *sql.Rows) (*models.APIAccessList, error) {
	m := models.MakeAPIAccessList()

	var jsonPerms2Share string

	var jsonAPIAccessListEntriesRbacRule string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonAPIAccessListEntriesRbacRule,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAPIAccessListEntriesRbacRule), &m.APIAccessListEntries.RbacRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListAPIAccessList(tx *sql.Tx) ([]*models.APIAccessList, error) {
	result := models.MakeAPIAccessListSlice()
	rows, err := tx.Query(listAPIAccessListQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAPIAccessList(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAPIAccessList(tx *sql.Tx, uuid string) (*models.APIAccessList, error) {
	rows, err := tx.Query(showAPIAccessListQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAPIAccessList(rows)
	}
	return nil, nil
}

func UpdateAPIAccessList(tx *sql.Tx, uuid string, model *models.APIAccessList) error {
	return nil
}

func DeleteAPIAccessList(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAPIAccessListQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
