package db

// global_vrouter_config

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`forwarding_mode`,`vxlan_network_identifier_mode`,`enable_security_logging`,`fq_name`,`display_name`,`linklocal_service_entry`,`encapsulation`,`owner_access`,`global_access`,`share`,`owner`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`flow_aging_timeout`,`flow_export_rate`,`uuid`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalVrouterConfigQuery = "update `global_vrouter_config` set `forwarding_mode` = ?,`vxlan_network_identifier_mode` = ?,`enable_security_logging` = ?,`fq_name` = ?,`display_name` = ?,`linklocal_service_entry` = ?,`encapsulation` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`flow_aging_timeout` = ?,`flow_export_rate` = ?,`uuid` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`key_value_pair` = ?;"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config` where uuid = ?"
const listGlobalVrouterConfigQuery = "select `forwarding_mode`,`vxlan_network_identifier_mode`,`enable_security_logging`,`fq_name`,`display_name`,`linklocal_service_entry`,`encapsulation`,`owner_access`,`global_access`,`share`,`owner`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`flow_aging_timeout`,`flow_export_rate`,`uuid`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair` from `global_vrouter_config`"
const showGlobalVrouterConfigQuery = "select `forwarding_mode`,`vxlan_network_identifier_mode`,`enable_security_logging`,`fq_name`,`display_name`,`linklocal_service_entry`,`encapsulation`,`owner_access`,`global_access`,`share`,`owner`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`flow_aging_timeout`,`flow_export_rate`,`uuid`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair` from `global_vrouter_config` where uuid = ?"

func CreateGlobalVrouterConfig(tx *sql.Tx, model *models.GlobalVrouterConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ForwardingMode),
		string(model.VxlanNetworkIdentifierMode),
		bool(model.EnableSecurityLogging),
		util.MustJSON(model.FQName),
		string(model.DisplayName),
		util.MustJSON(model.LinklocalServices.LinklocalServiceEntry),
		util.MustJSON(model.EncapsulationPriorities.Encapsulation),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		util.MustJSON(model.FlowAgingTimeoutList.FlowAgingTimeout),
		int(model.FlowExportRate),
		string(model.UUID),
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
		bool(model.IDPerms.UserVisible),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanGlobalVrouterConfig(rows *sql.Rows) (*models.GlobalVrouterConfig, error) {
	m := models.MakeGlobalVrouterConfig()

	var jsonFQName string

	var jsonLinklocalServicesLinklocalServiceEntry string

	var jsonEncapsulationPrioritiesEncapsulation string

	var jsonPerms2Share string

	var jsonFlowAgingTimeoutListFlowAgingTimeout string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ForwardingMode,
		&m.VxlanNetworkIdentifierMode,
		&m.EnableSecurityLogging,
		&jsonFQName,
		&m.DisplayName,
		&jsonLinklocalServicesLinklocalServiceEntry,
		&jsonEncapsulationPrioritiesEncapsulation,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&jsonFlowAgingTimeoutListFlowAgingTimeout,
		&m.FlowExportRate,
		&m.UUID,
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
		&m.IDPerms.UserVisible,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonLinklocalServicesLinklocalServiceEntry), &m.LinklocalServices.LinklocalServiceEntry)

	json.Unmarshal([]byte(jsonEncapsulationPrioritiesEncapsulation), &m.EncapsulationPriorities.Encapsulation)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFlowAgingTimeoutListFlowAgingTimeout), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createGlobalVrouterConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["forwarding_mode"]; ok {
		results = append(results, "forwarding_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["vxlan_network_identifier_mode"]; ok {
		results = append(results, "vxlan_network_identifier_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListGlobalVrouterConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.GlobalVrouterConfig, error) {
	result := models.MakeGlobalVrouterConfigSlice()
	whereQuery, values := createGlobalVrouterConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listGlobalVrouterConfigQuery)
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
		m, _ := scanGlobalVrouterConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowGlobalVrouterConfig(tx *sql.Tx, uuid string) (*models.GlobalVrouterConfig, error) {
	rows, err := tx.Query(showGlobalVrouterConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanGlobalVrouterConfig(rows)
	}
	return nil, nil
}

func UpdateGlobalVrouterConfig(tx *sql.Tx, uuid string, model *models.GlobalVrouterConfig) error {
	return nil
}

func DeleteGlobalVrouterConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalVrouterConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
