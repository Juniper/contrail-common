package db

// service_appliance_set

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`service_appliance_ha_mode`,`service_appliance_driver`,`fq_name`,`service_appliance_set_properties_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceSetQuery = "update `service_appliance_set` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`service_appliance_ha_mode` = ?,`service_appliance_driver` = ?,`fq_name` = ?,`service_appliance_set_properties_key_value_pair` = ?;"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set` where uuid = ?"
const listServiceApplianceSetQuery = "select `key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`service_appliance_ha_mode`,`service_appliance_driver`,`fq_name`,`service_appliance_set_properties_key_value_pair` from `service_appliance_set`"
const showServiceApplianceSetQuery = "select `key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`service_appliance_ha_mode`,`service_appliance_driver`,`fq_name`,`service_appliance_set_properties_key_value_pair` from `service_appliance_set` where uuid = ?"

func CreateServiceApplianceSet(tx *sql.Tx, model *models.ServiceApplianceSet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		string(model.ServiceApplianceHaMode),
		string(model.ServiceApplianceDriver),
		util.MustJSON(model.FQName),
		util.MustJSON(model.ServiceApplianceSetProperties.KeyValuePair))
	return err
}

func scanServiceApplianceSet(rows *sql.Rows) (*models.ServiceApplianceSet, error) {
	m := models.MakeServiceApplianceSet()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonServiceApplianceSetPropertiesKeyValuePair string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName,
		&m.ServiceApplianceHaMode,
		&m.ServiceApplianceDriver,
		&jsonFQName,
		&jsonServiceApplianceSetPropertiesKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonServiceApplianceSetPropertiesKeyValuePair), &m.ServiceApplianceSetProperties.KeyValuePair)

	return m, nil
}

func createServiceApplianceSetWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["service_appliance_ha_mode"]; ok {
		results = append(results, "service_appliance_ha_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["service_appliance_driver"]; ok {
		results = append(results, "service_appliance_driver = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceApplianceSet(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceApplianceSet, error) {
	result := models.MakeServiceApplianceSetSlice()
	whereQuery, values := createServiceApplianceSetWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceApplianceSetQuery)
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
