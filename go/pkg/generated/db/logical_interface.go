package db

// logical_interface

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLogicalInterfaceQuery = "insert into `logical_interface` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`logical_interface_vlan_tag`,`logical_interface_type`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalInterfaceQuery = "update `logical_interface` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`logical_interface_vlan_tag` = ?,`logical_interface_type` = ?,`uuid` = ?,`fq_name` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?;"
const deleteLogicalInterfaceQuery = "delete from `logical_interface` where uuid = ?"
const listLogicalInterfaceQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`logical_interface_vlan_tag`,`logical_interface_type`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `logical_interface`"
const showLogicalInterfaceQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`logical_interface_vlan_tag`,`logical_interface_type`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `logical_interface` where uuid = ?"

func CreateLogicalInterface(tx *sql.Tx, model *models.LogicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		int(model.LogicalInterfaceVlanTag),
		string(model.LogicalInterfaceType),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified))
	return err
}

func scanLogicalInterface(rows *sql.Rows) (*models.LogicalInterface, error) {
	m := models.MakeLogicalInterface()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.LogicalInterfaceVlanTag,
		&m.LogicalInterfaceType,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
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

func ListLogicalInterface(tx *sql.Tx) ([]*models.LogicalInterface, error) {
	result := models.MakeLogicalInterfaceSlice()
	rows, err := tx.Query(listLogicalInterfaceQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanLogicalInterface(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLogicalInterface(tx *sql.Tx, uuid string) (*models.LogicalInterface, error) {
	rows, err := tx.Query(showLogicalInterfaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLogicalInterface(rows)
	}
	return nil, nil
}

func UpdateLogicalInterface(tx *sql.Tx, uuid string, model *models.LogicalInterface) error {
	return nil
}

func DeleteLogicalInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLogicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
