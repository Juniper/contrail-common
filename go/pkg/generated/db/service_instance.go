package db

// service_instance

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceInstanceQuery = "insert into `service_instance` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`service_instance_bindings`,`right_ip_address`,`availability_zone`,`auto_scale`,`max_instances`,`ha_mode`,`virtual_router_id`,`interface_list`,`left_ip_address`,`left_virtual_network`,`auto_policy`,`right_virtual_network`,`management_virtual_network`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceInstanceQuery = "update `service_instance` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`service_instance_bindings` = ?,`right_ip_address` = ?,`availability_zone` = ?,`auto_scale` = ?,`max_instances` = ?,`ha_mode` = ?,`virtual_router_id` = ?,`interface_list` = ?,`left_ip_address` = ?,`left_virtual_network` = ?,`auto_policy` = ?,`right_virtual_network` = ?,`management_virtual_network` = ?,`uuid` = ?,`fq_name` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?;"
const deleteServiceInstanceQuery = "delete from `service_instance` where uuid = ?"
const listServiceInstanceQuery = "select `key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`service_instance_bindings`,`right_ip_address`,`availability_zone`,`auto_scale`,`max_instances`,`ha_mode`,`virtual_router_id`,`interface_list`,`left_ip_address`,`left_virtual_network`,`auto_policy`,`right_virtual_network`,`management_virtual_network`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name` from `service_instance`"
const showServiceInstanceQuery = "select `key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`service_instance_bindings`,`right_ip_address`,`availability_zone`,`auto_scale`,`max_instances`,`ha_mode`,`virtual_router_id`,`interface_list`,`left_ip_address`,`left_virtual_network`,`auto_policy`,`right_virtual_network`,`management_virtual_network`,`uuid`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name` from `service_instance` where uuid = ?"

func CreateServiceInstance(tx *sql.Tx, model *models.ServiceInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		util.MustJSON(model.ServiceInstanceBindings),
		string(model.ServiceInstanceProperties.RightIPAddress),
		string(model.ServiceInstanceProperties.AvailabilityZone),
		bool(model.ServiceInstanceProperties.ScaleOut.AutoScale),
		int(model.ServiceInstanceProperties.ScaleOut.MaxInstances),
		string(model.ServiceInstanceProperties.HaMode),
		string(model.ServiceInstanceProperties.VirtualRouterID),
		util.MustJSON(model.ServiceInstanceProperties.InterfaceList),
		string(model.ServiceInstanceProperties.LeftIPAddress),
		string(model.ServiceInstanceProperties.LeftVirtualNetwork),
		bool(model.ServiceInstanceProperties.AutoPolicy),
		string(model.ServiceInstanceProperties.RightVirtualNetwork),
		string(model.ServiceInstanceProperties.ManagementVirtualNetwork),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName))
	return err
}

func scanServiceInstance(rows *sql.Rows) (*models.ServiceInstance, error) {
	m := models.MakeServiceInstance()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonServiceInstanceBindings string

	var jsonServiceInstancePropertiesInterfaceList string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonServiceInstanceBindings,
		&m.ServiceInstanceProperties.RightIPAddress,
		&m.ServiceInstanceProperties.AvailabilityZone,
		&m.ServiceInstanceProperties.ScaleOut.AutoScale,
		&m.ServiceInstanceProperties.ScaleOut.MaxInstances,
		&m.ServiceInstanceProperties.HaMode,
		&m.ServiceInstanceProperties.VirtualRouterID,
		&jsonServiceInstancePropertiesInterfaceList,
		&m.ServiceInstanceProperties.LeftIPAddress,
		&m.ServiceInstanceProperties.LeftVirtualNetwork,
		&m.ServiceInstanceProperties.AutoPolicy,
		&m.ServiceInstanceProperties.RightVirtualNetwork,
		&m.ServiceInstanceProperties.ManagementVirtualNetwork,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonServiceInstanceBindings), &m.ServiceInstanceBindings)

	json.Unmarshal([]byte(jsonServiceInstancePropertiesInterfaceList), &m.ServiceInstanceProperties.InterfaceList)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListServiceInstance(tx *sql.Tx) ([]*models.ServiceInstance, error) {
	result := models.MakeServiceInstanceSlice()
	rows, err := tx.Query(listServiceInstanceQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceInstance(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceInstance(tx *sql.Tx, uuid string) (*models.ServiceInstance, error) {
	rows, err := tx.Query(showServiceInstanceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceInstance(rows)
	}
	return nil, nil
}

func UpdateServiceInstance(tx *sql.Tx, uuid string, model *models.ServiceInstance) error {
	return nil
}

func DeleteServiceInstance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
