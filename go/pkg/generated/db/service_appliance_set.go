package db

// service_appliance_set

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`service_appliance_set_properties_key_value_pair`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`service_appliance_ha_mode`,`service_appliance_driver`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceSetQuery = "update `service_appliance_set` set `key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`service_appliance_set_properties_key_value_pair` = ?,`fq_name` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`service_appliance_ha_mode` = ?,`service_appliance_driver` = ?;"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set` where uuid = ?"
const listServiceApplianceSetQuery = "select `key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`service_appliance_set_properties_key_value_pair`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`service_appliance_ha_mode`,`service_appliance_driver` from `service_appliance_set`"
const showServiceApplianceSetQuery = "select `key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`service_appliance_set_properties_key_value_pair`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`service_appliance_ha_mode`,`service_appliance_driver` from `service_appliance_set` where uuid = ?"

func CreateServiceApplianceSet(tx *sql.Tx, model *models.ServiceApplianceSet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.ServiceApplianceSetProperties.KeyValuePair),
		util.MustJSON(model.FQName),
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
		string(model.ServiceApplianceHaMode),
		string(model.ServiceApplianceDriver))
	return err
}

func scanServiceApplianceSet(rows *sql.Rows) (*models.ServiceApplianceSet, error) {
	m := models.MakeServiceApplianceSet()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonServiceApplianceSetPropertiesKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonServiceApplianceSetPropertiesKeyValuePair,
		&jsonFQName,
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
		&m.ServiceApplianceHaMode,
		&m.ServiceApplianceDriver); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonServiceApplianceSetPropertiesKeyValuePair), &m.ServiceApplianceSetProperties.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListServiceApplianceSet(tx *sql.Tx) ([]*models.ServiceApplianceSet, error) {
	result := models.MakeServiceApplianceSetSlice()
	rows, err := tx.Query(listServiceApplianceSetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceApplianceSet(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceApplianceSet(tx *sql.Tx, uuid string) (*models.ServiceApplianceSet, error) {
	rows, err := tx.Query(showServiceApplianceSetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceApplianceSet(rows)
	}
	return nil, nil
}

func UpdateServiceApplianceSet(tx *sql.Tx, uuid string, model *models.ServiceApplianceSet) error {
	return nil
}

func DeleteServiceApplianceSet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
