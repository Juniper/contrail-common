package db
// alias_ip

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertAliasIPQuery = "insert into `alias_ip` (`fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`alias_ip_address`,`alias_ip_address_family`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAliasIPQuery = "update `alias_ip` set `fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`alias_ip_address` = ?,`alias_ip_address_family` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?;"
const deleteAliasIPQuery = "delete from `alias_ip`"
const selectAliasIPQuery = "select `fq_name`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`alias_ip_address`,`alias_ip_address_family`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid` from `alias_ip`"

func CreateAliasIP(tx *sql.Tx, model *models.AliasIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAliasIPQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
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
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.AliasIPAddress,
    model.AliasIPAddressFamily,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID)
    return err
}

func ListAliasIP(tx *sql.Tx) ([]*models.AliasIP, error) {
    result := models.MakeAliasIPSlice()
    rows, err := tx.Query(selectAliasIPQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeAliasIP()
            if err := rows.Scan(&m.FQName,
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
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.AliasIPAddress,
                &m.AliasIPAddressFamily,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
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

func ShowAliasIP(db *sql.DB, id string, model *models.AliasIP) error {
    return nil
}

func UpdateAliasIP(db *sql.DB, id string, model *models.AliasIP) error {
    return nil
}

func DeleteAliasIP(db *sql.DB, id string) error {
    return nil
}