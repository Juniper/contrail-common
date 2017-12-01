package db
// service_connection_module

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertServiceConnectionModuleQuery = "insert into `service_connection_module` (`service_type`,`e2_service`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceConnectionModuleQuery = "update `service_connection_module` set `service_type` = ?,`e2_service` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteServiceConnectionModuleQuery = "delete from `service_connection_module`"
const selectServiceConnectionModuleQuery = "select `service_type`,`e2_service`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `service_connection_module`"

func CreateServiceConnectionModule(tx *sql.Tx, model *models.ServiceConnectionModule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceConnectionModuleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ServiceType,
    model.E2Service,
    model.UUID,
    model.FQName,
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
    model.IDPerms.Creator,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner)
    return err
}

func ListServiceConnectionModule(tx *sql.Tx) ([]*models.ServiceConnectionModule, error) {
    result := models.MakeServiceConnectionModuleSlice()
    rows, err := tx.Query(selectServiceConnectionModuleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeServiceConnectionModule()
            if err := rows.Scan(&m.ServiceType,
                &m.E2Service,
                &m.UUID,
                &m.FQName,
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
                &m.IDPerms.Creator,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowServiceConnectionModule(db *sql.DB, id string, model *models.ServiceConnectionModule) error {
    return nil
}

func UpdateServiceConnectionModule(db *sql.DB, id string, model *models.ServiceConnectionModule) error {
    return nil
}

func DeleteServiceConnectionModule(db *sql.DB, id string) error {
    return nil
}