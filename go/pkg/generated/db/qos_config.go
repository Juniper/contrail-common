package db

// qos_config

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertQosConfigQuery = "insert into `qos_config` (`qos_config_type`,`mpls_exp_entries`,`fq_name`,`uuid`,`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`vlan_priority_entries`,`default_forwarding_class_id`,`dscp_entries`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosConfigQuery = "update `qos_config` set `qos_config_type` = ?,`mpls_exp_entries` = ?,`fq_name` = ?,`uuid` = ?,`display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`vlan_priority_entries` = ?,`default_forwarding_class_id` = ?,`dscp_entries` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?;"
const deleteQosConfigQuery = "delete from `qos_config` where uuid = ?"
const listQosConfigQuery = "select `qos_config_type`,`mpls_exp_entries`,`fq_name`,`uuid`,`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`vlan_priority_entries`,`default_forwarding_class_id`,`dscp_entries`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description` from `qos_config`"
const showQosConfigQuery = "select `qos_config_type`,`mpls_exp_entries`,`fq_name`,`uuid`,`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`vlan_priority_entries`,`default_forwarding_class_id`,`dscp_entries`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description` from `qos_config` where uuid = ?"

func CreateQosConfig(tx *sql.Tx, model *models.QosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.QosConfigType),
		util.MustJSON(model.MPLSExpEntries),
		util.MustJSON(model.FQName),
		string(model.UUID),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		util.MustJSON(model.VlanPriorityEntries),
		int(model.DefaultForwardingClassID),
		util.MustJSON(model.DSCPEntries),
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
		string(model.IDPerms.Description))
	return err
}

func scanQosConfig(rows *sql.Rows) (*models.QosConfig, error) {
	m := models.MakeQosConfig()

	var jsonMPLSExpEntries string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonVlanPriorityEntries string

	var jsonDSCPEntries string

	if err := rows.Scan(&m.QosConfigType,
		&jsonMPLSExpEntries,
		&jsonFQName,
		&m.UUID,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonVlanPriorityEntries,
		&m.DefaultForwardingClassID,
		&jsonDSCPEntries,
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
		&m.IDPerms.Description); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonMPLSExpEntries), &m.MPLSExpEntries)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonVlanPriorityEntries), &m.VlanPriorityEntries)

	json.Unmarshal([]byte(jsonDSCPEntries), &m.DSCPEntries)

	return m, nil
}

func ListQosConfig(tx *sql.Tx) ([]*models.QosConfig, error) {
	result := models.MakeQosConfigSlice()
	rows, err := tx.Query(listQosConfigQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanQosConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowQosConfig(tx *sql.Tx, uuid string) (*models.QosConfig, error) {
	rows, err := tx.Query(showQosConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanQosConfig(rows)
	}
	return nil, nil
}

func UpdateQosConfig(tx *sql.Tx, uuid string, model *models.QosConfig) error {
	return nil
}

func DeleteQosConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
