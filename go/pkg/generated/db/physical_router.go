package db

// physical_router

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertPhysicalRouterQuery = "insert into `physical_router` (`owner_access`,`global_access`,`share`,`owner`,`username`,`password`,`physical_router_product_name`,`physical_router_lldp`,`physical_router_snmp`,`physical_router_dataplane_ip`,`display_name`,`physical_router_management_ip`,`v3_engine_id`,`v3_authentication_password`,`v3_engine_time`,`timeout`,`v3_privacy_password`,`retries`,`local_port`,`v3_privacy_protocol`,`v3_context`,`version`,`v3_security_level`,`v3_context_engine_id`,`v3_security_engine_id`,`v3_security_name`,`v3_engine_boots`,`v3_authentication_protocol`,`v2_community`,`physical_router_role`,`physical_router_vendor_name`,`physical_router_image_uri`,`server_ip`,`resource`,`server_port`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`physical_router_vnc_managed`,`physical_router_loopback_ip`,`service_port`,`key_value_pair`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalRouterQuery = "update `physical_router` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`username` = ?,`password` = ?,`physical_router_product_name` = ?,`physical_router_lldp` = ?,`physical_router_snmp` = ?,`physical_router_dataplane_ip` = ?,`display_name` = ?,`physical_router_management_ip` = ?,`v3_engine_id` = ?,`v3_authentication_password` = ?,`v3_engine_time` = ?,`timeout` = ?,`v3_privacy_password` = ?,`retries` = ?,`local_port` = ?,`v3_privacy_protocol` = ?,`v3_context` = ?,`version` = ?,`v3_security_level` = ?,`v3_context_engine_id` = ?,`v3_security_engine_id` = ?,`v3_security_name` = ?,`v3_engine_boots` = ?,`v3_authentication_protocol` = ?,`v2_community` = ?,`physical_router_role` = ?,`physical_router_vendor_name` = ?,`physical_router_image_uri` = ?,`server_ip` = ?,`resource` = ?,`server_port` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`physical_router_vnc_managed` = ?,`physical_router_loopback_ip` = ?,`service_port` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?;"
const deletePhysicalRouterQuery = "delete from `physical_router` where uuid = ?"
const listPhysicalRouterQuery = "select `owner_access`,`global_access`,`share`,`owner`,`username`,`password`,`physical_router_product_name`,`physical_router_lldp`,`physical_router_snmp`,`physical_router_dataplane_ip`,`display_name`,`physical_router_management_ip`,`v3_engine_id`,`v3_authentication_password`,`v3_engine_time`,`timeout`,`v3_privacy_password`,`retries`,`local_port`,`v3_privacy_protocol`,`v3_context`,`version`,`v3_security_level`,`v3_context_engine_id`,`v3_security_engine_id`,`v3_security_name`,`v3_engine_boots`,`v3_authentication_protocol`,`v2_community`,`physical_router_role`,`physical_router_vendor_name`,`physical_router_image_uri`,`server_ip`,`resource`,`server_port`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`physical_router_vnc_managed`,`physical_router_loopback_ip`,`service_port`,`key_value_pair`,`uuid`,`fq_name` from `physical_router`"
const showPhysicalRouterQuery = "select `owner_access`,`global_access`,`share`,`owner`,`username`,`password`,`physical_router_product_name`,`physical_router_lldp`,`physical_router_snmp`,`physical_router_dataplane_ip`,`display_name`,`physical_router_management_ip`,`v3_engine_id`,`v3_authentication_password`,`v3_engine_time`,`timeout`,`v3_privacy_password`,`retries`,`local_port`,`v3_privacy_protocol`,`v3_context`,`version`,`v3_security_level`,`v3_context_engine_id`,`v3_security_engine_id`,`v3_security_name`,`v3_engine_boots`,`v3_authentication_protocol`,`v2_community`,`physical_router_role`,`physical_router_vendor_name`,`physical_router_image_uri`,`server_ip`,`resource`,`server_port`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`physical_router_vnc_managed`,`physical_router_loopback_ip`,`service_port`,`key_value_pair`,`uuid`,`fq_name` from `physical_router` where uuid = ?"

