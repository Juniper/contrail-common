package db
// dashboard

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertDashboardQuery = "insert into `dashboard` (`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`container_config`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDashboardQuery = "update `dashboard` set `fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`container_config` = ?,`uuid` = ?;"
const deleteDashboardQuery = "delete from `dashboard`"
const selectDashboardQuery = "select `fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`container_config`,`uuid` from `dashboard`"

func CreateDashboard(tx *sql.Tx, model *models.Dashboard) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDashboardQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.IDPerms.Created,
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
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.ContainerConfig,
    model.UUID)
    return err
}

func ListDashboard(tx *sql.Tx) ([]*models.Dashboard, error) {
    result := models.MakeDashboardSlice()
    rows, err := tx.Query(selectDashboardQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeDashboard()
            if err := rows.Scan(&m.FQName,
                &m.IDPerms.Created,
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
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.ContainerConfig,
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

func ShowDashboard(db *sql.DB, id string, model *models.Dashboard) error {
    return nil
}

func UpdateDashboard(db *sql.DB, id string, model *models.Dashboard) error {
    return nil
}

func DeleteDashboard(db *sql.DB, id string) error {
    return nil
}