package db
// policy_management

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertPolicyManagementQuery = "insert into `policy_management` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePolicyManagementQuery = "update `policy_management` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?;"
const deletePolicyManagementQuery = "delete from `policy_management`"
const selectPolicyManagementQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator` from `policy_management`"

func CreatePolicyManagement(tx *sql.Tx, model *models.PolicyManagement) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPolicyManagementQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator)
    return err
}

func ListPolicyManagement(tx *sql.Tx) ([]*models.PolicyManagement, error) {
    result := models.MakePolicyManagementSlice()
    rows, err := tx.Query(selectPolicyManagementQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakePolicyManagement()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowPolicyManagement(db *sql.DB, id string, model *models.PolicyManagement) error {
    return nil
}

func UpdatePolicyManagement(db *sql.DB, id string, model *models.PolicyManagement) error {
    return nil
}

func DeletePolicyManagement(db *sql.DB, id string) error {
    return nil
}