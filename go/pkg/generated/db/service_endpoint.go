package db

// service_endpoint

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertServiceEndpointQuery = "insert into `service_endpoint` (`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceEndpointQuery = "update `service_endpoint` set `key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?;"
const deleteServiceEndpointQuery = "delete from `service_endpoint` where uuid = ?"
const listServiceEndpointQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name` from `service_endpoint`"
const showServiceEndpointQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`display_name` from `service_endpoint` where uuid = ?"

func CreateServiceEndpoint(tx *sql.Tx, model *models.ServiceEndpoint) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceEndpointQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
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

func scanServiceEndpoint(rows *sql.Rows) (*models.ServiceEndpoint, error) {
	m := models.MakeServiceEndpoint()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
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

func ListServiceEndpoint(tx *sql.Tx) ([]*models.ServiceEndpoint, error) {
	result := models.MakeServiceEndpointSlice()
	rows, err := tx.Query(listServiceEndpointQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanServiceEndpoint(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceEndpoint(tx *sql.Tx, uuid string) (*models.ServiceEndpoint, error) {
	rows, err := tx.Query(showServiceEndpointQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceEndpoint(rows)
	}
	return nil, nil
}

func UpdateServiceEndpoint(tx *sql.Tx, uuid string, model *models.ServiceEndpoint) error {
	return nil
}

func DeleteServiceEndpoint(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceEndpointQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
