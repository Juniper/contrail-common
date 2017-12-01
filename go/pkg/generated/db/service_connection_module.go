package db

// service_connection_module

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceConnectionModuleQuery = "insert into `service_connection_module` (`uuid`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`service_type`,`e2_service`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceConnectionModuleQuery = "update `service_connection_module` set `uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`service_type` = ?,`e2_service` = ?;"
const deleteServiceConnectionModuleQuery = "delete from `service_connection_module` where uuid = ?"
const listServiceConnectionModuleQuery = "select `uuid`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`service_type`,`e2_service` from `service_connection_module`"
const showServiceConnectionModuleQuery = "select `uuid`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`service_type`,`e2_service` from `service_connection_module` where uuid = ?"

func CreateServiceConnectionModule(tx *sql.Tx, model *models.ServiceConnectionModule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceConnectionModuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		util.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.ServiceType),
		string(model.E2Service))
	return err
}

func scanServiceConnectionModule(rows *sql.Rows) (*models.ServiceConnectionModule, error) {
	m := models.MakeServiceConnectionModule()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.ServiceType,
		&m.E2Service); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListServiceConnectionModule(tx *sql.Tx) ([]*models.ServiceConnectionModule, error) {
	result := models.MakeServiceConnectionModuleSlice()
	rows, err := tx.Query(listServiceConnectionModuleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceConnectionModule(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceConnectionModule(tx *sql.Tx, uuid string) (*models.ServiceConnectionModule, error) {
	rows, err := tx.Query(showServiceConnectionModuleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceConnectionModule(rows)
	}
	return nil, nil
}

func UpdateServiceConnectionModule(tx *sql.Tx, uuid string, model *models.ServiceConnectionModule) error {
	return nil
}

func DeleteServiceConnectionModule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceConnectionModuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
