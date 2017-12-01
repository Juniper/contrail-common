package db

// kuberunetes_node

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertKuberunetesNodeQuery = "insert into `kuberunetes_node` (`owner_access`,`global_access`,`share`,`owner`,`provisioning_state`,`provisioning_progress_stage`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`provisioning_progress`,`provisioning_start_time`,`uuid`,`fq_name`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateKuberunetesNodeQuery = "update `kuberunetes_node` set `owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`provisioning_state` = ?,`provisioning_progress_stage` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`provisioning_progress` = ?,`provisioning_start_time` = ?,`uuid` = ?,`fq_name` = ?,`provisioning_log` = ?;"
const deleteKuberunetesNodeQuery = "delete from `kuberunetes_node` where uuid = ?"
const listKuberunetesNodeQuery = "select `owner_access`,`global_access`,`share`,`owner`,`provisioning_state`,`provisioning_progress_stage`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`provisioning_progress`,`provisioning_start_time`,`uuid`,`fq_name`,`provisioning_log` from `kuberunetes_node`"
const showKuberunetesNodeQuery = "select `owner_access`,`global_access`,`share`,`owner`,`provisioning_state`,`provisioning_progress_stage`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`provisioning_progress`,`provisioning_start_time`,`uuid`,`fq_name`,`provisioning_log` from `kuberunetes_node` where uuid = ?"

func CreateKuberunetesNode(tx *sql.Tx, model *models.KuberunetesNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKuberunetesNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.ProvisioningState),
		string(model.ProvisioningProgressStage),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.ProvisioningProgress),
		string(model.ProvisioningStartTime),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.ProvisioningLog))
	return err
}

func scanKuberunetesNode(rows *sql.Rows) (*models.KuberunetesNode, error) {
	m := models.MakeKuberunetesNode()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.ProvisioningState,
		&m.ProvisioningProgressStage,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningProgress,
		&m.ProvisioningStartTime,
		&m.UUID,
		&jsonFQName,
		&m.ProvisioningLog); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListKuberunetesNode(tx *sql.Tx) ([]*models.KuberunetesNode, error) {
	result := models.MakeKuberunetesNodeSlice()
	rows, err := tx.Query(listKuberunetesNodeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanKuberunetesNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowKuberunetesNode(tx *sql.Tx, uuid string) (*models.KuberunetesNode, error) {
	rows, err := tx.Query(showKuberunetesNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanKuberunetesNode(rows)
	}
	return nil, nil
}

func UpdateKuberunetesNode(tx *sql.Tx, uuid string, model *models.KuberunetesNode) error {
	return nil
}

func DeleteKuberunetesNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteKuberunetesNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
