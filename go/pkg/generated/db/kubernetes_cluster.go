package db

// kubernetes_cluster

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertKubernetesClusterQuery = "insert into `kubernetes_cluster` (`kuberunetes_dashboard`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`display_name`,`contrail_cluster_id`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateKubernetesClusterQuery = "update `kubernetes_cluster` set `kuberunetes_dashboard` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`contrail_cluster_id` = ?;"
const deleteKubernetesClusterQuery = "delete from `kubernetes_cluster` where uuid = ?"
const listKubernetesClusterQuery = "select `kuberunetes_dashboard`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`display_name`,`contrail_cluster_id` from `kubernetes_cluster`"
const showKubernetesClusterQuery = "select `kuberunetes_dashboard`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`display_name`,`contrail_cluster_id` from `kubernetes_cluster` where uuid = ?"

func CreateKubernetesCluster(tx *sql.Tx, model *models.KubernetesCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKubernetesClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.KuberunetesDashboard),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.Created),
		string(model.DisplayName),
		string(model.ContrailClusterID))
	return err
}

func scanKubernetesCluster(rows *sql.Rows) (*models.KubernetesCluster, error) {
	m := models.MakeKubernetesCluster()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.KuberunetesDashboard,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Created,
		&m.DisplayName,
		&m.ContrailClusterID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListKubernetesCluster(tx *sql.Tx) ([]*models.KubernetesCluster, error) {
	result := models.MakeKubernetesClusterSlice()
	rows, err := tx.Query(listKubernetesClusterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanKubernetesCluster(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowKubernetesCluster(tx *sql.Tx, uuid string) (*models.KubernetesCluster, error) {
	rows, err := tx.Query(showKubernetesClusterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanKubernetesCluster(rows)
	}
	return nil, nil
}

func UpdateKubernetesCluster(tx *sql.Tx, uuid string, model *models.KubernetesCluster) error {
	return nil
}

func DeleteKubernetesCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteKubernetesClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
