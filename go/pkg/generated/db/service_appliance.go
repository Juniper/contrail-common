package db

// service_appliance

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceApplianceQuery = "insert into `service_appliance` (`display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`username`,`password`,`service_appliance_ip_address`,`service_appliance_properties_key_value_pair`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceQuery = "update `service_appliance` set `display_name` = ?,`key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`fq_name` = ?,`username` = ?,`password` = ?,`service_appliance_ip_address` = ?,`service_appliance_properties_key_value_pair` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`uuid` = ?;"
const deleteServiceApplianceQuery = "delete from `service_appliance` where uuid = ?"
const listServiceApplianceQuery = "select `display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`username`,`password`,`service_appliance_ip_address`,`service_appliance_properties_key_value_pair`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`uuid` from `service_appliance`"
const showServiceApplianceQuery = "select `display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`username`,`password`,`service_appliance_ip_address`,`service_appliance_properties_key_value_pair`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`uuid` from `service_appliance` where uuid = ?"

func CreateServiceAppliance(tx *sql.Tx, model *models.ServiceAppliance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		util.MustJSON(model.FQName),
		string(model.ServiceApplianceUserCredentials.Username),
		string(model.ServiceApplianceUserCredentials.Password),
		string(model.ServiceApplianceIPAddress),
		util.MustJSON(model.ServiceApplianceProperties.KeyValuePair),
		string(model.IDPerms.Creator),
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
		string(model.UUID))
	return err
}

func scanServiceAppliance(rows *sql.Rows) (*models.ServiceAppliance, error) {
	m := models.MakeServiceAppliance()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonServiceAppliancePropertiesKeyValuePair string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonFQName,
		&m.ServiceApplianceUserCredentials.Username,
		&m.ServiceApplianceUserCredentials.Password,
		&m.ServiceApplianceIPAddress,
		&jsonServiceAppliancePropertiesKeyValuePair,
		&m.IDPerms.Creator,
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
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonServiceAppliancePropertiesKeyValuePair), &m.ServiceApplianceProperties.KeyValuePair)

	return m, nil
}

func ListServiceAppliance(tx *sql.Tx) ([]*models.ServiceAppliance, error) {
	result := models.MakeServiceApplianceSlice()
	rows, err := tx.Query(listServiceApplianceQuery)
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
