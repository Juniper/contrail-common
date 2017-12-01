package db
// service_instance

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceInstanceQuery = "insert into `service_instance` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`service_instance_bindings`,`left_virtual_network`,`auto_policy`,`right_ip_address`,`auto_scale`,`max_instances`,`ha_mode`,`interface_list`,`right_virtual_network`,`availability_zone`,`management_virtual_network`,`virtual_router_id`,`left_ip_address`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceInstanceQuery = "update `service_instance` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`service_instance_bindings` = ?,`left_virtual_network` = ?,`auto_policy` = ?,`right_ip_address` = ?,`auto_scale` = ?,`max_instances` = ?,`ha_mode` = ?,`interface_list` = ?,`right_virtual_network` = ?,`availability_zone` = ?,`management_virtual_network` = ?,`virtual_router_id` = ?,`left_ip_address` = ?,`display_name` = ?;"
const deleteServiceInstanceQuery = "delete from `service_instance`"
const selectServiceInstanceQuery = "select `key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`service_instance_bindings`,`left_virtual_network`,`auto_policy`,`right_ip_address`,`auto_scale`,`max_instances`,`ha_mode`,`interface_list`,`right_virtual_network`,`availability_zone`,`management_virtual_network`,`virtual_router_id`,`left_ip_address`,`display_name` from `service_instance`"

func CreateServiceInstance(tx *sql.Tx, model *models.ServiceInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID,
    model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.ServiceInstanceBindings,
    model.ServiceInstanceProperties.LeftVirtualNetwork,
    model.ServiceInstanceProperties.AutoPolicy,
    model.ServiceInstanceProperties.RightIPAddress,
    model.ServiceInstanceProperties.ScaleOut.AutoScale,
    model.ServiceInstanceProperties.ScaleOut.MaxInstances,
    model.ServiceInstanceProperties.HaMode,
    model.ServiceInstanceProperties.InterfaceList,
    model.ServiceInstanceProperties.RightVirtualNetwork,
    model.ServiceInstanceProperties.AvailabilityZone,
    model.ServiceInstanceProperties.ManagementVirtualNetwork,
    model.ServiceInstanceProperties.VirtualRouterID,
    model.ServiceInstanceProperties.LeftIPAddress,
    model.DisplayName)
    return err
}

func ListServiceInstance(tx *sql.Tx) ([]*models.ServiceInstance, error) {
    result := models.MakeServiceInstanceSlice()
    rows, err := tx.Query(selectServiceInstanceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceInstance()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Creator,
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
                &m.ServiceInstanceBindings,
                &m.ServiceInstanceProperties.LeftVirtualNetwork,
                &m.ServiceInstanceProperties.AutoPolicy,
                &m.ServiceInstanceProperties.RightIPAddress,
                &m.ServiceInstanceProperties.ScaleOut.AutoScale,
                &m.ServiceInstanceProperties.ScaleOut.MaxInstances,
                &m.ServiceInstanceProperties.HaMode,
                &m.ServiceInstanceProperties.InterfaceList,
                &m.ServiceInstanceProperties.RightVirtualNetwork,
                &m.ServiceInstanceProperties.AvailabilityZone,
                &m.ServiceInstanceProperties.ManagementVirtualNetwork,
                &m.ServiceInstanceProperties.VirtualRouterID,
                &m.ServiceInstanceProperties.LeftIPAddress,
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

func ShowServiceInstance(db *sql.DB, id string, model *models.ServiceInstance) error {
    return nil
}

func UpdateServiceInstance(db *sql.DB, id string, model *models.ServiceInstance) error {
    return nil
}

func DeleteServiceInstance(db *sql.DB, id string) error {
    return nil
}