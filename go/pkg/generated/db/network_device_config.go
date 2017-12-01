package db

// network_device_config

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertNetworkDeviceConfigQuery = "insert into `network_device_config` (`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkDeviceConfigQuery = "update `network_device_config` set `fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?;"
const deleteNetworkDeviceConfigQuery = "delete from `network_device_config` where uuid = ?"
const listNetworkDeviceConfigQuery = "select `fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid` from `network_device_config`"
const showNetworkDeviceConfigQuery = "select `fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid` from `network_device_config` where uuid = ?"

func CreateNetworkDeviceConfig(tx *sql.Tx, model *models.NetworkDeviceConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkDeviceConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
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
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID))
	return err
}

func scanNetworkDeviceConfig(rows *sql.Rows) (*models.NetworkDeviceConfig, error) {
	m := models.MakeNetworkDeviceConfig()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
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
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListNetworkDeviceConfig(tx *sql.Tx) ([]*models.NetworkDeviceConfig, error) {
	result := models.MakeNetworkDeviceConfigSlice()
	rows, err := tx.Query(listNetworkDeviceConfigQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanNetworkDeviceConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNetworkDeviceConfig(tx *sql.Tx, uuid string) (*models.NetworkDeviceConfig, error) {
	rows, err := tx.Query(showNetworkDeviceConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNetworkDeviceConfig(rows)
	}
	return nil, nil
}

func UpdateNetworkDeviceConfig(tx *sql.Tx, uuid string, model *models.NetworkDeviceConfig) error {
	return nil
}

func DeleteNetworkDeviceConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkDeviceConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
