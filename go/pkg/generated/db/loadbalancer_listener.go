package db
// loadbalancer_listener

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertLoadbalancerListenerQuery = "insert into `loadbalancer_listener` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`admin_state`,`sni_containers`,`protocol_port`,`default_tls_container`,`protocol`,`connection_limit`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerListenerQuery = "update `loadbalancer_listener` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`admin_state` = ?,`sni_containers` = ?,`protocol_port` = ?,`default_tls_container` = ?,`protocol` = ?,`connection_limit` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteLoadbalancerListenerQuery = "delete from `loadbalancer_listener`"
const selectLoadbalancerListenerQuery = "select `description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`admin_state`,`sni_containers`,`protocol_port`,`default_tls_container`,`protocol`,`connection_limit`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name` from `loadbalancer_listener`"

func CreateLoadbalancerListener(tx *sql.Tx, model *models.LoadbalancerListener) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerListenerQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.LoadbalancerListenerProperties.AdminState,
    model.LoadbalancerListenerProperties.SniContainers,
    model.LoadbalancerListenerProperties.ProtocolPort,
    model.LoadbalancerListenerProperties.DefaultTLSContainer,
    model.LoadbalancerListenerProperties.Protocol,
    model.LoadbalancerListenerProperties.ConnectionLimit,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.UUID,
    model.FQName)
    return err
}

func ListLoadbalancerListener(tx *sql.Tx) ([]*models.LoadbalancerListener, error) {
    result := models.MakeLoadbalancerListenerSlice()
    rows, err := tx.Query(selectLoadbalancerListenerQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeLoadbalancerListener()
            if err := rows.Scan(&m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Enable,
                &m.LoadbalancerListenerProperties.AdminState,
                &m.LoadbalancerListenerProperties.SniContainers,
                &m.LoadbalancerListenerProperties.ProtocolPort,
                &m.LoadbalancerListenerProperties.DefaultTLSContainer,
                &m.LoadbalancerListenerProperties.Protocol,
                &m.LoadbalancerListenerProperties.ConnectionLimit,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.UUID,
                &m.FQName); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowLoadbalancerListener(db *sql.DB, id string, model *models.LoadbalancerListener) error {
    return nil
}

func UpdateLoadbalancerListener(db *sql.DB, id string, model *models.LoadbalancerListener) error {
    return nil
}

func DeleteLoadbalancerListener(db *sql.DB, id string) error {
    return nil
}