package db

// floating_ip

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertFloatingIPQuery = "insert into `floating_ip` (`floating_ip_port_mappings_enable`,`floating_ip_fixed_ip_address`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`display_name`,`floating_ip_traffic_direction`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPQuery = "update `floating_ip` set `floating_ip_port_mappings_enable` = ?,`floating_ip_fixed_ip_address` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`floating_ip_address_family` = ?,`floating_ip_port_mappings` = ?,`floating_ip_is_virtual_ip` = ?,`floating_ip_address` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`floating_ip_traffic_direction` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteFloatingIPQuery = "delete from `floating_ip` where uuid = ?"
const listFloatingIPQuery = "select `floating_ip_port_mappings_enable`,`floating_ip_fixed_ip_address`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`display_name`,`floating_ip_traffic_direction`,`uuid`,`fq_name` from `floating_ip`"
const showFloatingIPQuery = "select `floating_ip_port_mappings_enable`,`floating_ip_fixed_ip_address`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`display_name`,`floating_ip_traffic_direction`,`uuid`,`fq_name` from `floating_ip` where uuid = ?"

func CreateFloatingIP(tx *sql.Tx, model *models.FloatingIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.FloatingIPPortMappingsEnable),
		string(model.FloatingIPFixedIPAddress),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.FloatingIPAddressFamily),
		util.MustJSON(model.FloatingIPPortMappings),
		bool(model.FloatingIPIsVirtualIP),
		string(model.FloatingIPAddress),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		string(model.FloatingIPTrafficDirection),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanFloatingIP(rows *sql.Rows) (*models.FloatingIP, error) {
	m := models.MakeFloatingIP()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFloatingIPPortMappings string

	var jsonFQName string

	if err := rows.Scan(&m.FloatingIPPortMappingsEnable,
		&m.FloatingIPFixedIPAddress,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.FloatingIPAddressFamily,
		&jsonFloatingIPPortMappings,
		&m.FloatingIPIsVirtualIP,
		&m.FloatingIPAddress,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&m.FloatingIPTrafficDirection,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFloatingIPPortMappings), &m.FloatingIPPortMappings)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListFloatingIP(tx *sql.Tx) ([]*models.FloatingIP, error) {
	result := models.MakeFloatingIPSlice()
	rows, err := tx.Query(listFloatingIPQuery)
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
