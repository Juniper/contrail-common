package db
// kuberunetes_node

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertKuberunetesNodeQuery = "insert into `kuberunetes_node` (`provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`owner`,`owner_access`,`global_access`,`share`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`provisioning_log`,`provisioning_progress_stage`,`uuid`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateKuberunetesNodeQuery = "update `kuberunetes_node` set `provisioning_progress` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`provisioning_log` = ?,`provisioning_progress_stage` = ?,`uuid` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteKuberunetesNodeQuery = "delete from `kuberunetes_node`"
const selectKuberunetesNodeQuery = "select `provisioning_progress`,`provisioning_start_time`,`provisioning_state`,`owner`,`owner_access`,`global_access`,`share`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`provisioning_log`,`provisioning_progress_stage`,`uuid`,`display_name`,`key_value_pair` from `kuberunetes_node`"

func CreateKuberunetesNode(tx *sql.Tx, model *models.KuberunetesNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKuberunetesNodeQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningProgress,
    model.ProvisioningStartTime,
    model.ProvisioningState,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.FQName,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.ProvisioningLog,
    model.ProvisioningProgressStage,
    model.UUID,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListKuberunetesNode(tx *sql.Tx) ([]*models.KuberunetesNode, error) {
    result := models.MakeKuberunetesNodeSlice()
    rows, err := tx.Query(selectKuberunetesNodeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeKuberunetesNode()
            if err := rows.Scan(&m.ProvisioningProgress,
                &m.ProvisioningStartTime,
                &m.ProvisioningState,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.FQName,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.ProvisioningLog,
                &m.ProvisioningProgressStage,
                &m.UUID,
                &m.DisplayName,
                &m.Annotations.KeyValuePair); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowKuberunetesNode(db *sql.DB, id string, model *models.KuberunetesNode) error {
    return nil
}

func UpdateKuberunetesNode(db *sql.DB, id string, model *models.KuberunetesNode) error {
    return nil
}

func DeleteKuberunetesNode(db *sql.DB, id string) error {
    return nil
}