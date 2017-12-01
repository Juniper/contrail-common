package db
// kubernetes_cluster

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertKubernetesClusterQuery = "insert into `kubernetes_cluster` (`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`contrail_cluster_id`,`kuberunetes_dashboard`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateKubernetesClusterQuery = "update `kubernetes_cluster` set `key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`contrail_cluster_id` = ?,`kuberunetes_dashboard` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?;"
const deleteKubernetesClusterQuery = "delete from `kubernetes_cluster`"
const selectKubernetesClusterQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`contrail_cluster_id`,`kuberunetes_dashboard`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name` from `kubernetes_cluster`"

func CreateKubernetesCluster(tx *sql.Tx, model *models.KubernetesCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKubernetesClusterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.ContrailClusterID,
    model.KuberunetesDashboard,
    model.FQName,
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
    model.IDPerms.Creator,
    model.DisplayName)
    return err
}

func ListKubernetesCluster(tx *sql.Tx) ([]*models.KubernetesCluster, error) {
    result := models.MakeKubernetesClusterSlice()
    rows, err := tx.Query(selectKubernetesClusterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeKubernetesCluster()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.UUID,
                &m.ContrailClusterID,
                &m.KuberunetesDashboard,
                &m.FQName,
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
                &m.IDPerms.Creator,
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

func ShowKubernetesCluster(db *sql.DB, id string, model *models.KubernetesCluster) error {
    return nil
}

func UpdateKubernetesCluster(db *sql.DB, id string, model *models.KubernetesCluster) error {
    return nil
}

func DeleteKubernetesCluster(db *sql.DB, id string) error {
    return nil
}