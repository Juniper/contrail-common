package db
// node

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertNodeQuery = "insert into `node` (`private_machine_state`,`private_power_management_ip`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`ip_address`,`type`,`password`,`aws_instance_type`,`hostname`,`username`,`aws_ami`,`private_power_management_username`,`uuid`,`fq_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`mac_address`,`ssh_key`,`gcp_image`,`gcp_machine_type`,`private_machine_properties`,`private_power_management_password`,`key_value_pair`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNodeQuery = "update `node` set `private_machine_state` = ?,`private_power_management_ip` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`ip_address` = ?,`type` = ?,`password` = ?,`aws_instance_type` = ?,`hostname` = ?,`username` = ?,`aws_ami` = ?,`private_power_management_username` = ?,`uuid` = ?,`fq_name` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`mac_address` = ?,`ssh_key` = ?,`gcp_image` = ?,`gcp_machine_type` = ?,`private_machine_properties` = ?,`private_power_management_password` = ?,`key_value_pair` = ?,`display_name` = ?;"
const deleteNodeQuery = "delete from `node`"
const selectNodeQuery = "select `private_machine_state`,`private_power_management_ip`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`ip_address`,`type`,`password`,`aws_instance_type`,`hostname`,`username`,`aws_ami`,`private_power_management_username`,`uuid`,`fq_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`mac_address`,`ssh_key`,`gcp_image`,`gcp_machine_type`,`private_machine_properties`,`private_power_management_password`,`key_value_pair`,`display_name` from `node`"

func CreateNode(tx *sql.Tx, model *models.Node) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNodeQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.PrivateMachineState,
    model.PrivatePowerManagementIP,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IPAddress,
    model.Type,
    model.Password,
    model.AwsInstanceType,
    model.Hostname,
    model.Username,
    model.AwsAmi,
    model.PrivatePowerManagementUsername,
    model.UUID,
    model.FQName,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.MacAddress,
    model.SSHKey,
    model.GCPImage,
    model.GCPMachineType,
    model.PrivateMachineProperties,
    model.PrivatePowerManagementPassword,
    model.Annotations.KeyValuePair,
    model.DisplayName)
    return err
}

func ListNode(tx *sql.Tx) ([]*models.Node, error) {
    result := models.MakeNodeSlice()
    rows, err := tx.Query(selectNodeQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeNode()
            if err := rows.Scan(&m.PrivateMachineState,
                &m.PrivatePowerManagementIP,
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
                &m.IPAddress,
                &m.Type,
                &m.Password,
                &m.AwsInstanceType,
                &m.Hostname,
                &m.Username,
                &m.AwsAmi,
                &m.PrivatePowerManagementUsername,
                &m.UUID,
                &m.FQName,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.MacAddress,
                &m.SSHKey,
                &m.GCPImage,
                &m.GCPMachineType,
                &m.PrivateMachineProperties,
                &m.PrivatePowerManagementPassword,
                &m.Annotations.KeyValuePair,
                &m.DisplayName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowNode(db *sql.DB, id string, model *models.Node) error {
    return nil
}

func UpdateNode(db *sql.DB, id string, model *models.Node) error {
    return nil
}

func DeleteNode(db *sql.DB, id string) error {
    return nil
}