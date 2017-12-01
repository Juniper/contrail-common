package db
// bridge_domain

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertBridgeDomainQuery = "insert into `bridge_domain` (`isid`,`mac_learning_enabled`,`mac_move_limit_action`,`mac_move_time_window`,`mac_move_limit`,`mac_limit`,`mac_limit_action`,`uuid`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`mac_aging_time`,`fq_name`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBridgeDomainQuery = "update `bridge_domain` set `isid` = ?,`mac_learning_enabled` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`uuid` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`mac_aging_time` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteBridgeDomainQuery = "delete from `bridge_domain`"
const selectBridgeDomainQuery = "select `isid`,`mac_learning_enabled`,`mac_move_limit_action`,`mac_move_time_window`,`mac_move_limit`,`mac_limit`,`mac_limit_action`,`uuid`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`mac_aging_time`,`fq_name`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share` from `bridge_domain`"

func CreateBridgeDomain(tx *sql.Tx, model *models.BridgeDomain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBridgeDomainQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.Isid,
    model.MacLearningEnabled,
    model.MacMoveControl.MacMoveLimitAction,
    model.MacMoveControl.MacMoveTimeWindow,
    model.MacMoveControl.MacMoveLimit,
    model.MacLimitControl.MacLimit,
    model.MacLimitControl.MacLimitAction,
    model.UUID,
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
    model.MacAgingTime,
    model.FQName,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share)
    return err
}

func ListBridgeDomain(tx *sql.Tx) ([]*models.BridgeDomain, error) {
    result := models.MakeBridgeDomainSlice()
    rows, err := tx.Query(selectBridgeDomainQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeBridgeDomain()
            if err := rows.Scan(&m.Isid,
                &m.MacLearningEnabled,
                &m.MacMoveControl.MacMoveLimitAction,
                &m.MacMoveControl.MacMoveTimeWindow,
                &m.MacMoveControl.MacMoveLimit,
                &m.MacLimitControl.MacLimit,
                &m.MacLimitControl.MacLimitAction,
                &m.UUID,
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
                &m.MacAgingTime,
                &m.FQName,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowBridgeDomain(db *sql.DB, id string, model *models.BridgeDomain) error {
    return nil
}

func UpdateBridgeDomain(db *sql.DB, id string, model *models.BridgeDomain) error {
    return nil
}

func DeleteBridgeDomain(db *sql.DB, id string) error {
    return nil
}