package db
// controller_node_role

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertControllerNodeRoleQuery = "insert into `controller_node_role` (`storage_management_bond_interface_members`,`owner`,`owner_access`,`global_access`,`share`,`provisioning_progress`,`capacity_drives`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`provisioning_progress_stage`,`provisioning_state`,`internalapi_bond_interface_members`,`uuid`,`provisioning_start_time`,`performance_drives`,`fq_name`,`key_value_pair`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateControllerNodeRoleQuery = "update `controller_node_role` set `storage_management_bond_interface_members` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_progress` = ?,`capacity_drives` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`provisioning_progress_stage` = ?,`provisioning_state` = ?,`internalapi_bond_interface_members` = ?,`uuid` = ?,`provisioning_start_time` = ?,`performance_drives` = ?,`fq_name` = ?,`key_value_pair` = ?,`provisioning_log` = ?;"
const deleteControllerNodeRoleQuery = "delete from `controller_node_role`"
const selectControllerNodeRoleQuery = "select `storage_management_bond_interface_members`,`owner`,`owner_access`,`global_access`,`share`,`provisioning_progress`,`capacity_drives`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`provisioning_progress_stage`,`provisioning_state`,`internalapi_bond_interface_members`,`uuid`,`provisioning_start_time`,`performance_drives`,`fq_name`,`key_value_pair`,`provisioning_log` from `controller_node_role`"

func CreateControllerNodeRole(tx *sql.Tx, model *models.ControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertControllerNodeRoleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.StorageManagementBondInterfaceMembers,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.ProvisioningProgress,
    model.CapacityDrives,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.DisplayName,
    model.ProvisioningProgressStage,
    model.ProvisioningState,
    model.InternalapiBondInterfaceMembers,
    model.UUID,
    model.ProvisioningStartTime,
    model.PerformanceDrives,
    model.FQName,
    model.Annotations.KeyValuePair,
    model.ProvisioningLog)
    return err
}

func ListControllerNodeRole(tx *sql.Tx) ([]*models.ControllerNodeRole, error) {
    result := models.MakeControllerNodeRoleSlice()
    rows, err := tx.Query(selectControllerNodeRoleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeControllerNodeRole()
            if err := rows.Scan(&m.StorageManagementBondInterfaceMembers,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.ProvisioningProgress,
                &m.CapacityDrives,
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
                &m.ProvisioningProgressStage,
                &m.ProvisioningState,
                &m.InternalapiBondInterfaceMembers,
                &m.UUID,
                &m.ProvisioningStartTime,
                &m.PerformanceDrives,
                &m.FQName,
                &m.Annotations.KeyValuePair,
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

func ShowControllerNodeRole(db *sql.DB, id string, model *models.ControllerNodeRole) error {
    return nil
}

func UpdateControllerNodeRole(db *sql.DB, id string, model *models.ControllerNodeRole) error {
    return nil
}

func DeleteControllerNodeRole(db *sql.DB, id string) error {
    return nil
}