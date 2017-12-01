package db
// subnet

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertSubnetQuery = "insert into `subnet` (`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`ip_prefix`,`ip_prefix_len`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSubnetQuery = "update `subnet` set `key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?;"
const deleteSubnetQuery = "delete from `subnet`"
const selectSubnetQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`ip_prefix`,`ip_prefix_len`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name` from `subnet`"

func CreateSubnet(tx *sql.Tx, model *models.Subnet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSubnetQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.FQName,
    model.SubnetIPPrefix.IPPrefix,
    model.SubnetIPPrefix.IPPrefixLen,
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
    model.IDPerms.Enable,
    model.DisplayName)
    return err
}

func ListSubnet(tx *sql.Tx) ([]*models.Subnet, error) {
    result := models.MakeSubnetSlice()
    rows, err := tx.Query(selectSubnetQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeSubnet()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.UUID,
                &m.FQName,
                &m.SubnetIPPrefix.IPPrefix,
                &m.SubnetIPPrefix.IPPrefixLen,
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
                &m.IDPerms.Enable,
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

func ShowSubnet(db *sql.DB, id string, model *models.Subnet) error {
    return nil
}

func UpdateSubnet(db *sql.DB, id string, model *models.Subnet) error {
    return nil
}

func DeleteSubnet(db *sql.DB, id string) error {
    return nil
}