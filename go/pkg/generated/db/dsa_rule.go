package db

// dsa_rule

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertDsaRuleQuery = "insert into `dsa_rule` (`uuid`,`subscriber`,`ep_version`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`fq_name`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDsaRuleQuery = "update `dsa_rule` set `uuid` = ?,`subscriber` = ?,`ep_version` = ?,`ep_id` = ?,`ep_type` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`fq_name` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteDsaRuleQuery = "delete from `dsa_rule` where uuid = ?"
const listDsaRuleQuery = "select `uuid`,`subscriber`,`ep_version`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`fq_name`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `dsa_rule`"
const showDsaRuleQuery = "select `uuid`,`subscriber`,`ep_version`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`fq_name`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `dsa_rule` where uuid = ?"

func CreateDsaRule(tx *sql.Tx, model *models.DsaRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDsaRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		util.MustJSON(model.DsaRuleEntry.Subscriber),
		string(model.DsaRuleEntry.Publisher.EpVersion),
		string(model.DsaRuleEntry.Publisher.EpID),
		string(model.DsaRuleEntry.Publisher.EpType),
		string(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefix),
		int(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))
	return err
}

func scanDsaRule(rows *sql.Rows) (*models.DsaRule, error) {
	m := models.MakeDsaRule()

	var jsonDsaRuleEntrySubscriber string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonDsaRuleEntrySubscriber,
		&m.DsaRuleEntry.Publisher.EpVersion,
		&m.DsaRuleEntry.Publisher.EpID,
		&m.DsaRuleEntry.Publisher.EpType,
		&m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix,
		&m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen,
		&jsonFQName,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonDsaRuleEntrySubscriber), &m.DsaRuleEntry.Subscriber)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListDsaRule(tx *sql.Tx) ([]*models.DsaRule, error) {
	result := models.MakeDsaRuleSlice()
	rows, err := tx.Query(listDsaRuleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanDsaRule(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDsaRule(tx *sql.Tx, uuid string) (*models.DsaRule, error) {
	rows, err := tx.Query(showDsaRuleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDsaRule(rows)
	}
	return nil, nil
}

func UpdateDsaRule(tx *sql.Tx, uuid string, model *models.DsaRule) error {
	return nil
}

func DeleteDsaRule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDsaRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
