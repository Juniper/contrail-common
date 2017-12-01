package db
// address_group

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAddressGroupQuery = "insert into `address_group` (`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`address_group_prefix`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAddressGroupQuery = "update `address_group` set `key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`address_group_prefix` = ?;"
const deleteAddressGroupQuery = "delete from `address_group`"
const selectAddressGroupQuery = "select `key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`address_group_prefix` from `address_group`"

func CreateAddressGroup(tx *sql.Tx, model *models.AddressGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAddressGroupQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.DisplayName,
    model.AddressGroupPrefix)
    return err
}

func ListAddressGroup(tx *sql.Tx) ([]*models.AddressGroup, error) {
    result := models.MakeAddressGroupSlice()
    rows, err := tx.Query(selectAddressGroupQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAddressGroup()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.DisplayName,
                &m.AddressGroupPrefix); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowAddressGroup(db *sql.DB, id string, model *models.AddressGroup) error {
    return nil
}

func UpdateAddressGroup(db *sql.DB, id string, model *models.AddressGroup) error {
    return nil
}

func DeleteAddressGroup(db *sql.DB, id string) error {
    return nil
}