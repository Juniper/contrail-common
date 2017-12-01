package db

// vpn_group

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVPNGroupQuery = "insert into `vpn_group` (`type`,`fq_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_log`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVPNGroupQuery = "update `vpn_group` set `type` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_log` = ?,`uuid` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_progress_stage` = ?,`provisioning_state` = ?;"
const deleteVPNGroupQuery = "delete from `vpn_group` where uuid = ?"
const listVPNGroupQuery = "select `type`,`fq_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_log`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`provisioning_state` from `vpn_group`"
const showVPNGroupQuery = "select `type`,`fq_name`,`provisioning_progress`,`provisioning_start_time`,`provisioning_log`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`provisioning_state` from `vpn_group` where uuid = ?"

func CreateVPNGroup(tx *sql.Tx, model *models.VPNGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVPNGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Type),
		util.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningLog),
		string(model.UUID),
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
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningState))
	return err
}

func scanVPNGroup(rows *sql.Rows) (*models.VPNGroup, error) {
	m := models.MakeVPNGroup()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.Type,
		&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.ProvisioningLog,
		&m.UUID,
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
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ProvisioningProgressStage,
		&m.ProvisioningState); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListVPNGroup(tx *sql.Tx) ([]*models.VPNGroup, error) {
	result := models.MakeVPNGroupSlice()
	rows, err := tx.Query(listVPNGroupQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVPNGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVPNGroup(tx *sql.Tx, uuid string) (*models.VPNGroup, error) {
	rows, err := tx.Query(showVPNGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVPNGroup(rows)
	}
	return nil, nil
}

func UpdateVPNGroup(tx *sql.Tx, uuid string, model *models.VPNGroup) error {
	return nil
}

func DeleteVPNGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVPNGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
