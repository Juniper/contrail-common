package db
// loadbalancer_member

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLoadbalancerMemberQuery = "insert into `loadbalancer_member` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`status_description`,`weight`,`admin_state`,`address`,`protocol_port`,`status`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerMemberQuery = "update `loadbalancer_member` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`status_description` = ?,`weight` = ?,`admin_state` = ?,`address` = ?,`protocol_port` = ?,`status` = ?;"
const deleteLoadbalancerMemberQuery = "delete from `loadbalancer_member`"
const selectLoadbalancerMemberQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`status_description`,`weight`,`admin_state`,`address`,`protocol_port`,`status` from `loadbalancer_member`"

func CreateLoadbalancerMember(tx *sql.Tx, model *models.LoadbalancerMember) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerMemberQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.IDPerms.Creator,
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
    model.LoadbalancerMemberProperties.StatusDescription,
    model.LoadbalancerMemberProperties.Weight,
    model.LoadbalancerMemberProperties.AdminState,
    model.LoadbalancerMemberProperties.Address,
    model.LoadbalancerMemberProperties.ProtocolPort,
    model.LoadbalancerMemberProperties.Status)
    return err
}

func ListLoadbalancerMember(tx *sql.Tx) ([]*models.LoadbalancerMember, error) {
    result := models.MakeLoadbalancerMemberSlice()
    rows, err := tx.Query(selectLoadbalancerMemberQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLoadbalancerMember()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Creator,
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
                &m.LoadbalancerMemberProperties.StatusDescription,
                &m.LoadbalancerMemberProperties.Weight,
                &m.LoadbalancerMemberProperties.AdminState,
                &m.LoadbalancerMemberProperties.Address,
                &m.LoadbalancerMemberProperties.ProtocolPort,
                &m.LoadbalancerMemberProperties.Status); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowLoadbalancerMember(db *sql.DB, id string, model *models.LoadbalancerMember) error {
    return nil
}

func UpdateLoadbalancerMember(db *sql.DB, id string, model *models.LoadbalancerMember) error {
    return nil
}

func DeleteLoadbalancerMember(db *sql.DB, id string) error {
    return nil
}