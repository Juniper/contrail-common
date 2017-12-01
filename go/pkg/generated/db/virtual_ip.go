package db

// virtual_ip

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertVirtualIPQuery = "insert into `virtual_ip` (`persistence_cookie_name`,`connection_limit`,`persistence_type`,`protocol_port`,`protocol`,`subnet_id`,`address`,`status`,`status_description`,`admin_state`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualIPQuery = "update `virtual_ip` set `persistence_cookie_name` = ?,`connection_limit` = ?,`persistence_type` = ?,`protocol_port` = ?,`protocol` = ?,`subnet_id` = ?,`address` = ?,`status` = ?,`status_description` = ?,`admin_state` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteVirtualIPQuery = "delete from `virtual_ip` where uuid = ?"
const listVirtualIPQuery = "select `persistence_cookie_name`,`connection_limit`,`persistence_type`,`protocol_port`,`protocol`,`subnet_id`,`address`,`status`,`status_description`,`admin_state`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid` from `virtual_ip`"
const showVirtualIPQuery = "select `persistence_cookie_name`,`connection_limit`,`persistence_type`,`protocol_port`,`protocol`,`subnet_id`,`address`,`status`,`status_description`,`admin_state`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid` from `virtual_ip` where uuid = ?"

func CreateVirtualIP(tx *sql.Tx, model *models.VirtualIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.VirtualIPProperties.PersistenceCookieName),
		int(model.VirtualIPProperties.ConnectionLimit),
		string(model.VirtualIPProperties.PersistenceType),
		int(model.VirtualIPProperties.ProtocolPort),
		string(model.VirtualIPProperties.Protocol),
		string(model.VirtualIPProperties.SubnetID),
		string(model.VirtualIPProperties.Address),
		string(model.VirtualIPProperties.Status),
		string(model.VirtualIPProperties.StatusDescription),
		bool(model.VirtualIPProperties.AdminState),
		util.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID))
	return err
}

func scanVirtualIP(rows *sql.Rows) (*models.VirtualIP, error) {
	m := models.MakeVirtualIP()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.VirtualIPProperties.PersistenceCookieName,
		&m.VirtualIPProperties.ConnectionLimit,
		&m.VirtualIPProperties.PersistenceType,
		&m.VirtualIPProperties.ProtocolPort,
		&m.VirtualIPProperties.Protocol,
		&m.VirtualIPProperties.SubnetID,
		&m.VirtualIPProperties.Address,
		&m.VirtualIPProperties.Status,
		&m.VirtualIPProperties.StatusDescription,
		&m.VirtualIPProperties.AdminState,
		&jsonFQName,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createVirtualIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["persistence_cookie_name"]; ok {
		results = append(results, "persistence_cookie_name = ?")
		values = append(values, value)
	}

	if value, ok := where["persistence_type"]; ok {
		results = append(results, "persistence_type = ?")
		values = append(values, value)
	}

	if value, ok := where["protocol"]; ok {
		results = append(results, "protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["subnet_id"]; ok {
		results = append(results, "subnet_id = ?")
		values = append(values, value)
	}

	if value, ok := where["address"]; ok {
		results = append(results, "address = ?")
		values = append(values, value)
	}

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
		values = append(values, value)
	}

	if value, ok := where["status_description"]; ok {
		results = append(results, "status_description = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualIP, error) {
	result := models.MakeVirtualIPSlice()
	whereQuery, values := createVirtualIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualIPQuery)
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
