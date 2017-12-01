package db
// location

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLocationQuery = "insert into `location` (`provisioning_progress`,`provisioning_start_time`,`private_redhat_pool_id`,`gcp_region`,`gcp_subnet`,`aws_access_key`,`uuid`,`private_ospd_vm_name`,`private_ospd_vm_ram_mb`,`private_redhat_subscription_key`,`gcp_asn`,`aws_subnet`,`private_ospd_user_name`,`private_ospd_vm_disk_gb`,`private_redhat_subscription_user`,`owner_access`,`global_access`,`share`,`owner`,`private_ospd_vm_vcpus`,`provisioning_log`,`key_value_pair`,`fq_name`,`type`,`private_ntp_hosts`,`private_ospd_package_url`,`private_ospd_user_password`,`display_name`,`private_redhat_subscription_pasword`,`gcp_account_info`,`aws_secret_key`,`provisioning_progress_stage`,`private_dns_servers`,`aws_region`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLocationQuery = "update `location` set `provisioning_progress` = ?,`provisioning_start_time` = ?,`private_redhat_pool_id` = ?,`gcp_region` = ?,`gcp_subnet` = ?,`aws_access_key` = ?,`uuid` = ?,`private_ospd_vm_name` = ?,`private_ospd_vm_ram_mb` = ?,`private_redhat_subscription_key` = ?,`gcp_asn` = ?,`aws_subnet` = ?,`private_ospd_user_name` = ?,`private_ospd_vm_disk_gb` = ?,`private_redhat_subscription_user` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`private_ospd_vm_vcpus` = ?,`provisioning_log` = ?,`key_value_pair` = ?,`fq_name` = ?,`type` = ?,`private_ntp_hosts` = ?,`private_ospd_package_url` = ?,`private_ospd_user_password` = ?,`display_name` = ?,`private_redhat_subscription_pasword` = ?,`gcp_account_info` = ?,`aws_secret_key` = ?,`provisioning_progress_stage` = ?,`private_dns_servers` = ?,`aws_region` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`provisioning_state` = ?;"
const deleteLocationQuery = "delete from `location`"
const selectLocationQuery = "select `provisioning_progress`,`provisioning_start_time`,`private_redhat_pool_id`,`gcp_region`,`gcp_subnet`,`aws_access_key`,`uuid`,`private_ospd_vm_name`,`private_ospd_vm_ram_mb`,`private_redhat_subscription_key`,`gcp_asn`,`aws_subnet`,`private_ospd_user_name`,`private_ospd_vm_disk_gb`,`private_redhat_subscription_user`,`owner_access`,`global_access`,`share`,`owner`,`private_ospd_vm_vcpus`,`provisioning_log`,`key_value_pair`,`fq_name`,`type`,`private_ntp_hosts`,`private_ospd_package_url`,`private_ospd_user_password`,`display_name`,`private_redhat_subscription_pasword`,`gcp_account_info`,`aws_secret_key`,`provisioning_progress_stage`,`private_dns_servers`,`aws_region`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`provisioning_state` from `location`"

func CreateLocation(tx *sql.Tx, model *models.Location) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLocationQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ProvisioningProgress,
    model.ProvisioningStartTime,
    model.PrivateRedhatPoolID,
    model.GCPRegion,
    model.GCPSubnet,
    model.AwsAccessKey,
    model.UUID,
    model.PrivateOspdVMName,
    model.PrivateOspdVMRAMMB,
    model.PrivateRedhatSubscriptionKey,
    model.GCPAsn,
    model.AwsSubnet,
    model.PrivateOspdUserName,
    model.PrivateOspdVMDiskGB,
    model.PrivateRedhatSubscriptionUser,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.PrivateOspdVMVcpus,
    model.ProvisioningLog,
    model.Annotations.KeyValuePair,
    model.FQName,
    model.Type,
    model.PrivateNTPHosts,
    model.PrivateOspdPackageURL,
    model.PrivateOspdUserPassword,
    model.DisplayName,
    model.PrivateRedhatSubscriptionPasword,
    model.GCPAccountInfo,
    model.AwsSecretKey,
    model.ProvisioningProgressStage,
    model.PrivateDNSServers,
    model.AwsRegion,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.ProvisioningState)
    return err
}

func ListLocation(tx *sql.Tx) ([]*models.Location, error) {
    result := models.MakeLocationSlice()
    rows, err := tx.Query(selectLocationQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLocation()
            if err := rows.Scan(&m.ProvisioningProgress,
                &m.ProvisioningStartTime,
                &m.PrivateRedhatPoolID,
                &m.GCPRegion,
                &m.GCPSubnet,
                &m.AwsAccessKey,
                &m.UUID,
                &m.PrivateOspdVMName,
                &m.PrivateOspdVMRAMMB,
                &m.PrivateRedhatSubscriptionKey,
                &m.GCPAsn,
                &m.AwsSubnet,
                &m.PrivateOspdUserName,
                &m.PrivateOspdVMDiskGB,
                &m.PrivateRedhatSubscriptionUser,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.PrivateOspdVMVcpus,
                &m.ProvisioningLog,
                &m.Annotations.KeyValuePair,
                &m.FQName,
                &m.Type,
                &m.PrivateNTPHosts,
                &m.PrivateOspdPackageURL,
                &m.PrivateOspdUserPassword,
                &m.DisplayName,
                &m.PrivateRedhatSubscriptionPasword,
                &m.GCPAccountInfo,
                &m.AwsSecretKey,
                &m.ProvisioningProgressStage,
                &m.PrivateDNSServers,
                &m.AwsRegion,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.ProvisioningState); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowLocation(db *sql.DB, id string, model *models.Location) error {
    return nil
}

func UpdateLocation(db *sql.DB, id string, model *models.Location) error {
    return nil
}

func DeleteLocation(db *sql.DB, id string) error {
    return nil
}