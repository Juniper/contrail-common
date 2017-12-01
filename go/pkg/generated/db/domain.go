package db
// domain

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertDomainQuery = "insert into `domain` (`virtual_network_limit`,`security_group_limit`,`project_limit`,`display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDomainQuery = "update `domain` set `virtual_network_limit` = ?,`security_group_limit` = ?,`project_limit` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?;"
const deleteDomainQuery = "delete from `domain`"
const selectDomainQuery = "select `virtual_network_limit`,`security_group_limit`,`project_limit`,`display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable` from `domain`"

func CreateDomain(tx *sql.Tx, model *models.Domain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDomainQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DomainLimits.VirtualNetworkLimit,
    model.DomainLimits.SecurityGroupLimit,
    model.DomainLimits.ProjectLimit,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable)
    return err
}

func ListDomain(tx *sql.Tx) ([]*models.Domain, error) {
    result := models.MakeDomainSlice()
    rows, err := tx.Query(selectDomainQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeDomain()
            if err := rows.Scan(&m.DomainLimits.VirtualNetworkLimit,
                &m.DomainLimits.SecurityGroupLimit,
                &m.DomainLimits.ProjectLimit,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowDomain(db *sql.DB, id string, model *models.Domain) error {
    return nil
}

func UpdateDomain(db *sql.DB, id string, model *models.Domain) error {
    return nil
}

func DeleteDomain(db *sql.DB, id string) error {
    return nil
}