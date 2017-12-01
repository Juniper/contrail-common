package db

// dashboard

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertDashboardQuery = "insert into `dashboard` (`container_config`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDashboardQuery = "update `dashboard` set `container_config` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteDashboardQuery = "delete from `dashboard` where uuid = ?"
const listDashboardQuery = "select `container_config`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair` from `dashboard`"
const showDashboardQuery = "select `container_config`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair` from `dashboard` where uuid = ?"

func CreateDashboard(tx *sql.Tx, model *models.Dashboard) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDashboardQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ContainerConfig),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanDashboard(rows *sql.Rows) (*models.Dashboard, error) {
	m := models.MakeDashboard()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ContainerConfig,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListDashboard(tx *sql.Tx) ([]*models.Dashboard, error) {
	result := models.MakeDashboardSlice()
	rows, err := tx.Query(listDashboardQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanDashboard(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDashboard(tx *sql.Tx, uuid string) (*models.Dashboard, error) {
	rows, err := tx.Query(showDashboardQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDashboard(rows)
	}
	return nil, nil
}

func UpdateDashboard(tx *sql.Tx, uuid string, model *models.Dashboard) error {
	return nil
}

func DeleteDashboard(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDashboardQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
