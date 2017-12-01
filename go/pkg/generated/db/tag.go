package db

// tag

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertTagQuery = "insert into `tag` (`tag_value`,`uuid`,`fq_name`,`owner`,`owner_access`,`global_access`,`share`,`tag_type_name`,`tag_id`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateTagQuery = "update `tag` set `tag_value` = ?,`uuid` = ?,`fq_name` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`tag_type_name` = ?,`tag_id` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteTagQuery = "delete from `tag` where uuid = ?"
const listTagQuery = "select `tag_value`,`uuid`,`fq_name`,`owner`,`owner_access`,`global_access`,`share`,`tag_type_name`,`tag_id`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair` from `tag`"
const showTagQuery = "select `tag_value`,`uuid`,`fq_name`,`owner`,`owner_access`,`global_access`,`share`,`tag_type_name`,`tag_id`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair` from `tag` where uuid = ?"

func CreateTag(tx *sql.Tx, model *models.Tag) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertTagQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.TagValue),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.TagTypeName),
		string(model.TagID),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanTag(rows *sql.Rows) (*models.Tag, error) {
	m := models.MakeTag()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.TagValue,
		&m.UUID,
		&jsonFQName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.TagTypeName,
		&m.TagID,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListTag(tx *sql.Tx) ([]*models.Tag, error) {
	result := models.MakeTagSlice()
	rows, err := tx.Query(listTagQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanTag(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowTag(tx *sql.Tx, uuid string) (*models.Tag, error) {
	rows, err := tx.Query(showTagQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanTag(rows)
	}
	return nil, nil
}

func UpdateTag(tx *sql.Tx, uuid string, model *models.Tag) error {
	return nil
}

func DeleteTag(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteTagQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
