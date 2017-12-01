package db

// forwarding_class

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertForwardingClassQuery = "insert into `forwarding_class` (`forwarding_class_vlan_priority`,`forwarding_class_mpls_exp`,`forwarding_class_id`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`forwarding_class_dscp`,`key_value_pair`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateForwardingClassQuery = "update `forwarding_class` set `forwarding_class_vlan_priority` = ?,`forwarding_class_mpls_exp` = ?,`forwarding_class_id` = ?,`display_name` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`fq_name` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`forwarding_class_dscp` = ?,`key_value_pair` = ?,`uuid` = ?;"
const deleteForwardingClassQuery = "delete from `forwarding_class` where uuid = ?"
const listForwardingClassQuery = "select `forwarding_class_vlan_priority`,`forwarding_class_mpls_exp`,`forwarding_class_id`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`forwarding_class_dscp`,`key_value_pair`,`uuid` from `forwarding_class`"
const showForwardingClassQuery = "select `forwarding_class_vlan_priority`,`forwarding_class_mpls_exp`,`forwarding_class_id`,`display_name`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`forwarding_class_dscp`,`key_value_pair`,`uuid` from `forwarding_class` where uuid = ?"

func CreateForwardingClass(tx *sql.Tx, model *models.ForwardingClass) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertForwardingClassQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.ForwardingClassVlanPriority),
		int(model.ForwardingClassMPLSExp),
		int(model.ForwardingClassID),
		string(model.DisplayName),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		int(model.ForwardingClassDSCP),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID))
	return err
}

func scanForwardingClass(rows *sql.Rows) (*models.ForwardingClass, error) {
	m := models.MakeForwardingClass()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ForwardingClassVlanPriority,
		&m.ForwardingClassMPLSExp,
		&m.ForwardingClassID,
		&m.DisplayName,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.ForwardingClassDSCP,
		&jsonAnnotationsKeyValuePair,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func createForwardingClassWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListForwardingClass(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ForwardingClass, error) {
	result := models.MakeForwardingClassSlice()
	whereQuery, values := createForwardingClassWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listForwardingClassQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
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
