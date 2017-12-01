package db

// floating_ip

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertFloatingIPQuery = "insert into `floating_ip` (`owner_access`,`global_access`,`share`,`owner`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`floating_ip_fixed_ip_address`,`display_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`uuid`,`fq_name`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_port_mappings_enable`,`floating_ip_traffic_direction`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPQuery = "update `floating_ip` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`floating_ip_is_virtual_ip` = ?,`floating_ip_address` = ?,`floating_ip_fixed_ip_address` = ?,`display_name` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`floating_ip_address_family` = ?,`floating_ip_port_mappings` = ?,`floating_ip_port_mappings_enable` = ?,`floating_ip_traffic_direction` = ?;"
const deleteFloatingIPQuery = "delete from `floating_ip` where uuid = ?"
const listFloatingIPQuery = "select `owner_access`,`global_access`,`share`,`owner`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`floating_ip_fixed_ip_address`,`display_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`uuid`,`fq_name`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_port_mappings_enable`,`floating_ip_traffic_direction` from `floating_ip`"
const showFloatingIPQuery = "select `owner_access`,`global_access`,`share`,`owner`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`floating_ip_fixed_ip_address`,`display_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`uuid`,`fq_name`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_port_mappings_enable`,`floating_ip_traffic_direction` from `floating_ip` where uuid = ?"

func CreateFloatingIP(tx *sql.Tx, model *models.FloatingIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		bool(model.FloatingIPIsVirtualIP),
		string(model.FloatingIPAddress),
		string(model.FloatingIPFixedIPAddress),
		string(model.DisplayName),
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
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.FloatingIPAddressFamily),
		util.MustJSON(model.FloatingIPPortMappings),
		bool(model.FloatingIPPortMappingsEnable),
		string(model.FloatingIPTrafficDirection))
	return err
}

func scanFloatingIP(rows *sql.Rows) (*models.FloatingIP, error) {
	m := models.MakeFloatingIP()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonFloatingIPPortMappings string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.FloatingIPIsVirtualIP,
		&m.FloatingIPAddress,
		&m.FloatingIPFixedIPAddress,
		&m.DisplayName,
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
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName,
		&m.FloatingIPAddressFamily,
		&jsonFloatingIPPortMappings,
		&m.FloatingIPPortMappingsEnable,
		&m.FloatingIPTrafficDirection); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonFloatingIPPortMappings), &m.FloatingIPPortMappings)

	return m, nil
}

func createFloatingIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_address"]; ok {
		results = append(results, "floating_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_fixed_ip_address"]; ok {
		results = append(results, "floating_ip_fixed_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_address_family"]; ok {
		results = append(results, "floating_ip_address_family = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_traffic_direction"]; ok {
		results = append(results, "floating_ip_traffic_direction = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListFloatingIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.FloatingIP, error) {
	result := models.MakeFloatingIPSlice()
	whereQuery, values := createFloatingIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listFloatingIPQuery)
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
		m, _ := scanFloatingIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFloatingIP(tx *sql.Tx, uuid string) (*models.FloatingIP, error) {
	rows, err := tx.Query(showFloatingIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFloatingIP(rows)
	}
	return nil, nil
}

func UpdateFloatingIP(tx *sql.Tx, uuid string, model *models.FloatingIP) error {
	return nil
}

func DeleteFloatingIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFloatingIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
