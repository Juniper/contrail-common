package db

// project

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertProjectQuery = "insert into `project` (`subnet`,`virtual_DNS`,`service_template`,`floating_ip`,`network_ipam`,`virtual_ip`,`virtual_network`,`security_group`,`virtual_machine_interface`,`loadbalancer_member`,`instance_ip`,`floating_ip_pool`,`virtual_router`,`virtual_DNS_record`,`logical_router`,`service_instance`,`bgp_router`,`loadbalancer_healthmonitor`,`network_policy`,`security_group_rule`,`defaults`,`security_logging_object`,`global_vrouter_config`,`loadbalancer_pool`,`route_table`,`access_control_list`,`key_value_pair`,`uuid`,`fq_name`,`display_name`,`alarm_enable`,`owner_access`,`global_access`,`share`,`owner`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`vxlan_routing`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProjectQuery = "update `project` set `subnet` = ?,`virtual_DNS` = ?,`service_template` = ?,`floating_ip` = ?,`network_ipam` = ?,`virtual_ip` = ?,`virtual_network` = ?,`security_group` = ?,`virtual_machine_interface` = ?,`loadbalancer_member` = ?,`instance_ip` = ?,`floating_ip_pool` = ?,`virtual_router` = ?,`virtual_DNS_record` = ?,`logical_router` = ?,`service_instance` = ?,`bgp_router` = ?,`loadbalancer_healthmonitor` = ?,`network_policy` = ?,`security_group_rule` = ?,`defaults` = ?,`security_logging_object` = ?,`global_vrouter_config` = ?,`loadbalancer_pool` = ?,`route_table` = ?,`access_control_list` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?,`alarm_enable` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`vxlan_routing` = ?;"
const deleteProjectQuery = "delete from `project` where uuid = ?"
const listProjectQuery = "select `subnet`,`virtual_DNS`,`service_template`,`floating_ip`,`network_ipam`,`virtual_ip`,`virtual_network`,`security_group`,`virtual_machine_interface`,`loadbalancer_member`,`instance_ip`,`floating_ip_pool`,`virtual_router`,`virtual_DNS_record`,`logical_router`,`service_instance`,`bgp_router`,`loadbalancer_healthmonitor`,`network_policy`,`security_group_rule`,`defaults`,`security_logging_object`,`global_vrouter_config`,`loadbalancer_pool`,`route_table`,`access_control_list`,`key_value_pair`,`uuid`,`fq_name`,`display_name`,`alarm_enable`,`owner_access`,`global_access`,`share`,`owner`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`vxlan_routing` from `project`"
const showProjectQuery = "select `subnet`,`virtual_DNS`,`service_template`,`floating_ip`,`network_ipam`,`virtual_ip`,`virtual_network`,`security_group`,`virtual_machine_interface`,`loadbalancer_member`,`instance_ip`,`floating_ip_pool`,`virtual_router`,`virtual_DNS_record`,`logical_router`,`service_instance`,`bgp_router`,`loadbalancer_healthmonitor`,`network_policy`,`security_group_rule`,`defaults`,`security_logging_object`,`global_vrouter_config`,`loadbalancer_pool`,`route_table`,`access_control_list`,`key_value_pair`,`uuid`,`fq_name`,`display_name`,`alarm_enable`,`owner_access`,`global_access`,`share`,`owner`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`vxlan_routing` from `project` where uuid = ?"

func CreateProject(tx *sql.Tx, model *models.Project) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Quota.Subnet),
		int(model.Quota.VirtualDNS),
		int(model.Quota.ServiceTemplate),
		int(model.Quota.FloatingIP),
		int(model.Quota.NetworkIpam),
		int(model.Quota.VirtualIP),
		int(model.Quota.VirtualNetwork),
		int(model.Quota.SecurityGroup),
		int(model.Quota.VirtualMachineInterface),
		int(model.Quota.LoadbalancerMember),
		int(model.Quota.InstanceIP),
		int(model.Quota.FloatingIPPool),
		int(model.Quota.VirtualRouter),
		int(model.Quota.VirtualDNSRecord),
		int(model.Quota.LogicalRouter),
		int(model.Quota.ServiceInstance),
		int(model.Quota.BGPRouter),
		int(model.Quota.LoadbalancerHealthmonitor),
		int(model.Quota.NetworkPolicy),
		int(model.Quota.SecurityGroupRule),
		int(model.Quota.Defaults),
		int(model.Quota.SecurityLoggingObject),
		int(model.Quota.GlobalVrouterConfig),
		int(model.Quota.LoadbalancerPool),
		int(model.Quota.RouteTable),
		int(model.Quota.AccessControlList),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.DisplayName),
		bool(model.AlarmEnable),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.VxlanRouting))
	return err
}

func scanProject(rows *sql.Rows) (*models.Project, error) {
	m := models.MakeProject()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&m.Quota.Subnet,
		&m.Quota.VirtualDNS,
		&m.Quota.ServiceTemplate,
		&m.Quota.FloatingIP,
		&m.Quota.NetworkIpam,
		&m.Quota.VirtualIP,
		&m.Quota.VirtualNetwork,
		&m.Quota.SecurityGroup,
		&m.Quota.VirtualMachineInterface,
		&m.Quota.LoadbalancerMember,
		&m.Quota.InstanceIP,
		&m.Quota.FloatingIPPool,
		&m.Quota.VirtualRouter,
		&m.Quota.VirtualDNSRecord,
		&m.Quota.LogicalRouter,
		&m.Quota.ServiceInstance,
		&m.Quota.BGPRouter,
		&m.Quota.LoadbalancerHealthmonitor,
		&m.Quota.NetworkPolicy,
		&m.Quota.SecurityGroupRule,
		&m.Quota.Defaults,
		&m.Quota.SecurityLoggingObject,
		&m.Quota.GlobalVrouterConfig,
		&m.Quota.LoadbalancerPool,
		&m.Quota.RouteTable,
		&m.Quota.AccessControlList,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName,
		&m.DisplayName,
		&m.AlarmEnable,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
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
		&m.IDPerms.Creator,
		&m.VxlanRouting); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createProjectWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListProject(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Project, error) {
	result := models.MakeProjectSlice()
	whereQuery, values := createProjectWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listProjectQuery)
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
		m, _ := scanProject(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowProject(tx *sql.Tx, uuid string) (*models.Project, error) {
	rows, err := tx.Query(showProjectQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanProject(rows)
	}
	return nil, nil
}

func UpdateProject(tx *sql.Tx, uuid string, model *models.Project) error {
	return nil
}

func DeleteProject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteProjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
