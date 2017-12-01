package db

// security_group

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertSecurityGroupQuery = "insert into `security_group` (`uuid`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`policy_rule`,`configured_security_group_id`,`security_group_id`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityGroupQuery = "update `security_group` set `uuid` = ?,`display_name` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`policy_rule` = ?,`configured_security_group_id` = ?,`security_group_id` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`key_value_pair` = ?;"
const deleteSecurityGroupQuery = "delete from `security_group` where uuid = ?"
const listSecurityGroupQuery = "select `uuid`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`policy_rule`,`configured_security_group_id`,`security_group_id`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`key_value_pair` from `security_group`"
const showSecurityGroupQuery = "select `uuid`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`policy_rule`,`configured_security_group_id`,`security_group_id`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`key_value_pair` from `security_group` where uuid = ?"

func CreateSecurityGroup(tx *sql.Tx, model *models.SecurityGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		string(model.DisplayName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		util.MustJSON(model.SecurityGroupEntries.PolicyRule),
		int(model.ConfiguredSecurityGroupID),
		int(model.SecurityGroupID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanSecurityGroup(rows *sql.Rows) (*models.SecurityGroup, error) {
	m := models.MakeSecurityGroup()

	var jsonPerms2Share string

	var jsonSecurityGroupEntriesPolicyRule string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.UUID,
		&m.DisplayName,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonSecurityGroupEntriesPolicyRule,
		&m.ConfiguredSecurityGroupID,
		&m.SecurityGroupID,
		&jsonFQName,
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
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonSecurityGroupEntriesPolicyRule), &m.SecurityGroupEntries.PolicyRule)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListSecurityGroup(tx *sql.Tx) ([]*models.SecurityGroup, error) {
	result := models.MakeSecurityGroupSlice()
	rows, err := tx.Query(listSecurityGroupQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanSecurityGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowSecurityGroup(tx *sql.Tx, uuid string) (*models.SecurityGroup, error) {
	rows, err := tx.Query(showSecurityGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanSecurityGroup(rows)
	}
	return nil, nil
}

func UpdateSecurityGroup(tx *sql.Tx, uuid string, model *models.SecurityGroup) error {
	return nil
}

func DeleteSecurityGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSecurityGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
