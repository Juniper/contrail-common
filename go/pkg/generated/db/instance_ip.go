package db
// instance_ip

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertInstanceIPQuery = "insert into `instance_ip` (`subnet_uuid`,`instance_ip_family`,`instance_ip_mode`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`instance_ip_local_ip`,`service_instance_ip`,`key_value_pair`,`instance_ip_address`,`ip_prefix_len`,`ip_prefix`,`instance_ip_secondary`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`service_health_check_ip`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInstanceIPQuery = "update `instance_ip` set `subnet_uuid` = ?,`instance_ip_family` = ?,`instance_ip_mode` = ?,`display_name` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`instance_ip_local_ip` = ?,`service_instance_ip` = ?,`key_value_pair` = ?,`instance_ip_address` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`instance_ip_secondary` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`service_health_check_ip` = ?;"
const deleteInstanceIPQuery = "delete from `instance_ip`"
const selectInstanceIPQuery = "select `subnet_uuid`,`instance_ip_family`,`instance_ip_mode`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`instance_ip_local_ip`,`service_instance_ip`,`key_value_pair`,`instance_ip_address`,`ip_prefix_len`,`ip_prefix`,`instance_ip_secondary`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`service_health_check_ip` from `instance_ip`"

func CreateInstanceIP(tx *sql.Tx, model *models.InstanceIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInstanceIPQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.SubnetUUID,
    model.InstanceIPFamily,
    model.InstanceIPMode,
    model.DisplayName,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID,
    model.FQName,
    model.InstanceIPLocalIP,
    model.ServiceInstanceIP,
    model.Annotations.KeyValuePair,
    model.InstanceIPAddress,
    model.SecondaryIPTrackingIP.IPPrefixLen,
    model.SecondaryIPTrackingIP.IPPrefix,
    model.InstanceIPSecondary,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.ServiceHealthCheckIP)
    return err
}

func ListInstanceIP(tx *sql.Tx) ([]*models.InstanceIP, error) {
    result := models.MakeInstanceIPSlice()
    rows, err := tx.Query(selectInstanceIPQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeInstanceIP()
            if err := rows.Scan(&m.SubnetUUID,
                &m.InstanceIPFamily,
                &m.InstanceIPMode,
                &m.DisplayName,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.UUID,
                &m.FQName,
                &m.InstanceIPLocalIP,
                &m.ServiceInstanceIP,
                &m.Annotations.KeyValuePair,
                &m.InstanceIPAddress,
                &m.SecondaryIPTrackingIP.IPPrefixLen,
                &m.SecondaryIPTrackingIP.IPPrefix,
                &m.InstanceIPSecondary,
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
                &m.IDPerms.Creator,
                &m.ServiceHealthCheckIP); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowInstanceIP(db *sql.DB, id string, model *models.InstanceIP) error {
    return nil
}

func UpdateInstanceIP(db *sql.DB, id string, model *models.InstanceIP) error {
    return nil
}

func DeleteInstanceIP(db *sql.DB, id string) error {
    return nil
}