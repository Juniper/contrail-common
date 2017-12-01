package db

// qos_config

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertQosConfigQuery = "insert into `qos_config` (`vlan_priority_entries`,`default_forwarding_class_id`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`qos_config_type`,`mpls_exp_entries`,`fq_name`,`dscp_entries`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosConfigQuery = "update `qos_config` set `vlan_priority_entries` = ?,`default_forwarding_class_id` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`qos_config_type` = ?,`mpls_exp_entries` = ?,`fq_name` = ?,`dscp_entries` = ?,`key_value_pair` = ?;"
const deleteQosConfigQuery = "delete from `qos_config` where uuid = ?"
const listQosConfigQuery = "select `vlan_priority_entries`,`default_forwarding_class_id`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`qos_config_type`,`mpls_exp_entries`,`fq_name`,`dscp_entries`,`key_value_pair` from `qos_config`"
const showQosConfigQuery = "select `vlan_priority_entries`,`default_forwarding_class_id`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`qos_config_type`,`mpls_exp_entries`,`fq_name`,`dscp_entries`,`key_value_pair` from `qos_config` where uuid = ?"

func CreateQosConfig(tx *sql.Tx, model *models.QosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.VlanPriorityEntries),
		int(model.DefaultForwardingClassID),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
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
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		string(model.QosConfigType),
		util.MustJSON(model.MPLSExpEntries),
		util.MustJSON(model.FQName),
		util.MustJSON(model.DSCPEntries),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanQosConfig(rows *sql.Rows) (*models.QosConfig, error) {
	m := models.MakeQosConfig()

	var jsonVlanPriorityEntries string

	var jsonPerms2Share string

	var jsonMPLSExpEntries string

	var jsonFQName string

	var jsonDSCPEntries string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonVlanPriorityEntries,
		&m.DefaultForwardingClassID,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
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
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&m.QosConfigType,
		&jsonMPLSExpEntries,
		&jsonFQName,
		&jsonDSCPEntries,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonVlanPriorityEntries), &m.VlanPriorityEntries)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonMPLSExpEntries), &m.MPLSExpEntries)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonDSCPEntries), &m.DSCPEntries)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createQosConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["qos_config_type"]; ok {
		results = append(results, "qos_config_type = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListQosConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.QosConfig, error) {
	result := models.MakeQosConfigSlice()
	whereQuery, values := createQosConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listQosConfigQuery)
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
