package db

// kuberunetes_node

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertKuberunetesNodeQuery = "insert into `kuberunetes_node` (`provisioning_start_time`,`key_value_pair`,`display_name`,`provisioning_progress_stage`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`provisioning_log`,`provisioning_progress`,`provisioning_state`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateKuberunetesNodeQuery = "update `kuberunetes_node` set `provisioning_start_time` = ?,`key_value_pair` = ?,`display_name` = ?,`provisioning_progress_stage` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_state` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteKuberunetesNodeQuery = "delete from `kuberunetes_node` where uuid = ?"
const listKuberunetesNodeQuery = "select `provisioning_start_time`,`key_value_pair`,`display_name`,`provisioning_progress_stage`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`provisioning_log`,`provisioning_progress`,`provisioning_state`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name` from `kuberunetes_node`"
const showKuberunetesNodeQuery = "select `provisioning_start_time`,`key_value_pair`,`display_name`,`provisioning_progress_stage`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`provisioning_log`,`provisioning_progress`,`provisioning_state`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name` from `kuberunetes_node` where uuid = ?"

func CreateKuberunetesNode(tx *sql.Tx, model *models.KuberunetesNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKuberunetesNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ProvisioningStartTime),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.DisplayName),
		string(model.ProvisioningProgressStage),
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
		string(model.IDPerms.LastModified),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningState),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanKuberunetesNode(rows *sql.Rows) (*models.KuberunetesNode, error) {
	m := models.MakeKuberunetesNode()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.ProvisioningStartTime,
		&jsonAnnotationsKeyValuePair,
		&m.DisplayName,
		&m.ProvisioningProgressStage,
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
		&m.IDPerms.LastModified,
		&m.ProvisioningLog,
		&m.ProvisioningProgress,
		&m.ProvisioningState,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createKuberunetesNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListKuberunetesNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.KuberunetesNode, error) {
	result := models.MakeKuberunetesNodeSlice()
	whereQuery, values := createKuberunetesNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listKuberunetesNodeQuery)
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
