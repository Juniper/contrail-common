package db

// contrail_cluster

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertContrailClusterQuery = "insert into `contrail_cluster` (`display_name`,`contrail_webui`,`data_ttl`,`flow_ttl`,`default_gateway`,`default_vrouter_bond_interface`,`uuid`,`fq_name`,`share`,`owner`,`owner_access`,`global_access`,`config_audit_ttl`,`default_vrouter_bond_interface_members`,`key_value_pair`,`statistics_ttl`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailClusterQuery = "update `contrail_cluster` set `display_name` = ?,`contrail_webui` = ?,`data_ttl` = ?,`flow_ttl` = ?,`default_gateway` = ?,`default_vrouter_bond_interface` = ?,`uuid` = ?,`fq_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`config_audit_ttl` = ?,`default_vrouter_bond_interface_members` = ?,`key_value_pair` = ?,`statistics_ttl` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteContrailClusterQuery = "delete from `contrail_cluster` where uuid = ?"
const listContrailClusterQuery = "select `display_name`,`contrail_webui`,`data_ttl`,`flow_ttl`,`default_gateway`,`default_vrouter_bond_interface`,`uuid`,`fq_name`,`share`,`owner`,`owner_access`,`global_access`,`config_audit_ttl`,`default_vrouter_bond_interface_members`,`key_value_pair`,`statistics_ttl`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible` from `contrail_cluster`"
const showContrailClusterQuery = "select `display_name`,`contrail_webui`,`data_ttl`,`flow_ttl`,`default_gateway`,`default_vrouter_bond_interface`,`uuid`,`fq_name`,`share`,`owner`,`owner_access`,`global_access`,`config_audit_ttl`,`default_vrouter_bond_interface_members`,`key_value_pair`,`statistics_ttl`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible` from `contrail_cluster` where uuid = ?"

func CreateContrailCluster(tx *sql.Tx, model *models.ContrailCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		string(model.ContrailWebui),
		string(model.DataTTL),
		string(model.FlowTTL),
		string(model.DefaultGateway),
		string(model.DefaultVrouterBondInterface),
		string(model.UUID),
		util.MustJSON(model.FQName),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.ConfigAuditTTL),
		string(model.DefaultVrouterBondInterfaceMembers),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.StatisticsTTL),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible))
	return err
}

func scanContrailCluster(rows *sql.Rows) (*models.ContrailCluster, error) {
	m := models.MakeContrailCluster()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.DisplayName,
		&m.ContrailWebui,
		&m.DataTTL,
		&m.FlowTTL,
		&m.DefaultGateway,
		&m.DefaultVrouterBondInterface,
		&m.UUID,
		&jsonFQName,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.ConfigAuditTTL,
		&m.DefaultVrouterBondInterfaceMembers,
		&jsonAnnotationsKeyValuePair,
		&m.StatisticsTTL,
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
		&m.IDPerms.UserVisible); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListContrailCluster(tx *sql.Tx) ([]*models.ContrailCluster, error) {
	result := models.MakeContrailClusterSlice()
	rows, err := tx.Query(listContrailClusterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanContrailCluster(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowContrailCluster(tx *sql.Tx, uuid string) (*models.ContrailCluster, error) {
	rows, err := tx.Query(showContrailClusterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanContrailCluster(rows)
	}
	return nil, nil
}

func UpdateContrailCluster(tx *sql.Tx, uuid string, model *models.ContrailCluster) error {
	return nil
}

func DeleteContrailCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
