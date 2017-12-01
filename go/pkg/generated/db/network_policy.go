package db

// network_policy

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertNetworkPolicyQuery = "insert into `network_policy` (`policy_rule`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkPolicyQuery = "update `network_policy` set `policy_rule` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteNetworkPolicyQuery = "delete from `network_policy` where uuid = ?"
const listNetworkPolicyQuery = "select `policy_rule`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `network_policy`"
const showNetworkPolicyQuery = "select `policy_rule`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `network_policy` where uuid = ?"

func CreateNetworkPolicy(tx *sql.Tx, model *models.NetworkPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.NetworkPolicyEntries.PolicyRule),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share))
	return err
}

func scanNetworkPolicy(rows *sql.Rows) (*models.NetworkPolicy, error) {
	m := models.MakeNetworkPolicy()

	var jsonNetworkPolicyEntriesPolicyRule string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonNetworkPolicyEntriesPolicyRule,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonNetworkPolicyEntriesPolicyRule), &m.NetworkPolicyEntries.PolicyRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListNetworkPolicy(tx *sql.Tx) ([]*models.NetworkPolicy, error) {
	result := models.MakeNetworkPolicySlice()
	rows, err := tx.Query(listNetworkPolicyQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanNetworkPolicy(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNetworkPolicy(tx *sql.Tx, uuid string) (*models.NetworkPolicy, error) {
	rows, err := tx.Query(showNetworkPolicyQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNetworkPolicy(rows)
	}
	return nil, nil
}

func UpdateNetworkPolicy(tx *sql.Tx, uuid string, model *models.NetworkPolicy) error {
	return nil
}

func DeleteNetworkPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkPolicyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
