package db

// forwarding_class

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertForwardingClassQuery = "insert into `forwarding_class` (`global_access`,`share`,`owner`,`owner_access`,`fq_name`,`forwarding_class_mpls_exp`,`display_name`,`key_value_pair`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid`,`forwarding_class_dscp`,`forwarding_class_vlan_priority`,`forwarding_class_id`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateForwardingClassQuery = "update `forwarding_class` set `global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`fq_name` = ?,`forwarding_class_mpls_exp` = ?,`display_name` = ?,`key_value_pair` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`uuid` = ?,`forwarding_class_dscp` = ?,`forwarding_class_vlan_priority` = ?,`forwarding_class_id` = ?;"
const deleteForwardingClassQuery = "delete from `forwarding_class` where uuid = ?"
const listForwardingClassQuery = "select `global_access`,`share`,`owner`,`owner_access`,`fq_name`,`forwarding_class_mpls_exp`,`display_name`,`key_value_pair`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid`,`forwarding_class_dscp`,`forwarding_class_vlan_priority`,`forwarding_class_id` from `forwarding_class`"
const showForwardingClassQuery = "select `global_access`,`share`,`owner`,`owner_access`,`fq_name`,`forwarding_class_mpls_exp`,`display_name`,`key_value_pair`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid`,`forwarding_class_dscp`,`forwarding_class_vlan_priority`,`forwarding_class_id` from `forwarding_class` where uuid = ?"

func CreateForwardingClass(tx *sql.Tx, model *models.ForwardingClass) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertForwardingClassQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		util.MustJSON(model.FQName),
		int(model.ForwardingClassMPLSExp),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.UUID),
		int(model.ForwardingClassDSCP),
		int(model.ForwardingClassVlanPriority),
		int(model.ForwardingClassID))
	return err
}

func scanForwardingClass(rows *sql.Rows) (*models.ForwardingClass, error) {
	m := models.MakeForwardingClass()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonFQName,
		&m.ForwardingClassMPLSExp,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
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
		&m.IDPerms.LastModified,
		&m.UUID,
		&m.ForwardingClassDSCP,
		&m.ForwardingClassVlanPriority,
		&m.ForwardingClassID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListForwardingClass(tx *sql.Tx) ([]*models.ForwardingClass, error) {
	result := models.MakeForwardingClassSlice()
	rows, err := tx.Query(listForwardingClassQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanForwardingClass(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowForwardingClass(tx *sql.Tx, uuid string) (*models.ForwardingClass, error) {
	rows, err := tx.Query(showForwardingClassQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanForwardingClass(rows)
	}
	return nil, nil
}

func UpdateForwardingClass(tx *sql.Tx, uuid string, model *models.ForwardingClass) error {
	return nil
}

func DeleteForwardingClass(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteForwardingClassQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
