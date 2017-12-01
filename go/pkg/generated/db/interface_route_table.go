package db
// interface_route_table

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertInterfaceRouteTableQuery = "insert into `interface_route_table` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInterfaceRouteTableQuery = "update `interface_route_table` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`route` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?;"
const deleteInterfaceRouteTableQuery = "delete from `interface_route_table`"
const selectInterfaceRouteTableQuery = "select `key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`route`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name` from `interface_route_table`"

func CreateInterfaceRouteTable(tx *sql.Tx, model *models.InterfaceRouteTable) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInterfaceRouteTableQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.InterfaceRouteTableRoutes.Route,
    model.UUID,
    model.FQName,
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
    model.IDPerms.Creator,
    model.DisplayName)
    return err
}

func ListInterfaceRouteTable(tx *sql.Tx) ([]*models.InterfaceRouteTable, error) {
    result := models.MakeInterfaceRouteTableSlice()
    rows, err := tx.Query(selectInterfaceRouteTableQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeInterfaceRouteTable()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.InterfaceRouteTableRoutes.Route,
                &m.UUID,
                &m.FQName,
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

func ShowInterfaceRouteTable(db *sql.DB, id string, model *models.InterfaceRouteTable) error {
    return nil
}

func UpdateInterfaceRouteTable(db *sql.DB, id string, model *models.InterfaceRouteTable) error {
    return nil
}

func DeleteInterfaceRouteTable(db *sql.DB, id string) error {
    return nil
}