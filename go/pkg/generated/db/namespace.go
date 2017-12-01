package db

// namespace

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertNamespaceQuery = "insert into `namespace` (`owner_access`,`global_access`,`share`,`owner`,`uuid`,`ip_prefix`,`ip_prefix_len`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNamespaceQuery = "update `namespace` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteNamespaceQuery = "delete from `namespace` where uuid = ?"
const listNamespaceQuery = "select `owner_access`,`global_access`,`share`,`owner`,`uuid`,`ip_prefix`,`ip_prefix_len`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`key_value_pair` from `namespace`"
const showNamespaceQuery = "select `owner_access`,`global_access`,`share`,`owner`,`uuid`,`ip_prefix`,`ip_prefix_len`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`key_value_pair` from `namespace` where uuid = ?"

func CreateNamespace(tx *sql.Tx, model *models.Namespace) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNamespaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		string(model.NamespaceCidr.IPPrefix),
		int(model.NamespaceCidr.IPPrefixLen),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanNamespace(rows *sql.Rows) (*models.Namespace, error) {
	m := models.MakeNamespace()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&m.NamespaceCidr.IPPrefix,
		&m.NamespaceCidr.IPPrefixLen,
		&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListNamespace(tx *sql.Tx) ([]*models.Namespace, error) {
	result := models.MakeNamespaceSlice()
	rows, err := tx.Query(listNamespaceQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanNamespace(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNamespace(tx *sql.Tx, uuid string) (*models.Namespace, error) {
	rows, err := tx.Query(showNamespaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNamespace(rows)
	}
	return nil, nil
}

func UpdateNamespace(tx *sql.Tx, uuid string, model *models.Namespace) error {
	return nil
}

func DeleteNamespace(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNamespaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
