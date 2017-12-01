package db

// bgp_router

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertBGPRouterQuery = "insert into `bgp_router` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPRouterQuery = "update `bgp_router` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteBGPRouterQuery = "delete from `bgp_router` where uuid = ?"
const listBGPRouterQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair` from `bgp_router`"
const showBGPRouterQuery = "select `owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair` from `bgp_router` where uuid = ?"

func CreateBGPRouter(tx *sql.Tx, model *models.BGPRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPRouterQuery)
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
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanBGPRouter(rows *sql.Rows) (*models.BGPRouter, error) {
	m := models.MakeBGPRouter()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListBGPRouter(tx *sql.Tx) ([]*models.BGPRouter, error) {
	result := models.MakeBGPRouterSlice()
	rows, err := tx.Query(listBGPRouterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanBGPRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBGPRouter(tx *sql.Tx, uuid string) (*models.BGPRouter, error) {
	rows, err := tx.Query(showBGPRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBGPRouter(rows)
	}
	return nil, nil
}

func UpdateBGPRouter(tx *sql.Tx, uuid string, model *models.BGPRouter) error {
	return nil
}

func DeleteBGPRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
