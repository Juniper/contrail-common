package db
// config_node

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertConfigNodeQuery = "insert into `config_node` (`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`display_name`,`config_node_ip_address`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateConfigNodeQuery = "update `config_node` set `share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`display_name` = ?,`config_node_ip_address` = ?,`key_value_pair` = ?;"
const deleteConfigNodeQuery = "delete from `config_node`"
const selectConfigNodeQuery = "select `share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`display_name`,`config_node_ip_address`,`key_value_pair` from `config_node`"

func CreateConfigNode(tx *sql.Tx, model *models.ConfigNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertConfigNodeQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.DisplayName,
    model.ConfigNodeIPAddress,
    model.Annotations.KeyValuePair)
    return err
}

func ListConfigNode(tx *sql.Tx) ([]*models.ConfigNode, error) {
    result := models.MakeConfigNodeSlice()
    rows, err := tx.Query(selectConfigNodeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeConfigNode()
            if err := rows.Scan(&m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.DisplayName,
                &m.ConfigNodeIPAddress,
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

func ShowConfigNode(db *sql.DB, id string, model *models.ConfigNode) error {
    return nil
}

func UpdateConfigNode(db *sql.DB, id string, model *models.ConfigNode) error {
    return nil
}

func DeleteConfigNode(db *sql.DB, id string) error {
    return nil
}