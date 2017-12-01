package db

// e2_service_provider

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertE2ServiceProviderQuery = "insert into `e2_service_provider` (`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`e2_service_provider_promiscuous`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateE2ServiceProviderQuery = "update `e2_service_provider` set `fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`e2_service_provider_promiscuous` = ?;"
const deleteE2ServiceProviderQuery = "delete from `e2_service_provider` where uuid = ?"
const listE2ServiceProviderQuery = "select `fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`e2_service_provider_promiscuous` from `e2_service_provider`"
const showE2ServiceProviderQuery = "select `fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`e2_service_provider_promiscuous` from `e2_service_provider` where uuid = ?"

func CreateE2ServiceProvider(tx *sql.Tx, model *models.E2ServiceProvider) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertE2ServiceProviderQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		bool(model.E2ServiceProviderPromiscuous))
	return err
}

func scanE2ServiceProvider(rows *sql.Rows) (*models.E2ServiceProvider, error) {
	m := models.MakeE2ServiceProvider()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
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
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&m.E2ServiceProviderPromiscuous); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListE2ServiceProvider(tx *sql.Tx) ([]*models.E2ServiceProvider, error) {
	result := models.MakeE2ServiceProviderSlice()
	rows, err := tx.Query(listE2ServiceProviderQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanE2ServiceProvider(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowE2ServiceProvider(tx *sql.Tx, uuid string) (*models.E2ServiceProvider, error) {
	rows, err := tx.Query(showE2ServiceProviderQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanE2ServiceProvider(rows)
	}
	return nil, nil
}

func UpdateE2ServiceProvider(tx *sql.Tx, uuid string, model *models.E2ServiceProvider) error {
	return nil
}

func DeleteE2ServiceProvider(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteE2ServiceProviderQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
