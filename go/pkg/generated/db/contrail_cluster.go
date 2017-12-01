package db

// contrail_cluster

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertContrailClusterQuery = "insert into `contrail_cluster` (`contrail_webui`,`data_ttl`,`default_vrouter_bond_interface_members`,`flow_ttl`,`owner`,`owner_access`,`global_access`,`share`,`config_audit_ttl`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`default_gateway`,`statistics_ttl`,`display_name`,`default_vrouter_bond_interface`,`fq_name`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailClusterQuery = "update `contrail_cluster` set `contrail_webui` = ?,`data_ttl` = ?,`default_vrouter_bond_interface_members` = ?,`flow_ttl` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`config_audit_ttl` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`key_value_pair` = ?,`default_gateway` = ?,`statistics_ttl` = ?,`display_name` = ?,`default_vrouter_bond_interface` = ?,`fq_name` = ?,`uuid` = ?;"
const deleteContrailClusterQuery = "delete from `contrail_cluster` where uuid = ?"
const listContrailClusterQuery = "select `contrail_webui`,`data_ttl`,`default_vrouter_bond_interface_members`,`flow_ttl`,`owner`,`owner_access`,`global_access`,`share`,`config_audit_ttl`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`default_gateway`,`statistics_ttl`,`display_name`,`default_vrouter_bond_interface`,`fq_name`,`uuid` from `contrail_cluster`"
const showContrailClusterQuery = "select `contrail_webui`,`data_ttl`,`default_vrouter_bond_interface_members`,`flow_ttl`,`owner`,`owner_access`,`global_access`,`share`,`config_audit_ttl`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`default_gateway`,`statistics_ttl`,`display_name`,`default_vrouter_bond_interface`,`fq_name`,`uuid` from `contrail_cluster` where uuid = ?"

func CreateContrailCluster(tx *sql.Tx, model *models.ContrailCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ContrailWebui),
		string(model.DataTTL),
		string(model.DefaultVrouterBondInterfaceMembers),
		string(model.FlowTTL),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.ConfigAuditTTL),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.DefaultGateway),
		string(model.StatisticsTTL),
		string(model.DisplayName),
		string(model.DefaultVrouterBondInterface),
		util.MustJSON(model.FQName),
		string(model.UUID))
	return err
}

func scanContrailCluster(rows *sql.Rows) (*models.ContrailCluster, error) {
	m := models.MakeContrailCluster()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.ContrailWebui,
		&m.DataTTL,
		&m.DefaultVrouterBondInterfaceMembers,
		&m.FlowTTL,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ConfigAuditTTL,
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
		&m.IDPerms.LastModified,
		&jsonAnnotationsKeyValuePair,
		&m.DefaultGateway,
		&m.StatisticsTTL,
		&m.DisplayName,
		&m.DefaultVrouterBondInterface,
		&jsonFQName,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createContrailClusterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["contrail_webui"]; ok {
		results = append(results, "contrail_webui = ?")
		values = append(values, value)
	}

	if value, ok := where["data_ttl"]; ok {
		results = append(results, "data_ttl = ?")
		values = append(values, value)
	}

	if value, ok := where["default_vrouter_bond_interface_members"]; ok {
		results = append(results, "default_vrouter_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["flow_ttl"]; ok {
		results = append(results, "flow_ttl = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["config_audit_ttl"]; ok {
		results = append(results, "config_audit_ttl = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["default_gateway"]; ok {
		results = append(results, "default_gateway = ?")
		values = append(values, value)
	}

	if value, ok := where["statistics_ttl"]; ok {
		results = append(results, "statistics_ttl = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["default_vrouter_bond_interface"]; ok {
		results = append(results, "default_vrouter_bond_interface = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListContrailCluster(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ContrailCluster, error) {
	result := models.MakeContrailClusterSlice()
	whereQuery, values := createContrailClusterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listContrailClusterQuery)
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
