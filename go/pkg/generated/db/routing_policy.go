package db
// routing_policy

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertRoutingPolicyQuery = "insert into `routing_policy` (`fq_name`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateRoutingPolicyQuery = "update `routing_policy` set `fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?;"
const deleteRoutingPolicyQuery = "delete from `routing_policy`"
const selectRoutingPolicyQuery = "select `fq_name`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid` from `routing_policy`"

func CreateRoutingPolicy(tx *sql.Tx, model *models.RoutingPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertRoutingPolicyQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID)
    return err
}

func ListRoutingPolicy(tx *sql.Tx) ([]*models.RoutingPolicy, error) {
    result := models.MakeRoutingPolicySlice()
    rows, err := tx.Query(selectRoutingPolicyQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeRoutingPolicy()
            if err := rows.Scan(&m.FQName,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
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

func ShowRoutingPolicy(db *sql.DB, id string, model *models.RoutingPolicy) error {
    return nil
}

func UpdateRoutingPolicy(db *sql.DB, id string, model *models.RoutingPolicy) error {
    return nil
}

func DeleteRoutingPolicy(db *sql.DB, id string) error {
    return nil
}