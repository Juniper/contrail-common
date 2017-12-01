package db
// api_access_list

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAPIAccessListQuery = "insert into `api_access_list` (`rbac_rule`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAPIAccessListQuery = "update `api_access_list` set `rbac_rule` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteAPIAccessListQuery = "delete from `api_access_list`"
const selectAPIAccessListQuery = "select `rbac_rule`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair` from `api_access_list`"

func CreateAPIAccessList(tx *sql.Tx, model *models.APIAccessList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAPIAccessListQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.APIAccessListEntries.RbacRule,
    model.Perms2.Share,
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
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListAPIAccessList(tx *sql.Tx) ([]*models.APIAccessList, error) {
    result := models.MakeAPIAccessListSlice()
    rows, err := tx.Query(selectAPIAccessListQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAPIAccessList()
            if err := rows.Scan(&m.APIAccessListEntries.RbacRule,
                &m.Perms2.Share,
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
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
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

func ShowAPIAccessList(db *sql.DB, id string, model *models.APIAccessList) error {
    return nil
}

func UpdateAPIAccessList(db *sql.DB, id string, model *models.APIAccessList) error {
    return nil
}

func DeleteAPIAccessList(db *sql.DB, id string) error {
    return nil
}