package db
// service_group

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceGroupQuery = "insert into `service_group` (`service_group_firewall_service_list`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceGroupQuery = "update `service_group` set `service_group_firewall_service_list` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteServiceGroupQuery = "delete from `service_group`"
const selectServiceGroupQuery = "select `service_group_firewall_service_list`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `service_group`"

func CreateServiceGroup(tx *sql.Tx, model *models.ServiceGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceGroupQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ServiceGroupFirewallServiceList,
    model.UUID,
    model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share)
    return err
}

func ListServiceGroup(tx *sql.Tx) ([]*models.ServiceGroup, error) {
    result := models.MakeServiceGroupSlice()
    rows, err := tx.Query(selectServiceGroupQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceGroup()
            if err := rows.Scan(&m.ServiceGroupFirewallServiceList,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceGroup(db *sql.DB, id string, model *models.ServiceGroup) error {
    return nil
}

func UpdateServiceGroup(db *sql.DB, id string, model *models.ServiceGroup) error {
    return nil
}

func DeleteServiceGroup(db *sql.DB, id string) error {
    return nil
}