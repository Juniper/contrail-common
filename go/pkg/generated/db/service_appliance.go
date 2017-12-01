package db

// service_appliance

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertServiceApplianceQuery = "insert into `service_appliance` (`uuid`,`global_access`,`share`,`owner`,`owner_access`,`password`,`username`,`service_appliance_ip_address`,`key_value_pair`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceQuery = "update `service_appliance` set `uuid` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`password` = ?,`username` = ?,`service_appliance_ip_address` = ?,`key_value_pair` = ?,`fq_name` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`annotations_key_value_pair` = ?;"
const deleteServiceApplianceQuery = "delete from `service_appliance` where uuid = ?"
const listServiceApplianceQuery = "select `uuid`,`global_access`,`share`,`owner`,`owner_access`,`password`,`username`,`service_appliance_ip_address`,`key_value_pair`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`annotations_key_value_pair` from `service_appliance`"
const showServiceApplianceQuery = "select `uuid`,`global_access`,`share`,`owner`,`owner_access`,`password`,`username`,`service_appliance_ip_address`,`key_value_pair`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`annotations_key_value_pair` from `service_appliance` where uuid = ?"

func CreateServiceAppliance(tx *sql.Tx, model *models.ServiceAppliance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.ServiceApplianceUserCredentials.Password),
		string(model.ServiceApplianceUserCredentials.Username),
		string(model.ServiceApplianceIPAddress),
		util.MustJSON(model.ServiceApplianceProperties.KeyValuePair),
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
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanServiceAppliance(rows *sql.Rows) (*models.ServiceAppliance, error) {
	m := models.MakeServiceAppliance()

	var jsonPerms2Share string

	var jsonServiceAppliancePropertiesKeyValuePair string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.UUID,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.ServiceApplianceUserCredentials.Password,
		&m.ServiceApplianceUserCredentials.Username,
		&m.ServiceApplianceIPAddress,
		&jsonServiceAppliancePropertiesKeyValuePair,
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
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonServiceAppliancePropertiesKeyValuePair), &m.ServiceApplianceProperties.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createServiceApplianceWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["password"]; ok {
		results = append(results, "password = ?")
		values = append(values, value)
	}

	if value, ok := where["username"]; ok {
		results = append(results, "username = ?")
		values = append(values, value)
	}

	if value, ok := where["service_appliance_ip_address"]; ok {
		results = append(results, "service_appliance_ip_address = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceAppliance(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceAppliance, error) {
	result := models.MakeServiceApplianceSlice()
	whereQuery, values := createServiceApplianceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceApplianceQuery)
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
		m, _ := scanServiceAppliance(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceAppliance(tx *sql.Tx, uuid string) (*models.ServiceAppliance, error) {
	rows, err := tx.Query(showServiceApplianceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceAppliance(rows)
	}
	return nil, nil
}

func UpdateServiceAppliance(tx *sql.Tx, uuid string, model *models.ServiceAppliance) error {
	return nil
}

func DeleteServiceAppliance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceApplianceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
