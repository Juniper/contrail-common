package db
// global_system_config

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertGlobalSystemConfigQuery = "insert into `global_system_config` (`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`bgp_always_compare_med`,`mac_limit`,`mac_limit_action`,`subnet`,`autonomous_system`,`display_name`,`port_start`,`port_end`,`alarm_enable`,`user_defined_log_statistics`,`key_value_pair`,`uuid`,`plugin_property`,`mac_aging_time`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`config_version`,`ibgp_auto_mesh`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalSystemConfigQuery = "update `global_system_config` set `user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`bgp_always_compare_med` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`subnet` = ?,`autonomous_system` = ?,`display_name` = ?,`port_start` = ?,`port_end` = ?,`alarm_enable` = ?,`user_defined_log_statistics` = ?,`key_value_pair` = ?,`uuid` = ?,`plugin_property` = ?,`mac_aging_time` = ?,`bgp_helper_enable` = ?,`xmpp_helper_enable` = ?,`restart_time` = ?,`long_lived_restart_time` = ?,`graceful_restart_parameters_enable` = ?,`end_of_rib_timeout` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`config_version` = ?,`ibgp_auto_mesh` = ?,`fq_name` = ?;"
const deleteGlobalSystemConfigQuery = "delete from `global_system_config`"
const selectGlobalSystemConfigQuery = "select `user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`bgp_always_compare_med`,`mac_limit`,`mac_limit_action`,`subnet`,`autonomous_system`,`display_name`,`port_start`,`port_end`,`alarm_enable`,`user_defined_log_statistics`,`key_value_pair`,`uuid`,`plugin_property`,`mac_aging_time`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`config_version`,`ibgp_auto_mesh`,`fq_name` from `global_system_config`"

func CreateGlobalSystemConfig(tx *sql.Tx, model *models.GlobalSystemConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalSystemConfigQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.MacMoveControl.MacMoveTimeWindow,
    model.MacMoveControl.MacMoveLimit,
    model.MacMoveControl.MacMoveLimitAction,
    model.BGPAlwaysCompareMed,
    model.MacLimitControl.MacLimit,
    model.MacLimitControl.MacLimitAction,
    model.IPFabricSubnets.Subnet,
    model.AutonomousSystem,
    model.DisplayName,
    model.BgpaasParameters.PortStart,
    model.BgpaasParameters.PortEnd,
    model.AlarmEnable,
    model.UserDefinedLogStatistics,
    model.Annotations.KeyValuePair,
    model.UUID,
    model.PluginTuning.PluginProperty,
    model.MacAgingTime,
    model.GracefulRestartParameters.BGPHelperEnable,
    model.GracefulRestartParameters.XMPPHelperEnable,
    model.GracefulRestartParameters.RestartTime,
    model.GracefulRestartParameters.LongLivedRestartTime,
    model.GracefulRestartParameters.Enable,
    model.GracefulRestartParameters.EndOfRibTimeout,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.ConfigVersion,
    model.IbgpAutoMesh,
    model.FQName)
    return err
}

func ListGlobalSystemConfig(tx *sql.Tx) ([]*models.GlobalSystemConfig, error) {
    result := models.MakeGlobalSystemConfigSlice()
    rows, err := tx.Query(selectGlobalSystemConfigQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeGlobalSystemConfig()
            if err := rows.Scan(&m.IDPerms.UserVisible,
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
                &m.MacMoveControl.MacMoveTimeWindow,
                &m.MacMoveControl.MacMoveLimit,
                &m.MacMoveControl.MacMoveLimitAction,
                &m.BGPAlwaysCompareMed,
                &m.MacLimitControl.MacLimit,
                &m.MacLimitControl.MacLimitAction,
                &m.IPFabricSubnets.Subnet,
                &m.AutonomousSystem,
                &m.DisplayName,
                &m.BgpaasParameters.PortStart,
                &m.BgpaasParameters.PortEnd,
                &m.AlarmEnable,
                &m.UserDefinedLogStatistics,
                &m.Annotations.KeyValuePair,
                &m.UUID,
                &m.PluginTuning.PluginProperty,
                &m.MacAgingTime,
                &m.GracefulRestartParameters.BGPHelperEnable,
                &m.GracefulRestartParameters.XMPPHelperEnable,
                &m.GracefulRestartParameters.RestartTime,
                &m.GracefulRestartParameters.LongLivedRestartTime,
                &m.GracefulRestartParameters.Enable,
                &m.GracefulRestartParameters.EndOfRibTimeout,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.ConfigVersion,
                &m.IbgpAutoMesh,
                &m.FQName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowGlobalSystemConfig(db *sql.DB, id string, model *models.GlobalSystemConfig) error {
    return nil
}

func UpdateGlobalSystemConfig(db *sql.DB, id string, model *models.GlobalSystemConfig) error {
    return nil
}

func DeleteGlobalSystemConfig(db *sql.DB, id string) error {
    return nil
}