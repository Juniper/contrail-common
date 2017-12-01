package db
// openstack_storage_node_role

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertOpenstackStorageNodeRoleQuery = "insert into `openstack_storage_node_role` (`provisioning_start_time`,`provisioning_state`,`uuid`,`fq_name`,`display_name`,`key_value_pair`,`journal_drives`,`provisioning_progress_stage`,`osd_drives`,`storage_access_bond_interface_members`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`provisioning_progress`,`storage_backend_bond_interface_members`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackStorageNodeRoleQuery = "update `openstack_storage_node_role` set `provisioning_start_time` = ?,`provisioning_state` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?,`journal_drives` = ?,`provisioning_progress_stage` = ?,`osd_drives` = ?,`storage_access_bond_interface_members` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`provisioning_progress` = ?,`storage_backend_bond_interface_members` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_log` = ?;"
const deleteOpenstackStorageNodeRoleQuery = "delete from `openstack_storage_node_role`"
const selectOpenstackStorageNodeRoleQuery = "select `provisioning_start_time`,`provisioning_state`,`uuid`,`fq_name`,`display_name`,`key_value_pair`,`journal_drives`,`provisioning_progress_stage`,`osd_drives`,`storage_access_bond_interface_members`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`provisioning_progress`,`storage_backend_bond_interface_members`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_log` from `openstack_storage_node_role`"

func CreateOpenstackStorageNodeRole(tx *sql.Tx, model *models.OpenstackStorageNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackStorageNodeRoleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningStartTime,
    model.ProvisioningState,
    model.UUID,
    model.FQName,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.JournalDrives,
    model.ProvisioningProgressStage,
    model.OsdDrives,
    model.StorageAccessBondInterfaceMembers,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.ProvisioningProgress,
    model.StorageBackendBondInterfaceMembers,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.ProvisioningLog)
    return err
}

func ListOpenstackStorageNodeRole(tx *sql.Tx) ([]*models.OpenstackStorageNodeRole, error) {
    result := models.MakeOpenstackStorageNodeRoleSlice()
    rows, err := tx.Query(selectOpenstackStorageNodeRoleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeOpenstackStorageNodeRole()
            if err := rows.Scan(&m.ProvisioningStartTime,
                &m.ProvisioningState,
                &m.UUID,
                &m.FQName,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.JournalDrives,
                &m.ProvisioningProgressStage,
                &m.OsdDrives,
                &m.StorageAccessBondInterfaceMembers,
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
                &m.IDPerms.Creator,
                &m.ProvisioningProgress,
                &m.StorageBackendBondInterfaceMembers,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.ProvisioningLog); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowOpenstackStorageNodeRole(db *sql.DB, id string, model *models.OpenstackStorageNodeRole) error {
    return nil
}

func UpdateOpenstackStorageNodeRole(db *sql.DB, id string, model *models.OpenstackStorageNodeRole) error {
    return nil
}

func DeleteOpenstackStorageNodeRole(db *sql.DB, id string) error {
    return nil
}