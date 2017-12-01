package db

// service_template

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertServiceTemplateQuery = "insert into `service_template` (`service_mode`,`flavor`,`service_type`,`interface_type`,`image_name`,`availability_zone_enable`,`instance_data`,`service_virtualization_type`,`service_scaling`,`vrouter_instance_type`,`ordered_interfaces`,`version`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceTemplateQuery = "update `service_template` set `service_mode` = ?,`flavor` = ?,`service_type` = ?,`interface_type` = ?,`image_name` = ?,`availability_zone_enable` = ?,`instance_data` = ?,`service_virtualization_type` = ?,`service_scaling` = ?,`vrouter_instance_type` = ?,`ordered_interfaces` = ?,`version` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteServiceTemplateQuery = "delete from `service_template` where uuid = ?"
const listServiceTemplateQuery = "select `service_mode`,`flavor`,`service_type`,`interface_type`,`image_name`,`availability_zone_enable`,`instance_data`,`service_virtualization_type`,`service_scaling`,`vrouter_instance_type`,`ordered_interfaces`,`version`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `service_template`"
const showServiceTemplateQuery = "select `service_mode`,`flavor`,`service_type`,`interface_type`,`image_name`,`availability_zone_enable`,`instance_data`,`service_virtualization_type`,`service_scaling`,`vrouter_instance_type`,`ordered_interfaces`,`version`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `service_template` where uuid = ?"

func CreateServiceTemplate(tx *sql.Tx, model *models.ServiceTemplate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceTemplateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ServiceTemplateProperties.ServiceMode),
		string(model.ServiceTemplateProperties.Flavor),
		string(model.ServiceTemplateProperties.ServiceType),
		util.MustJSON(model.ServiceTemplateProperties.InterfaceType),
		string(model.ServiceTemplateProperties.ImageName),
		bool(model.ServiceTemplateProperties.AvailabilityZoneEnable),
		string(model.ServiceTemplateProperties.InstanceData),
		string(model.ServiceTemplateProperties.ServiceVirtualizationType),
		bool(model.ServiceTemplateProperties.ServiceScaling),
		string(model.ServiceTemplateProperties.VrouterInstanceType),
		bool(model.ServiceTemplateProperties.OrderedInterfaces),
		int(model.ServiceTemplateProperties.Version),
		string(model.UUID),
		util.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share))
	return err
}

func scanServiceTemplate(rows *sql.Rows) (*models.ServiceTemplate, error) {
	m := models.MakeServiceTemplate()

	var jsonServiceTemplatePropertiesInterfaceType string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ServiceTemplateProperties.ServiceMode,
		&m.ServiceTemplateProperties.Flavor,
		&m.ServiceTemplateProperties.ServiceType,
		&jsonServiceTemplatePropertiesInterfaceType,
		&m.ServiceTemplateProperties.ImageName,
		&m.ServiceTemplateProperties.AvailabilityZoneEnable,
		&m.ServiceTemplateProperties.InstanceData,
		&m.ServiceTemplateProperties.ServiceVirtualizationType,
		&m.ServiceTemplateProperties.ServiceScaling,
		&m.ServiceTemplateProperties.VrouterInstanceType,
		&m.ServiceTemplateProperties.OrderedInterfaces,
		&m.ServiceTemplateProperties.Version,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonServiceTemplatePropertiesInterfaceType), &m.ServiceTemplateProperties.InterfaceType)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createServiceTemplateWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["service_mode"]; ok {
		results = append(results, "service_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["flavor"]; ok {
		results = append(results, "flavor = ?")
		values = append(values, value)
	}

	if value, ok := where["service_type"]; ok {
		results = append(results, "service_type = ?")
		values = append(values, value)
	}

	if value, ok := where["image_name"]; ok {
		results = append(results, "image_name = ?")
		values = append(values, value)
	}

	if value, ok := where["instance_data"]; ok {
		results = append(results, "instance_data = ?")
		values = append(values, value)
	}

	if value, ok := where["service_virtualization_type"]; ok {
		results = append(results, "service_virtualization_type = ?")
		values = append(values, value)
	}

	if value, ok := where["vrouter_instance_type"]; ok {
		results = append(results, "vrouter_instance_type = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceTemplate(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceTemplate, error) {
	result := models.MakeServiceTemplateSlice()
	whereQuery, values := createServiceTemplateWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceTemplateQuery)
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
