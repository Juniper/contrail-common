package db

// location

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertLocationQuery = "insert into `location` (`private_ospd_vm_ram_mb`,`private_redhat_subscription_key`,`gcp_account_info`,`aws_subnet`,`uuid`,`provisioning_log`,`private_redhat_subscription_user`,`aws_access_key`,`provisioning_progress`,`private_dns_servers`,`private_ospd_package_url`,`private_ospd_user_name`,`private_ospd_user_password`,`provisioning_progress_stage`,`private_redhat_subscription_pasword`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`private_ntp_hosts`,`gcp_asn`,`gcp_subnet`,`aws_secret_key`,`key_value_pair`,`private_ospd_vm_disk_gb`,`private_ospd_vm_vcpus`,`private_redhat_pool_id`,`aws_region`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`provisioning_state`,`type`,`private_ospd_vm_name`,`gcp_region`,`fq_name`,`provisioning_start_time`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLocationQuery = "update `location` set `private_ospd_vm_ram_mb` = ?,`private_redhat_subscription_key` = ?,`gcp_account_info` = ?,`aws_subnet` = ?,`uuid` = ?,`provisioning_log` = ?,`private_redhat_subscription_user` = ?,`aws_access_key` = ?,`provisioning_progress` = ?,`private_dns_servers` = ?,`private_ospd_package_url` = ?,`private_ospd_user_name` = ?,`private_ospd_user_password` = ?,`provisioning_progress_stage` = ?,`private_redhat_subscription_pasword` = ?,`display_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`private_ntp_hosts` = ?,`gcp_asn` = ?,`gcp_subnet` = ?,`aws_secret_key` = ?,`key_value_pair` = ?,`private_ospd_vm_disk_gb` = ?,`private_ospd_vm_vcpus` = ?,`private_redhat_pool_id` = ?,`aws_region` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`provisioning_state` = ?,`type` = ?,`private_ospd_vm_name` = ?,`gcp_region` = ?,`fq_name` = ?,`provisioning_start_time` = ?;"
const deleteLocationQuery = "delete from `location` where uuid = ?"
const listLocationQuery = "select `private_ospd_vm_ram_mb`,`private_redhat_subscription_key`,`gcp_account_info`,`aws_subnet`,`uuid`,`provisioning_log`,`private_redhat_subscription_user`,`aws_access_key`,`provisioning_progress`,`private_dns_servers`,`private_ospd_package_url`,`private_ospd_user_name`,`private_ospd_user_password`,`provisioning_progress_stage`,`private_redhat_subscription_pasword`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`private_ntp_hosts`,`gcp_asn`,`gcp_subnet`,`aws_secret_key`,`key_value_pair`,`private_ospd_vm_disk_gb`,`private_ospd_vm_vcpus`,`private_redhat_pool_id`,`aws_region`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`provisioning_state`,`type`,`private_ospd_vm_name`,`gcp_region`,`fq_name`,`provisioning_start_time` from `location`"
const showLocationQuery = "select `private_ospd_vm_ram_mb`,`private_redhat_subscription_key`,`gcp_account_info`,`aws_subnet`,`uuid`,`provisioning_log`,`private_redhat_subscription_user`,`aws_access_key`,`provisioning_progress`,`private_dns_servers`,`private_ospd_package_url`,`private_ospd_user_name`,`private_ospd_user_password`,`provisioning_progress_stage`,`private_redhat_subscription_pasword`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`private_ntp_hosts`,`gcp_asn`,`gcp_subnet`,`aws_secret_key`,`key_value_pair`,`private_ospd_vm_disk_gb`,`private_ospd_vm_vcpus`,`private_redhat_pool_id`,`aws_region`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`provisioning_state`,`type`,`private_ospd_vm_name`,`gcp_region`,`fq_name`,`provisioning_start_time` from `location` where uuid = ?"

func CreateLocation(tx *sql.Tx, model *models.Location) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLocationQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.PrivateOspdVMRAMMB),
		string(model.PrivateRedhatSubscriptionKey),
		string(model.GCPAccountInfo),
		string(model.AwsSubnet),
		string(model.UUID),
		string(model.ProvisioningLog),
		string(model.PrivateRedhatSubscriptionUser),
		string(model.AwsAccessKey),
		int(model.ProvisioningProgress),
		string(model.PrivateDNSServers),
		string(model.PrivateOspdPackageURL),
		string(model.PrivateOspdUserName),
		string(model.PrivateOspdUserPassword),
		string(model.ProvisioningProgressStage),
		string(model.PrivateRedhatSubscriptionPasword),
		string(model.DisplayName),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.PrivateNTPHosts),
		int(model.GCPAsn),
		string(model.GCPSubnet),
		string(model.AwsSecretKey),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.PrivateOspdVMDiskGB),
		string(model.PrivateOspdVMVcpus),
		string(model.PrivateRedhatPoolID),
		string(model.AwsRegion),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.ProvisioningState),
		string(model.Type),
		string(model.PrivateOspdVMName),
		string(model.GCPRegion),
		util.MustJSON(model.FQName),
		string(model.ProvisioningStartTime))
	return err
}

func scanLocation(rows *sql.Rows) (*models.Location, error) {
	m := models.MakeLocation()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.PrivateOspdVMRAMMB,
		&m.PrivateRedhatSubscriptionKey,
		&m.GCPAccountInfo,
		&m.AwsSubnet,
		&m.UUID,
		&m.ProvisioningLog,
		&m.PrivateRedhatSubscriptionUser,
		&m.AwsAccessKey,
		&m.ProvisioningProgress,
		&m.PrivateDNSServers,
		&m.PrivateOspdPackageURL,
		&m.PrivateOspdUserName,
		&m.PrivateOspdUserPassword,
		&m.ProvisioningProgressStage,
		&m.PrivateRedhatSubscriptionPasword,
		&m.DisplayName,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.PrivateNTPHosts,
		&m.GCPAsn,
		&m.GCPSubnet,
		&m.AwsSecretKey,
		&jsonAnnotationsKeyValuePair,
		&m.PrivateOspdVMDiskGB,
		&m.PrivateOspdVMVcpus,
		&m.PrivateRedhatPoolID,
		&m.AwsRegion,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.ProvisioningState,
		&m.Type,
		&m.PrivateOspdVMName,
		&m.GCPRegion,
		&jsonFQName,
		&m.ProvisioningStartTime); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListLocation(tx *sql.Tx) ([]*models.Location, error) {
	result := models.MakeLocationSlice()
	rows, err := tx.Query(listLocationQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanLocation(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLocation(tx *sql.Tx, uuid string) (*models.Location, error) {
	rows, err := tx.Query(showLocationQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLocation(rows)
	}
	return nil, nil
}

func UpdateLocation(tx *sql.Tx, uuid string, model *models.Location) error {
	return nil
}

func DeleteLocation(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLocationQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
