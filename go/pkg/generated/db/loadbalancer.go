package db
// loadbalancer

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLoadbalancerQuery = "insert into `loadbalancer` (`loadbalancer_provider`,`display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`status`,`provisioning_status`,`admin_state`,`vip_address`,`vip_subnet_id`,`operating_status`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerQuery = "update `loadbalancer` set `loadbalancer_provider` = ?,`display_name` = ?,`key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`status` = ?,`provisioning_status` = ?,`admin_state` = ?,`vip_address` = ?,`vip_subnet_id` = ?,`operating_status` = ?;"
const deleteLoadbalancerQuery = "delete from `loadbalancer`"
const selectLoadbalancerQuery = "select `loadbalancer_provider`,`display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`status`,`provisioning_status`,`admin_state`,`vip_address`,`vip_subnet_id`,`operating_status` from `loadbalancer`"

func CreateLoadbalancer(tx *sql.Tx, model *models.Loadbalancer) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.LoadbalancerProvider,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID,
    model.FQName,
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
    model.IDPerms.Creator,
    model.LoadbalancerProperties.Status,
    model.LoadbalancerProperties.ProvisioningStatus,
    model.LoadbalancerProperties.AdminState,
    model.LoadbalancerProperties.VipAddress,
    model.LoadbalancerProperties.VipSubnetID,
    model.LoadbalancerProperties.OperatingStatus)
    return err
}

func ListLoadbalancer(tx *sql.Tx) ([]*models.Loadbalancer, error) {
    result := models.MakeLoadbalancerSlice()
    rows, err := tx.Query(selectLoadbalancerQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLoadbalancer()
            if err := rows.Scan(&m.LoadbalancerProvider,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.UUID,
                &m.FQName,
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
                &m.LoadbalancerProperties.Status,
                &m.LoadbalancerProperties.ProvisioningStatus,
                &m.LoadbalancerProperties.AdminState,
                &m.LoadbalancerProperties.VipAddress,
                &m.LoadbalancerProperties.VipSubnetID,
                &m.LoadbalancerProperties.OperatingStatus); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowLoadbalancer(db *sql.DB, id string, model *models.Loadbalancer) error {
    return nil
}

func UpdateLoadbalancer(db *sql.DB, id string, model *models.Loadbalancer) error {
    return nil
}

func DeleteLoadbalancer(db *sql.DB, id string) error {
    return nil
}