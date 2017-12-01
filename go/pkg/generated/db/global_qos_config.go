package db

// global_qos_config

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertGlobalQosConfigQuery = "insert into `global_qos_config` (`analytics`,`dns`,`control`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalQosConfigQuery = "update `global_qos_config` set `analytics` = ?,`dns` = ?,`control` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteGlobalQosConfigQuery = "delete from `global_qos_config` where uuid = ?"
const listGlobalQosConfigQuery = "select `analytics`,`dns`,`control`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `global_qos_config`"
const showGlobalQosConfigQuery = "select `analytics`,`dns`,`control`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `global_qos_config` where uuid = ?"

func CreateGlobalQosConfig(tx *sql.Tx, model *models.GlobalQosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.ControlTrafficDSCP.Analytics),
		int(model.ControlTrafficDSCP.DNS),
		int(model.ControlTrafficDSCP.Control),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	return err
}

func scanGlobalQosConfig(rows *sql.Rows) (*models.GlobalQosConfig, error) {
	m := models.MakeGlobalQosConfig()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ControlTrafficDSCP.Analytics,
		&m.ControlTrafficDSCP.DNS,
		&m.ControlTrafficDSCP.Control,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListGlobalQosConfig(tx *sql.Tx) ([]*models.GlobalQosConfig, error) {
	result := models.MakeGlobalQosConfigSlice()
	rows, err := tx.Query(listGlobalQosConfigQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanGlobalQosConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowGlobalQosConfig(tx *sql.Tx, uuid string) (*models.GlobalQosConfig, error) {
	rows, err := tx.Query(showGlobalQosConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanGlobalQosConfig(rows)
	}
	return nil, nil
}

func UpdateGlobalQosConfig(tx *sql.Tx, uuid string, model *models.GlobalQosConfig) error {
	return nil
}

func DeleteGlobalQosConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
