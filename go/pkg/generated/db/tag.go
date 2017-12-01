package db
// tag

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertTagQuery = "insert into `tag` (`uuid`,`tag_id`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`display_name`,`key_value_pair`,`tag_type_name`,`tag_value`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateTagQuery = "update `tag` set `uuid` = ?,`tag_id` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`display_name` = ?,`key_value_pair` = ?,`tag_type_name` = ?,`tag_value` = ?,`fq_name` = ?;"
const deleteTagQuery = "delete from `tag`"
const selectTagQuery = "select `uuid`,`tag_id`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`display_name`,`key_value_pair`,`tag_type_name`,`tag_value`,`fq_name` from `tag`"

func CreateTag(tx *sql.Tx, model *models.Tag) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertTagQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.UUID,
    model.TagID,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.TagTypeName,
    model.TagValue,
    model.FQName)
    return err
}

func ListTag(tx *sql.Tx) ([]*models.Tag, error) {
    result := models.MakeTagSlice()
    rows, err := tx.Query(selectTagQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeTag()
            if err := rows.Scan(&m.UUID,
                &m.TagID,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.TagTypeName,
                &m.TagValue,
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

func ShowTag(db *sql.DB, id string, model *models.Tag) error {
    return nil
}

func UpdateTag(db *sql.DB, id string, model *models.Tag) error {
    return nil
}

func DeleteTag(db *sql.DB, id string) error {
    return nil
}