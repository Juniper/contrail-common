package db
// physical_router

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertPhysicalRouterQuery = "insert into `physical_router` (`physical_router_vendor_name`,`physical_router_vnc_managed`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`uuid`,`username`,`password`,`physical_router_lldp`,`resource`,`server_port`,`server_ip`,`fq_name`,`physical_router_management_ip`,`physical_router_loopback_ip`,`physical_router_image_uri`,`physical_router_dataplane_ip`,`service_port`,`display_name`,`key_value_pair`,`physical_router_role`,`physical_router_product_name`,`physical_router_snmp`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`v3_authentication_protocol`,`v3_security_engine_id`,`retries`,`v3_engine_time`,`v2_community`,`v3_engine_boots`,`v3_security_name`,`v3_security_level`,`v3_authentication_password`,`v3_context_engine_id`,`timeout`,`local_port`,`v3_privacy_protocol`,`v3_privacy_password`,`version`,`v3_context`,`v3_engine_id`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalRouterQuery = "update `physical_router` set `physical_router_vendor_name` = ?,`physical_router_vnc_managed` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`uuid` = ?,`username` = ?,`password` = ?,`physical_router_lldp` = ?,`resource` = ?,`server_port` = ?,`server_ip` = ?,`fq_name` = ?,`physical_router_management_ip` = ?,`physical_router_loopback_ip` = ?,`physical_router_image_uri` = ?,`physical_router_dataplane_ip` = ?,`service_port` = ?,`display_name` = ?,`key_value_pair` = ?,`physical_router_role` = ?,`physical_router_product_name` = ?,`physical_router_snmp` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`v3_authentication_protocol` = ?,`v3_security_engine_id` = ?,`retries` = ?,`v3_engine_time` = ?,`v2_community` = ?,`v3_engine_boots` = ?,`v3_security_name` = ?,`v3_security_level` = ?,`v3_authentication_password` = ?,`v3_context_engine_id` = ?,`timeout` = ?,`local_port` = ?,`v3_privacy_protocol` = ?,`v3_privacy_password` = ?,`version` = ?,`v3_context` = ?,`v3_engine_id` = ?;"
const deletePhysicalRouterQuery = "delete from `physical_router`"
const selectPhysicalRouterQuery = "select `physical_router_vendor_name`,`physical_router_vnc_managed`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`uuid`,`username`,`password`,`physical_router_lldp`,`resource`,`server_port`,`server_ip`,`fq_name`,`physical_router_management_ip`,`physical_router_loopback_ip`,`physical_router_image_uri`,`physical_router_dataplane_ip`,`service_port`,`display_name`,`key_value_pair`,`physical_router_role`,`physical_router_product_name`,`physical_router_snmp`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`v3_authentication_protocol`,`v3_security_engine_id`,`retries`,`v3_engine_time`,`v2_community`,`v3_engine_boots`,`v3_security_name`,`v3_security_level`,`v3_authentication_password`,`v3_context_engine_id`,`timeout`,`local_port`,`v3_privacy_protocol`,`v3_privacy_password`,`version`,`v3_context`,`v3_engine_id` from `physical_router`"

func CreatePhysicalRouter(tx *sql.Tx, model *models.PhysicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalRouterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.PhysicalRouterVendorName,
    model.PhysicalRouterVNCManaged,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.UUID,
    model.PhysicalRouterUserCredentials.Username,
    model.PhysicalRouterUserCredentials.Password,
    model.PhysicalRouterLLDP,
    model.TelemetryInfo.Resource,
    model.TelemetryInfo.ServerPort,
    model.TelemetryInfo.ServerIP,
    model.FQName,
    model.PhysicalRouterManagementIP,
    model.PhysicalRouterLoopbackIP,
    model.PhysicalRouterImageURI,
    model.PhysicalRouterDataplaneIP,
    model.PhysicalRouterJunosServicePorts.ServicePort,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.PhysicalRouterRole,
    model.PhysicalRouterProductName,
    model.PhysicalRouterSNMP,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol,
    model.PhysicalRouterSNMPCredentials.V3SecurityEngineID,
    model.PhysicalRouterSNMPCredentials.Retries,
    model.PhysicalRouterSNMPCredentials.V3EngineTime,
    model.PhysicalRouterSNMPCredentials.V2Community,
    model.PhysicalRouterSNMPCredentials.V3EngineBoots,
    model.PhysicalRouterSNMPCredentials.V3SecurityName,
    model.PhysicalRouterSNMPCredentials.V3SecurityLevel,
    model.PhysicalRouterSNMPCredentials.V3AuthenticationPassword,
    model.PhysicalRouterSNMPCredentials.V3ContextEngineID,
    model.PhysicalRouterSNMPCredentials.Timeout,
    model.PhysicalRouterSNMPCredentials.LocalPort,
    model.PhysicalRouterSNMPCredentials.V3PrivacyProtocol,
    model.PhysicalRouterSNMPCredentials.V3PrivacyPassword,
    model.PhysicalRouterSNMPCredentials.Version,
    model.PhysicalRouterSNMPCredentials.V3Context,
    model.PhysicalRouterSNMPCredentials.V3EngineID)
    return err
}

func ListPhysicalRouter(tx *sql.Tx) ([]*models.PhysicalRouter, error) {
    result := models.MakePhysicalRouterSlice()
    rows, err := tx.Query(selectPhysicalRouterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakePhysicalRouter()
            if err := rows.Scan(&m.PhysicalRouterVendorName,
                &m.PhysicalRouterVNCManaged,
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
                &m.UUID,
                &m.PhysicalRouterUserCredentials.Username,
                &m.PhysicalRouterUserCredentials.Password,
                &m.PhysicalRouterLLDP,
                &m.TelemetryInfo.Resource,
                &m.TelemetryInfo.ServerPort,
                &m.TelemetryInfo.ServerIP,
                &m.FQName,
                &m.PhysicalRouterManagementIP,
                &m.PhysicalRouterLoopbackIP,
                &m.PhysicalRouterImageURI,
                &m.PhysicalRouterDataplaneIP,
                &m.PhysicalRouterJunosServicePorts.ServicePort,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.PhysicalRouterRole,
                &m.PhysicalRouterProductName,
                &m.PhysicalRouterSNMP,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol,
                &m.PhysicalRouterSNMPCredentials.V3SecurityEngineID,
                &m.PhysicalRouterSNMPCredentials.Retries,
                &m.PhysicalRouterSNMPCredentials.V3EngineTime,
                &m.PhysicalRouterSNMPCredentials.V2Community,
                &m.PhysicalRouterSNMPCredentials.V3EngineBoots,
                &m.PhysicalRouterSNMPCredentials.V3SecurityName,
                &m.PhysicalRouterSNMPCredentials.V3SecurityLevel,
                &m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword,
                &m.PhysicalRouterSNMPCredentials.V3ContextEngineID,
                &m.PhysicalRouterSNMPCredentials.Timeout,
                &m.PhysicalRouterSNMPCredentials.LocalPort,
                &m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol,
                &m.PhysicalRouterSNMPCredentials.V3PrivacyPassword,
                &m.PhysicalRouterSNMPCredentials.Version,
                &m.PhysicalRouterSNMPCredentials.V3Context,
                &m.PhysicalRouterSNMPCredentials.V3EngineID); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowPhysicalRouter(db *sql.DB, id string, model *models.PhysicalRouter) error {
    return nil
}

func UpdatePhysicalRouter(db *sql.DB, id string, model *models.PhysicalRouter) error {
    return nil
}

func DeletePhysicalRouter(db *sql.DB, id string) error {
    return nil
}