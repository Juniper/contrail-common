package db
// namespace

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertNamespaceQuery = "insert into `namespace` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`ip_prefix`,`ip_prefix_len`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNamespaceQuery = "update `namespace` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?;"
const deleteNamespaceQuery = "delete from `namespace`"
const selectNamespaceQuery = "select `key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`ip_prefix`,`ip_prefix_len`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name` from `namespace`"

func CreateNamespace(tx *sql.Tx, model *models.Namespace) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNamespaceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.NamespaceCidr.IPPrefix,
    model.NamespaceCidr.IPPrefixLen,
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
    model.DisplayName)
    return err
}

func ListNamespace(tx *sql.Tx) ([]*models.Namespace, error) {
    result := models.MakeNamespaceSlice()
    rows, err := tx.Query(selectNamespaceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeNamespace()
            if err := rows.Scan(&m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.NamespaceCidr.IPPrefix,
                &m.NamespaceCidr.IPPrefixLen,
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
                &m.DisplayName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowNamespace(db *sql.DB, id string, model *models.Namespace) error {
    return nil
}

func UpdateNamespace(db *sql.DB, id string, model *models.Namespace) error {
    return nil
}

func DeleteNamespace(db *sql.DB, id string) error {
    return nil
}