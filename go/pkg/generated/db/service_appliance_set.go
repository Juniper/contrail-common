package db
// service_appliance_set

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`fq_name`,`key_value_pair`,`service_appliance_ha_mode`,`service_appliance_driver`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceSetQuery = "update `service_appliance_set` set `fq_name` = ?,`key_value_pair` = ?,`service_appliance_ha_mode` = ?,`service_appliance_driver` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`annotations_key_value_pair` = ?;"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set`"
const selectServiceApplianceSetQuery = "select `fq_name`,`key_value_pair`,`service_appliance_ha_mode`,`service_appliance_driver`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`annotations_key_value_pair` from `service_appliance_set`"

func CreateServiceApplianceSet(tx *sql.Tx, model *models.ServiceApplianceSet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.ServiceApplianceSetProperties.KeyValuePair,
    model.ServiceApplianceHaMode,
    model.ServiceApplianceDriver,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListServiceApplianceSet(tx *sql.Tx) ([]*models.ServiceApplianceSet, error) {
    result := models.MakeServiceApplianceSetSlice()
    rows, err := tx.Query(selectServiceApplianceSetQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceApplianceSet()
            if err := rows.Scan(&m.FQName,
                &m.ServiceApplianceSetProperties.KeyValuePair,
                &m.ServiceApplianceHaMode,
                &m.ServiceApplianceDriver,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.IDPerms.Creator,
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
                &m.DisplayName,
                &m.Annotations.KeyValuePair); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceApplianceSet(db *sql.DB, id string, model *models.ServiceApplianceSet) error {
    return nil
}

func UpdateServiceApplianceSet(db *sql.DB, id string, model *models.ServiceApplianceSet) error {
    return nil
}

func DeleteServiceApplianceSet(db *sql.DB, id string) error {
    return nil
}