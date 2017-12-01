package db

// loadbalancer_member

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLoadbalancerMemberQuery = "insert into `loadbalancer_member` (`display_name`,`admin_state`,`address`,`protocol_port`,`status`,`status_description`,`weight`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerMemberQuery = "update `loadbalancer_member` set `display_name` = ?,`admin_state` = ?,`address` = ?,`protocol_port` = ?,`status` = ?,`status_description` = ?,`weight` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?;"
const deleteLoadbalancerMemberQuery = "delete from `loadbalancer_member` where uuid = ?"
const listLoadbalancerMemberQuery = "select `display_name`,`admin_state`,`address`,`protocol_port`,`status`,`status_description`,`weight`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `loadbalancer_member`"
const showLoadbalancerMemberQuery = "select `display_name`,`admin_state`,`address`,`protocol_port`,`status`,`status_description`,`weight`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified` from `loadbalancer_member` where uuid = ?"

func CreateLoadbalancerMember(tx *sql.Tx, model *models.LoadbalancerMember) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		bool(model.LoadbalancerMemberProperties.AdminState),
		string(model.LoadbalancerMemberProperties.Address),
		int(model.LoadbalancerMemberProperties.ProtocolPort),
		string(model.LoadbalancerMemberProperties.Status),
		string(model.LoadbalancerMemberProperties.StatusDescription),
		int(model.LoadbalancerMemberProperties.Weight),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified))
	return err
}

func scanLoadbalancerMember(rows *sql.Rows) (*models.LoadbalancerMember, error) {
	m := models.MakeLoadbalancerMember()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&m.LoadbalancerMemberProperties.AdminState,
		&m.LoadbalancerMemberProperties.Address,
		&m.LoadbalancerMemberProperties.ProtocolPort,
		&m.LoadbalancerMemberProperties.Status,
		&m.LoadbalancerMemberProperties.StatusDescription,
		&m.LoadbalancerMemberProperties.Weight,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListLoadbalancerMember(tx *sql.Tx) ([]*models.LoadbalancerMember, error) {
	result := models.MakeLoadbalancerMemberSlice()
	rows, err := tx.Query(listLoadbalancerMemberQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanLoadbalancerMember(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerMember(tx *sql.Tx, uuid string) (*models.LoadbalancerMember, error) {
	rows, err := tx.Query(showLoadbalancerMemberQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerMember(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerMember(tx *sql.Tx, uuid string, model *models.LoadbalancerMember) error {
	return nil
}

func DeleteLoadbalancerMember(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
