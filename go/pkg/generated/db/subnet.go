package db

// subnet

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertSubnetQuery = "insert into `subnet` (`fq_name`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`ip_prefix`,`ip_prefix_len`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSubnetQuery = "update `subnet` set `fq_name` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?;"
const deleteSubnetQuery = "delete from `subnet` where uuid = ?"
const listSubnetQuery = "select `fq_name`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`ip_prefix`,`ip_prefix_len`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid` from `subnet`"
const showSubnetQuery = "select `fq_name`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`ip_prefix`,`ip_prefix_len`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid` from `subnet` where uuid = ?"

func CreateSubnet(tx *sql.Tx, model *models.Subnet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSubnetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
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
		string(model.IDPerms.LastModified),
		string(model.SubnetIPPrefix.IPPrefix),
		int(model.SubnetIPPrefix.IPPrefixLen),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID))
	return err
}

func scanSubnet(rows *sql.Rows) (*models.Subnet, error) {
	m := models.MakeSubnet()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
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
		&m.IDPerms.LastModified,
		&m.SubnetIPPrefix.IPPrefix,
		&m.SubnetIPPrefix.IPPrefixLen,
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

func ListSubnet(tx *sql.Tx) ([]*models.Subnet, error) {
	result := models.MakeSubnetSlice()
	rows, err := tx.Query(listSubnetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanSubnet(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowSubnet(tx *sql.Tx, uuid string) (*models.Subnet, error) {
	rows, err := tx.Query(showSubnetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanSubnet(rows)
	}
	return nil, nil
}

func UpdateSubnet(tx *sql.Tx, uuid string, model *models.Subnet) error {
	return nil
}

func DeleteSubnet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSubnetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
