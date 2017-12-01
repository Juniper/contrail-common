package db

// global_system_config

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertGlobalSystemConfigQuery = "insert into `global_system_config` (`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`config_version`,`alarm_enable`,`mac_aging_time`,`user_defined_log_statistics`,`enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`subnet`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`plugin_property`,`ibgp_auto_mesh`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`id_perms_enable`,`description`,`bgp_always_compare_med`,`autonomous_system`,`mac_limit`,`mac_limit_action`,`display_name`,`uuid`,`port_start`,`port_end`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalSystemConfigQuery = "update `global_system_config` set `key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`config_version` = ?,`alarm_enable` = ?,`mac_aging_time` = ?,`user_defined_log_statistics` = ?,`enable` = ?,`end_of_rib_timeout` = ?,`bgp_helper_enable` = ?,`xmpp_helper_enable` = ?,`restart_time` = ?,`long_lived_restart_time` = ?,`subnet` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`plugin_property` = ?,`ibgp_auto_mesh` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`id_perms_enable` = ?,`description` = ?,`bgp_always_compare_med` = ?,`autonomous_system` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`display_name` = ?,`uuid` = ?,`port_start` = ?,`port_end` = ?;"
const deleteGlobalSystemConfigQuery = "delete from `global_system_config` where uuid = ?"
const listGlobalSystemConfigQuery = "select `key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`config_version`,`alarm_enable`,`mac_aging_time`,`user_defined_log_statistics`,`enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`subnet`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`plugin_property`,`ibgp_auto_mesh`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`id_perms_enable`,`description`,`bgp_always_compare_med`,`autonomous_system`,`mac_limit`,`mac_limit_action`,`display_name`,`uuid`,`port_start`,`port_end` from `global_system_config`"
const showGlobalSystemConfigQuery = "select `key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`config_version`,`alarm_enable`,`mac_aging_time`,`user_defined_log_statistics`,`enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`subnet`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`plugin_property`,`ibgp_auto_mesh`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`id_perms_enable`,`description`,`bgp_always_compare_med`,`autonomous_system`,`mac_limit`,`mac_limit_action`,`display_name`,`uuid`,`port_start`,`port_end` from `global_system_config` where uuid = ?"

func CreateGlobalSystemConfig(tx *sql.Tx, model *models.GlobalSystemConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalSystemConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.ConfigVersion),
		bool(model.AlarmEnable),
		int(model.MacAgingTime),
		util.MustJSON(model.UserDefinedLogStatistics),
		bool(model.GracefulRestartParameters.Enable),
		int(model.GracefulRestartParameters.EndOfRibTimeout),
		bool(model.GracefulRestartParameters.BGPHelperEnable),
		bool(model.GracefulRestartParameters.XMPPHelperEnable),
		int(model.GracefulRestartParameters.RestartTime),
		int(model.GracefulRestartParameters.LongLivedRestartTime),
		util.MustJSON(model.IPFabricSubnets.Subnet),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		util.MustJSON(model.PluginTuning.PluginProperty),
		bool(model.IbgpAutoMesh),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		bool(model.BGPAlwaysCompareMed),
		int(model.AutonomousSystem),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		string(model.DisplayName),
		string(model.UUID),
		int(model.BgpaasParameters.PortStart),
		int(model.BgpaasParameters.PortEnd))
	return err
}

func scanGlobalSystemConfig(rows *sql.Rows) (*models.GlobalSystemConfig, error) {
	m := models.MakeGlobalSystemConfig()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonUserDefinedLogStatistics string

	var jsonIPFabricSubnetsSubnet string

	var jsonPluginTuningPluginProperty string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.ConfigVersion,
		&m.AlarmEnable,
		&m.MacAgingTime,
		&jsonUserDefinedLogStatistics,
		&m.GracefulRestartParameters.Enable,
		&m.GracefulRestartParameters.EndOfRibTimeout,
		&m.GracefulRestartParameters.BGPHelperEnable,
		&m.GracefulRestartParameters.XMPPHelperEnable,
		&m.GracefulRestartParameters.RestartTime,
		&m.GracefulRestartParameters.LongLivedRestartTime,
		&jsonIPFabricSubnetsSubnet,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&jsonPluginTuningPluginProperty,
		&m.IbgpAutoMesh,
		&jsonFQName,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.BGPAlwaysCompareMed,
		&m.AutonomousSystem,
		&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.DisplayName,
		&m.UUID,
		&m.BgpaasParameters.PortStart,
		&m.BgpaasParameters.PortEnd); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonUserDefinedLogStatistics), &m.UserDefinedLogStatistics)

	json.Unmarshal([]byte(jsonIPFabricSubnetsSubnet), &m.IPFabricSubnets.Subnet)

	json.Unmarshal([]byte(jsonPluginTuningPluginProperty), &m.PluginTuning.PluginProperty)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListGlobalSystemConfig(tx *sql.Tx) ([]*models.GlobalSystemConfig, error) {
	result := models.MakeGlobalSystemConfigSlice()
	rows, err := tx.Query(listGlobalSystemConfigQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanGlobalSystemConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowGlobalSystemConfig(tx *sql.Tx, uuid string) (*models.GlobalSystemConfig, error) {
	rows, err := tx.Query(showGlobalSystemConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanGlobalSystemConfig(rows)
	}
	return nil, nil
}

func UpdateGlobalSystemConfig(tx *sql.Tx, uuid string, model *models.GlobalSystemConfig) error {
	return nil
}

func DeleteGlobalSystemConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalSystemConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
