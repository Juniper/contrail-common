package db

// node

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertNodeQuery = "insert into `node` (`username`,`gcp_image`,`mac_address`,`password`,`private_machine_properties`,`owner_access`,`global_access`,`share`,`owner`,`hostname`,`type`,`gcp_machine_type`,`private_machine_state`,`private_power_management_username`,`ip_address`,`ssh_key`,`aws_ami`,`aws_instance_type`,`private_power_management_ip`,`private_power_management_password`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNodeQuery = "update `node` set `username` = ?,`gcp_image` = ?,`mac_address` = ?,`password` = ?,`private_machine_properties` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`hostname` = ?,`type` = ?,`gcp_machine_type` = ?,`private_machine_state` = ?,`private_power_management_username` = ?,`ip_address` = ?,`ssh_key` = ?,`aws_ami` = ?,`aws_instance_type` = ?,`private_power_management_ip` = ?,`private_power_management_password` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteNodeQuery = "delete from `node` where uuid = ?"
const listNodeQuery = "select `username`,`gcp_image`,`mac_address`,`password`,`private_machine_properties`,`owner_access`,`global_access`,`share`,`owner`,`hostname`,`type`,`gcp_machine_type`,`private_machine_state`,`private_power_management_username`,`ip_address`,`ssh_key`,`aws_ami`,`aws_instance_type`,`private_power_management_ip`,`private_power_management_password`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`display_name`,`key_value_pair` from `node`"
const showNodeQuery = "select `username`,`gcp_image`,`mac_address`,`password`,`private_machine_properties`,`owner_access`,`global_access`,`share`,`owner`,`hostname`,`type`,`gcp_machine_type`,`private_machine_state`,`private_power_management_username`,`ip_address`,`ssh_key`,`aws_ami`,`aws_instance_type`,`private_power_management_ip`,`private_power_management_password`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`display_name`,`key_value_pair` from `node` where uuid = ?"

func CreateNode(tx *sql.Tx, model *models.Node) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Username),
		string(model.GCPImage),
		string(model.MacAddress),
		string(model.Password),
		string(model.PrivateMachineProperties),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.Hostname),
		string(model.Type),
		string(model.GCPMachineType),
		string(model.PrivateMachineState),
		string(model.PrivatePowerManagementUsername),
		string(model.IPAddress),
		string(model.SSHKey),
		string(model.AwsAmi),
		string(model.AwsInstanceType),
		string(model.PrivatePowerManagementIP),
		string(model.PrivatePowerManagementPassword),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanNode(rows *sql.Rows) (*models.Node, error) {
	m := models.MakeNode()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Username,
		&m.GCPImage,
		&m.MacAddress,
		&m.Password,
		&m.PrivateMachineProperties,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Hostname,
		&m.Type,
		&m.GCPMachineType,
		&m.PrivateMachineState,
		&m.PrivatePowerManagementUsername,
		&m.IPAddress,
		&m.SSHKey,
		&m.AwsAmi,
		&m.AwsInstanceType,
		&m.PrivatePowerManagementIP,
		&m.PrivatePowerManagementPassword,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListNode(tx *sql.Tx) ([]*models.Node, error) {
	result := models.MakeNodeSlice()
	rows, err := tx.Query(listNodeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNode(tx *sql.Tx, uuid string) (*models.Node, error) {
	rows, err := tx.Query(showNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNode(rows)
	}
	return nil, nil
}

func UpdateNode(tx *sql.Tx, uuid string, model *models.Node) error {
	return nil
}

func DeleteNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
