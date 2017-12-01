package db
// logical_router

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLogicalRouterQuery = "insert into `logical_router` (`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`vxlan_network_identifier`,`route_target`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalRouterQuery = "update `logical_router` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`vxlan_network_identifier` = ?,`route_target` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteLogicalRouterQuery = "delete from `logical_router`"
const selectLogicalRouterQuery = "select `owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`vxlan_network_identifier`,`route_target`,`display_name`,`key_value_pair` from `logical_router`"

func CreateLogicalRouter(tx *sql.Tx, model *models.LogicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalRouterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
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
    model.VxlanNetworkIdentifier,
    model.ConfiguredRouteTargetList.RouteTarget,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListLogicalRouter(tx *sql.Tx) ([]*models.LogicalRouter, error) {
    result := models.MakeLogicalRouterSlice()
    rows, err := tx.Query(selectLogicalRouterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLogicalRouter()
            if err := rows.Scan(&m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
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
                &m.VxlanNetworkIdentifier,
                &m.ConfiguredRouteTargetList.RouteTarget,
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

func ShowLogicalRouter(db *sql.DB, id string, model *models.LogicalRouter) error {
    return nil
}

func UpdateLogicalRouter(db *sql.DB, id string, model *models.LogicalRouter) error {
    return nil
}

func DeleteLogicalRouter(db *sql.DB, id string) error {
    return nil
}