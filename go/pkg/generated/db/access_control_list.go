package db
// access_control_list

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAccessControlListQuery = "insert into `access_control_list` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`access_control_list_hash`,`dynamic`,`acl_rule`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAccessControlListQuery = "update `access_control_list` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`access_control_list_hash` = ?,`dynamic` = ?,`acl_rule` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`display_name` = ?;"
const deleteAccessControlListQuery = "delete from `access_control_list`"
const selectAccessControlListQuery = "select `key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`access_control_list_hash`,`dynamic`,`acl_rule`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`display_name` from `access_control_list`"

func CreateAccessControlList(tx *sql.Tx, model *models.AccessControlList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAccessControlListQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.AccessControlListHash,
    model.AccessControlListEntries.Dynamic,
    model.AccessControlListEntries.ACLRule,
    model.IDPerms.Enable,
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
    model.DisplayName)
    return err
}

func ListAccessControlList(tx *sql.Tx) ([]*models.AccessControlList, error) {
    result := models.MakeAccessControlListSlice()
    rows, err := tx.Query(selectAccessControlListQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAccessControlList()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.AccessControlListHash,
                &m.AccessControlListEntries.Dynamic,
                &m.AccessControlListEntries.ACLRule,
                &m.IDPerms.Enable,
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

func ShowAccessControlList(db *sql.DB, id string, model *models.AccessControlList) error {
    return nil
}

func UpdateAccessControlList(db *sql.DB, id string, model *models.AccessControlList) error {
    return nil
}

func DeleteAccessControlList(db *sql.DB, id string) error {
    return nil
}