package db
// bgp_as_a_service

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`global_access`,`share`,`owner`,`owner_access`,`uuid`,`bgpaas_session_attributes`,`bgpaas_suppress_route_advertisement`,`bgpaas_ip_address`,`key_value_pair`,`display_name`,`fq_name`,`bgpaas_shared`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`autonomous_system`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPAsAServiceQuery = "update `bgp_as_a_service` set `global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`bgpaas_session_attributes` = ?,`bgpaas_suppress_route_advertisement` = ?,`bgpaas_ip_address` = ?,`key_value_pair` = ?,`display_name` = ?,`fq_name` = ?,`bgpaas_shared` = ?,`bgpaas_ipv4_mapped_ipv6_nexthop` = ?,`autonomous_system` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?;"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service`"
const selectBGPAsAServiceQuery = "select `global_access`,`share`,`owner`,`owner_access`,`uuid`,`bgpaas_session_attributes`,`bgpaas_suppress_route_advertisement`,`bgpaas_ip_address`,`key_value_pair`,`display_name`,`fq_name`,`bgpaas_shared`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`autonomous_system`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description` from `bgp_as_a_service`"

func CreateBGPAsAService(tx *sql.Tx, model *models.BGPAsAService) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.BgpaasSessionAttributes,
    model.BgpaasSuppressRouteAdvertisement,
    model.BgpaasIPAddress,
    model.Annotations.KeyValuePair,
    model.DisplayName,
    model.FQName,
    model.BgpaasShared,
    model.BgpaasIpv4MappedIpv6Nexthop,
    model.AutonomousSystem,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.IDPerms.Description)
    return err
}

func ListBGPAsAService(tx *sql.Tx) ([]*models.BGPAsAService, error) {
    result := models.MakeBGPAsAServiceSlice()
    rows, err := tx.Query(selectBGPAsAServiceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeBGPAsAService()
            if err := rows.Scan(&m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.UUID,
                &m.BgpaasSessionAttributes,
                &m.BgpaasSuppressRouteAdvertisement,
                &m.BgpaasIPAddress,
                &m.Annotations.KeyValuePair,
                &m.DisplayName,
                &m.FQName,
                &m.BgpaasShared,
                &m.BgpaasIpv4MappedIpv6Nexthop,
                &m.AutonomousSystem,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Enable,
                &m.IDPerms.Description); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowBGPAsAService(db *sql.DB, id string, model *models.BGPAsAService) error {
    return nil
}

func UpdateBGPAsAService(db *sql.DB, id string, model *models.BGPAsAService) error {
    return nil
}

func DeleteBGPAsAService(db *sql.DB, id string) error {
    return nil
}