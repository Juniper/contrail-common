package db

// security_logging_object

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertSecurityLoggingObjectQuery = "insert into `security_logging_object` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`rule`,`security_logging_object_rate`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityLoggingObjectQuery = "update `security_logging_object` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`display_name` = ?,`rule` = ?,`security_logging_object_rate` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteSecurityLoggingObjectQuery = "delete from `security_logging_object` where uuid = ?"
const listSecurityLoggingObjectQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`rule`,`security_logging_object_rate`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name` from `security_logging_object`"
const showSecurityLoggingObjectQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`rule`,`security_logging_object_rate`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name` from `security_logging_object` where uuid = ?"

func CreateSecurityLoggingObject(tx *sql.Tx, model *models.SecurityLoggingObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Description),
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
		string(model.DisplayName),
		util.MustJSON(model.SecurityLoggingObjectRules.Rule),
		int(model.SecurityLoggingObjectRate),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanSecurityLoggingObject(rows *sql.Rows) (*models.SecurityLoggingObject, error) {
	m := models.MakeSecurityLoggingObject()

	var jsonSecurityLoggingObjectRulesRule string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Description,
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
		&m.DisplayName,
		&jsonSecurityLoggingObjectRulesRule,
		&m.SecurityLoggingObjectRate,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonSecurityLoggingObjectRulesRule), &m.SecurityLoggingObjectRules.Rule)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListSecurityLoggingObject(tx *sql.Tx) ([]*models.SecurityLoggingObject, error) {
	result := models.MakeSecurityLoggingObjectSlice()
	rows, err := tx.Query(listSecurityLoggingObjectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanSecurityLoggingObject(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowSecurityLoggingObject(tx *sql.Tx, uuid string) (*models.SecurityLoggingObject, error) {
	rows, err := tx.Query(showSecurityLoggingObjectQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanSecurityLoggingObject(rows)
	}
	return nil, nil
}

func UpdateSecurityLoggingObject(tx *sql.Tx, uuid string, model *models.SecurityLoggingObject) error {
	return nil
}

func DeleteSecurityLoggingObject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSecurityLoggingObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
