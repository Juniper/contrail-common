package db

// openstack_cluster

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`provisioning_start_time`,`contrail_cluster_id`,`openstack_webui`,`public_gateway`,`public_ip`,`uuid`,`key_value_pair`,`provisioning_log`,`admin_password`,`default_journal_drives`,`default_storage_access_bond_interface_members`,`external_net_cidr`,`fq_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_state`,`default_capacity_drives`,`default_storage_backend_bond_interface_members`,`external_allocation_pool_end`,`provisioning_progress_stage`,`provisioning_progress`,`default_osd_drives`,`default_performance_drives`,`external_allocation_pool_start`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackClusterQuery = "update `openstack_cluster` set `group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`provisioning_start_time` = ?,`contrail_cluster_id` = ?,`openstack_webui` = ?,`public_gateway` = ?,`public_ip` = ?,`uuid` = ?,`key_value_pair` = ?,`provisioning_log` = ?,`admin_password` = ?,`default_journal_drives` = ?,`default_storage_access_bond_interface_members` = ?,`external_net_cidr` = ?,`fq_name` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_state` = ?,`default_capacity_drives` = ?,`default_storage_backend_bond_interface_members` = ?,`external_allocation_pool_end` = ?,`provisioning_progress_stage` = ?,`provisioning_progress` = ?,`default_osd_drives` = ?,`default_performance_drives` = ?,`external_allocation_pool_start` = ?;"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster` where uuid = ?"
const listOpenstackClusterQuery = "select `group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`provisioning_start_time`,`contrail_cluster_id`,`openstack_webui`,`public_gateway`,`public_ip`,`uuid`,`key_value_pair`,`provisioning_log`,`admin_password`,`default_journal_drives`,`default_storage_access_bond_interface_members`,`external_net_cidr`,`fq_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_state`,`default_capacity_drives`,`default_storage_backend_bond_interface_members`,`external_allocation_pool_end`,`provisioning_progress_stage`,`provisioning_progress`,`default_osd_drives`,`default_performance_drives`,`external_allocation_pool_start` from `openstack_cluster`"
const showOpenstackClusterQuery = "select `group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`provisioning_start_time`,`contrail_cluster_id`,`openstack_webui`,`public_gateway`,`public_ip`,`uuid`,`key_value_pair`,`provisioning_log`,`admin_password`,`default_journal_drives`,`default_storage_access_bond_interface_members`,`external_net_cidr`,`fq_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_state`,`default_capacity_drives`,`default_storage_backend_bond_interface_members`,`external_allocation_pool_end`,`provisioning_progress_stage`,`provisioning_progress`,`default_osd_drives`,`default_performance_drives`,`external_allocation_pool_start` from `openstack_cluster` where uuid = ?"

func CreateOpenstackCluster(tx *sql.Tx, model *models.OpenstackCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		string(model.ProvisioningStartTime),
		string(model.ContrailClusterID),
		string(model.OpenstackWebui),
		string(model.PublicGateway),
		string(model.PublicIP),
		string(model.UUID),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningLog),
		string(model.AdminPassword),
		string(model.DefaultJournalDrives),
		string(model.DefaultStorageAccessBondInterfaceMembers),
		string(model.ExternalNetCidr),
		util.MustJSON(model.FQName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.ProvisioningState),
		string(model.DefaultCapacityDrives),
		string(model.DefaultStorageBackendBondInterfaceMembers),
		string(model.ExternalAllocationPoolEnd),
		string(model.ProvisioningProgressStage),
		int(model.ProvisioningProgress),
		string(model.DefaultOsdDrives),
		string(model.DefaultPerformanceDrives),
		string(model.ExternalAllocationPoolStart))
	return err
}

func scanOpenstackCluster(rows *sql.Rows) (*models.OpenstackCluster, error) {
	m := models.MakeOpenstackCluster()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&m.ProvisioningStartTime,
		&m.ContrailClusterID,
		&m.OpenstackWebui,
		&m.PublicGateway,
		&m.PublicIP,
		&m.UUID,
		&jsonAnnotationsKeyValuePair,
		&m.ProvisioningLog,
		&m.AdminPassword,
		&m.DefaultJournalDrives,
		&m.DefaultStorageAccessBondInterfaceMembers,
		&m.ExternalNetCidr,
		&jsonFQName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ProvisioningState,
		&m.DefaultCapacityDrives,
		&m.DefaultStorageBackendBondInterfaceMembers,
		&m.ExternalAllocationPoolEnd,
		&m.ProvisioningProgressStage,
		&m.ProvisioningProgress,
		&m.DefaultOsdDrives,
		&m.DefaultPerformanceDrives,
		&m.ExternalAllocationPoolStart); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createOpenstackClusterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["contrail_cluster_id"]; ok {
		results = append(results, "contrail_cluster_id = ?")
		values = append(values, value)
	}

	if value, ok := where["openstack_webui"]; ok {
		results = append(results, "openstack_webui = ?")
		values = append(values, value)
	}

	if value, ok := where["public_gateway"]; ok {
		results = append(results, "public_gateway = ?")
		values = append(values, value)
	}

	if value, ok := where["public_ip"]; ok {
		results = append(results, "public_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["admin_password"]; ok {
		results = append(results, "admin_password = ?")
		values = append(values, value)
	}

	if value, ok := where["default_journal_drives"]; ok {
		results = append(results, "default_journal_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["default_storage_access_bond_interface_members"]; ok {
		results = append(results, "default_storage_access_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["external_net_cidr"]; ok {
		results = append(results, "external_net_cidr = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["default_capacity_drives"]; ok {
		results = append(results, "default_capacity_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["default_storage_backend_bond_interface_members"]; ok {
		results = append(results, "default_storage_backend_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["external_allocation_pool_end"]; ok {
		results = append(results, "external_allocation_pool_end = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["default_osd_drives"]; ok {
		results = append(results, "default_osd_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["default_performance_drives"]; ok {
		results = append(results, "default_performance_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["external_allocation_pool_start"]; ok {
		results = append(results, "external_allocation_pool_start = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListOpenstackCluster(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.OpenstackCluster, error) {
	result := models.MakeOpenstackClusterSlice()
	whereQuery, values := createOpenstackClusterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listOpenstackClusterQuery)
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
