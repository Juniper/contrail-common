package db

// location

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertLocationQuery = "insert into `location` (`private_redhat_pool_id`,`private_redhat_subscription_pasword`,`gcp_asn`,`gcp_subnet`,`provisioning_progress_stage`,`private_dns_servers`,`private_ospd_user_name`,`private_ospd_user_password`,`aws_region`,`private_ospd_package_url`,`private_ospd_vm_name`,`aws_access_key`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`aws_secret_key`,`provisioning_state`,`aws_subnet`,`fq_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_progress`,`private_ntp_hosts`,`private_ospd_vm_ram_mb`,`gcp_region`,`provisioning_start_time`,`private_redhat_subscription_user`,`uuid`,`provisioning_log`,`type`,`gcp_account_info`,`display_name`,`key_value_pair`,`private_ospd_vm_disk_gb`,`private_ospd_vm_vcpus`,`private_redhat_subscription_key`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLocationQuery = "update `location` set `private_redhat_pool_id` = ?,`private_redhat_subscription_pasword` = ?,`gcp_asn` = ?,`gcp_subnet` = ?,`provisioning_progress_stage` = ?,`private_dns_servers` = ?,`private_ospd_user_name` = ?,`private_ospd_user_password` = ?,`aws_region` = ?,`private_ospd_package_url` = ?,`private_ospd_vm_name` = ?,`aws_access_key` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`aws_secret_key` = ?,`provisioning_state` = ?,`aws_subnet` = ?,`fq_name` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`provisioning_progress` = ?,`private_ntp_hosts` = ?,`private_ospd_vm_ram_mb` = ?,`gcp_region` = ?,`provisioning_start_time` = ?,`private_redhat_subscription_user` = ?,`uuid` = ?,`provisioning_log` = ?,`type` = ?,`gcp_account_info` = ?,`display_name` = ?,`key_value_pair` = ?,`private_ospd_vm_disk_gb` = ?,`private_ospd_vm_vcpus` = ?,`private_redhat_subscription_key` = ?;"
const deleteLocationQuery = "delete from `location` where uuid = ?"
const listLocationQuery = "select `private_redhat_pool_id`,`private_redhat_subscription_pasword`,`gcp_asn`,`gcp_subnet`,`provisioning_progress_stage`,`private_dns_servers`,`private_ospd_user_name`,`private_ospd_user_password`,`aws_region`,`private_ospd_package_url`,`private_ospd_vm_name`,`aws_access_key`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`aws_secret_key`,`provisioning_state`,`aws_subnet`,`fq_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_progress`,`private_ntp_hosts`,`private_ospd_vm_ram_mb`,`gcp_region`,`provisioning_start_time`,`private_redhat_subscription_user`,`uuid`,`provisioning_log`,`type`,`gcp_account_info`,`display_name`,`key_value_pair`,`private_ospd_vm_disk_gb`,`private_ospd_vm_vcpus`,`private_redhat_subscription_key` from `location`"
const showLocationQuery = "select `private_redhat_pool_id`,`private_redhat_subscription_pasword`,`gcp_asn`,`gcp_subnet`,`provisioning_progress_stage`,`private_dns_servers`,`private_ospd_user_name`,`private_ospd_user_password`,`aws_region`,`private_ospd_package_url`,`private_ospd_vm_name`,`aws_access_key`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`aws_secret_key`,`provisioning_state`,`aws_subnet`,`fq_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`provisioning_progress`,`private_ntp_hosts`,`private_ospd_vm_ram_mb`,`gcp_region`,`provisioning_start_time`,`private_redhat_subscription_user`,`uuid`,`provisioning_log`,`type`,`gcp_account_info`,`display_name`,`key_value_pair`,`private_ospd_vm_disk_gb`,`private_ospd_vm_vcpus`,`private_redhat_subscription_key` from `location` where uuid = ?"

