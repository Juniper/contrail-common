package db
// e2_service_provider

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertE2ServiceProviderQuery = "insert into `e2_service_provider` (`e2_service_provider_promiscuous`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateE2ServiceProviderQuery = "update `e2_service_provider` set `e2_service_provider_promiscuous` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteE2ServiceProviderQuery = "delete from `e2_service_provider`"
const selectE2ServiceProviderQuery = "select `e2_service_provider_promiscuous`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `e2_service_provider`"

func CreateE2ServiceProvider(tx *sql.Tx, model *models.E2ServiceProvider) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertE2ServiceProviderQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.E2ServiceProviderPromiscuous,
    model.UUID,
    model.FQName,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess)
    return err
}

func ListE2ServiceProvider(tx *sql.Tx) ([]*models.E2ServiceProvider, error) {
    result := models.MakeE2ServiceProviderSlice()
    rows, err := tx.Query(selectE2ServiceProviderQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeE2ServiceProvider()
            if err := rows.Scan(&m.E2ServiceProviderPromiscuous,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowE2ServiceProvider(db *sql.DB, id string, model *models.E2ServiceProvider) error {
    return nil
}

func UpdateE2ServiceProvider(db *sql.DB, id string, model *models.E2ServiceProvider) error {
    return nil
}

func DeleteE2ServiceProvider(db *sql.DB, id string) error {
    return nil
}