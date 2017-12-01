package db

// dsa_rule

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertDsaRuleQuery = "insert into `dsa_rule` (`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`subscriber`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`ep_version`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDsaRuleQuery = "update `dsa_rule` set `created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`subscriber` = ?,`ep_id` = ?,`ep_type` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`ep_version` = ?;"
const deleteDsaRuleQuery = "delete from `dsa_rule` where uuid = ?"
const listDsaRuleQuery = "select `created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`subscriber`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`ep_version` from `dsa_rule`"
const showDsaRuleQuery = "select `created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`subscriber`,`ep_id`,`ep_type`,`ip_prefix`,`ip_prefix_len`,`ep_version` from `dsa_rule` where uuid = ?"

func CreateDsaRule(tx *sql.Tx, model *models.DsaRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDsaRuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		util.MustJSON(model.DsaRuleEntry.Subscriber),
		string(model.DsaRuleEntry.Publisher.EpID),
		string(model.DsaRuleEntry.Publisher.EpType),
		string(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefix),
		int(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen),
		string(model.DsaRuleEntry.Publisher.EpVersion))
	return err
}

func scanDsaRule(rows *sql.Rows) (*models.DsaRule, error) {
	m := models.MakeDsaRule()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonDsaRuleEntrySubscriber string

	if err := rows.Scan(&m.IDPerms.Created,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
		&jsonDsaRuleEntrySubscriber,
		&m.DsaRuleEntry.Publisher.EpID,
		&m.DsaRuleEntry.Publisher.EpType,
		&m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix,
		&m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen,
		&m.DsaRuleEntry.Publisher.EpVersion); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonDsaRuleEntrySubscriber), &m.DsaRuleEntry.Subscriber)

	return m, nil
}

func createDsaRuleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["ep_id"]; ok {
		results = append(results, "ep_id = ?")
		values = append(values, value)
	}

	if value, ok := where["ep_type"]; ok {
		results = append(results, "ep_type = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["ep_version"]; ok {
		results = append(results, "ep_version = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListDsaRule(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.DsaRule, error) {
	result := models.MakeDsaRuleSlice()
	whereQuery, values := createDsaRuleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listDsaRuleQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
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
