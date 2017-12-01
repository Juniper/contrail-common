package db

// global_system_config

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertGlobalSystemConfigQuery = "insert into `global_system_config` (`ibgp_auto_mesh`,`mac_aging_time`,`user_defined_log_statistics`,`mac_limit_action`,`mac_limit`,`config_version`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`plugin_property`,`fq_name`,`port_end`,`port_start`,`alarm_enable`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`bgp_always_compare_med`,`subnet`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`autonomous_system`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalSystemConfigQuery = "update `global_system_config` set `ibgp_auto_mesh` = ?,`mac_aging_time` = ?,`user_defined_log_statistics` = ?,`mac_limit_action` = ?,`mac_limit` = ?,`config_version` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`plugin_property` = ?,`fq_name` = ?,`port_end` = ?,`port_start` = ?,`alarm_enable` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`bgp_always_compare_med` = ?,`subnet` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`graceful_restart_parameters_enable` = ?,`end_of_rib_timeout` = ?,`bgp_helper_enable` = ?,`xmpp_helper_enable` = ?,`restart_time` = ?,`long_lived_restart_time` = ?,`autonomous_system` = ?;"
const deleteGlobalSystemConfigQuery = "delete from `global_system_config` where uuid = ?"
const listGlobalSystemConfigQuery = "select `ibgp_auto_mesh`,`mac_aging_time`,`user_defined_log_statistics`,`mac_limit_action`,`mac_limit`,`config_version`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`plugin_property`,`fq_name`,`port_end`,`port_start`,`alarm_enable`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`bgp_always_compare_med`,`subnet`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`autonomous_system` from `global_system_config`"
const showGlobalSystemConfigQuery = "select `ibgp_auto_mesh`,`mac_aging_time`,`user_defined_log_statistics`,`mac_limit_action`,`mac_limit`,`config_version`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`plugin_property`,`fq_name`,`port_end`,`port_start`,`alarm_enable`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`bgp_always_compare_med`,`subnet`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`graceful_restart_parameters_enable`,`end_of_rib_timeout`,`bgp_helper_enable`,`xmpp_helper_enable`,`restart_time`,`long_lived_restart_time`,`autonomous_system` from `global_system_config` where uuid = ?"

func CreateGlobalSystemConfig(tx *sql.Tx, model *models.GlobalSystemConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalSystemConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.IbgpAutoMesh),
		int(model.MacAgingTime),
		util.MustJSON(model.UserDefinedLogStatistics),
		string(model.MacLimitControl.MacLimitAction),
		int(model.MacLimitControl.MacLimit),
		string(model.ConfigVersion),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		util.MustJSON(model.PluginTuning.PluginProperty),
		util.MustJSON(model.FQName),
		int(model.BgpaasParameters.PortEnd),
		int(model.BgpaasParameters.PortStart),
		bool(model.AlarmEnable),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		bool(model.BGPAlwaysCompareMed),
		util.MustJSON(model.IPFabricSubnets.Subnet),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		bool(model.GracefulRestartParameters.Enable),
		int(model.GracefulRestartParameters.EndOfRibTimeout),
		bool(model.GracefulRestartParameters.BGPHelperEnable),
		bool(model.GracefulRestartParameters.XMPPHelperEnable),
		int(model.GracefulRestartParameters.RestartTime),
		int(model.GracefulRestartParameters.LongLivedRestartTime),
		int(model.AutonomousSystem))
	return err
}

func scanGlobalSystemConfig(rows *sql.Rows) (*models.GlobalSystemConfig, error) {
	m := models.MakeGlobalSystemConfig()

	var jsonUserDefinedLogStatistics string

	var jsonPluginTuningPluginProperty string

	var jsonFQName string

	var jsonPerms2Share string

	var jsonIPFabricSubnetsSubnet string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.IbgpAutoMesh,
		&m.MacAgingTime,
		&jsonUserDefinedLogStatistics,
		&m.MacLimitControl.MacLimitAction,
		&m.MacLimitControl.MacLimit,
		&m.ConfigVersion,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.MacMoveControl.MacMoveTimeWindow,
		&jsonPluginTuningPluginProperty,
		&jsonFQName,
		&m.BgpaasParameters.PortEnd,
		&m.BgpaasParameters.PortStart,
		&m.AlarmEnable,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&m.BGPAlwaysCompareMed,
		&jsonIPFabricSubnetsSubnet,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.GracefulRestartParameters.Enable,
		&m.GracefulRestartParameters.EndOfRibTimeout,
		&m.GracefulRestartParameters.BGPHelperEnable,
		&m.GracefulRestartParameters.XMPPHelperEnable,
		&m.GracefulRestartParameters.RestartTime,
		&m.GracefulRestartParameters.LongLivedRestartTime,
		&m.AutonomousSystem); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonUserDefinedLogStatistics), &m.UserDefinedLogStatistics)

	json.Unmarshal([]byte(jsonPluginTuningPluginProperty), &m.PluginTuning.PluginProperty)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonIPFabricSubnetsSubnet), &m.IPFabricSubnets.Subnet)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createGlobalSystemConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["mac_limit_action"]; ok {
		results = append(results, "mac_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["config_version"]; ok {
		results = append(results, "config_version = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_move_limit_action"]; ok {
		results = append(results, "mac_move_limit_action = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListGlobalSystemConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.GlobalSystemConfig, error) {
	result := models.MakeGlobalSystemConfigSlice()
	whereQuery, values := createGlobalSystemConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listGlobalSystemConfigQuery)
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
