package db
// forwarding_class

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertForwardingClassQuery = "insert into `forwarding_class` (`forwarding_class_dscp`,`forwarding_class_mpls_exp`,`forwarding_class_id`,`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`forwarding_class_vlan_priority`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateForwardingClassQuery = "update `forwarding_class` set `forwarding_class_dscp` = ?,`forwarding_class_mpls_exp` = ?,`forwarding_class_id` = ?,`display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`forwarding_class_vlan_priority` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?;"
const deleteForwardingClassQuery = "delete from `forwarding_class`"
const selectForwardingClassQuery = "select `forwarding_class_dscp`,`forwarding_class_mpls_exp`,`forwarding_class_id`,`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`forwarding_class_vlan_priority`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created` from `forwarding_class`"

func CreateForwardingClass(tx *sql.Tx, model *models.ForwardingClass) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertForwardingClassQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ForwardingClassDSCP,
    model.ForwardingClassMPLSExp,
    model.ForwardingClassID,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.ForwardingClassVlanPriority,
    model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created)
    return err
}

func ListForwardingClass(tx *sql.Tx) ([]*models.ForwardingClass, error) {
    result := models.MakeForwardingClassSlice()
    rows, err := tx.Query(selectForwardingClassQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeForwardingClass()
            if err := rows.Scan(&m.ForwardingClassDSCP,
                &m.ForwardingClassMPLSExp,
                &m.ForwardingClassID,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.ForwardingClassVlanPriority,
                &m.FQName,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowForwardingClass(db *sql.DB, id string, model *models.ForwardingClass) error {
    return nil
}

func UpdateForwardingClass(db *sql.DB, id string, model *models.ForwardingClass) error {
    return nil
}

func DeleteForwardingClass(db *sql.DB, id string) error {
    return nil
}