package db
// service_appliance

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceApplianceQuery = "insert into `service_appliance` (`key_value_pair`,`fq_name`,`password`,`username`,`service_appliance_ip_address`,`uuid`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`service_appliance_properties_key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceQuery = "update `service_appliance` set `key_value_pair` = ?,`fq_name` = ?,`password` = ?,`username` = ?,`service_appliance_ip_address` = ?,`uuid` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`display_name` = ?,`service_appliance_properties_key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteServiceApplianceQuery = "delete from `service_appliance`"
const selectServiceApplianceQuery = "select `key_value_pair`,`fq_name`,`password`,`username`,`service_appliance_ip_address`,`uuid`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`service_appliance_properties_key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `service_appliance`"

func CreateServiceAppliance(tx *sql.Tx, model *models.ServiceAppliance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.FQName,
    model.ServiceApplianceUserCredentials.Password,
    model.ServiceApplianceUserCredentials.Username,
    model.ServiceApplianceIPAddress,
    model.UUID,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.DisplayName,
    model.ServiceApplianceProperties.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess)
    return err
}

func ListServiceAppliance(tx *sql.Tx) ([]*models.ServiceAppliance, error) {
    result := models.MakeServiceApplianceSlice()
    rows, err := tx.Query(selectServiceApplianceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceAppliance()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.FQName,
                &m.ServiceApplianceUserCredentials.Password,
                &m.ServiceApplianceUserCredentials.Username,
                &m.ServiceApplianceIPAddress,
                &m.UUID,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Enable,
                &m.DisplayName,
                &m.ServiceApplianceProperties.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceAppliance(db *sql.DB, id string, model *models.ServiceAppliance) error {
    return nil
}

func UpdateServiceAppliance(db *sql.DB, id string, model *models.ServiceAppliance) error {
    return nil
}

func DeleteServiceAppliance(db *sql.DB, id string) error {
    return nil
}