package db

// physical_router

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertPhysicalRouterQuery = "insert into `physical_router` (`username`,`password`,`physical_router_loopback_ip`,`physical_router_image_uri`,`v3_security_engine_id`,`v3_engine_time`,`v3_authentication_password`,`v3_context_engine_id`,`version`,`v3_privacy_protocol`,`v3_privacy_password`,`local_port`,`v3_engine_id`,`v3_security_level`,`v2_community`,`v3_security_name`,`v3_engine_boots`,`timeout`,`retries`,`v3_context`,`v3_authentication_protocol`,`physical_router_role`,`physical_router_snmp`,`physical_router_dataplane_ip`,`display_name`,`physical_router_management_ip`,`physical_router_vnc_managed`,`physical_router_lldp`,`resource`,`server_port`,`server_ip`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`physical_router_vendor_name`,`physical_router_product_name`,`service_port`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalRouterQuery = "update `physical_router` set `username` = ?,`password` = ?,`physical_router_loopback_ip` = ?,`physical_router_image_uri` = ?,`v3_security_engine_id` = ?,`v3_engine_time` = ?,`v3_authentication_password` = ?,`v3_context_engine_id` = ?,`version` = ?,`v3_privacy_protocol` = ?,`v3_privacy_password` = ?,`local_port` = ?,`v3_engine_id` = ?,`v3_security_level` = ?,`v2_community` = ?,`v3_security_name` = ?,`v3_engine_boots` = ?,`timeout` = ?,`retries` = ?,`v3_context` = ?,`v3_authentication_protocol` = ?,`physical_router_role` = ?,`physical_router_snmp` = ?,`physical_router_dataplane_ip` = ?,`display_name` = ?,`physical_router_management_ip` = ?,`physical_router_vnc_managed` = ?,`physical_router_lldp` = ?,`resource` = ?,`server_port` = ?,`server_ip` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`physical_router_vendor_name` = ?,`physical_router_product_name` = ?,`service_port` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deletePhysicalRouterQuery = "delete from `physical_router` where uuid = ?"
const listPhysicalRouterQuery = "select `username`,`password`,`physical_router_loopback_ip`,`physical_router_image_uri`,`v3_security_engine_id`,`v3_engine_time`,`v3_authentication_password`,`v3_context_engine_id`,`version`,`v3_privacy_protocol`,`v3_privacy_password`,`local_port`,`v3_engine_id`,`v3_security_level`,`v2_community`,`v3_security_name`,`v3_engine_boots`,`timeout`,`retries`,`v3_context`,`v3_authentication_protocol`,`physical_router_role`,`physical_router_snmp`,`physical_router_dataplane_ip`,`display_name`,`physical_router_management_ip`,`physical_router_vnc_managed`,`physical_router_lldp`,`resource`,`server_port`,`server_ip`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`physical_router_vendor_name`,`physical_router_product_name`,`service_port`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible` from `physical_router`"
const showPhysicalRouterQuery = "select `username`,`password`,`physical_router_loopback_ip`,`physical_router_image_uri`,`v3_security_engine_id`,`v3_engine_time`,`v3_authentication_password`,`v3_context_engine_id`,`version`,`v3_privacy_protocol`,`v3_privacy_password`,`local_port`,`v3_engine_id`,`v3_security_level`,`v2_community`,`v3_security_name`,`v3_engine_boots`,`timeout`,`retries`,`v3_context`,`v3_authentication_protocol`,`physical_router_role`,`physical_router_snmp`,`physical_router_dataplane_ip`,`display_name`,`physical_router_management_ip`,`physical_router_vnc_managed`,`physical_router_lldp`,`resource`,`server_port`,`server_ip`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`physical_router_vendor_name`,`physical_router_product_name`,`service_port`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible` from `physical_router` where uuid = ?"

