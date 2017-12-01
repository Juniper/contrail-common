package db

// openstack_storage_node_role

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertOpenstackStorageNodeRoleQuery = "insert into `openstack_storage_node_role` (`display_name`,`provisioning_state`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`provisioning_progress_stage`,`storage_backend_bond_interface_members`,`uuid`,`fq_name`,`osd_drives`,`storage_access_bond_interface_members`,`provisioning_progress`,`journal_drives`,`provisioning_start_time`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackStorageNodeRoleQuery = "update `openstack_storage_node_role` set `display_name` = ?,`provisioning_state` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`provisioning_progress_stage` = ?,`storage_backend_bond_interface_members` = ?,`uuid` = ?,`fq_name` = ?,`osd_drives` = ?,`storage_access_bond_interface_members` = ?,`provisioning_progress` = ?,`journal_drives` = ?,`provisioning_start_time` = ?,`provisioning_log` = ?;"
const deleteOpenstackStorageNodeRoleQuery = "delete from `openstack_storage_node_role` where uuid = ?"
const listOpenstackStorageNodeRoleQuery = "select `display_name`,`provisioning_state`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`provisioning_progress_stage`,`storage_backend_bond_interface_members`,`uuid`,`fq_name`,`osd_drives`,`storage_access_bond_interface_members`,`provisioning_progress`,`journal_drives`,`provisioning_start_time`,`provisioning_log` from `openstack_storage_node_role`"
const showOpenstackStorageNodeRoleQuery = "select `display_name`,`provisioning_state`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`provisioning_progress_stage`,`storage_backend_bond_interface_members`,`uuid`,`fq_name`,`osd_drives`,`storage_access_bond_interface_members`,`provisioning_progress`,`journal_drives`,`provisioning_start_time`,`provisioning_log` from `openstack_storage_node_role` where uuid = ?"

func CreateOpenstackStorageNodeRole(tx *sql.Tx, model *models.OpenstackStorageNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackStorageNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		string(model.ProvisioningState),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.ProvisioningProgressStage),
		string(model.StorageBackendBondInterfaceMembers),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.OsdDrives),
		string(model.StorageAccessBondInterfaceMembers),
		int(model.ProvisioningProgress),
		string(model.JournalDrives),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningLog))
	return err
}

func scanOpenstackStorageNodeRole(rows *sql.Rows) (*models.OpenstackStorageNodeRole, error) {
	m := models.MakeOpenstackStorageNodeRole()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&m.ProvisioningState,
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
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.ProvisioningProgressStage,
		&m.StorageBackendBondInterfaceMembers,
		&m.UUID,
		&jsonFQName,
		&m.OsdDrives,
		&m.StorageAccessBondInterfaceMembers,
		&m.ProvisioningProgress,
		&m.JournalDrives,
		&m.ProvisioningStartTime,
		&m.ProvisioningLog); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createOpenstackStorageNodeRoleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["storage_backend_bond_interface_members"]; ok {
		results = append(results, "storage_backend_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["osd_drives"]; ok {
		results = append(results, "osd_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["storage_access_bond_interface_members"]; ok {
		results = append(results, "storage_access_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["journal_drives"]; ok {
		results = append(results, "journal_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListOpenstackStorageNodeRole(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.OpenstackStorageNodeRole, error) {
	result := models.MakeOpenstackStorageNodeRoleSlice()
	whereQuery, values := createOpenstackStorageNodeRoleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listOpenstackStorageNodeRoleQuery)
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
		m, _ := scanOpenstackStorageNodeRole(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackStorageNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackStorageNodeRole, error) {
	rows, err := tx.Query(showOpenstackStorageNodeRoleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackStorageNodeRole(rows)
	}
	return nil, nil
}

func UpdateOpenstackStorageNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackStorageNodeRole) error {
	return nil
}

func DeleteOpenstackStorageNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackStorageNodeRoleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
