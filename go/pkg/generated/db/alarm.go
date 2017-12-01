package db

// alarm

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAlarmQuery = "insert into `alarm` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`alarm_rules`,`alarm_severity`,`key_value_pair`,`uuid`,`uve_key`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAlarmQuery = "update `alarm` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`alarm_rules` = ?,`alarm_severity` = ?,`key_value_pair` = ?,`uuid` = ?,`uve_key` = ?,`fq_name` = ?;"
const deleteAlarmQuery = "delete from `alarm` where uuid = ?"
const listAlarmQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`alarm_rules`,`alarm_severity`,`key_value_pair`,`uuid`,`uve_key`,`fq_name` from `alarm`"
const showAlarmQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`alarm_rules`,`alarm_severity`,`key_value_pair`,`uuid`,`uve_key`,`fq_name` from `alarm` where uuid = ?"

func CreateAlarm(tx *sql.Tx, model *models.Alarm) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAlarmQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		util.MustJSON(model.AlarmRules),
		int(model.AlarmSeverity),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		util.MustJSON(model.UveKeys.UveKey),
		util.MustJSON(model.FQName))
	return err
}

func scanAlarm(rows *sql.Rows) (*models.Alarm, error) {
	m := models.MakeAlarm()

	var jsonPerms2Share string

	var jsonAlarmRules string

	var jsonAnnotationsKeyValuePair string

	var jsonUveKeysUveKey string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonAlarmRules,
		&m.AlarmSeverity,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonUveKeysUveKey,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAlarmRules), &m.AlarmRules)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonUveKeysUveKey), &m.UveKeys.UveKey)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListAlarm(tx *sql.Tx) ([]*models.Alarm, error) {
	result := models.MakeAlarmSlice()
	rows, err := tx.Query(listAlarmQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAlarm(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAlarm(tx *sql.Tx, uuid string) (*models.Alarm, error) {
	rows, err := tx.Query(showAlarmQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAlarm(rows)
	}
	return nil, nil
}

func UpdateAlarm(tx *sql.Tx, uuid string, model *models.Alarm) error {
	return nil
}

func DeleteAlarm(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAlarmQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
