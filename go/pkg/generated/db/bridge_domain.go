package db

// bridge_domain

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertBridgeDomainQuery = "insert into `bridge_domain` (`mac_limit`,`mac_limit_action`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`key_value_pair`,`uuid`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`mac_aging_time`,`isid`,`mac_learning_enabled`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBridgeDomainQuery = "update `bridge_domain` set `mac_limit` = ?,`mac_limit_action` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`fq_name` = ?,`key_value_pair` = ?,`uuid` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`mac_aging_time` = ?,`isid` = ?,`mac_learning_enabled` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`display_name` = ?;"
const deleteBridgeDomainQuery = "delete from `bridge_domain` where uuid = ?"
const listBridgeDomainQuery = "select `mac_limit`,`mac_limit_action`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`key_value_pair`,`uuid`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`mac_aging_time`,`isid`,`mac_learning_enabled`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`display_name` from `bridge_domain`"
const showBridgeDomainQuery = "select `mac_limit`,`mac_limit_action`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`key_value_pair`,`uuid`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`mac_aging_time`,`isid`,`mac_learning_enabled`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`display_name` from `bridge_domain` where uuid = ?"

func CreateBridgeDomain(tx *sql.Tx, model *models.BridgeDomain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBridgeDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		util.MustJSON(model.FQName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
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
		string(model.IDPerms.Creator),
		int(model.MacAgingTime),
		int(model.Isid),
		bool(model.MacLearningEnabled),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		string(model.DisplayName))
	return err
}

func scanBridgeDomain(rows *sql.Rows) (*models.BridgeDomain, error) {
	m := models.MakeBridgeDomain()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonFQName,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
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
		&m.IDPerms.Creator,
		&m.MacAgingTime,
		&m.Isid,
		&m.MacLearningEnabled,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createBridgeDomainWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["mac_limit_action"]; ok {
		results = append(results, "mac_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
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

	if value, ok := where["mac_move_limit_action"]; ok {
		results = append(results, "mac_move_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListBridgeDomain(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.BridgeDomain, error) {
	result := models.MakeBridgeDomainSlice()
	whereQuery, values := createBridgeDomainWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listBridgeDomainQuery)
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
