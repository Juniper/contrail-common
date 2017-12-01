package db

// service_object

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertServiceObjectQuery = "insert into `service_object` (`uuid`,`fq_name`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceObjectQuery = "update `service_object` set `uuid` = ?,`fq_name` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteServiceObjectQuery = "delete from `service_object` where uuid = ?"
const listServiceObjectQuery = "select `uuid`,`fq_name`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `service_object`"
const showServiceObjectQuery = "select `uuid`,`fq_name`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner` from `service_object` where uuid = ?"

func CreateServiceObject(tx *sql.Tx, model *models.ServiceObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	return err
}

func scanServiceObject(rows *sql.Rows) (*models.ServiceObject, error) {
	m := models.MakeServiceObject()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func createServiceObjectWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceObject(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceObject, error) {
	result := models.MakeServiceObjectSlice()
	whereQuery, values := createServiceObjectWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceObjectQuery)
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
		m, _ := scanServiceObject(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceObject(tx *sql.Tx, uuid string) (*models.ServiceObject, error) {
	rows, err := tx.Query(showServiceObjectQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceObject(rows)
	}
	return nil, nil
}

func UpdateServiceObject(tx *sql.Tx, uuid string, model *models.ServiceObject) error {
	return nil
}

func DeleteServiceObject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
