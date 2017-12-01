package db

// service_group

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceGroupQuery = "insert into `service_group` (`service_group_firewall_service_list`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceGroupQuery = "update `service_group` set `service_group_firewall_service_list` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?;"
const deleteServiceGroupQuery = "delete from `service_group` where uuid = ?"
const listServiceGroupQuery = "select `service_group_firewall_service_list`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name` from `service_group`"
const showServiceGroupQuery = "select `service_group_firewall_service_list`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name` from `service_group` where uuid = ?"

func CreateServiceGroup(tx *sql.Tx, model *models.ServiceGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.ServiceGroupFirewallServiceList),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
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
		string(model.DisplayName))
	return err
}

func scanServiceGroup(rows *sql.Rows) (*models.ServiceGroup, error) {
	m := models.MakeServiceGroup()

	var jsonServiceGroupFirewallServiceList string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonServiceGroupFirewallServiceList,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Enable,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonServiceGroupFirewallServiceList), &m.ServiceGroupFirewallServiceList)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListServiceGroup(tx *sql.Tx) ([]*models.ServiceGroup, error) {
	result := models.MakeServiceGroupSlice()
	rows, err := tx.Query(listServiceGroupQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceGroup(tx *sql.Tx, uuid string) (*models.ServiceGroup, error) {
	rows, err := tx.Query(showServiceGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceGroup(rows)
	}
	return nil, nil
}

func UpdateServiceGroup(tx *sql.Tx, uuid string, model *models.ServiceGroup) error {
	return nil
}

func DeleteServiceGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
