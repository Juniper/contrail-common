package db

// bgpvpn

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertBGPVPNQuery = "insert into `bgpvpn` (`fq_name`,`route_target`,`export_route_target_list_route_target`,`bgpvpn_type`,`key_value_pair`,`display_name`,`import_route_target_list_route_target`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPVPNQuery = "update `bgpvpn` set `fq_name` = ?,`route_target` = ?,`export_route_target_list_route_target` = ?,`bgpvpn_type` = ?,`key_value_pair` = ?,`display_name` = ?,`import_route_target_list_route_target` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteBGPVPNQuery = "delete from `bgpvpn` where uuid = ?"
const listBGPVPNQuery = "select `fq_name`,`route_target`,`export_route_target_list_route_target`,`bgpvpn_type`,`key_value_pair`,`display_name`,`import_route_target_list_route_target`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible` from `bgpvpn`"
const showBGPVPNQuery = "select `fq_name`,`route_target`,`export_route_target_list_route_target`,`bgpvpn_type`,`key_value_pair`,`display_name`,`import_route_target_list_route_target`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible` from `bgpvpn` where uuid = ?"

func CreateBGPVPN(tx *sql.Tx, model *models.BGPVPN) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPVPNQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
		util.MustJSON(model.RouteTargetList.RouteTarget),
		util.MustJSON(model.ExportRouteTargetList.RouteTarget),
		string(model.BGPVPNType),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.DisplayName),
		util.MustJSON(model.ImportRouteTargetList.RouteTarget),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
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
		bool(model.IDPerms.UserVisible))
	return err
}

func scanBGPVPN(rows *sql.Rows) (*models.BGPVPN, error) {
	m := models.MakeBGPVPN()

	var jsonFQName string

	var jsonRouteTargetListRouteTarget string

	var jsonExportRouteTargetListRouteTarget string

	var jsonAnnotationsKeyValuePair string

	var jsonImportRouteTargetListRouteTarget string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&jsonRouteTargetListRouteTarget,
		&jsonExportRouteTargetListRouteTarget,
		&m.BGPVPNType,
		&jsonAnnotationsKeyValuePair,
		&m.DisplayName,
		&jsonImportRouteTargetListRouteTarget,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
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
		&m.IDPerms.UserVisible); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonRouteTargetListRouteTarget), &m.RouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonExportRouteTargetListRouteTarget), &m.ExportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonImportRouteTargetListRouteTarget), &m.ImportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListBGPVPN(tx *sql.Tx) ([]*models.BGPVPN, error) {
	result := models.MakeBGPVPNSlice()
	rows, err := tx.Query(listBGPVPNQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanBGPVPN(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBGPVPN(tx *sql.Tx, uuid string) (*models.BGPVPN, error) {
	rows, err := tx.Query(showBGPVPNQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBGPVPN(rows)
	}
	return nil, nil
}

func UpdateBGPVPN(tx *sql.Tx, uuid string, model *models.BGPVPN) error {
	return nil
}

func DeleteBGPVPN(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPVPNQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
