package db

// alarm

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertAlarmQuery = "insert into `alarm` (`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`key_value_pair`,`alarm_rules`,`uve_key`,`alarm_severity`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAlarmQuery = "update `alarm` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`fq_name` = ?,`key_value_pair` = ?,`alarm_rules` = ?,`uve_key` = ?,`alarm_severity` = ?,`uuid` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteAlarmQuery = "delete from `alarm` where uuid = ?"
const listAlarmQuery = "select `owner_access`,`global_access`,`share`,`owner`,`fq_name`,`key_value_pair`,`alarm_rules`,`uve_key`,`alarm_severity`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name` from `alarm`"
const showAlarmQuery = "select `owner_access`,`global_access`,`share`,`owner`,`fq_name`,`key_value_pair`,`alarm_rules`,`uve_key`,`alarm_severity`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name` from `alarm` where uuid = ?"

func CreateAlarm(tx *sql.Tx, model *models.Alarm) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAlarmQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		util.MustJSON(model.FQName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.AlarmRules),
		util.MustJSON(model.UveKeys.UveKey),
		int(model.AlarmSeverity),
		string(model.UUID),
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
		string(model.IDPerms.Description),
		string(model.DisplayName))
	return err
}

func scanAlarm(rows *sql.Rows) (*models.Alarm, error) {
	m := models.MakeAlarm()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonAlarmRules string

	var jsonUveKeysUveKey string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonFQName,
		&jsonAnnotationsKeyValuePair,
		&jsonAlarmRules,
		&jsonUveKeysUveKey,
		&m.AlarmSeverity,
		&m.UUID,
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
		&m.IDPerms.Description,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonAlarmRules), &m.AlarmRules)

	json.Unmarshal([]byte(jsonUveKeysUveKey), &m.UveKeys.UveKey)

	return m, nil
}

func createAlarmWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListAlarm(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Alarm, error) {
	result := models.MakeAlarmSlice()
	whereQuery, values := createAlarmWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAlarmQuery)
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
