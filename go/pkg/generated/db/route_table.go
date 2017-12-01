package db

// route_table

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertRouteTableQuery = "insert into `route_table` (`owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteTableQuery = "update `route_table` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`route` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteRouteTableQuery = "delete from `route_table` where uuid = ?"
const listRouteTableQuery = "select `owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair` from `route_table`"
const showRouteTableQuery = "select `owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair` from `route_table` where uuid = ?"

func CreateRouteTable(tx *sql.Tx, model *models.RouteTable) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		util.MustJSON(model.Routes.Route),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanRouteTable(rows *sql.Rows) (*models.RouteTable, error) {
	m := models.MakeRouteTable()

	var jsonPerms2Share string

	var jsonRoutesRoute string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonRoutesRoute,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonRoutesRoute), &m.Routes.Route)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createRouteTableWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListRouteTable(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.RouteTable, error) {
	result := models.MakeRouteTableSlice()
	whereQuery, values := createRouteTableWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listRouteTableQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanRouteTable(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRouteTable(tx *sql.Tx, uuid string) (*models.RouteTable, error) {
	rows, err := tx.Query(showRouteTableQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRouteTable(rows)
	}
	return nil, nil
}

func UpdateRouteTable(tx *sql.Tx, uuid string, model *models.RouteTable) error {
	return nil
}

func DeleteRouteTable(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
