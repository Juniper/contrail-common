package db

// instance_ip

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertInstanceIPQuery = "insert into `instance_ip` (`service_health_check_ip`,`ip_prefix`,`ip_prefix_len`,`instance_ip_address`,`instance_ip_local_ip`,`instance_ip_secondary`,`share`,`owner`,`owner_access`,`global_access`,`instance_ip_mode`,`service_instance_ip`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`subnet_uuid`,`instance_ip_family`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInstanceIPQuery = "update `instance_ip` set `service_health_check_ip` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`instance_ip_address` = ?,`instance_ip_local_ip` = ?,`instance_ip_secondary` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`instance_ip_mode` = ?,`service_instance_ip` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`subnet_uuid` = ?,`instance_ip_family` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteInstanceIPQuery = "delete from `instance_ip` where uuid = ?"
const listInstanceIPQuery = "select `service_health_check_ip`,`ip_prefix`,`ip_prefix_len`,`instance_ip_address`,`instance_ip_local_ip`,`instance_ip_secondary`,`share`,`owner`,`owner_access`,`global_access`,`instance_ip_mode`,`service_instance_ip`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`subnet_uuid`,`instance_ip_family`,`uuid`,`fq_name` from `instance_ip`"
const showInstanceIPQuery = "select `service_health_check_ip`,`ip_prefix`,`ip_prefix_len`,`instance_ip_address`,`instance_ip_local_ip`,`instance_ip_secondary`,`share`,`owner`,`owner_access`,`global_access`,`instance_ip_mode`,`service_instance_ip`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`subnet_uuid`,`instance_ip_family`,`uuid`,`fq_name` from `instance_ip` where uuid = ?"

func CreateInstanceIP(tx *sql.Tx, model *models.InstanceIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInstanceIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.ServiceHealthCheckIP),
		string(model.SecondaryIPTrackingIP.IPPrefix),
		int(model.SecondaryIPTrackingIP.IPPrefixLen),
		string(model.InstanceIPAddress),
		bool(model.InstanceIPLocalIP),
		bool(model.InstanceIPSecondary),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.InstanceIPMode),
		bool(model.ServiceInstanceIP),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.SubnetUUID),
		string(model.InstanceIPFamily),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanInstanceIP(rows *sql.Rows) (*models.InstanceIP, error) {
	m := models.MakeInstanceIP()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.ServiceHealthCheckIP,
		&m.SecondaryIPTrackingIP.IPPrefix,
		&m.SecondaryIPTrackingIP.IPPrefixLen,
		&m.InstanceIPAddress,
		&m.InstanceIPLocalIP,
		&m.InstanceIPSecondary,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.InstanceIPMode,
		&m.ServiceInstanceIP,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.SubnetUUID,
		&m.InstanceIPFamily,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListInstanceIP(tx *sql.Tx) ([]*models.InstanceIP, error) {
	result := models.MakeInstanceIPSlice()
	rows, err := tx.Query(listInstanceIPQuery)
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
