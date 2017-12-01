package db

// service_instance

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertServiceInstanceQuery = "insert into `service_instance` (`ha_mode`,`virtual_router_id`,`interface_list`,`right_virtual_network`,`management_virtual_network`,`left_ip_address`,`left_virtual_network`,`auto_policy`,`right_ip_address`,`availability_zone`,`auto_scale`,`max_instances`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`service_instance_bindings`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceInstanceQuery = "update `service_instance` set `ha_mode` = ?,`virtual_router_id` = ?,`interface_list` = ?,`right_virtual_network` = ?,`management_virtual_network` = ?,`left_ip_address` = ?,`left_virtual_network` = ?,`auto_policy` = ?,`right_ip_address` = ?,`availability_zone` = ?,`auto_scale` = ?,`max_instances` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`service_instance_bindings` = ?;"
const deleteServiceInstanceQuery = "delete from `service_instance` where uuid = ?"
const listServiceInstanceQuery = "select `ha_mode`,`virtual_router_id`,`interface_list`,`right_virtual_network`,`management_virtual_network`,`left_ip_address`,`left_virtual_network`,`auto_policy`,`right_ip_address`,`availability_zone`,`auto_scale`,`max_instances`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`service_instance_bindings` from `service_instance`"
const showServiceInstanceQuery = "select `ha_mode`,`virtual_router_id`,`interface_list`,`right_virtual_network`,`management_virtual_network`,`left_ip_address`,`left_virtual_network`,`auto_policy`,`right_ip_address`,`availability_zone`,`auto_scale`,`max_instances`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`,`service_instance_bindings` from `service_instance` where uuid = ?"

func CreateServiceInstance(tx *sql.Tx, model *models.ServiceInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ServiceInstanceProperties.HaMode),
		string(model.ServiceInstanceProperties.VirtualRouterID),
		util.MustJSON(model.ServiceInstanceProperties.InterfaceList),
		string(model.ServiceInstanceProperties.RightVirtualNetwork),
		string(model.ServiceInstanceProperties.ManagementVirtualNetwork),
		string(model.ServiceInstanceProperties.LeftIPAddress),
		string(model.ServiceInstanceProperties.LeftVirtualNetwork),
		bool(model.ServiceInstanceProperties.AutoPolicy),
		string(model.ServiceInstanceProperties.RightIPAddress),
		string(model.ServiceInstanceProperties.AvailabilityZone),
		bool(model.ServiceInstanceProperties.ScaleOut.AutoScale),
		int(model.ServiceInstanceProperties.ScaleOut.MaxInstances),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.DisplayName),
		util.MustJSON(model.ServiceInstanceBindings))
	return err
}

func scanServiceInstance(rows *sql.Rows) (*models.ServiceInstance, error) {
	m := models.MakeServiceInstance()

	var jsonServiceInstancePropertiesInterfaceList string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonServiceInstanceBindings string

	if err := rows.Scan(&m.ServiceInstanceProperties.HaMode,
		&m.ServiceInstanceProperties.VirtualRouterID,
		&jsonServiceInstancePropertiesInterfaceList,
		&m.ServiceInstanceProperties.RightVirtualNetwork,
		&m.ServiceInstanceProperties.ManagementVirtualNetwork,
		&m.ServiceInstanceProperties.LeftIPAddress,
		&m.ServiceInstanceProperties.LeftVirtualNetwork,
		&m.ServiceInstanceProperties.AutoPolicy,
		&m.ServiceInstanceProperties.RightIPAddress,
		&m.ServiceInstanceProperties.AvailabilityZone,
		&m.ServiceInstanceProperties.ScaleOut.AutoScale,
		&m.ServiceInstanceProperties.ScaleOut.MaxInstances,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
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
		&m.DisplayName,
		&jsonServiceInstanceBindings); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonServiceInstancePropertiesInterfaceList), &m.ServiceInstanceProperties.InterfaceList)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonServiceInstanceBindings), &m.ServiceInstanceBindings)

	return m, nil
}

func createServiceInstanceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["ha_mode"]; ok {
		results = append(results, "ha_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_router_id"]; ok {
		results = append(results, "virtual_router_id = ?")
		values = append(values, value)
	}

	if value, ok := where["right_virtual_network"]; ok {
		results = append(results, "right_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["management_virtual_network"]; ok {
		results = append(results, "management_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["left_ip_address"]; ok {
		results = append(results, "left_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["left_virtual_network"]; ok {
		results = append(results, "left_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["right_ip_address"]; ok {
		results = append(results, "right_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["availability_zone"]; ok {
		results = append(results, "availability_zone = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceInstance(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceInstance, error) {
	result := models.MakeServiceInstanceSlice()
	whereQuery, values := createServiceInstanceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceInstanceQuery)
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
