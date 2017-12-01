package db
// database_node

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertDatabaseNodeQuery = "insert into `database_node` (`database_node_ip_address`,`uuid`,`fq_name`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDatabaseNodeQuery = "update `database_node` set `database_node_ip_address` = ?,`uuid` = ?,`fq_name` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteDatabaseNodeQuery = "delete from `database_node`"
const selectDatabaseNodeQuery = "select `database_node_ip_address`,`uuid`,`fq_name`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `database_node`"

func CreateDatabaseNode(tx *sql.Tx, model *models.DatabaseNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDatabaseNodeQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DatabaseNodeIPAddress,
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
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner)
    return err
}

func ListDatabaseNode(tx *sql.Tx) ([]*models.DatabaseNode, error) {
    result := models.MakeDatabaseNodeSlice()
    rows, err := tx.Query(selectDatabaseNodeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeDatabaseNode()
            if err := rows.Scan(&m.DatabaseNodeIPAddress,
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
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowDatabaseNode(db *sql.DB, id string, model *models.DatabaseNode) error {
    return nil
}

func UpdateDatabaseNode(db *sql.DB, id string, model *models.DatabaseNode) error {
    return nil
}

func DeleteDatabaseNode(db *sql.DB, id string) error {
    return nil
}