package db
// service_object

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceObjectQuery = "insert into `service_object` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceObjectQuery = "update `service_object` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?;"
const deleteServiceObjectQuery = "delete from `service_object`"
const selectServiceObjectQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator` from `service_object`"

func CreateServiceObject(tx *sql.Tx, model *models.ServiceObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceObjectQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator)
    return err
}

func ListServiceObject(tx *sql.Tx) ([]*models.ServiceObject, error) {
    result := models.MakeServiceObjectSlice()
    rows, err := tx.Query(selectServiceObjectQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceObject()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceObject(db *sql.DB, id string, model *models.ServiceObject) error {
    return nil
}

func UpdateServiceObject(db *sql.DB, id string, model *models.ServiceObject) error {
    return nil
}

func DeleteServiceObject(db *sql.DB, id string) error {
    return nil
}