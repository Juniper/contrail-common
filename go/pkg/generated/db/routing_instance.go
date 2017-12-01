package db
// routing_instance

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertRoutingInstanceQuery = "insert into `routing_instance` (`display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRoutingInstanceQuery = "update `routing_instance` set `display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteRoutingInstanceQuery = "delete from `routing_instance`"
const selectRoutingInstanceQuery = "select `display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible` from `routing_instance`"

func CreateRoutingInstance(tx *sql.Tx, model *models.RoutingInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingInstanceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible)
    return err
}

func ListRoutingInstance(tx *sql.Tx) ([]*models.RoutingInstance, error) {
    result := models.MakeRoutingInstanceSlice()
    rows, err := tx.Query(selectRoutingInstanceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeRoutingInstance()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.UUID,
                &m.FQName,
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
                &m.IDPerms.UserVisible); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowRoutingInstance(db *sql.DB, id string, model *models.RoutingInstance) error {
    return nil
}

func UpdateRoutingInstance(db *sql.DB, id string, model *models.RoutingInstance) error {
    return nil
}

func DeleteRoutingInstance(db *sql.DB, id string) error {
    return nil
}