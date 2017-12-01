package db
// bgpvpn

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertBGPVPNQuery = "insert into `bgpvpn` (`route_target`,`display_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`import_route_target_list_route_target`,`export_route_target_list_route_target`,`bgpvpn_type`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPVPNQuery = "update `bgpvpn` set `route_target` = ?,`display_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`import_route_target_list_route_target` = ?,`export_route_target_list_route_target` = ?,`bgpvpn_type` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteBGPVPNQuery = "delete from `bgpvpn`"
const selectBGPVPNQuery = "select `route_target`,`display_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`import_route_target_list_route_target`,`export_route_target_list_route_target`,`bgpvpn_type`,`uuid`,`fq_name` from `bgpvpn`"

func CreateBGPVPN(tx *sql.Tx, model *models.BGPVPN) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPVPNQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.RouteTargetList.RouteTarget,
    model.DisplayName,
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
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.ImportRouteTargetList.RouteTarget,
    model.ExportRouteTargetList.RouteTarget,
    model.BGPVPNType,
    model.UUID,
    model.FQName)
    return err
}

func ListBGPVPN(tx *sql.Tx) ([]*models.BGPVPN, error) {
    result := models.MakeBGPVPNSlice()
    rows, err := tx.Query(selectBGPVPNQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeBGPVPN()
            if err := rows.Scan(&m.RouteTargetList.RouteTarget,
                &m.DisplayName,
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
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.ImportRouteTargetList.RouteTarget,
                &m.ExportRouteTargetList.RouteTarget,
                &m.BGPVPNType,
                &m.UUID,
                &m.FQName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowBGPVPN(db *sql.DB, id string, model *models.BGPVPN) error {
    return nil
}

func UpdateBGPVPN(db *sql.DB, id string, model *models.BGPVPN) error {
    return nil
}

func DeleteBGPVPN(db *sql.DB, id string) error {
    return nil
}