func CreatePhysicalRouter(tx *sql.Tx, model *models.PhysicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.PhysicalRouterUserCredentials.Username),
		string(model.PhysicalRouterUserCredentials.Password),
		string(model.PhysicalRouterProductName),
		bool(model.PhysicalRouterLLDP),
		bool(model.PhysicalRouterSNMP),
		string(model.PhysicalRouterDataplaneIP),
		string(model.DisplayName),
		string(model.PhysicalRouterManagementIP),
		string(model.PhysicalRouterSNMPCredentials.V3EngineID),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationPassword),
		int(model.PhysicalRouterSNMPCredentials.V3EngineTime),
		int(model.PhysicalRouterSNMPCredentials.Timeout),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyPassword),
		int(model.PhysicalRouterSNMPCredentials.Retries),
		int(model.PhysicalRouterSNMPCredentials.LocalPort),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyProtocol),
		string(model.PhysicalRouterSNMPCredentials.V3Context),
		int(model.PhysicalRouterSNMPCredentials.Version),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityLevel),
		string(model.PhysicalRouterSNMPCredentials.V3ContextEngineID),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityEngineID),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityName),
		int(model.PhysicalRouterSNMPCredentials.V3EngineBoots),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol),
		string(model.PhysicalRouterSNMPCredentials.V2Community),
		string(model.PhysicalRouterRole),
		string(model.PhysicalRouterVendorName),
		string(model.PhysicalRouterImageURI),
		string(model.TelemetryInfo.ServerIP),
		util.MustJSON(model.TelemetryInfo.Resource),
		int(model.TelemetryInfo.ServerPort),
		bool(model.IDPerms.UserVisible),
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
		bool(model.PhysicalRouterVNCManaged),
		string(model.PhysicalRouterLoopbackIP),
		util.MustJSON(model.PhysicalRouterJunosServicePorts.ServicePort),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanPhysicalRouter(rows *sql.Rows) (*models.PhysicalRouter, error) {
	m := models.MakePhysicalRouter()

	var jsonPerms2Share string

	var jsonTelemetryInfoResource string

	var jsonPhysicalRouterJunosServicePortsServicePort string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.PhysicalRouterUserCredentials.Username,
		&m.PhysicalRouterUserCredentials.Password,
		&m.PhysicalRouterProductName,
		&m.PhysicalRouterLLDP,
		&m.PhysicalRouterSNMP,
		&m.PhysicalRouterDataplaneIP,
		&m.DisplayName,
		&m.PhysicalRouterManagementIP,
		&m.PhysicalRouterSNMPCredentials.V3EngineID,
		&m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword,
		&m.PhysicalRouterSNMPCredentials.V3EngineTime,
		&m.PhysicalRouterSNMPCredentials.Timeout,
		&m.PhysicalRouterSNMPCredentials.V3PrivacyPassword,
		&m.PhysicalRouterSNMPCredentials.Retries,
		&m.PhysicalRouterSNMPCredentials.LocalPort,
		&m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol,
		&m.PhysicalRouterSNMPCredentials.V3Context,
		&m.PhysicalRouterSNMPCredentials.Version,
		&m.PhysicalRouterSNMPCredentials.V3SecurityLevel,
		&m.PhysicalRouterSNMPCredentials.V3ContextEngineID,
		&m.PhysicalRouterSNMPCredentials.V3SecurityEngineID,
		&m.PhysicalRouterSNMPCredentials.V3SecurityName,
		&m.PhysicalRouterSNMPCredentials.V3EngineBoots,
		&m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol,
		&m.PhysicalRouterSNMPCredentials.V2Community,
		&m.PhysicalRouterRole,
		&m.PhysicalRouterVendorName,
		&m.PhysicalRouterImageURI,
		&m.TelemetryInfo.ServerIP,
		&jsonTelemetryInfoResource,
		&m.TelemetryInfo.ServerPort,
		&m.IDPerms.UserVisible,
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
		&m.PhysicalRouterVNCManaged,
		&m.PhysicalRouterLoopbackIP,
		&jsonPhysicalRouterJunosServicePortsServicePort,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonTelemetryInfoResource), &m.TelemetryInfo.Resource)

	json.Unmarshal([]byte(jsonPhysicalRouterJunosServicePortsServicePort), &m.PhysicalRouterJunosServicePorts.ServicePort)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createPhysicalRouterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["username"]; ok {
		results = append(results, "username = ?")
		values = append(values, value)
	}

	if value, ok := where["password"]; ok {
		results = append(results, "password = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_product_name"]; ok {
		results = append(results, "physical_router_product_name = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_dataplane_ip"]; ok {
		results = append(results, "physical_router_dataplane_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_management_ip"]; ok {
		results = append(results, "physical_router_management_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_engine_id"]; ok {
		results = append(results, "v3_engine_id = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_authentication_password"]; ok {
		results = append(results, "v3_authentication_password = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_privacy_password"]; ok {
		results = append(results, "v3_privacy_password = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_privacy_protocol"]; ok {
		results = append(results, "v3_privacy_protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_context"]; ok {
		results = append(results, "v3_context = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_security_level"]; ok {
		results = append(results, "v3_security_level = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_context_engine_id"]; ok {
		results = append(results, "v3_context_engine_id = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_security_engine_id"]; ok {
		results = append(results, "v3_security_engine_id = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_security_name"]; ok {
		results = append(results, "v3_security_name = ?")
		values = append(values, value)
	}

	if value, ok := where["v3_authentication_protocol"]; ok {
		results = append(results, "v3_authentication_protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["v2_community"]; ok {
		results = append(results, "v2_community = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_role"]; ok {
		results = append(results, "physical_router_role = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_vendor_name"]; ok {
		results = append(results, "physical_router_vendor_name = ?")
		values = append(values, value)
	}

	if value, ok := where["physical_router_image_uri"]; ok {
		results = append(results, "physical_router_image_uri = ?")
		values = append(values, value)
	}

	if value, ok := where["server_ip"]; ok {
		results = append(results, "server_ip = ?")
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

	if value, ok := where["physical_router_loopback_ip"]; ok {
		results = append(results, "physical_router_loopback_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListPhysicalRouter(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.PhysicalRouter, error) {
	result := models.MakePhysicalRouterSlice()
	whereQuery, values := createPhysicalRouterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listPhysicalRouterQuery)
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
