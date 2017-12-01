package db
// security_group

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertSecurityGroupQuery = "insert into `security_group` (`policy_rule`,`configured_security_group_id`,`security_group_id`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityGroupQuery = "update `security_group` set `policy_rule` = ?,`configured_security_group_id` = ?,`security_group_id` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`uuid` = ?;"
const deleteSecurityGroupQuery = "delete from `security_group`"
const selectSecurityGroupQuery = "select `policy_rule`,`configured_security_group_id`,`security_group_id`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name`,`uuid` from `security_group`"

func CreateSecurityGroup(tx *sql.Tx, model *models.SecurityGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityGroupQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.SecurityGroupEntries.PolicyRule,
    model.ConfiguredSecurityGroupID,
    model.SecurityGroupID,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.FQName,
    model.UUID)
    return err
}

func ListSecurityGroup(tx *sql.Tx) ([]*models.SecurityGroup, error) {
    result := models.MakeSecurityGroupSlice()
    rows, err := tx.Query(selectSecurityGroupQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeSecurityGroup()
            if err := rows.Scan(&m.SecurityGroupEntries.PolicyRule,
                &m.ConfiguredSecurityGroupID,
                &m.SecurityGroupID,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.FQName,
                &m.UUID); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowSecurityGroup(db *sql.DB, id string, model *models.SecurityGroup) error {
    return nil
}

func UpdateSecurityGroup(db *sql.DB, id string, model *models.SecurityGroup) error {
    return nil
}

func DeleteSecurityGroup(db *sql.DB, id string) error {
    return nil
}