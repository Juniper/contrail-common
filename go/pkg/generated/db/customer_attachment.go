package db
// customer_attachment

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertCustomerAttachmentQuery = "insert into `customer_attachment` (`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateCustomerAttachmentQuery = "update `customer_attachment` set `created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteCustomerAttachmentQuery = "delete from `customer_attachment`"
const selectCustomerAttachmentQuery = "select `created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name` from `customer_attachment`"

func CreateCustomerAttachment(tx *sql.Tx, model *models.CustomerAttachment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertCustomerAttachmentQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.FQName)
    return err
}

func ListCustomerAttachment(tx *sql.Tx) ([]*models.CustomerAttachment, error) {
    result := models.MakeCustomerAttachmentSlice()
    rows, err := tx.Query(selectCustomerAttachmentQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeCustomerAttachment()
            if err := rows.Scan(&m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
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

func ShowCustomerAttachment(db *sql.DB, id string, model *models.CustomerAttachment) error {
    return nil
}

func UpdateCustomerAttachment(db *sql.DB, id string, model *models.CustomerAttachment) error {
    return nil
}

func DeleteCustomerAttachment(db *sql.DB, id string) error {
    return nil
}