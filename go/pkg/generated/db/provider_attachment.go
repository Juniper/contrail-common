package db
// provider_attachment

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertProviderAttachmentQuery = "insert into `provider_attachment` (`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProviderAttachmentQuery = "update `provider_attachment` set `display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteProviderAttachmentQuery = "delete from `provider_attachment`"
const selectProviderAttachmentQuery = "select `display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible` from `provider_attachment`"

func CreateProviderAttachment(tx *sql.Tx, model *models.ProviderAttachment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProviderAttachmentQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible)
    return err
}

func ListProviderAttachment(tx *sql.Tx) ([]*models.ProviderAttachment, error) {
    result := models.MakeProviderAttachmentSlice()
    rows, err := tx.Query(selectProviderAttachmentQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeProviderAttachment()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowProviderAttachment(db *sql.DB, id string, model *models.ProviderAttachment) error {
    return nil
}

func UpdateProviderAttachment(db *sql.DB, id string, model *models.ProviderAttachment) error {
    return nil
}

func DeleteProviderAttachment(db *sql.DB, id string) error {
    return nil
}