package db
// logical_interface

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLogicalInterfaceQuery = "insert into `logical_interface` (`owner_access`,`global_access`,`share`,`owner`,`logical_interface_vlan_tag`,`logical_interface_type`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalInterfaceQuery = "update `logical_interface` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`logical_interface_vlan_tag` = ?,`logical_interface_type` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteLogicalInterfaceQuery = "delete from `logical_interface`"
const selectLogicalInterfaceQuery = "select `owner_access`,`global_access`,`share`,`owner`,`logical_interface_vlan_tag`,`logical_interface_type`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`display_name`,`key_value_pair` from `logical_interface`"

func CreateLogicalInterface(tx *sql.Tx, model *models.LogicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalInterfaceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.LogicalInterfaceVlanTag,
    model.LogicalInterfaceType,
    model.UUID,
    model.FQName,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListLogicalInterface(tx *sql.Tx) ([]*models.LogicalInterface, error) {
    result := models.MakeLogicalInterfaceSlice()
    rows, err := tx.Query(selectLogicalInterfaceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLogicalInterface()
            if err := rows.Scan(&m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.LogicalInterfaceVlanTag,
                &m.LogicalInterfaceType,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable,
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

func ShowLogicalInterface(db *sql.DB, id string, model *models.LogicalInterface) error {
    return nil
}

func UpdateLogicalInterface(db *sql.DB, id string, model *models.LogicalInterface) error {
    return nil
}

func DeleteLogicalInterface(db *sql.DB, id string) error {
    return nil
}