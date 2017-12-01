package db
// physical_interface

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertPhysicalInterfaceQuery = "insert into `physical_interface` (`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`ethernet_segment_identifier`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalInterfaceQuery = "update `physical_interface` set `uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`ethernet_segment_identifier` = ?;"
const deletePhysicalInterfaceQuery = "delete from `physical_interface`"
const selectPhysicalInterfaceQuery = "select `uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`ethernet_segment_identifier` from `physical_interface`"

func CreatePhysicalInterface(tx *sql.Tx, model *models.PhysicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalInterfaceQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.UUID,
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
    model.IDPerms.Enable,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.EthernetSegmentIdentifier)
    return err
}

func ListPhysicalInterface(tx *sql.Tx) ([]*models.PhysicalInterface, error) {
    result := models.MakePhysicalInterfaceSlice()
    rows, err := tx.Query(selectPhysicalInterfaceQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakePhysicalInterface()
            if err := rows.Scan(&m.UUID,
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
                &m.IDPerms.Enable,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.EthernetSegmentIdentifier); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowPhysicalInterface(db *sql.DB, id string, model *models.PhysicalInterface) error {
    return nil
}

func UpdatePhysicalInterface(db *sql.DB, id string, model *models.PhysicalInterface) error {
    return nil
}

func DeletePhysicalInterface(db *sql.DB, id string) error {
    return nil
}