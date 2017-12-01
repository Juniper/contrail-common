package db
// alias_ip_pool

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAliasIPPoolQuery = "insert into `alias_ip_pool` (`fq_name`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAliasIPPoolQuery = "update `alias_ip_pool` set `fq_name` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?;"
const deleteAliasIPPoolQuery = "delete from `alias_ip_pool`"
const selectAliasIPPoolQuery = "select `fq_name`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid` from `alias_ip_pool`"

func CreateAliasIPPool(tx *sql.Tx, model *models.AliasIPPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAliasIPPoolQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.UUID)
    return err
}

func ListAliasIPPool(tx *sql.Tx) ([]*models.AliasIPPool, error) {
    result := models.MakeAliasIPPoolSlice()
    rows, err := tx.Query(selectAliasIPPoolQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAliasIPPool()
            if err := rows.Scan(&m.FQName,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
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

func ShowAliasIPPool(db *sql.DB, id string, model *models.AliasIPPool) error {
    return nil
}

func UpdateAliasIPPool(db *sql.DB, id string, model *models.AliasIPPool) error {
    return nil
}

func DeleteAliasIPPool(db *sql.DB, id string) error {
    return nil
}