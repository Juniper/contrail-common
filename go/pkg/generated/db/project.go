package db

// project

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertProjectQuery = "insert into `project` (`alarm_enable`,`floating_ip`,`virtual_router`,`floating_ip_pool`,`network_ipam`,`loadbalancer_member`,`security_group_rule`,`loadbalancer_pool`,`loadbalancer_healthmonitor`,`virtual_ip`,`virtual_DNS_record`,`security_logging_object`,`subnet`,`network_policy`,`access_control_list`,`defaults`,`virtual_DNS`,`virtual_machine_interface`,`global_vrouter_config`,`route_table`,`service_instance`,`instance_ip`,`virtual_network`,`logical_router`,`service_template`,`bgp_router`,`security_group`,`owner`,`owner_access`,`global_access`,`share`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`vxlan_routing`,`key_value_pair`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProjectQuery = "update `project` set `alarm_enable` = ?,`floating_ip` = ?,`virtual_router` = ?,`floating_ip_pool` = ?,`network_ipam` = ?,`loadbalancer_member` = ?,`security_group_rule` = ?,`loadbalancer_pool` = ?,`loadbalancer_healthmonitor` = ?,`virtual_ip` = ?,`virtual_DNS_record` = ?,`security_logging_object` = ?,`subnet` = ?,`network_policy` = ?,`access_control_list` = ?,`defaults` = ?,`virtual_DNS` = ?,`virtual_machine_interface` = ?,`global_vrouter_config` = ?,`route_table` = ?,`service_instance` = ?,`instance_ip` = ?,`virtual_network` = ?,`logical_router` = ?,`service_template` = ?,`bgp_router` = ?,`security_group` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`vxlan_routing` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteProjectQuery = "delete from `project` where uuid = ?"
const listProjectQuery = "select `alarm_enable`,`floating_ip`,`virtual_router`,`floating_ip_pool`,`network_ipam`,`loadbalancer_member`,`security_group_rule`,`loadbalancer_pool`,`loadbalancer_healthmonitor`,`virtual_ip`,`virtual_DNS_record`,`security_logging_object`,`subnet`,`network_policy`,`access_control_list`,`defaults`,`virtual_DNS`,`virtual_machine_interface`,`global_vrouter_config`,`route_table`,`service_instance`,`instance_ip`,`virtual_network`,`logical_router`,`service_template`,`bgp_router`,`security_group`,`owner`,`owner_access`,`global_access`,`share`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`vxlan_routing`,`key_value_pair`,`uuid`,`fq_name` from `project`"
const showProjectQuery = "select `alarm_enable`,`floating_ip`,`virtual_router`,`floating_ip_pool`,`network_ipam`,`loadbalancer_member`,`security_group_rule`,`loadbalancer_pool`,`loadbalancer_healthmonitor`,`virtual_ip`,`virtual_DNS_record`,`security_logging_object`,`subnet`,`network_policy`,`access_control_list`,`defaults`,`virtual_DNS`,`virtual_machine_interface`,`global_vrouter_config`,`route_table`,`service_instance`,`instance_ip`,`virtual_network`,`logical_router`,`service_template`,`bgp_router`,`security_group`,`owner`,`owner_access`,`global_access`,`share`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`vxlan_routing`,`key_value_pair`,`uuid`,`fq_name` from `project` where uuid = ?"

func CreateProject(tx *sql.Tx, model *models.Project) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.AlarmEnable),
		int(model.Quota.FloatingIP),
		int(model.Quota.VirtualRouter),
		int(model.Quota.FloatingIPPool),
		int(model.Quota.NetworkIpam),
		int(model.Quota.LoadbalancerMember),
		int(model.Quota.SecurityGroupRule),
		int(model.Quota.LoadbalancerPool),
		int(model.Quota.LoadbalancerHealthmonitor),
		int(model.Quota.VirtualIP),
		int(model.Quota.VirtualDNSRecord),
		int(model.Quota.SecurityLoggingObject),
		int(model.Quota.Subnet),
		int(model.Quota.NetworkPolicy),
		int(model.Quota.AccessControlList),
		int(model.Quota.Defaults),
		int(model.Quota.VirtualDNS),
		int(model.Quota.VirtualMachineInterface),
		int(model.Quota.GlobalVrouterConfig),
		int(model.Quota.RouteTable),
		int(model.Quota.ServiceInstance),
		int(model.Quota.InstanceIP),
		int(model.Quota.VirtualNetwork),
		int(model.Quota.LogicalRouter),
		int(model.Quota.ServiceTemplate),
		int(model.Quota.BGPRouter),
		int(model.Quota.SecurityGroup),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		bool(model.VxlanRouting),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		util.MustJSON(model.FQName))
	return err
}

func scanProject(rows *sql.Rows) (*models.Project, error) {
	m := models.MakeProject()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.AlarmEnable,
		&m.Quota.FloatingIP,
		&m.Quota.VirtualRouter,
		&m.Quota.FloatingIPPool,
		&m.Quota.NetworkIpam,
		&m.Quota.LoadbalancerMember,
		&m.Quota.SecurityGroupRule,
		&m.Quota.LoadbalancerPool,
		&m.Quota.LoadbalancerHealthmonitor,
		&m.Quota.VirtualIP,
		&m.Quota.VirtualDNSRecord,
		&m.Quota.SecurityLoggingObject,
		&m.Quota.Subnet,
		&m.Quota.NetworkPolicy,
		&m.Quota.AccessControlList,
		&m.Quota.Defaults,
		&m.Quota.VirtualDNS,
		&m.Quota.VirtualMachineInterface,
		&m.Quota.GlobalVrouterConfig,
		&m.Quota.RouteTable,
		&m.Quota.ServiceInstance,
		&m.Quota.InstanceIP,
		&m.Quota.VirtualNetwork,
		&m.Quota.LogicalRouter,
		&m.Quota.ServiceTemplate,
		&m.Quota.BGPRouter,
		&m.Quota.SecurityGroup,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.DisplayName,
		&m.VxlanRouting,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListProject(tx *sql.Tx) ([]*models.Project, error) {
	result := models.MakeProjectSlice()
	rows, err := tx.Query(listProjectQuery)
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
