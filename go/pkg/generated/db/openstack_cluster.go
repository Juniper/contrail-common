package db
// openstack_cluster

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`provisioning_log`,`contrail_cluster_id`,`default_storage_backend_bond_interface_members`,`external_allocation_pool_end`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`provisioning_progress_stage`,`admin_password`,`default_performance_drives`,`default_storage_access_bond_interface_members`,`external_net_cidr`,`public_ip`,`provisioning_start_time`,`provisioning_state`,`default_journal_drives`,`external_allocation_pool_start`,`openstack_webui`,`public_gateway`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`default_capacity_drives`,`default_osd_drives`,`fq_name`,`provisioning_progress`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackClusterQuery = "update `openstack_cluster` set `provisioning_log` = ?,`contrail_cluster_id` = ?,`default_storage_backend_bond_interface_members` = ?,`external_allocation_pool_end` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`provisioning_progress_stage` = ?,`admin_password` = ?,`default_performance_drives` = ?,`default_storage_access_bond_interface_members` = ?,`external_net_cidr` = ?,`public_ip` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`default_journal_drives` = ?,`external_allocation_pool_start` = ?,`openstack_webui` = ?,`public_gateway` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`default_capacity_drives` = ?,`default_osd_drives` = ?,`fq_name` = ?,`provisioning_progress` = ?;"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster`"
const selectOpenstackClusterQuery = "select `provisioning_log`,`contrail_cluster_id`,`default_storage_backend_bond_interface_members`,`external_allocation_pool_end`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`provisioning_progress_stage`,`admin_password`,`default_performance_drives`,`default_storage_access_bond_interface_members`,`external_net_cidr`,`public_ip`,`provisioning_start_time`,`provisioning_state`,`default_journal_drives`,`external_allocation_pool_start`,`openstack_webui`,`public_gateway`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`default_capacity_drives`,`default_osd_drives`,`fq_name`,`provisioning_progress` from `openstack_cluster`"

func CreateOpenstackCluster(tx *sql.Tx, model *models.OpenstackCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackClusterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningLog,
    model.ContrailClusterID,
    model.DefaultStorageBackendBondInterfaceMembers,
    model.ExternalAllocationPoolEnd,
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
    model.IDPerms.UserVisible,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.ProvisioningProgressStage,
    model.AdminPassword,
    model.DefaultPerformanceDrives,
    model.DefaultStorageAccessBondInterfaceMembers,
    model.ExternalNetCidr,
    model.PublicIP,
    model.ProvisioningStartTime,
    model.ProvisioningState,
    model.DefaultJournalDrives,
    model.ExternalAllocationPoolStart,
    model.OpenstackWebui,
    model.PublicGateway,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.DefaultCapacityDrives,
    model.DefaultOsdDrives,
    model.FQName,
    model.ProvisioningProgress)
    return err
}

func ListOpenstackCluster(tx *sql.Tx) ([]*models.OpenstackCluster, error) {
    result := models.MakeOpenstackClusterSlice()
    rows, err := tx.Query(selectOpenstackClusterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeOpenstackCluster()
            if err := rows.Scan(&m.ProvisioningLog,
                &m.ContrailClusterID,
                &m.DefaultStorageBackendBondInterfaceMembers,
                &m.ExternalAllocationPoolEnd,
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
                &m.IDPerms.UserVisible,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.ProvisioningProgressStage,
                &m.AdminPassword,
                &m.DefaultPerformanceDrives,
                &m.DefaultStorageAccessBondInterfaceMembers,
                &m.ExternalNetCidr,
                &m.PublicIP,
                &m.ProvisioningStartTime,
                &m.ProvisioningState,
                &m.DefaultJournalDrives,
                &m.ExternalAllocationPoolStart,
                &m.OpenstackWebui,
                &m.PublicGateway,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.DefaultCapacityDrives,
                &m.DefaultOsdDrives,
                &m.FQName,
                &m.ProvisioningProgress); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowOpenstackCluster(db *sql.DB, id string, model *models.OpenstackCluster) error {
    return nil
}

func UpdateOpenstackCluster(db *sql.DB, id string, model *models.OpenstackCluster) error {
    return nil
}

func DeleteOpenstackCluster(db *sql.DB, id string) error {
    return nil
}