package db

// route_aggregate

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertRouteAggregateQuery = "insert into `route_aggregate` (`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteAggregateQuery = "update `route_aggregate` set `share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteRouteAggregateQuery = "delete from `route_aggregate` where uuid = ?"
const listRouteAggregateQuery = "select `share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair` from `route_aggregate`"
const showRouteAggregateQuery = "select `share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair` from `route_aggregate` where uuid = ?"

func CreateRouteAggregate(tx *sql.Tx, model *models.RouteAggregate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteAggregateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanRouteAggregate(rows *sql.Rows) (*models.RouteAggregate, error) {
	m := models.MakeRouteAggregate()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListRouteAggregate(tx *sql.Tx) ([]*models.RouteAggregate, error) {
	result := models.MakeRouteAggregateSlice()
	rows, err := tx.Query(listRouteAggregateQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanRouteAggregate(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowRouteAggregate(tx *sql.Tx, uuid string) (*models.RouteAggregate, error) {
	rows, err := tx.Query(showRouteAggregateQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanRouteAggregate(rows)
	}
	return nil, nil
}

func UpdateRouteAggregate(tx *sql.Tx, uuid string, model *models.RouteAggregate) error {
	return nil
}

func DeleteRouteAggregate(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteRouteAggregateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
