package db

// logical_router

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLogicalRouterQuery = "insert into `logical_router` (`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`vxlan_network_identifier`,`route_target`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalRouterQuery = "update `logical_router` set `fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`vxlan_network_identifier` = ?,`route_target` = ?,`uuid` = ?;"
const deleteLogicalRouterQuery = "delete from `logical_router` where uuid = ?"
const listLogicalRouterQuery = "select `fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`vxlan_network_identifier`,`route_target`,`uuid` from `logical_router`"
const showLogicalRouterQuery = "select `fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`vxlan_network_identifier`,`route_target`,`uuid` from `logical_router` where uuid = ?"

func CreateLogicalRouter(tx *sql.Tx, model *models.LogicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.VxlanNetworkIdentifier),
		util.MustJSON(model.ConfiguredRouteTargetList.RouteTarget),
		string(model.UUID))
	return err
}

func scanLogicalRouter(rows *sql.Rows) (*models.LogicalRouter, error) {
	m := models.MakeLogicalRouter()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonConfiguredRouteTargetListRouteTarget string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.VxlanNetworkIdentifier,
		&jsonConfiguredRouteTargetListRouteTarget,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonConfiguredRouteTargetListRouteTarget), &m.ConfiguredRouteTargetList.RouteTarget)

	return m, nil
}

func ListLogicalRouter(tx *sql.Tx) ([]*models.LogicalRouter, error) {
	result := models.MakeLogicalRouterSlice()
	rows, err := tx.Query(listLogicalRouterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanLogicalRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLogicalRouter(tx *sql.Tx, uuid string) (*models.LogicalRouter, error) {
	rows, err := tx.Query(showLogicalRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLogicalRouter(rows)
	}
	return nil, nil
}

func UpdateLogicalRouter(tx *sql.Tx, uuid string, model *models.LogicalRouter) error {
	return nil
}

func DeleteLogicalRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLogicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