func CreatePhysicalRouter(tx *sql.Tx, model *models.PhysicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.PhysicalRouterUserCredentials.Username),
		string(model.PhysicalRouterUserCredentials.Password),
		string(model.PhysicalRouterLoopbackIP),
		string(model.PhysicalRouterImageURI),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityEngineID),
		int(model.PhysicalRouterSNMPCredentials.V3EngineTime),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationPassword),
		string(model.PhysicalRouterSNMPCredentials.V3ContextEngineID),
		int(model.PhysicalRouterSNMPCredentials.Version),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyProtocol),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyPassword),
		int(model.PhysicalRouterSNMPCredentials.LocalPort),
		string(model.PhysicalRouterSNMPCredentials.V3EngineID),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityLevel),
		string(model.PhysicalRouterSNMPCredentials.V2Community),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityName),
		int(model.PhysicalRouterSNMPCredentials.V3EngineBoots),
		int(model.PhysicalRouterSNMPCredentials.Timeout),
		int(model.PhysicalRouterSNMPCredentials.Retries),
		string(model.PhysicalRouterSNMPCredentials.V3Context),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol),
		string(model.PhysicalRouterRole),
		bool(model.PhysicalRouterSNMP),
		string(model.PhysicalRouterDataplaneIP),
		string(model.DisplayName),
		string(model.PhysicalRouterManagementIP),
		bool(model.PhysicalRouterVNCManaged),
		bool(model.PhysicalRouterLLDP),
		util.MustJSON(model.TelemetryInfo.Resource),
		int(model.TelemetryInfo.ServerPort),
		string(model.TelemetryInfo.ServerIP),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.PhysicalRouterVendorName),
		string(model.PhysicalRouterProductName),
		util.MustJSON(model.PhysicalRouterJunosServicePorts.ServicePort),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		bool(model.IDPerms.UserVisible))
	return err
}

func scanPhysicalRouter(rows *sql.Rows) (*models.PhysicalRouter, error) {
	m := models.MakePhysicalRouter()

	var jsonTelemetryInfoResource string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonPhysicalRouterJunosServicePortsServicePort string

	var jsonFQName string

	if err := rows.Scan(&m.PhysicalRouterUserCredentials.Username,
		&m.PhysicalRouterUserCredentials.Password,
		&m.PhysicalRouterLoopbackIP,
		&m.PhysicalRouterImageURI,
		&m.PhysicalRouterSNMPCredentials.V3SecurityEngineID,
		&m.PhysicalRouterSNMPCredentials.V3EngineTime,
		&m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword,
		&m.PhysicalRouterSNMPCredentials.V3ContextEngineID,
		&m.PhysicalRouterSNMPCredentials.Version,
		&m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol,
		&m.PhysicalRouterSNMPCredentials.V3PrivacyPassword,
		&m.PhysicalRouterSNMPCredentials.LocalPort,
		&m.PhysicalRouterSNMPCredentials.V3EngineID,
		&m.PhysicalRouterSNMPCredentials.V3SecurityLevel,
		&m.PhysicalRouterSNMPCredentials.V2Community,
		&m.PhysicalRouterSNMPCredentials.V3SecurityName,
		&m.PhysicalRouterSNMPCredentials.V3EngineBoots,
		&m.PhysicalRouterSNMPCredentials.Timeout,
		&m.PhysicalRouterSNMPCredentials.Retries,
		&m.PhysicalRouterSNMPCredentials.V3Context,
		&m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol,
		&m.PhysicalRouterRole,
		&m.PhysicalRouterSNMP,
		&m.PhysicalRouterDataplaneIP,
		&m.DisplayName,
		&m.PhysicalRouterManagementIP,
		&m.PhysicalRouterVNCManaged,
		&m.PhysicalRouterLLDP,
		&jsonTelemetryInfoResource,
		&m.TelemetryInfo.ServerPort,
		&m.TelemetryInfo.ServerIP,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.PhysicalRouterVendorName,
		&m.PhysicalRouterProductName,
		&jsonPhysicalRouterJunosServicePortsServicePort,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.UserVisible); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonTelemetryInfoResource), &m.TelemetryInfo.Resource)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonPhysicalRouterJunosServicePortsServicePort), &m.PhysicalRouterJunosServicePorts.ServicePort)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListPhysicalRouter(tx *sql.Tx) ([]*models.PhysicalRouter, error) {
	result := models.MakePhysicalRouterSlice()
	rows, err := tx.Query(listPhysicalRouterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanPhysicalRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPhysicalRouter(tx *sql.Tx, uuid string) (*models.PhysicalRouter, error) {
	rows, err := tx.Query(showPhysicalRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPhysicalRouter(rows)
	}
	return nil, nil
}

func UpdatePhysicalRouter(tx *sql.Tx, uuid string, model *models.PhysicalRouter) error {
	return nil
}

func DeletePhysicalRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
