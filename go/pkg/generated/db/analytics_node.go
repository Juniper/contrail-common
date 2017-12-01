package db
// analytics_node

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAnalyticsNodeQuery = "insert into `analytics_node` (`key_value_pair`,`analytics_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAnalyticsNodeQuery = "update `analytics_node` set `key_value_pair` = ?,`analytics_node_ip_address` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?;"
const deleteAnalyticsNodeQuery = "delete from `analytics_node`"
const selectAnalyticsNodeQuery = "select `key_value_pair`,`analytics_node_ip_address`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name` from `analytics_node`"

func CreateAnalyticsNode(tx *sql.Tx, model *models.AnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAnalyticsNodeQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.AnalyticsNodeIPAddress,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.DisplayName)
    return err
}

func ListAnalyticsNode(tx *sql.Tx) ([]*models.AnalyticsNode, error) {
    result := models.MakeAnalyticsNodeSlice()
    rows, err := tx.Query(selectAnalyticsNodeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAnalyticsNode()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.AnalyticsNodeIPAddress,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
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

func ShowAnalyticsNode(db *sql.DB, id string, model *models.AnalyticsNode) error {
    return nil
}

func UpdateAnalyticsNode(db *sql.DB, id string, model *models.AnalyticsNode) error {
    return nil
}

func DeleteAnalyticsNode(db *sql.DB, id string) error {
    return nil
}