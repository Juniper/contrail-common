package db

// instance_ip

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertInstanceIPQuery = "insert into `instance_ip` (`service_instance_ip`,`instance_ip_secondary`,`service_health_check_ip`,`instance_ip_mode`,`instance_ip_address`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`instance_ip_family`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`instance_ip_local_ip`,`fq_name`,`key_value_pair`,`uuid`,`ip_prefix_len`,`ip_prefix`,`subnet_uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInstanceIPQuery = "update `instance_ip` set `service_instance_ip` = ?,`instance_ip_secondary` = ?,`service_health_check_ip` = ?,`instance_ip_mode` = ?,`instance_ip_address` = ?,`display_name` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`instance_ip_family` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`instance_ip_local_ip` = ?,`fq_name` = ?,`key_value_pair` = ?,`uuid` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`subnet_uuid` = ?;"
const deleteInstanceIPQuery = "delete from `instance_ip` where uuid = ?"
const listInstanceIPQuery = "select `service_instance_ip`,`instance_ip_secondary`,`service_health_check_ip`,`instance_ip_mode`,`instance_ip_address`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`instance_ip_family`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`instance_ip_local_ip`,`fq_name`,`key_value_pair`,`uuid`,`ip_prefix_len`,`ip_prefix`,`subnet_uuid` from `instance_ip`"
const showInstanceIPQuery = "select `service_instance_ip`,`instance_ip_secondary`,`service_health_check_ip`,`instance_ip_mode`,`instance_ip_address`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`instance_ip_family`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`instance_ip_local_ip`,`fq_name`,`key_value_pair`,`uuid`,`ip_prefix_len`,`ip_prefix`,`subnet_uuid` from `instance_ip` where uuid = ?"

func CreateInstanceIP(tx *sql.Tx, model *models.InstanceIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInstanceIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.ServiceInstanceIP),
		bool(model.InstanceIPSecondary),
		bool(model.ServiceHealthCheckIP),
		string(model.InstanceIPMode),
		string(model.InstanceIPAddress),
		string(model.DisplayName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.InstanceIPFamily),
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
		string(model.IDPerms.Description),
		bool(model.InstanceIPLocalIP),
		util.MustJSON(model.FQName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		int(model.SecondaryIPTrackingIP.IPPrefixLen),
		string(model.SecondaryIPTrackingIP.IPPrefix),
		string(model.SubnetUUID))
	return err
}

func scanInstanceIP(rows *sql.Rows) (*models.InstanceIP, error) {
	m := models.MakeInstanceIP()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ServiceInstanceIP,
		&m.InstanceIPSecondary,
		&m.ServiceHealthCheckIP,
		&m.InstanceIPMode,
		&m.InstanceIPAddress,
		&m.DisplayName,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.InstanceIPFamily,
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
		&m.IDPerms.Description,
		&m.InstanceIPLocalIP,
		&jsonFQName,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&m.SecondaryIPTrackingIP.IPPrefixLen,
		&m.SecondaryIPTrackingIP.IPPrefix,
		&m.SubnetUUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createInstanceIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["instance_ip_mode"]; ok {
		results = append(results, "instance_ip_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["instance_ip_address"]; ok {
		results = append(results, "instance_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["instance_ip_family"]; ok {
		results = append(results, "instance_ip_family = ?")
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["subnet_uuid"]; ok {
		results = append(results, "subnet_uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListInstanceIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.InstanceIP, error) {
	result := models.MakeInstanceIPSlice()
	whereQuery, values := createInstanceIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listInstanceIPQuery)
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
		m, _ := scanInstanceIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowInstanceIP(tx *sql.Tx, uuid string) (*models.InstanceIP, error) {
	rows, err := tx.Query(showInstanceIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanInstanceIP(rows)
	}
	return nil, nil
}

func UpdateInstanceIP(tx *sql.Tx, uuid string, model *models.InstanceIP) error {
	return nil
}

func DeleteInstanceIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteInstanceIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
