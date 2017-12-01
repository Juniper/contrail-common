package db
// vpn_group

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVPNGroupQuery = "insert into `vpn_group` (`fq_name`,`provisioning_progress_stage`,`provisioning_start_time`,`provisioning_state`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`provisioning_log`,`provisioning_progress`,`type`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVPNGroupQuery = "update `vpn_group` set `fq_name` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`display_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`type` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`key_value_pair` = ?;"
const deleteVPNGroupQuery = "delete from `vpn_group`"
const selectVPNGroupQuery = "select `fq_name`,`provisioning_progress_stage`,`provisioning_start_time`,`provisioning_state`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`provisioning_log`,`provisioning_progress`,`type`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair` from `vpn_group`"

func CreateVPNGroup(tx *sql.Tx, model *models.VPNGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVPNGroupQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.FQName,
    model.ProvisioningProgressStage,
    model.ProvisioningStartTime,
    model.ProvisioningState,
    model.DisplayName,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.ProvisioningLog,
    model.ProvisioningProgress,
    model.Type,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.Annotations.KeyValuePair)
    return err
}

func ListVPNGroup(tx *sql.Tx) ([]*models.VPNGroup, error) {
    result := models.MakeVPNGroupSlice()
    rows, err := tx.Query(selectVPNGroupQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVPNGroup()
            if err := rows.Scan(&m.FQName,
                &m.ProvisioningProgressStage,
                &m.ProvisioningStartTime,
                &m.ProvisioningState,
                &m.DisplayName,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.ProvisioningLog,
                &m.ProvisioningProgress,
                &m.Type,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
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

func ShowVPNGroup(db *sql.DB, id string, model *models.VPNGroup) error {
    return nil
}

func UpdateVPNGroup(db *sql.DB, id string, model *models.VPNGroup) error {
    return nil
}

func DeleteVPNGroup(db *sql.DB, id string) error {
    return nil
}