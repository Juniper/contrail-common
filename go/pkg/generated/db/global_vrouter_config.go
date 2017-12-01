package db

// global_vrouter_config

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`linklocal_service_entry`,`share`,`owner`,`owner_access`,`global_access`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`forwarding_mode`,`vxlan_network_identifier_mode`,`flow_aging_timeout`,`flow_export_rate`,`encapsulation`,`enable_security_logging`,`key_value_pair`,`fq_name`,`uuid`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalVrouterConfigQuery = "update `global_vrouter_config` set `linklocal_service_entry` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`ip_protocol` = ?,`forwarding_mode` = ?,`vxlan_network_identifier_mode` = ?,`flow_aging_timeout` = ?,`flow_export_rate` = ?,`encapsulation` = ?,`enable_security_logging` = ?,`key_value_pair` = ?,`fq_name` = ?,`uuid` = ?,`display_name` = ?;"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config` where uuid = ?"
const listGlobalVrouterConfigQuery = "select `linklocal_service_entry`,`share`,`owner`,`owner_access`,`global_access`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`forwarding_mode`,`vxlan_network_identifier_mode`,`flow_aging_timeout`,`flow_export_rate`,`encapsulation`,`enable_security_logging`,`key_value_pair`,`fq_name`,`uuid`,`display_name` from `global_vrouter_config`"
const showGlobalVrouterConfigQuery = "select `linklocal_service_entry`,`share`,`owner`,`owner_access`,`global_access`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`forwarding_mode`,`vxlan_network_identifier_mode`,`flow_aging_timeout`,`flow_export_rate`,`encapsulation`,`enable_security_logging`,`key_value_pair`,`fq_name`,`uuid`,`display_name` from `global_vrouter_config` where uuid = ?"

func CreateGlobalVrouterConfig(tx *sql.Tx, model *models.GlobalVrouterConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.LinklocalServices.LinklocalServiceEntry),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		string(model.ForwardingMode),
		string(model.VxlanNetworkIdentifierMode),
		util.MustJSON(model.FlowAgingTimeoutList.FlowAgingTimeout),
		int(model.FlowExportRate),
		util.MustJSON(model.EncapsulationPriorities.Encapsulation),
		bool(model.EnableSecurityLogging),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.FQName),
		string(model.UUID),
		string(model.DisplayName))
	return err
}

func scanGlobalVrouterConfig(rows *sql.Rows) (*models.GlobalVrouterConfig, error) {
	m := models.MakeGlobalVrouterConfig()

	var jsonLinklocalServicesLinklocalServiceEntry string

	var jsonPerms2Share string

	var jsonFlowAgingTimeoutListFlowAgingTimeout string

	var jsonEncapsulationPrioritiesEncapsulation string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&jsonLinklocalServicesLinklocalServiceEntry,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.IDPerms.LastModified,
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
		&m.EcmpHashingIncludeFields.SourceIP,
		&m.EcmpHashingIncludeFields.HashingConfigured,
		&m.EcmpHashingIncludeFields.SourcePort,
		&m.EcmpHashingIncludeFields.DestinationPort,
		&m.EcmpHashingIncludeFields.DestinationIP,
		&m.EcmpHashingIncludeFields.IPProtocol,
		&m.ForwardingMode,
		&m.VxlanNetworkIdentifierMode,
		&jsonFlowAgingTimeoutListFlowAgingTimeout,
		&m.FlowExportRate,
		&jsonEncapsulationPrioritiesEncapsulation,
		&m.EnableSecurityLogging,
		&jsonAnnotationsKeyValuePair,
		&jsonFQName,
		&m.UUID,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonLinklocalServicesLinklocalServiceEntry), &m.LinklocalServices.LinklocalServiceEntry)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFlowAgingTimeoutListFlowAgingTimeout), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	json.Unmarshal([]byte(jsonEncapsulationPrioritiesEncapsulation), &m.EncapsulationPriorities.Encapsulation)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListGlobalVrouterConfig(tx *sql.Tx) ([]*models.GlobalVrouterConfig, error) {
	result := models.MakeGlobalVrouterConfigSlice()
	rows, err := tx.Query(listGlobalVrouterConfigQuery)
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
