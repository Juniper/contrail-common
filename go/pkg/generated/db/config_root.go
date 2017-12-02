package db

// config_root

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertConfigRootQuery = "insert into `config_root` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateConfigRootQuery = "update `config_root` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?;"
const deleteConfigRootQuery = "delete from `config_root` where uuid = ?"
const listConfigRootQuery = "select `key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name` from `config_root`"
const showConfigRootQuery = "select `key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name` from `config_root` where uuid = ?"

func CreateConfigRoot(tx *sql.Tx, model *models.ConfigRoot) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertConfigRootQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.DisplayName))
	return err
}

func scanConfigRoot(rows *sql.Rows) (*models.ConfigRoot, error) {
	m := models.MakeConfigRoot()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createConfigRootWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListConfigRoot(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ConfigRoot, error) {
	result := models.MakeConfigRootSlice()
	whereQuery, values := createConfigRootWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listConfigRootQuery)
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
		m, _ := scanConfigRoot(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowConfigRoot(tx *sql.Tx, uuid string) (*models.ConfigRoot, error) {
	rows, err := tx.Query(showConfigRootQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanConfigRoot(rows)
	}
	return nil, nil
}

func UpdateConfigRoot(tx *sql.Tx, uuid string, model *models.ConfigRoot) error {
	return nil
}

func DeleteConfigRoot(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteConfigRootQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
