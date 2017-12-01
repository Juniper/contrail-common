package db

// loadbalancer

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertLoadbalancerQuery = "insert into `loadbalancer` (`uuid`,`provisioning_status`,`admin_state`,`vip_address`,`vip_subnet_id`,`operating_status`,`status`,`loadbalancer_provider`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerQuery = "update `loadbalancer` set `uuid` = ?,`provisioning_status` = ?,`admin_state` = ?,`vip_address` = ?,`vip_subnet_id` = ?,`operating_status` = ?,`status` = ?,`loadbalancer_provider` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?;"
const deleteLoadbalancerQuery = "delete from `loadbalancer` where uuid = ?"
const listLoadbalancerQuery = "select `uuid`,`provisioning_status`,`admin_state`,`vip_address`,`vip_subnet_id`,`operating_status`,`status`,`loadbalancer_provider`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access` from `loadbalancer`"
const showLoadbalancerQuery = "select `uuid`,`provisioning_status`,`admin_state`,`vip_address`,`vip_subnet_id`,`operating_status`,`status`,`loadbalancer_provider`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access` from `loadbalancer` where uuid = ?"

func CreateLoadbalancer(tx *sql.Tx, model *models.Loadbalancer) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		string(model.LoadbalancerProperties.ProvisioningStatus),
		bool(model.LoadbalancerProperties.AdminState),
		string(model.LoadbalancerProperties.VipAddress),
		string(model.LoadbalancerProperties.VipSubnetID),
		string(model.LoadbalancerProperties.OperatingStatus),
		string(model.LoadbalancerProperties.Status),
		string(model.LoadbalancerProvider),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess))
	return err
}

func scanLoadbalancer(rows *sql.Rows) (*models.Loadbalancer, error) {
	m := models.MakeLoadbalancer()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&m.LoadbalancerProperties.ProvisioningStatus,
		&m.LoadbalancerProperties.AdminState,
		&m.LoadbalancerProperties.VipAddress,
		&m.LoadbalancerProperties.VipSubnetID,
		&m.LoadbalancerProperties.OperatingStatus,
		&m.LoadbalancerProperties.Status,
		&m.LoadbalancerProvider,
		&jsonFQName,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createLoadbalancerWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_status"]; ok {
		results = append(results, "provisioning_status = ?")
		values = append(values, value)
	}

	if value, ok := where["vip_address"]; ok {
		results = append(results, "vip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["vip_subnet_id"]; ok {
		results = append(results, "vip_subnet_id = ?")
		values = append(values, value)
	}

	if value, ok := where["operating_status"]; ok {
		results = append(results, "operating_status = ?")
		values = append(values, value)
	}

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
		values = append(values, value)
	}

	if value, ok := where["loadbalancer_provider"]; ok {
		results = append(results, "loadbalancer_provider = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancer(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Loadbalancer, error) {
	result := models.MakeLoadbalancerSlice()
	whereQuery, values := createLoadbalancerWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerQuery)
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
		m, _ := scanLoadbalancer(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancer(tx *sql.Tx, uuid string) (*models.Loadbalancer, error) {
	rows, err := tx.Query(showLoadbalancerQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancer(rows)
	}
	return nil, nil
}

func UpdateLoadbalancer(tx *sql.Tx, uuid string, model *models.Loadbalancer) error {
	return nil
}

func DeleteLoadbalancer(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
