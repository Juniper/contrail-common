package db
// global_vrouter_config

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`fq_name`,`forwarding_mode`,`linklocal_service_entry`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`flow_aging_timeout`,`encapsulation`,`vxlan_network_identifier_mode`,`enable_security_logging`,`uuid`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`flow_export_rate`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalVrouterConfigQuery = "update `global_vrouter_config` set `last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`fq_name` = ?,`forwarding_mode` = ?,`linklocal_service_entry` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`flow_aging_timeout` = ?,`encapsulation` = ?,`vxlan_network_identifier_mode` = ?,`enable_security_logging` = ?,`uuid` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`flow_export_rate` = ?,`display_name` = ?;"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config`"
const selectGlobalVrouterConfigQuery = "select `last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`fq_name`,`forwarding_mode`,`linklocal_service_entry`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`flow_aging_timeout`,`encapsulation`,`vxlan_network_identifier_mode`,`enable_security_logging`,`uuid`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`destination_ip`,`flow_export_rate`,`display_name` from `global_vrouter_config`"

func CreateGlobalVrouterConfig(tx *sql.Tx, model *models.GlobalVrouterConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.FQName,
    model.ForwardingMode,
    model.LinklocalServices.LinklocalServiceEntry,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.FlowAgingTimeoutList.FlowAgingTimeout,
    model.EncapsulationPriorities.Encapsulation,
    model.VxlanNetworkIdentifierMode,
    model.EnableSecurityLogging,
    model.UUID,
    model.EcmpHashingIncludeFields.IPProtocol,
    model.EcmpHashingIncludeFields.SourceIP,
    model.EcmpHashingIncludeFields.HashingConfigured,
    model.EcmpHashingIncludeFields.SourcePort,
    model.EcmpHashingIncludeFields.DestinationPort,
    model.EcmpHashingIncludeFields.DestinationIP,
    model.FlowExportRate,
    model.DisplayName)
    return err
}

func ListGlobalVrouterConfig(tx *sql.Tx) ([]*models.GlobalVrouterConfig, error) {
    result := models.MakeGlobalVrouterConfigSlice()
    rows, err := tx.Query(selectGlobalVrouterConfigQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeGlobalVrouterConfig()
            if err := rows.Scan(&m.IDPerms.LastModified,
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
                &m.FQName,
                &m.ForwardingMode,
                &m.LinklocalServices.LinklocalServiceEntry,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.FlowAgingTimeoutList.FlowAgingTimeout,
                &m.EncapsulationPriorities.Encapsulation,
                &m.VxlanNetworkIdentifierMode,
                &m.EnableSecurityLogging,
                &m.UUID,
                &m.EcmpHashingIncludeFields.IPProtocol,
                &m.EcmpHashingIncludeFields.SourceIP,
                &m.EcmpHashingIncludeFields.HashingConfigured,
                &m.EcmpHashingIncludeFields.SourcePort,
                &m.EcmpHashingIncludeFields.DestinationPort,
                &m.EcmpHashingIncludeFields.DestinationIP,
                &m.FlowExportRate,
                &m.DisplayName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowGlobalVrouterConfig(db *sql.DB, id string, model *models.GlobalVrouterConfig) error {
    return nil
}

func UpdateGlobalVrouterConfig(db *sql.DB, id string, model *models.GlobalVrouterConfig) error {
    return nil
}

func DeleteGlobalVrouterConfig(db *sql.DB, id string) error {
    return nil
}