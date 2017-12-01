package db

// virtual_ip

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVirtualIPQuery = "insert into `virtual_ip` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`connection_limit`,`address`,`protocol_port`,`subnet_id`,`protocol`,`admin_state`,`persistence_cookie_name`,`status_description`,`persistence_type`,`status`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualIPQuery = "update `virtual_ip` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`connection_limit` = ?,`address` = ?,`protocol_port` = ?,`subnet_id` = ?,`protocol` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`status_description` = ?,`persistence_type` = ?,`status` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteVirtualIPQuery = "delete from `virtual_ip` where uuid = ?"
const listVirtualIPQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`connection_limit`,`address`,`protocol_port`,`subnet_id`,`protocol`,`admin_state`,`persistence_cookie_name`,`status_description`,`persistence_type`,`status`,`display_name`,`key_value_pair` from `virtual_ip`"
const showVirtualIPQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`connection_limit`,`address`,`protocol_port`,`subnet_id`,`protocol`,`admin_state`,`persistence_cookie_name`,`status_description`,`persistence_type`,`status`,`display_name`,`key_value_pair` from `virtual_ip` where uuid = ?"

func CreateVirtualIP(tx *sql.Tx, model *models.VirtualIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.Created),
		int(model.VirtualIPProperties.ConnectionLimit),
		string(model.VirtualIPProperties.Address),
		int(model.VirtualIPProperties.ProtocolPort),
		string(model.VirtualIPProperties.SubnetID),
		string(model.VirtualIPProperties.Protocol),
		bool(model.VirtualIPProperties.AdminState),
		string(model.VirtualIPProperties.PersistenceCookieName),
		string(model.VirtualIPProperties.StatusDescription),
		string(model.VirtualIPProperties.PersistenceType),
		string(model.VirtualIPProperties.Status),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanVirtualIP(rows *sql.Rows) (*models.VirtualIP, error) {
	m := models.MakeVirtualIP()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Created,
		&m.VirtualIPProperties.ConnectionLimit,
		&m.VirtualIPProperties.Address,
		&m.VirtualIPProperties.ProtocolPort,
		&m.VirtualIPProperties.SubnetID,
		&m.VirtualIPProperties.Protocol,
		&m.VirtualIPProperties.AdminState,
		&m.VirtualIPProperties.PersistenceCookieName,
		&m.VirtualIPProperties.StatusDescription,
		&m.VirtualIPProperties.PersistenceType,
		&m.VirtualIPProperties.Status,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListVirtualIP(tx *sql.Tx) ([]*models.VirtualIP, error) {
	result := models.MakeVirtualIPSlice()
	rows, err := tx.Query(listVirtualIPQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualIP(tx *sql.Tx, uuid string) (*models.VirtualIP, error) {
	rows, err := tx.Query(showVirtualIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualIP(rows)
	}
	return nil, nil
}

func UpdateVirtualIP(tx *sql.Tx, uuid string, model *models.VirtualIP) error {
	return nil
}

func DeleteVirtualIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
