package db

// service_template

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceTemplateQuery = "insert into `service_template` (`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`service_scaling`,`service_virtualization_type`,`vrouter_instance_type`,`service_mode`,`service_type`,`availability_zone_enable`,`image_name`,`version`,`flavor`,`instance_data`,`interface_type`,`ordered_interfaces`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceTemplateQuery = "update `service_template` set `key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`service_scaling` = ?,`service_virtualization_type` = ?,`vrouter_instance_type` = ?,`service_mode` = ?,`service_type` = ?,`availability_zone_enable` = ?,`image_name` = ?,`version` = ?,`flavor` = ?,`instance_data` = ?,`interface_type` = ?,`ordered_interfaces` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteServiceTemplateQuery = "delete from `service_template` where uuid = ?"
const listServiceTemplateQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`service_scaling`,`service_virtualization_type`,`vrouter_instance_type`,`service_mode`,`service_type`,`availability_zone_enable`,`image_name`,`version`,`flavor`,`instance_data`,`interface_type`,`ordered_interfaces`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name` from `service_template`"
const showServiceTemplateQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`service_scaling`,`service_virtualization_type`,`vrouter_instance_type`,`service_mode`,`service_type`,`availability_zone_enable`,`image_name`,`version`,`flavor`,`instance_data`,`interface_type`,`ordered_interfaces`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name` from `service_template` where uuid = ?"

func CreateServiceTemplate(tx *sql.Tx, model *models.ServiceTemplate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceTemplateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		bool(model.ServiceTemplateProperties.ServiceScaling),
		string(model.ServiceTemplateProperties.ServiceVirtualizationType),
		string(model.ServiceTemplateProperties.VrouterInstanceType),
		string(model.ServiceTemplateProperties.ServiceMode),
		string(model.ServiceTemplateProperties.ServiceType),
		bool(model.ServiceTemplateProperties.AvailabilityZoneEnable),
		string(model.ServiceTemplateProperties.ImageName),
		int(model.ServiceTemplateProperties.Version),
		string(model.ServiceTemplateProperties.Flavor),
		string(model.ServiceTemplateProperties.InstanceData),
		util.MustJSON(model.ServiceTemplateProperties.InterfaceType),
		bool(model.ServiceTemplateProperties.OrderedInterfaces),
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
		string(model.DisplayName))
	return err
}

func scanServiceTemplate(rows *sql.Rows) (*models.ServiceTemplate, error) {
	m := models.MakeServiceTemplate()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonServiceTemplatePropertiesInterfaceType string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
		&m.ServiceTemplateProperties.ServiceScaling,
		&m.ServiceTemplateProperties.ServiceVirtualizationType,
		&m.ServiceTemplateProperties.VrouterInstanceType,
		&m.ServiceTemplateProperties.ServiceMode,
		&m.ServiceTemplateProperties.ServiceType,
		&m.ServiceTemplateProperties.AvailabilityZoneEnable,
		&m.ServiceTemplateProperties.ImageName,
		&m.ServiceTemplateProperties.Version,
		&m.ServiceTemplateProperties.Flavor,
		&m.ServiceTemplateProperties.InstanceData,
		&jsonServiceTemplatePropertiesInterfaceType,
		&m.ServiceTemplateProperties.OrderedInterfaces,
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
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonServiceTemplatePropertiesInterfaceType), &m.ServiceTemplateProperties.InterfaceType)

	return m, nil
}

func ListServiceTemplate(tx *sql.Tx) ([]*models.ServiceTemplate, error) {
	result := models.MakeServiceTemplateSlice()
	rows, err := tx.Query(listServiceTemplateQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceTemplate(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceTemplate(tx *sql.Tx, uuid string) (*models.ServiceTemplate, error) {
	rows, err := tx.Query(showServiceTemplateQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceTemplate(rows)
	}
	return nil, nil
}

func UpdateServiceTemplate(tx *sql.Tx, uuid string, model *models.ServiceTemplate) error {
	return nil
}

func DeleteServiceTemplate(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceTemplateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
