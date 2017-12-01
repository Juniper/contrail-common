package db
// bgp_router

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertBGPRouterQuery = "insert into `bgp_router` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPRouterQuery = "update `bgp_router` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteBGPRouterQuery = "delete from `bgp_router`"
const selectBGPRouterQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name`,`key_value_pair` from `bgp_router`"

func CreateBGPRouter(tx *sql.Tx, model *models.BGPRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPRouterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
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
    model.IDPerms.Description,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListBGPRouter(tx *sql.Tx) ([]*models.BGPRouter, error) {
    result := models.MakeBGPRouterSlice()
    rows, err := tx.Query(selectBGPRouterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeBGPRouter()
            if err := rows.Scan(&m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
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
                &m.IDPerms.Description,
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

func ShowBGPRouter(db *sql.DB, id string, model *models.BGPRouter) error {
    return nil
}

func UpdateBGPRouter(db *sql.DB, id string, model *models.BGPRouter) error {
    return nil
}

func DeleteBGPRouter(db *sql.DB, id string) error {
    return nil
}