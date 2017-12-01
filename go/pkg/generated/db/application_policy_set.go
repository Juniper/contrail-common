package db
// application_policy_set

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertApplicationPolicySetQuery = "insert into `application_policy_set` (`uuid`,`fq_name`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`all_applications`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateApplicationPolicySetQuery = "update `application_policy_set` set `uuid` = ?,`fq_name` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`all_applications` = ?;"
const deleteApplicationPolicySetQuery = "delete from `application_policy_set`"
const selectApplicationPolicySetQuery = "select `uuid`,`fq_name`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`all_applications` from `application_policy_set`"

func CreateApplicationPolicySet(tx *sql.Tx, model *models.ApplicationPolicySet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertApplicationPolicySetQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.UUID,
    model.FQName,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.AllApplications)
    return err
}

func ListApplicationPolicySet(tx *sql.Tx) ([]*models.ApplicationPolicySet, error) {
    result := models.MakeApplicationPolicySetSlice()
    rows, err := tx.Query(selectApplicationPolicySetQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeApplicationPolicySet()
            if err := rows.Scan(&m.UUID,
                &m.FQName,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.AllApplications); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowApplicationPolicySet(db *sql.DB, id string, model *models.ApplicationPolicySet) error {
    return nil
}

func UpdateApplicationPolicySet(db *sql.DB, id string, model *models.ApplicationPolicySet) error {
    return nil
}

func DeleteApplicationPolicySet(db *sql.DB, id string) error {
    return nil
}