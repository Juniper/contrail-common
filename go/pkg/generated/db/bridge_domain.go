package db

// bridge_domain

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertBridgeDomainQuery = "insert into `bridge_domain` (`uuid`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`isid`,`mac_learning_enabled`,`mac_limit`,`mac_limit_action`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`,`mac_aging_time`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBridgeDomainQuery = "update `bridge_domain` set `uuid` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`isid` = ?,`mac_learning_enabled` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`fq_name` = ?,`mac_aging_time` = ?;"
const deleteBridgeDomainQuery = "delete from `bridge_domain` where uuid = ?"
const listBridgeDomainQuery = "select `uuid`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`isid`,`mac_learning_enabled`,`mac_limit`,`mac_limit_action`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`,`mac_aging_time` from `bridge_domain`"
const showBridgeDomainQuery = "select `uuid`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`isid`,`mac_learning_enabled`,`mac_limit`,`mac_limit_action`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`fq_name`,`mac_aging_time` from `bridge_domain` where uuid = ?"

func CreateBridgeDomain(tx *sql.Tx, model *models.BridgeDomain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBridgeDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.Isid),
		bool(model.MacLearningEnabled),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		util.MustJSON(model.FQName),
		int(model.MacAgingTime))
	return err
}

func scanBridgeDomain(rows *sql.Rows) (*models.BridgeDomain, error) {
	m := models.MakeBridgeDomain()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.UUID,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.Isid,
		&m.MacLearningEnabled,
		&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonFQName,
		&m.MacAgingTime); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListBridgeDomain(tx *sql.Tx) ([]*models.BridgeDomain, error) {
	result := models.MakeBridgeDomainSlice()
	rows, err := tx.Query(listBridgeDomainQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanBridgeDomain(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBridgeDomain(tx *sql.Tx, uuid string) (*models.BridgeDomain, error) {
	rows, err := tx.Query(showBridgeDomainQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBridgeDomain(rows)
	}
	return nil, nil
}

func UpdateBridgeDomain(tx *sql.Tx, uuid string, model *models.BridgeDomain) error {
	return nil
}

func DeleteBridgeDomain(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBridgeDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
