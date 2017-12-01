package db

// interface_route_table

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertInterfaceRouteTableQuery = "insert into `interface_route_table` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`route`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInterfaceRouteTableQuery = "update `interface_route_table` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`route` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteInterfaceRouteTableQuery = "delete from `interface_route_table` where uuid = ?"
const listInterfaceRouteTableQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`route`,`display_name`,`key_value_pair` from `interface_route_table`"
const showInterfaceRouteTableQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`route`,`display_name`,`key_value_pair` from `interface_route_table` where uuid = ?"

func CreateInterfaceRouteTable(tx *sql.Tx, model *models.InterfaceRouteTable) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInterfaceRouteTableQuery)
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
		string(model.IDPerms.Description),
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
		util.MustJSON(model.InterfaceRouteTableRoutes.Route),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanInterfaceRouteTable(rows *sql.Rows) (*models.InterfaceRouteTable, error) {
	m := models.MakeInterfaceRouteTable()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonInterfaceRouteTableRoutesRoute string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Description,
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
		&jsonInterfaceRouteTableRoutesRoute,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonInterfaceRouteTableRoutesRoute), &m.InterfaceRouteTableRoutes.Route)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListInterfaceRouteTable(tx *sql.Tx) ([]*models.InterfaceRouteTable, error) {
	result := models.MakeInterfaceRouteTableSlice()
	rows, err := tx.Query(listInterfaceRouteTableQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanInterfaceRouteTable(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowInterfaceRouteTable(tx *sql.Tx, uuid string) (*models.InterfaceRouteTable, error) {
	rows, err := tx.Query(showInterfaceRouteTableQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanInterfaceRouteTable(rows)
	}
	return nil, nil
}

func UpdateInterfaceRouteTable(tx *sql.Tx, uuid string, model *models.InterfaceRouteTable) error {
	return nil
}

func DeleteInterfaceRouteTable(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteInterfaceRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
