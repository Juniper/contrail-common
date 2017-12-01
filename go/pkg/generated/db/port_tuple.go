package db
// port_tuple

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertPortTupleQuery = "insert into `port_tuple` (`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePortTupleQuery = "update `port_tuple` set `share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deletePortTupleQuery = "delete from `port_tuple`"
const selectPortTupleQuery = "select `share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`display_name`,`key_value_pair` from `port_tuple`"

func CreatePortTuple(tx *sql.Tx, model *models.PortTuple) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPortTupleQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListPortTuple(tx *sql.Tx) ([]*models.PortTuple, error) {
    result := models.MakePortTupleSlice()
    rows, err := tx.Query(selectPortTupleQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakePortTuple()
            if err := rows.Scan(&m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.DisplayName,
                &m.Annotations.KeyValuePair); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowPortTuple(db *sql.DB, id string, model *models.PortTuple) error {
    return nil
}

func UpdatePortTuple(db *sql.DB, id string, model *models.PortTuple) error {
    return nil
}

func DeletePortTuple(db *sql.DB, id string) error {
    return nil
}