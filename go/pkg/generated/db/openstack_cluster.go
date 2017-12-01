package db

// openstack_cluster

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`openstack_webui`,`public_ip`,`fq_name`,`provisioning_start_time`,`provisioning_state`,`default_storage_backend_bond_interface_members`,`external_net_cidr`,`default_performance_drives`,`external_allocation_pool_end`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`provisioning_progress`,`default_journal_drives`,`default_osd_drives`,`default_storage_access_bond_interface_members`,`external_allocation_pool_start`,`provisioning_log`,`contrail_cluster_id`,`default_capacity_drives`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`key_value_pair`,`provisioning_progress_stage`,`admin_password`,`public_gateway`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackClusterQuery = "update `openstack_cluster` set `openstack_webui` = ?,`public_ip` = ?,`fq_name` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`default_storage_backend_bond_interface_members` = ?,`external_net_cidr` = ?,`default_performance_drives` = ?,`external_allocation_pool_end` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`provisioning_progress` = ?,`default_journal_drives` = ?,`default_osd_drives` = ?,`default_storage_access_bond_interface_members` = ?,`external_allocation_pool_start` = ?,`provisioning_log` = ?,`contrail_cluster_id` = ?,`default_capacity_drives` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`key_value_pair` = ?,`provisioning_progress_stage` = ?,`admin_password` = ?,`public_gateway` = ?;"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster` where uuid = ?"
const listOpenstackClusterQuery = "select `openstack_webui`,`public_ip`,`fq_name`,`provisioning_start_time`,`provisioning_state`,`default_storage_backend_bond_interface_members`,`external_net_cidr`,`default_performance_drives`,`external_allocation_pool_end`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`provisioning_progress`,`default_journal_drives`,`default_osd_drives`,`default_storage_access_bond_interface_members`,`external_allocation_pool_start`,`provisioning_log`,`contrail_cluster_id`,`default_capacity_drives`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`key_value_pair`,`provisioning_progress_stage`,`admin_password`,`public_gateway` from `openstack_cluster`"
const showOpenstackClusterQuery = "select `openstack_webui`,`public_ip`,`fq_name`,`provisioning_start_time`,`provisioning_state`,`default_storage_backend_bond_interface_members`,`external_net_cidr`,`default_performance_drives`,`external_allocation_pool_end`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`provisioning_progress`,`default_journal_drives`,`default_osd_drives`,`default_storage_access_bond_interface_members`,`external_allocation_pool_start`,`provisioning_log`,`contrail_cluster_id`,`default_capacity_drives`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`key_value_pair`,`provisioning_progress_stage`,`admin_password`,`public_gateway` from `openstack_cluster` where uuid = ?"

func CreateOpenstackCluster(tx *sql.Tx, model *models.OpenstackCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.OpenstackWebui),
		string(model.PublicIP),
		util.MustJSON(model.FQName),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.DefaultStorageBackendBondInterfaceMembers),
		string(model.ExternalNetCidr),
		string(model.DefaultPerformanceDrives),
		string(model.ExternalAllocationPoolEnd),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		int(model.ProvisioningProgress),
		string(model.DefaultJournalDrives),
		string(model.DefaultOsdDrives),
		string(model.DefaultStorageAccessBondInterfaceMembers),
		string(model.ExternalAllocationPoolStart),
		string(model.ProvisioningLog),
		string(model.ContrailClusterID),
		string(model.DefaultCapacityDrives),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningProgressStage),
		string(model.AdminPassword),
		string(model.PublicGateway))
	return err
}

func scanOpenstackCluster(rows *sql.Rows) (*models.OpenstackCluster, error) {
	m := models.MakeOpenstackCluster()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.OpenstackWebui,
		&m.PublicIP,
		&jsonFQName,
		&m.ProvisioningStartTime,
		&m.ProvisioningState,
		&m.DefaultStorageBackendBondInterfaceMembers,
		&m.ExternalNetCidr,
		&m.DefaultPerformanceDrives,
		&m.ExternalAllocationPoolEnd,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&m.ProvisioningProgress,
		&m.DefaultJournalDrives,
		&m.DefaultOsdDrives,
		&m.DefaultStorageAccessBondInterfaceMembers,
		&m.ExternalAllocationPoolStart,
		&m.ProvisioningLog,
		&m.ContrailClusterID,
		&m.DefaultCapacityDrives,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningProgressStage,
		&m.AdminPassword,
		&m.PublicGateway); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListOpenstackCluster(tx *sql.Tx) ([]*models.OpenstackCluster, error) {
	result := models.MakeOpenstackClusterSlice()
	rows, err := tx.Query(listOpenstackClusterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanOpenstackCluster(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackCluster(tx *sql.Tx, uuid string) (*models.OpenstackCluster, error) {
	rows, err := tx.Query(showOpenstackClusterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackCluster(rows)
	}
	return nil, nil
}

func UpdateOpenstackCluster(tx *sql.Tx, uuid string, model *models.OpenstackCluster) error {
	return nil
}

func DeleteOpenstackCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
