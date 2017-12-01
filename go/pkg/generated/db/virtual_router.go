package db

// virtual_router

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVirtualRouterQuery = "insert into `virtual_router` (`virtual_router_dpdk_enabled`,`key_value_pair`,`fq_name`,`virtual_router_type`,`virtual_router_ip_address`,`display_name`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualRouterQuery = "update `virtual_router` set `virtual_router_dpdk_enabled` = ?,`key_value_pair` = ?,`fq_name` = ?,`virtual_router_type` = ?,`virtual_router_ip_address` = ?,`display_name` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?;"
const deleteVirtualRouterQuery = "delete from `virtual_router` where uuid = ?"
const listVirtualRouterQuery = "select `virtual_router_dpdk_enabled`,`key_value_pair`,`fq_name`,`virtual_router_type`,`virtual_router_ip_address`,`display_name`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `virtual_router`"
const showVirtualRouterQuery = "select `virtual_router_dpdk_enabled`,`key_value_pair`,`fq_name`,`virtual_router_type`,`virtual_router_ip_address`,`display_name`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `virtual_router` where uuid = ?"

func CreateVirtualRouter(tx *sql.Tx, model *models.VirtualRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.VirtualRouterDPDKEnabled),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.FQName),
		string(model.VirtualRouterType),
		string(model.VirtualRouterIPAddress),
		string(model.DisplayName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified))
	return err
}

func scanVirtualRouter(rows *sql.Rows) (*models.VirtualRouter, error) {
	m := models.MakeVirtualRouter()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&m.VirtualRouterDPDKEnabled,
		&jsonAnnotationsKeyValuePair,
		&jsonFQName,
		&m.VirtualRouterType,
		&m.VirtualRouterIPAddress,
		&m.DisplayName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListVirtualRouter(tx *sql.Tx) ([]*models.VirtualRouter, error) {
	result := models.MakeVirtualRouterSlice()
	rows, err := tx.Query(listVirtualRouterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualRouter(tx *sql.Tx, uuid string) (*models.VirtualRouter, error) {
	rows, err := tx.Query(showVirtualRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualRouter(rows)
	}
	return nil, nil
}

func UpdateVirtualRouter(tx *sql.Tx, uuid string, model *models.VirtualRouter) error {
	return nil
}

func DeleteVirtualRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
