package db
// floating_ip_pool

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertFloatingIPPoolQuery = "insert into `floating_ip_pool` (`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`subnet_uuid`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPPoolQuery = "update `floating_ip_pool` set `uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`subnet_uuid` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteFloatingIPPoolQuery = "delete from `floating_ip_pool`"
const selectFloatingIPPoolQuery = "select `uuid`,`fq_name`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`subnet_uuid`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access` from `floating_ip_pool`"

func CreateFloatingIPPool(tx *sql.Tx, model *models.FloatingIPPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPPoolQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.UUID,
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
    model.FloatingIPPoolSubnets.SubnetUUID,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess)
    return err
}

func ListFloatingIPPool(tx *sql.Tx) ([]*models.FloatingIPPool, error) {
    result := models.MakeFloatingIPPoolSlice()
    rows, err := tx.Query(selectFloatingIPPoolQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeFloatingIPPool()
            if err := rows.Scan(&m.UUID,
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
                &m.FloatingIPPoolSubnets.SubnetUUID,
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

func ShowFloatingIPPool(db *sql.DB, id string, model *models.FloatingIPPool) error {
    return nil
}

func UpdateFloatingIPPool(db *sql.DB, id string, model *models.FloatingIPPool) error {
    return nil
}

func DeleteFloatingIPPool(db *sql.DB, id string) error {
    return nil
}