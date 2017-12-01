package db
// openstack_compute_node_role

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertOpenstackComputeNodeRoleQuery = "insert into `openstack_compute_node_role` (`vrouter_bond_interface`,`fq_name`,`provisioning_progress`,`provisioning_state`,`provisioning_log`,`vrouter_type`,`default_gateway`,`uuid`,`display_name`,`global_access`,`share`,`owner`,`owner_access`,`vrouter_bond_interface_members`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`key_value_pair`,`provisioning_progress_stage`,`provisioning_start_time`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackComputeNodeRoleQuery = "update `openstack_compute_node_role` set `vrouter_bond_interface` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_state` = ?,`provisioning_log` = ?,`vrouter_type` = ?,`default_gateway` = ?,`uuid` = ?,`display_name` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`vrouter_bond_interface_members` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`key_value_pair` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?;"
const deleteOpenstackComputeNodeRoleQuery = "delete from `openstack_compute_node_role`"
const selectOpenstackComputeNodeRoleQuery = "select `vrouter_bond_interface`,`fq_name`,`provisioning_progress`,`provisioning_state`,`provisioning_log`,`vrouter_type`,`default_gateway`,`uuid`,`display_name`,`global_access`,`share`,`owner`,`owner_access`,`vrouter_bond_interface_members`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`key_value_pair`,`provisioning_progress_stage`,`provisioning_start_time` from `openstack_compute_node_role`"

func CreateOpenstackComputeNodeRole(tx *sql.Tx, model *models.OpenstackComputeNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackComputeNodeRoleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.VrouterBondInterface,
    model.FQName,
    model.ProvisioningProgress,
    model.ProvisioningState,
    model.ProvisioningLog,
    model.VrouterType,
    model.DefaultGateway,
    model.UUID,
    model.DisplayName,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.VrouterBondInterfaceMembers,
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
    model.Annotations.KeyValuePair,
    model.ProvisioningProgressStage,
    model.ProvisioningStartTime)
    return err
}

func ListOpenstackComputeNodeRole(tx *sql.Tx) ([]*models.OpenstackComputeNodeRole, error) {
    result := models.MakeOpenstackComputeNodeRoleSlice()
    rows, err := tx.Query(selectOpenstackComputeNodeRoleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeOpenstackComputeNodeRole()
            if err := rows.Scan(&m.VrouterBondInterface,
                &m.FQName,
                &m.ProvisioningProgress,
                &m.ProvisioningState,
                &m.ProvisioningLog,
                &m.VrouterType,
                &m.DefaultGateway,
                &m.UUID,
                &m.DisplayName,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.VrouterBondInterfaceMembers,
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
                &m.Annotations.KeyValuePair,
                &m.ProvisioningProgressStage,
                &m.ProvisioningStartTime); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowOpenstackComputeNodeRole(db *sql.DB, id string, model *models.OpenstackComputeNodeRole) error {
    return nil
}

func UpdateOpenstackComputeNodeRole(db *sql.DB, id string, model *models.OpenstackComputeNodeRole) error {
    return nil
}

func DeleteOpenstackComputeNodeRole(db *sql.DB, id string) error {
    return nil
}