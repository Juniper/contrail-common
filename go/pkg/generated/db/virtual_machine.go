package db
// virtual_machine

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualMachineQuery = "insert into `virtual_machine` (`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineQuery = "update `virtual_machine` set `uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?;"
const deleteVirtualMachineQuery = "delete from `virtual_machine`"
const selectVirtualMachineQuery = "select `uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access` from `virtual_machine`"

func CreateVirtualMachine(tx *sql.Tx, model *models.VirtualMachine) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineQuery)
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
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Enable,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess)
    return err
}

func ListVirtualMachine(tx *sql.Tx) ([]*models.VirtualMachine, error) {
    result := models.MakeVirtualMachineSlice()
    rows, err := tx.Query(selectVirtualMachineQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualMachine()
            if err := rows.Scan(&m.UUID,
                &m.FQName,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Enable,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowVirtualMachine(db *sql.DB, id string, model *models.VirtualMachine) error {
    return nil
}

func UpdateVirtualMachine(db *sql.DB, id string, model *models.VirtualMachine) error {
    return nil
}

func DeleteVirtualMachine(db *sql.DB, id string) error {
    return nil
}