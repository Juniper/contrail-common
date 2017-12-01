package db

// contrail_analytics_node

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertContrailAnalyticsNodeQuery = "insert into `contrail_analytics_node` (`provisioning_state`,`global_access`,`share`,`owner`,`owner_access`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`provisioning_progress`,`provisioning_progress_stage`,`uuid`,`key_value_pair`,`provisioning_log`,`provisioning_start_time`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailAnalyticsNodeQuery = "update `contrail_analytics_node` set `provisioning_state` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`uuid` = ?,`key_value_pair` = ?,`provisioning_log` = ?,`provisioning_start_time` = ?;"
const deleteContrailAnalyticsNodeQuery = "delete from `contrail_analytics_node` where uuid = ?"
const listContrailAnalyticsNodeQuery = "select `provisioning_state`,`global_access`,`share`,`owner`,`owner_access`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`provisioning_progress`,`provisioning_progress_stage`,`uuid`,`key_value_pair`,`provisioning_log`,`provisioning_start_time` from `contrail_analytics_node`"
const showContrailAnalyticsNodeQuery = "select `provisioning_state`,`global_access`,`share`,`owner`,`owner_access`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`provisioning_progress`,`provisioning_progress_stage`,`uuid`,`key_value_pair`,`provisioning_log`,`provisioning_start_time` from `contrail_analytics_node` where uuid = ?"

func CreateContrailAnalyticsNode(tx *sql.Tx, model *models.ContrailAnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningState),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		util.MustJSON(model.FQName),
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
		string(model.DisplayName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.UUID),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		string(model.ProvisioningStartTime))
	return err
}

func scanContrailAnalyticsNode(rows *sql.Rows) (*models.ContrailAnalyticsNode, error) {
	m := models.MakeContrailAnalyticsNode()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ProvisioningState,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonFQName,
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
		&m.DisplayName,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.UUID,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningLog,
		&m.ProvisioningStartTime); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListContrailAnalyticsNode(tx *sql.Tx) ([]*models.ContrailAnalyticsNode, error) {
	result := models.MakeContrailAnalyticsNodeSlice()
	rows, err := tx.Query(listContrailAnalyticsNodeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanContrailAnalyticsNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailAnalyticsNode(tx *sql.Tx, uuid string) (*models.ContrailAnalyticsNode, error) {
	rows, err := tx.Query(showContrailAnalyticsNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailAnalyticsNode(rows)
	}
	return nil, nil
}

func UpdateContrailAnalyticsNode(tx *sql.Tx, uuid string, model *models.ContrailAnalyticsNode) error {
	return nil
}

func DeleteContrailAnalyticsNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailAnalyticsNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
