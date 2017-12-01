package db

// bgp_as_a_service

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`key_value_pair`,`bgpaas_session_attributes`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`autonomous_system`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`bgpaas_shared`,`bgpaas_suppress_route_advertisement`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPAsAServiceQuery = "update `bgp_as_a_service` set `key_value_pair` = ?,`bgpaas_session_attributes` = ?,`bgpaas_ipv4_mapped_ipv6_nexthop` = ?,`bgpaas_ip_address` = ?,`autonomous_system` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`bgpaas_shared` = ?,`bgpaas_suppress_route_advertisement` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`display_name` = ?;"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service` where uuid = ?"
const listBGPAsAServiceQuery = "select `key_value_pair`,`bgpaas_session_attributes`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`autonomous_system`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`bgpaas_shared`,`bgpaas_suppress_route_advertisement`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`display_name` from `bgp_as_a_service`"
const showBGPAsAServiceQuery = "select `key_value_pair`,`bgpaas_session_attributes`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`autonomous_system`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`bgpaas_shared`,`bgpaas_suppress_route_advertisement`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`display_name` from `bgp_as_a_service` where uuid = ?"

func CreateBGPAsAService(tx *sql.Tx, model *models.BGPAsAService) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.Annotations.KeyValuePair),
		string(model.BgpaasSessionAttributes),
		bool(model.BgpaasIpv4MappedIpv6Nexthop),
		string(model.BgpaasIPAddress),
		int(model.AutonomousSystem),
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
		bool(model.BgpaasShared),
		bool(model.BgpaasSuppressRouteAdvertisement),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.DisplayName))
	return err
}

func scanBGPAsAService(rows *sql.Rows) (*models.BGPAsAService, error) {
	m := models.MakeBGPAsAService()

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.BgpaasSessionAttributes,
		&m.BgpaasIpv4MappedIpv6Nexthop,
		&m.BgpaasIPAddress,
		&m.AutonomousSystem,
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
		&m.BgpaasShared,
		&m.BgpaasSuppressRouteAdvertisement,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListBGPAsAService(tx *sql.Tx) ([]*models.BGPAsAService, error) {
	result := models.MakeBGPAsAServiceSlice()
	rows, err := tx.Query(listBGPAsAServiceQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanBGPAsAService(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBGPAsAService(tx *sql.Tx, uuid string) (*models.BGPAsAService, error) {
	rows, err := tx.Query(showBGPAsAServiceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBGPAsAService(rows)
	}
	return nil, nil
}

func UpdateBGPAsAService(tx *sql.Tx, uuid string, model *models.BGPAsAService) error {
	return nil
}

func DeleteBGPAsAService(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPAsAServiceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
