package db
// widget

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertWidgetQuery = "insert into `widget` (`layout_config`,`content_config`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`container_config`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateWidgetQuery = "update `widget` set `layout_config` = ?,`content_config` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`container_config` = ?;"
const deleteWidgetQuery = "delete from `widget`"
const selectWidgetQuery = "select `layout_config`,`content_config`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`container_config` from `widget`"

func CreateWidget(tx *sql.Tx, model *models.Widget) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertWidgetQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.LayoutConfig,
    model.ContentConfig,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Description,
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
    model.ContainerConfig)
    return err
}

func ListWidget(tx *sql.Tx) ([]*models.Widget, error) {
    result := models.MakeWidgetSlice()
    rows, err := tx.Query(selectWidgetQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeWidget()
            if err := rows.Scan(&m.LayoutConfig,
                &m.ContentConfig,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Description,
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
                &m.ContainerConfig); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowWidget(db *sql.DB, id string, model *models.Widget) error {
    return nil
}

func UpdateWidget(db *sql.DB, id string, model *models.Widget) error {
    return nil
}

func DeleteWidget(db *sql.DB, id string) error {
    return nil
}