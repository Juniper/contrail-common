package db

// discovery_service_assignment

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertDiscoveryServiceAssignmentQuery = "insert into `discovery_service_assignment` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDiscoveryServiceAssignmentQuery = "update `discovery_service_assignment` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteDiscoveryServiceAssignmentQuery = "delete from `discovery_service_assignment` where uuid = ?"
const listDiscoveryServiceAssignmentQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair` from `discovery_service_assignment`"
const showDiscoveryServiceAssignmentQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair` from `discovery_service_assignment` where uuid = ?"

func CreateDiscoveryServiceAssignment(tx *sql.Tx, model *models.DiscoveryServiceAssignment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDiscoveryServiceAssignmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanDiscoveryServiceAssignment(rows *sql.Rows) (*models.DiscoveryServiceAssignment, error) {
	m := models.MakeDiscoveryServiceAssignment()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListDiscoveryServiceAssignment(tx *sql.Tx) ([]*models.DiscoveryServiceAssignment, error) {
	result := models.MakeDiscoveryServiceAssignmentSlice()
	rows, err := tx.Query(listDiscoveryServiceAssignmentQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanDiscoveryServiceAssignment(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDiscoveryServiceAssignment(tx *sql.Tx, uuid string) (*models.DiscoveryServiceAssignment, error) {
	rows, err := tx.Query(showDiscoveryServiceAssignmentQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDiscoveryServiceAssignment(rows)
	}
	return nil, nil
}

func UpdateDiscoveryServiceAssignment(tx *sql.Tx, uuid string, model *models.DiscoveryServiceAssignment) error {
	return nil
}

func DeleteDiscoveryServiceAssignment(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDiscoveryServiceAssignmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
