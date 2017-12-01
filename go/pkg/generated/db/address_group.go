package db

// address_group

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAddressGroupQuery = "insert into `address_group` (`key_value_pair`,`address_group_prefix`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAddressGroupQuery = "update `address_group` set `key_value_pair` = ?,`address_group_prefix` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?;"
const deleteAddressGroupQuery = "delete from `address_group` where uuid = ?"
const listAddressGroupQuery = "select `key_value_pair`,`address_group_prefix`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name` from `address_group`"
const showAddressGroupQuery = "select `key_value_pair`,`address_group_prefix`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name` from `address_group` where uuid = ?"

func CreateAddressGroup(tx *sql.Tx, model *models.AddressGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAddressGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.AddressGroupPrefix),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
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
		string(model.DisplayName))
	return err
}

func scanAddressGroup(rows *sql.Rows) (*models.AddressGroup, error) {
	m := models.MakeAddressGroup()

	var jsonAnnotationsKeyValuePair string

	var jsonAddressGroupPrefix string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&jsonAddressGroupPrefix,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
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
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonAddressGroupPrefix), &m.AddressGroupPrefix)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListAddressGroup(tx *sql.Tx) ([]*models.AddressGroup, error) {
	result := models.MakeAddressGroupSlice()
	rows, err := tx.Query(listAddressGroupQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAddressGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAddressGroup(tx *sql.Tx, uuid string) (*models.AddressGroup, error) {
	rows, err := tx.Query(showAddressGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAddressGroup(rows)
	}
	return nil, nil
}

func UpdateAddressGroup(tx *sql.Tx, uuid string, model *models.AddressGroup) error {
	return nil
}

func DeleteAddressGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAddressGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
