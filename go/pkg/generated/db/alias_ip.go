package db

// alias_ip

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertAliasIPQuery = "insert into `alias_ip` (`alias_ip_address_family`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`alias_ip_address`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAliasIPQuery = "update `alias_ip` set `alias_ip_address_family` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`alias_ip_address` = ?;"
const deleteAliasIPQuery = "delete from `alias_ip` where uuid = ?"
const listAliasIPQuery = "select `alias_ip_address_family`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`alias_ip_address` from `alias_ip`"
const showAliasIPQuery = "select `alias_ip_address_family`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`alias_ip_address` from `alias_ip` where uuid = ?"

func CreateAliasIP(tx *sql.Tx, model *models.AliasIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAliasIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.AliasIPAddressFamily),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.AliasIPAddress))
	return err
}

func scanAliasIP(rows *sql.Rows) (*models.AliasIP, error) {
	m := models.MakeAliasIP()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.AliasIPAddressFamily,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.AliasIPAddress); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListAliasIP(tx *sql.Tx) ([]*models.AliasIP, error) {
	result := models.MakeAliasIPSlice()
	rows, err := tx.Query(listAliasIPQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanAliasIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAliasIP(tx *sql.Tx, uuid string) (*models.AliasIP, error) {
	rows, err := tx.Query(showAliasIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAliasIP(rows)
	}
	return nil, nil
}

func UpdateAliasIP(tx *sql.Tx, uuid string, model *models.AliasIP) error {
	return nil
}

func DeleteAliasIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAliasIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
