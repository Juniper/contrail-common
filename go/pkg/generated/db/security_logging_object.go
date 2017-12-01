package db
// security_logging_object

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertSecurityLoggingObjectQuery = "insert into `security_logging_object` (`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`rule`,`security_logging_object_rate`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityLoggingObjectQuery = "update `security_logging_object` set `fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`rule` = ?,`security_logging_object_rate` = ?;"
const deleteSecurityLoggingObjectQuery = "delete from `security_logging_object`"
const selectSecurityLoggingObjectQuery = "select `fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`rule`,`security_logging_object_rate` from `security_logging_object`"

func CreateSecurityLoggingObject(tx *sql.Tx, model *models.SecurityLoggingObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityLoggingObjectQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.SecurityLoggingObjectRules.Rule,
    model.SecurityLoggingObjectRate)
    return err
}

func ListSecurityLoggingObject(tx *sql.Tx) ([]*models.SecurityLoggingObject, error) {
    result := models.MakeSecurityLoggingObjectSlice()
    rows, err := tx.Query(selectSecurityLoggingObjectQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeSecurityLoggingObject()
            if err := rows.Scan(&m.FQName,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Enable,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.SecurityLoggingObjectRules.Rule,
                &m.SecurityLoggingObjectRate); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowSecurityLoggingObject(db *sql.DB, id string, model *models.SecurityLoggingObject) error {
    return nil
}

func UpdateSecurityLoggingObject(db *sql.DB, id string, model *models.SecurityLoggingObject) error {
    return nil
}

func DeleteSecurityLoggingObject(db *sql.DB, id string) error {
    return nil
}