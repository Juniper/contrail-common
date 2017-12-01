package db
// alarm

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAlarmQuery = "insert into `alarm` (`alarm_severity`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uve_key`,`uuid`,`fq_name`,`alarm_rules`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAlarmQuery = "update `alarm` set `alarm_severity` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uve_key` = ?,`uuid` = ?,`fq_name` = ?,`alarm_rules` = ?;"
const deleteAlarmQuery = "delete from `alarm`"
const selectAlarmQuery = "select `alarm_severity`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uve_key`,`uuid`,`fq_name`,`alarm_rules` from `alarm`"

func CreateAlarm(tx *sql.Tx, model *models.Alarm) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAlarmQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.AlarmSeverity,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UveKeys.UveKey,
    model.UUID,
    model.FQName,
    model.AlarmRules)
    return err
}

func ListAlarm(tx *sql.Tx) ([]*models.Alarm, error) {
    result := models.MakeAlarmSlice()
    rows, err := tx.Query(selectAlarmQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAlarm()
            if err := rows.Scan(&m.AlarmSeverity,
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
                &m.IDPerms.LastModified,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UveKeys.UveKey,
                &m.UUID,
                &m.FQName,
                &m.AlarmRules); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowAlarm(db *sql.DB, id string, model *models.Alarm) error {
    return nil
}

func UpdateAlarm(db *sql.DB, id string, model *models.Alarm) error {
    return nil
}

func DeleteAlarm(db *sql.DB, id string) error {
    return nil
}