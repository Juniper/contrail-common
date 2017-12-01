package db
// route_target

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertRouteTargetQuery = "insert into `route_target` (`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteTargetQuery = "update `route_target` set `key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?;"
const deleteRouteTargetQuery = "delete from `route_target`"
const selectRouteTargetQuery = "select `key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name` from `route_target`"

func CreateRouteTarget(tx *sql.Tx, model *models.RouteTarget) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteTargetQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DisplayName)
    return err
}

func ListRouteTarget(tx *sql.Tx) ([]*models.RouteTarget, error) {
    result := models.MakeRouteTargetSlice()
    rows, err := tx.Query(selectRouteTargetQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeRouteTarget()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
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

func ShowRouteTarget(db *sql.DB, id string, model *models.RouteTarget) error {
    return nil
}

func UpdateRouteTarget(db *sql.DB, id string, model *models.RouteTarget) error {
    return nil
}

func DeleteRouteTarget(db *sql.DB, id string) error {
    return nil
}