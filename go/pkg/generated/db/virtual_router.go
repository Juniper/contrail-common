package db
// virtual_router

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualRouterQuery = "insert into `virtual_router` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`virtual_router_dpdk_enabled`,`virtual_router_type`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`virtual_router_ip_address`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualRouterQuery = "update `virtual_router` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`virtual_router_dpdk_enabled` = ?,`virtual_router_type` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`virtual_router_ip_address` = ?,`uuid` = ?;"
const deleteVirtualRouterQuery = "delete from `virtual_router`"
const selectVirtualRouterQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`virtual_router_dpdk_enabled`,`virtual_router_type`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`virtual_router_ip_address`,`uuid` from `virtual_router`"

func CreateVirtualRouter(tx *sql.Tx, model *models.VirtualRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualRouterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.VirtualRouterDPDKEnabled,
    model.VirtualRouterType,
    model.FQName,
    model.IDPerms.Creator,
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
    model.VirtualRouterIPAddress,
    model.UUID)
    return err
}

func ListVirtualRouter(tx *sql.Tx) ([]*models.VirtualRouter, error) {
    result := models.MakeVirtualRouterSlice()
    rows, err := tx.Query(selectVirtualRouterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualRouter()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.VirtualRouterDPDKEnabled,
                &m.VirtualRouterType,
                &m.FQName,
                &m.IDPerms.Creator,
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
                &m.VirtualRouterIPAddress,
                &m.UUID); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowVirtualRouter(db *sql.DB, id string, model *models.VirtualRouter) error {
    return nil
}

func UpdateVirtualRouter(db *sql.DB, id string, model *models.VirtualRouter) error {
    return nil
}

func DeleteVirtualRouter(db *sql.DB, id string) error {
    return nil
}