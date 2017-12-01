package db
// route_table

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertRouteTableQuery = "insert into `route_table` (`owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteTableQuery = "update `route_table` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`route` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteRouteTableQuery = "delete from `route_table`"
const selectRouteTableQuery = "select `owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair` from `route_table`"

func CreateRouteTable(tx *sql.Tx, model *models.RouteTable) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteTableQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Routes.Route,
    model.UUID,
    model.FQName,
    model.IDPerms.Creator,
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
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListRouteTable(tx *sql.Tx) ([]*models.RouteTable, error) {
    result := models.MakeRouteTableSlice()
    rows, err := tx.Query(selectRouteTableQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeRouteTable()
            if err := rows.Scan(&m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Routes.Route,
                &m.UUID,
                &m.FQName,
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

func ShowRouteTable(db *sql.DB, id string, model *models.RouteTable) error {
    return nil
}

func UpdateRouteTable(db *sql.DB, id string, model *models.RouteTable) error {
    return nil
}

func DeleteRouteTable(db *sql.DB, id string) error {
    return nil
}