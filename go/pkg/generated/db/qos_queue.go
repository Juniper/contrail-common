package db

// qos_queue

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertQosQueueQuery = "insert into `qos_queue` (`min_bandwidth`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`qos_queue_identifier`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`max_bandwidth`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosQueueQuery = "update `qos_queue` set `min_bandwidth` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`qos_queue_identifier` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`max_bandwidth` = ?;"
const deleteQosQueueQuery = "delete from `qos_queue` where uuid = ?"
const listQosQueueQuery = "select `min_bandwidth`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`qos_queue_identifier`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`max_bandwidth` from `qos_queue`"
const showQosQueueQuery = "select `min_bandwidth`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`qos_queue_identifier`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`max_bandwidth` from `qos_queue` where uuid = ?"

func CreateQosQueue(tx *sql.Tx, model *models.QosQueue) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosQueueQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.MinBandwidth),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.Created),
		string(model.DisplayName),
		int(model.QosQueueIdentifier),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		int(model.MaxBandwidth))
	return err
}

func scanQosQueue(rows *sql.Rows) (*models.QosQueue, error) {
	m := models.MakeQosQueue()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.MinBandwidth,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Created,
		&m.DisplayName,
		&m.QosQueueIdentifier,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.MaxBandwidth); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListQosQueue(tx *sql.Tx) ([]*models.QosQueue, error) {
	result := models.MakeQosQueueSlice()
	rows, err := tx.Query(listQosQueueQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanQosQueue(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowQosQueue(tx *sql.Tx, uuid string) (*models.QosQueue, error) {
	rows, err := tx.Query(showQosQueueQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanQosQueue(rows)
	}
	return nil, nil
}

func UpdateQosQueue(tx *sql.Tx, uuid string, model *models.QosQueue) error {
	return nil
}

func DeleteQosQueue(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteQosQueueQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
