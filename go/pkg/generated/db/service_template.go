package db
// service_template

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceTemplateQuery = "insert into `service_template` (`key_value_pair`,`ordered_interfaces`,`version`,`service_type`,`vrouter_instance_type`,`service_mode`,`service_virtualization_type`,`image_name`,`availability_zone_enable`,`instance_data`,`interface_type`,`service_scaling`,`flavor`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceTemplateQuery = "update `service_template` set `key_value_pair` = ?,`ordered_interfaces` = ?,`version` = ?,`service_type` = ?,`vrouter_instance_type` = ?,`service_mode` = ?,`service_virtualization_type` = ?,`image_name` = ?,`availability_zone_enable` = ?,`instance_data` = ?,`interface_type` = ?,`service_scaling` = ?,`flavor` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?;"
const deleteServiceTemplateQuery = "delete from `service_template`"
const selectServiceTemplateQuery = "select `key_value_pair`,`ordered_interfaces`,`version`,`service_type`,`vrouter_instance_type`,`service_mode`,`service_virtualization_type`,`image_name`,`availability_zone_enable`,`instance_data`,`interface_type`,`service_scaling`,`flavor`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name` from `service_template`"

func CreateServiceTemplate(tx *sql.Tx, model *models.ServiceTemplate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceTemplateQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.ServiceTemplateProperties.OrderedInterfaces,
    model.ServiceTemplateProperties.Version,
    model.ServiceTemplateProperties.ServiceType,
    model.ServiceTemplateProperties.VrouterInstanceType,
    model.ServiceTemplateProperties.ServiceMode,
    model.ServiceTemplateProperties.ServiceVirtualizationType,
    model.ServiceTemplateProperties.ImageName,
    model.ServiceTemplateProperties.AvailabilityZoneEnable,
    model.ServiceTemplateProperties.InstanceData,
    model.ServiceTemplateProperties.InterfaceType,
    model.ServiceTemplateProperties.ServiceScaling,
    model.ServiceTemplateProperties.Flavor,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID,
    model.FQName,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.DisplayName)
    return err
}

func ListServiceTemplate(tx *sql.Tx) ([]*models.ServiceTemplate, error) {
    result := models.MakeServiceTemplateSlice()
    rows, err := tx.Query(selectServiceTemplateQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceTemplate()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.ServiceTemplateProperties.OrderedInterfaces,
                &m.ServiceTemplateProperties.Version,
                &m.ServiceTemplateProperties.ServiceType,
                &m.ServiceTemplateProperties.VrouterInstanceType,
                &m.ServiceTemplateProperties.ServiceMode,
                &m.ServiceTemplateProperties.ServiceVirtualizationType,
                &m.ServiceTemplateProperties.ImageName,
                &m.ServiceTemplateProperties.AvailabilityZoneEnable,
                &m.ServiceTemplateProperties.InstanceData,
                &m.ServiceTemplateProperties.InterfaceType,
                &m.ServiceTemplateProperties.ServiceScaling,
                &m.ServiceTemplateProperties.Flavor,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.UUID,
                &m.FQName,
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
                &m.IDPerms.Creator,
                &m.DisplayName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceTemplate(db *sql.DB, id string, model *models.ServiceTemplate) error {
    return nil
}

func UpdateServiceTemplate(db *sql.DB, id string, model *models.ServiceTemplate) error {
    return nil
}

func DeleteServiceTemplate(db *sql.DB, id string) error {
    return nil
}