package db

// virtual_router

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertVirtualRouterQuery = "insert into `virtual_router` (`uuid`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`virtual_router_dpdk_enabled`,`virtual_router_type`,`virtual_router_ip_address`,`fq_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualRouterQuery = "update `virtual_router` set `uuid` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`virtual_router_dpdk_enabled` = ?,`virtual_router_type` = ?,`virtual_router_ip_address` = ?,`fq_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteVirtualRouterQuery = "delete from `virtual_router` where uuid = ?"
const listVirtualRouterQuery = "select `uuid`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`virtual_router_dpdk_enabled`,`virtual_router_type`,`virtual_router_ip_address`,`fq_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `virtual_router`"
const showVirtualRouterQuery = "select `uuid`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`virtual_router_dpdk_enabled`,`virtual_router_type`,`virtual_router_ip_address`,`fq_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `virtual_router` where uuid = ?"

func CreateVirtualRouter(tx *sql.Tx, model *models.VirtualRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
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
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		bool(model.VirtualRouterDPDKEnabled),
		string(model.VirtualRouterType),
		string(model.VirtualRouterIPAddress),
		util.MustJSON(model.FQName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	return err
}

func scanVirtualRouter(rows *sql.Rows) (*models.VirtualRouter, error) {
	m := models.MakeVirtualRouter()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
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
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&m.VirtualRouterDPDKEnabled,
		&m.VirtualRouterType,
		&m.VirtualRouterIPAddress,
		&jsonFQName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createVirtualRouterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_router_type"]; ok {
		results = append(results, "virtual_router_type = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_router_ip_address"]; ok {
		results = append(results, "virtual_router_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualRouter(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualRouter, error) {
	result := models.MakeVirtualRouterSlice()
	whereQuery, values := createVirtualRouterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualRouterQuery)
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
