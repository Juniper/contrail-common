package db

// widget

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertWidgetQuery = "insert into `widget` (`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`container_config`,`content_config`,`layout_config`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateWidgetQuery = "update `widget` set `display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`fq_name` = ?,`container_config` = ?,`content_config` = ?,`layout_config` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`uuid` = ?;"
const deleteWidgetQuery = "delete from `widget` where uuid = ?"
const listWidgetQuery = "select `display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`container_config`,`content_config`,`layout_config`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid` from `widget`"
const showWidgetQuery = "select `display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`fq_name`,`container_config`,`content_config`,`layout_config`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`uuid` from `widget` where uuid = ?"

func CreateWidget(tx *sql.Tx, model *models.Widget) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertWidgetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.FQName),
		string(model.ContainerConfig),
		string(model.ContentConfig),
		string(model.LayoutConfig),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.UUID))
	return err
}

func scanWidget(rows *sql.Rows) (*models.Widget, error) {
	m := models.MakeWidget()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonFQName,
		&m.ContainerConfig,
		&m.ContentConfig,
		&m.LayoutConfig,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListWidget(tx *sql.Tx) ([]*models.Widget, error) {
	result := models.MakeWidgetSlice()
	rows, err := tx.Query(listWidgetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanWidget(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowWidget(tx *sql.Tx, uuid string) (*models.Widget, error) {
	rows, err := tx.Query(showWidgetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanWidget(rows)
	}
	return nil, nil
}

func UpdateWidget(tx *sql.Tx, uuid string, model *models.Widget) error {
	return nil
}

func DeleteWidget(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteWidgetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
