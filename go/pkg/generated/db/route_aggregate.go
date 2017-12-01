package db
// route_aggregate

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertRouteAggregateQuery = "insert into `route_aggregate` (`fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRouteAggregateQuery = "update `route_aggregate` set `fq_name` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?;"
const deleteRouteAggregateQuery = "delete from `route_aggregate`"
const selectRouteAggregateQuery = "select `fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid` from `route_aggregate`"

func CreateRouteAggregate(tx *sql.Tx, model *models.RouteAggregate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRouteAggregateQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
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
    model.IDPerms.UserVisible,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID)
    return err
}

func ListRouteAggregate(tx *sql.Tx) ([]*models.RouteAggregate, error) {
    result := models.MakeRouteAggregateSlice()
    rows, err := tx.Query(selectRouteAggregateQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeRouteAggregate()
            if err := rows.Scan(&m.FQName,
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
                &m.IDPerms.UserVisible,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowRouteAggregate(db *sql.DB, id string, model *models.RouteAggregate) error {
    return nil
}

func UpdateRouteAggregate(db *sql.DB, id string, model *models.RouteAggregate) error {
    return nil
}

func DeleteRouteAggregate(db *sql.DB, id string) error {
    return nil
}