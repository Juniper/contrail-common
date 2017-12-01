package db

// domain

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertDomainQuery = "insert into `domain` (`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`project_limit`,`virtual_network_limit`,`security_group_limit`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDomainQuery = "update `domain` set `share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`project_limit` = ?,`virtual_network_limit` = ?,`security_group_limit` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteDomainQuery = "delete from `domain` where uuid = ?"
const listDomainQuery = "select `share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`project_limit`,`virtual_network_limit`,`security_group_limit`,`display_name`,`key_value_pair` from `domain`"
const showDomainQuery = "select `share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`project_limit`,`virtual_network_limit`,`security_group_limit`,`display_name`,`key_value_pair` from `domain` where uuid = ?"

func CreateDomain(tx *sql.Tx, model *models.Domain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
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
		int(model.DomainLimits.ProjectLimit),
		int(model.DomainLimits.VirtualNetworkLimit),
		int(model.DomainLimits.SecurityGroupLimit),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanDomain(rows *sql.Rows) (*models.Domain, error) {
	m := models.MakeDomain()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
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
		&m.DomainLimits.ProjectLimit,
		&m.DomainLimits.VirtualNetworkLimit,
		&m.DomainLimits.SecurityGroupLimit,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListDomain(tx *sql.Tx) ([]*models.Domain, error) {
	result := models.MakeDomainSlice()
	rows, err := tx.Query(listDomainQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanDomain(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDomain(tx *sql.Tx, uuid string) (*models.Domain, error) {
	rows, err := tx.Query(showDomainQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDomain(rows)
	}
	return nil, nil
}

func UpdateDomain(tx *sql.Tx, uuid string, model *models.Domain) error {
	return nil
}

func DeleteDomain(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
