package db
// dsa_rule

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertDsaRuleQuery = "insert into `dsa_rule` (`key_value_pair`,`subscriber`,`ep_version`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDsaRuleQuery = "update `dsa_rule` set `key_value_pair` = ?,`subscriber` = ?,`ep_version` = ?,`ep_id` = ?,`ep_type` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteDsaRuleQuery = "delete from `dsa_rule`"
const selectDsaRuleQuery = "select `key_value_pair`,`subscriber`,`ep_version`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name` from `dsa_rule`"

func CreateDsaRule(tx *sql.Tx, model *models.DsaRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDsaRuleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.DsaRuleEntry.Subscriber,
    model.DsaRuleEntry.Publisher.EpVersion,
    model.DsaRuleEntry.Publisher.EpID,
    model.DsaRuleEntry.Publisher.EpType,
    model.DsaRuleEntry.Publisher.EpPrefix.IPPrefix,
    model.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen,
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
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.DisplayName)
    return err
}

func ListDsaRule(tx *sql.Tx) ([]*models.DsaRule, error) {
    result := models.MakeDsaRuleSlice()
    rows, err := tx.Query(selectDsaRuleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeDsaRule()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.DsaRuleEntry.Subscriber,
                &m.DsaRuleEntry.Publisher.EpVersion,
                &m.DsaRuleEntry.Publisher.EpID,
                &m.DsaRuleEntry.Publisher.EpType,
                &m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix,
                &m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen,
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
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
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

func ShowDsaRule(db *sql.DB, id string, model *models.DsaRule) error {
    return nil
}

func UpdateDsaRule(db *sql.DB, id string, model *models.DsaRule) error {
    return nil
}

func DeleteDsaRule(db *sql.DB, id string) error {
    return nil
}