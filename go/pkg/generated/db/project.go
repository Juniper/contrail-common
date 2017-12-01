package db
// project

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertProjectQuery = "insert into `project` (`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`vxlan_routing`,`alarm_enable`,`service_template`,`virtual_router`,`global_vrouter_config`,`defaults`,`floating_ip`,`service_instance`,`instance_ip`,`virtual_network`,`security_logging_object`,`virtual_DNS`,`bgp_router`,`route_table`,`loadbalancer_pool`,`virtual_DNS_record`,`security_group_rule`,`access_control_list`,`virtual_ip`,`network_policy`,`loadbalancer_healthmonitor`,`virtual_machine_interface`,`subnet`,`security_group`,`floating_ip_pool`,`network_ipam`,`logical_router`,`loadbalancer_member`,`uuid`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProjectQuery = "update `project` set `creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`vxlan_routing` = ?,`alarm_enable` = ?,`service_template` = ?,`virtual_router` = ?,`global_vrouter_config` = ?,`defaults` = ?,`floating_ip` = ?,`service_instance` = ?,`instance_ip` = ?,`virtual_network` = ?,`security_logging_object` = ?,`virtual_DNS` = ?,`bgp_router` = ?,`route_table` = ?,`loadbalancer_pool` = ?,`virtual_DNS_record` = ?,`security_group_rule` = ?,`access_control_list` = ?,`virtual_ip` = ?,`network_policy` = ?,`loadbalancer_healthmonitor` = ?,`virtual_machine_interface` = ?,`subnet` = ?,`security_group` = ?,`floating_ip_pool` = ?,`network_ipam` = ?,`logical_router` = ?,`loadbalancer_member` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteProjectQuery = "delete from `project`"
const selectProjectQuery = "select `creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`vxlan_routing`,`alarm_enable`,`service_template`,`virtual_router`,`global_vrouter_config`,`defaults`,`floating_ip`,`service_instance`,`instance_ip`,`virtual_network`,`security_logging_object`,`virtual_DNS`,`bgp_router`,`route_table`,`loadbalancer_pool`,`virtual_DNS_record`,`security_group_rule`,`access_control_list`,`virtual_ip`,`network_policy`,`loadbalancer_healthmonitor`,`virtual_machine_interface`,`subnet`,`security_group`,`floating_ip_pool`,`network_ipam`,`logical_router`,`loadbalancer_member`,`uuid`,`fq_name`,`display_name`,`key_value_pair` from `project`"

func CreateProject(tx *sql.Tx, model *models.Project) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProjectQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.VxlanRouting,
    model.AlarmEnable,
    model.Quota.ServiceTemplate,
    model.Quota.VirtualRouter,
    model.Quota.GlobalVrouterConfig,
    model.Quota.Defaults,
    model.Quota.FloatingIP,
    model.Quota.ServiceInstance,
    model.Quota.InstanceIP,
    model.Quota.VirtualNetwork,
    model.Quota.SecurityLoggingObject,
    model.Quota.VirtualDNS,
    model.Quota.BGPRouter,
    model.Quota.RouteTable,
    model.Quota.LoadbalancerPool,
    model.Quota.VirtualDNSRecord,
    model.Quota.SecurityGroupRule,
    model.Quota.AccessControlList,
    model.Quota.VirtualIP,
    model.Quota.NetworkPolicy,
    model.Quota.LoadbalancerHealthmonitor,
    model.Quota.VirtualMachineInterface,
    model.Quota.Subnet,
    model.Quota.SecurityGroup,
    model.Quota.FloatingIPPool,
    model.Quota.NetworkIpam,
    model.Quota.LogicalRouter,
    model.Quota.LoadbalancerMember,
    model.UUID,
    model.FQName,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListProject(tx *sql.Tx) ([]*models.Project, error) {
    result := models.MakeProjectSlice()
    rows, err := tx.Query(selectProjectQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeProject()
            if err := rows.Scan(&m.IDPerms.Creator,
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
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.VxlanRouting,
                &m.AlarmEnable,
                &m.Quota.ServiceTemplate,
                &m.Quota.VirtualRouter,
                &m.Quota.GlobalVrouterConfig,
                &m.Quota.Defaults,
                &m.Quota.FloatingIP,
                &m.Quota.ServiceInstance,
                &m.Quota.InstanceIP,
                &m.Quota.VirtualNetwork,
                &m.Quota.SecurityLoggingObject,
                &m.Quota.VirtualDNS,
                &m.Quota.BGPRouter,
                &m.Quota.RouteTable,
                &m.Quota.LoadbalancerPool,
                &m.Quota.VirtualDNSRecord,
                &m.Quota.SecurityGroupRule,
                &m.Quota.AccessControlList,
                &m.Quota.VirtualIP,
                &m.Quota.NetworkPolicy,
                &m.Quota.LoadbalancerHealthmonitor,
                &m.Quota.VirtualMachineInterface,
                &m.Quota.Subnet,
                &m.Quota.SecurityGroup,
                &m.Quota.FloatingIPPool,
                &m.Quota.NetworkIpam,
                &m.Quota.LogicalRouter,
                &m.Quota.LoadbalancerMember,
                &m.UUID,
                &m.FQName,
                &m.DisplayName,
                &m.Annotations.KeyValuePair); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowProject(db *sql.DB, id string, model *models.Project) error {
    return nil
}

func UpdateProject(db *sql.DB, id string, model *models.Project) error {
    return nil
}

func DeleteProject(db *sql.DB, id string) error {
    return nil
}