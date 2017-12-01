package db
// network_policy

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertNetworkPolicyQuery = "insert into `network_policy` (`policy_rule`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkPolicyQuery = "update `network_policy` set `policy_rule` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteNetworkPolicyQuery = "delete from `network_policy`"
const selectNetworkPolicyQuery = "select `policy_rule`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name` from `network_policy`"

func CreateNetworkPolicy(tx *sql.Tx, model *models.NetworkPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkPolicyQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.NetworkPolicyEntries.PolicyRule,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
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
    model.IDPerms.Description,
    model.DisplayName)
    return err
}

func ListNetworkPolicy(tx *sql.Tx) ([]*models.NetworkPolicy, error) {
    result := models.MakeNetworkPolicySlice()
    rows, err := tx.Query(selectNetworkPolicyQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeNetworkPolicy()
            if err := rows.Scan(&m.NetworkPolicyEntries.PolicyRule,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
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
                &m.IDPerms.Description,
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

func ShowNetworkPolicy(db *sql.DB, id string, model *models.NetworkPolicy) error {
    return nil
}

func UpdateNetworkPolicy(db *sql.DB, id string, model *models.NetworkPolicy) error {
    return nil
}

func DeleteNetworkPolicy(db *sql.DB, id string) error {
    return nil
}