package db

// loadbalancer

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLoadbalancerQuery = "insert into `loadbalancer` (`owner`,`owner_access`,`global_access`,`share`,`vip_subnet_id`,`operating_status`,`status`,`provisioning_status`,`admin_state`,`vip_address`,`loadbalancer_provider`,`uuid`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerQuery = "update `loadbalancer` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`vip_subnet_id` = ?,`operating_status` = ?,`status` = ?,`provisioning_status` = ?,`admin_state` = ?,`vip_address` = ?,`loadbalancer_provider` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteLoadbalancerQuery = "delete from `loadbalancer` where uuid = ?"
const listLoadbalancerQuery = "select `owner`,`owner_access`,`global_access`,`share`,`vip_subnet_id`,`operating_status`,`status`,`provisioning_status`,`admin_state`,`vip_address`,`loadbalancer_provider`,`uuid`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair` from `loadbalancer`"
const showLoadbalancerQuery = "select `owner`,`owner_access`,`global_access`,`share`,`vip_subnet_id`,`operating_status`,`status`,`provisioning_status`,`admin_state`,`vip_address`,`loadbalancer_provider`,`uuid`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair` from `loadbalancer` where uuid = ?"

func CreateLoadbalancer(tx *sql.Tx, model *models.Loadbalancer) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.LoadbalancerProperties.VipSubnetID),
		string(model.LoadbalancerProperties.OperatingStatus),
		string(model.LoadbalancerProperties.Status),
		string(model.LoadbalancerProperties.ProvisioningStatus),
		bool(model.LoadbalancerProperties.AdminState),
		string(model.LoadbalancerProperties.VipAddress),
		string(model.LoadbalancerProvider),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanLoadbalancer(rows *sql.Rows) (*models.Loadbalancer, error) {
	m := models.MakeLoadbalancer()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.LoadbalancerProperties.VipSubnetID,
		&m.LoadbalancerProperties.OperatingStatus,
		&m.LoadbalancerProperties.Status,
		&m.LoadbalancerProperties.ProvisioningStatus,
		&m.LoadbalancerProperties.AdminState,
		&m.LoadbalancerProperties.VipAddress,
		&m.LoadbalancerProvider,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListLoadbalancer(tx *sql.Tx) ([]*models.Loadbalancer, error) {
	result := models.MakeLoadbalancerSlice()
	rows, err := tx.Query(listLoadbalancerQuery)
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