func CreateLocation(tx *sql.Tx, model *models.Location) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLocationQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.PrivateRedhatPoolID),
		string(model.PrivateRedhatSubscriptionPasword),
		int(model.GCPAsn),
		string(model.GCPSubnet),
		string(model.ProvisioningProgressStage),
		string(model.PrivateDNSServers),
		string(model.PrivateOspdUserName),
		string(model.PrivateOspdUserPassword),
		string(model.AwsRegion),
		string(model.PrivateOspdPackageURL),
		string(model.PrivateOspdVMName),
		string(model.AwsAccessKey),
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
		string(model.IDPerms.Description),
		string(model.AwsSecretKey),
		string(model.ProvisioningState),
		string(model.AwsSubnet),
		util.MustJSON(model.FQName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.ProvisioningProgress),
		string(model.PrivateNTPHosts),
		string(model.PrivateOspdVMRAMMB),
		string(model.GCPRegion),
		string(model.ProvisioningStartTime),
		string(model.PrivateRedhatSubscriptionUser),
		string(model.UUID),
		string(model.ProvisioningLog),
		string(model.Type),
		string(model.GCPAccountInfo),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.PrivateOspdVMDiskGB),
		string(model.PrivateOspdVMVcpus),
		string(model.PrivateRedhatSubscriptionKey))
	return err
}

func scanLocation(rows *sql.Rows) (*models.Location, error) {
	m := models.MakeLocation()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.PrivateRedhatPoolID,
		&m.PrivateRedhatSubscriptionPasword,
		&m.GCPAsn,
		&m.GCPSubnet,
		&m.ProvisioningProgressStage,
		&m.PrivateDNSServers,
		&m.PrivateOspdUserName,
		&m.PrivateOspdUserPassword,
		&m.AwsRegion,
		&m.PrivateOspdPackageURL,
		&m.PrivateOspdVMName,
		&m.AwsAccessKey,
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
		&m.IDPerms.Description,
		&m.AwsSecretKey,
		&m.ProvisioningState,
		&m.AwsSubnet,
		&jsonFQName,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.ProvisioningProgress,
		&m.PrivateNTPHosts,
		&m.PrivateOspdVMRAMMB,
		&m.GCPRegion,
		&m.ProvisioningStartTime,
		&m.PrivateRedhatSubscriptionUser,
		&m.UUID,
		&m.ProvisioningLog,
		&m.Type,
		&m.GCPAccountInfo,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.PrivateOspdVMDiskGB,
		&m.PrivateOspdVMVcpus,
		&m.PrivateRedhatSubscriptionKey); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createLocationWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["private_redhat_pool_id"]; ok {
		results = append(results, "private_redhat_pool_id = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_subscription_pasword"]; ok {
		results = append(results, "private_redhat_subscription_pasword = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_subnet"]; ok {
		results = append(results, "gcp_subnet = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["private_dns_servers"]; ok {
		results = append(results, "private_dns_servers = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_user_name"]; ok {
		results = append(results, "private_ospd_user_name = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_user_password"]; ok {
		results = append(results, "private_ospd_user_password = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_region"]; ok {
		results = append(results, "aws_region = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_package_url"]; ok {
		results = append(results, "private_ospd_package_url = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_name"]; ok {
		results = append(results, "private_ospd_vm_name = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_access_key"]; ok {
		results = append(results, "aws_access_key = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_secret_key"]; ok {
		results = append(results, "aws_secret_key = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_subnet"]; ok {
		results = append(results, "aws_subnet = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ntp_hosts"]; ok {
		results = append(results, "private_ntp_hosts = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_ram_mb"]; ok {
		results = append(results, "private_ospd_vm_ram_mb = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_region"]; ok {
		results = append(results, "gcp_region = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_subscription_user"]; ok {
		results = append(results, "private_redhat_subscription_user = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["type"]; ok {
		results = append(results, "type = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_account_info"]; ok {
		results = append(results, "gcp_account_info = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_disk_gb"]; ok {
		results = append(results, "private_ospd_vm_disk_gb = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_vcpus"]; ok {
		results = append(results, "private_ospd_vm_vcpus = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_subscription_key"]; ok {
		results = append(results, "private_redhat_subscription_key = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLocation(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Location, error) {
	result := models.MakeLocationSlice()
	whereQuery, values := createLocationWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLocationQuery)
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
