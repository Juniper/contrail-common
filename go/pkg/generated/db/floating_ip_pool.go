package db

// floating_ip_pool

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertFloatingIPPoolQuery = "insert into `floating_ip_pool` (`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`subnet_uuid`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPPoolQuery = "update `floating_ip_pool` set `key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`subnet_uuid` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteFloatingIPPoolQuery = "delete from `floating_ip_pool` where uuid = ?"
const listFloatingIPPoolQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`subnet_uuid`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name` from `floating_ip_pool`"
const showFloatingIPPoolQuery = "select `key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`subnet_uuid`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`display_name` from `floating_ip_pool` where uuid = ?"

func CreateFloatingIPPool(tx *sql.Tx, model *models.FloatingIPPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		util.MustJSON(model.FloatingIPPoolSubnets.SubnetUUID),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName))
	return err
}

func scanFloatingIPPool(rows *sql.Rows) (*models.FloatingIPPool, error) {
	m := models.MakeFloatingIPPool()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFloatingIPPoolSubnetsSubnetUUID string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonFloatingIPPoolSubnetsSubnetUUID,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFloatingIPPoolSubnetsSubnetUUID), &m.FloatingIPPoolSubnets.SubnetUUID)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListFloatingIPPool(tx *sql.Tx) ([]*models.FloatingIPPool, error) {
	result := models.MakeFloatingIPPoolSlice()
	rows, err := tx.Query(listFloatingIPPoolQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanFloatingIPPool(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFloatingIPPool(tx *sql.Tx, uuid string) (*models.FloatingIPPool, error) {
	rows, err := tx.Query(showFloatingIPPoolQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFloatingIPPool(rows)
	}
	return nil, nil
}

func UpdateFloatingIPPool(tx *sql.Tx, uuid string, model *models.FloatingIPPool) error {
	return nil
}

func DeleteFloatingIPPool(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFloatingIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
