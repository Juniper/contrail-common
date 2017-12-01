package db

// loadbalancer_listener

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLoadbalancerListenerQuery = "insert into `loadbalancer_listener` (`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`admin_state`,`sni_containers`,`protocol_port`,`default_tls_container`,`protocol`,`connection_limit`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerListenerQuery = "update `loadbalancer_listener` set `uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`admin_state` = ?,`sni_containers` = ?,`protocol_port` = ?,`default_tls_container` = ?,`protocol` = ?,`connection_limit` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteLoadbalancerListenerQuery = "delete from `loadbalancer_listener` where uuid = ?"
const listLoadbalancerListenerQuery = "select `uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`admin_state`,`sni_containers`,`protocol_port`,`default_tls_container`,`protocol`,`connection_limit`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `loadbalancer_listener`"
const showLoadbalancerListenerQuery = "select `uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`admin_state`,`sni_containers`,`protocol_port`,`default_tls_container`,`protocol`,`connection_limit`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `loadbalancer_listener` where uuid = ?"

func CreateLoadbalancerListener(tx *sql.Tx, model *models.LoadbalancerListener) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerListenerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		bool(model.LoadbalancerListenerProperties.AdminState),
		util.MustJSON(model.LoadbalancerListenerProperties.SniContainers),
		int(model.LoadbalancerListenerProperties.ProtocolPort),
		string(model.LoadbalancerListenerProperties.DefaultTLSContainer),
		string(model.LoadbalancerListenerProperties.Protocol),
		int(model.LoadbalancerListenerProperties.ConnectionLimit),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share))
	return err
}

func scanLoadbalancerListener(rows *sql.Rows) (*models.LoadbalancerListener, error) {
	m := models.MakeLoadbalancerListener()

	var jsonFQName string

	var jsonLoadbalancerListenerPropertiesSniContainers string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&m.LoadbalancerListenerProperties.AdminState,
		&jsonLoadbalancerListenerPropertiesSniContainers,
		&m.LoadbalancerListenerProperties.ProtocolPort,
		&m.LoadbalancerListenerProperties.DefaultTLSContainer,
		&m.LoadbalancerListenerProperties.Protocol,
		&m.LoadbalancerListenerProperties.ConnectionLimit,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonLoadbalancerListenerPropertiesSniContainers), &m.LoadbalancerListenerProperties.SniContainers)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListLoadbalancerListener(tx *sql.Tx) ([]*models.LoadbalancerListener, error) {
	result := models.MakeLoadbalancerListenerSlice()
	rows, err := tx.Query(listLoadbalancerListenerQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanLoadbalancerListener(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerListener(tx *sql.Tx, uuid string) (*models.LoadbalancerListener, error) {
	rows, err := tx.Query(showLoadbalancerListenerQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerListener(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerListener(tx *sql.Tx, uuid string, model *models.LoadbalancerListener) error {
	return nil
}

func DeleteLoadbalancerListener(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerListenerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
