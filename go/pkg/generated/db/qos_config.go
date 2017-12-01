package db
// qos_config

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertQosConfigQuery = "insert into `qos_config` (`dscp_entries`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`qos_config_type`,`vlan_priority_entries`,`default_forwarding_class_id`,`fq_name`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`mpls_exp_entries`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosConfigQuery = "update `qos_config` set `dscp_entries` = ?,`display_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`qos_config_type` = ?,`vlan_priority_entries` = ?,`default_forwarding_class_id` = ?,`fq_name` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`key_value_pair` = ?,`mpls_exp_entries` = ?;"
const deleteQosConfigQuery = "delete from `qos_config`"
const selectQosConfigQuery = "select `dscp_entries`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`qos_config_type`,`vlan_priority_entries`,`default_forwarding_class_id`,`fq_name`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`mpls_exp_entries` from `qos_config`"

func CreateQosConfig(tx *sql.Tx, model *models.QosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosConfigQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DSCPEntries,
    model.DisplayName,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.QosConfigType,
    model.VlanPriorityEntries,
    model.DefaultForwardingClassID,
    model.FQName,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.Annotations.KeyValuePair,
    model.MPLSExpEntries)
    return err
}

func ListQosConfig(tx *sql.Tx) ([]*models.QosConfig, error) {
    result := models.MakeQosConfigSlice()
    rows, err := tx.Query(selectQosConfigQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeQosConfig()
            if err := rows.Scan(&m.DSCPEntries,
                &m.DisplayName,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.QosConfigType,
                &m.VlanPriorityEntries,
                &m.DefaultForwardingClassID,
                &m.FQName,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.Annotations.KeyValuePair,
                &m.MPLSExpEntries); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowQosConfig(db *sql.DB, id string, model *models.QosConfig) error {
    return nil
}

func UpdateQosConfig(db *sql.DB, id string, model *models.QosConfig) error {
    return nil
}

func DeleteQosConfig(db *sql.DB, id string) error {
    return nil
}