package db

// bgp_as_a_service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
	"strings"
)

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`display_name`,`key_value_pair`,`bgpaas_session_attributes`,`bgpaas_suppress_route_advertisement`,`autonomous_system`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`bgpaas_shared`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPAsAServiceQuery = "update `bgp_as_a_service` set `display_name` = ?,`key_value_pair` = ?,`bgpaas_session_attributes` = ?,`bgpaas_suppress_route_advertisement` = ?,`autonomous_system` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`bgpaas_shared` = ?,`bgpaas_ipv4_mapped_ipv6_nexthop` = ?,`bgpaas_ip_address` = ?,`fq_name` = ?;"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service` where uuid = ?"
const listBGPAsAServiceQuery = "select `display_name`,`key_value_pair`,`bgpaas_session_attributes`,`bgpaas_suppress_route_advertisement`,`autonomous_system`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`bgpaas_shared`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`fq_name` from `bgp_as_a_service`"
const showBGPAsAServiceQuery = "select `display_name`,`key_value_pair`,`bgpaas_session_attributes`,`bgpaas_suppress_route_advertisement`,`autonomous_system`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`bgpaas_shared`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`fq_name` from `bgp_as_a_service` where uuid = ?"

func CreateBGPAsAService(tx *sql.Tx, model *models.BGPAsAService) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		string(model.BgpaasSessionAttributes),
		bool(model.BgpaasSuppressRouteAdvertisement),
		int(model.AutonomousSystem),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
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
		bool(model.BgpaasShared),
		bool(model.BgpaasIpv4MappedIpv6Nexthop),
		string(model.BgpaasIPAddress),
		util.MustJSON(model.FQName))
	return err
}

func scanBGPAsAService(rows *sql.Rows) (*models.BGPAsAService, error) {
	m := models.MakeBGPAsAService()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.BgpaasSessionAttributes,
		&m.BgpaasSuppressRouteAdvertisement,
		&m.AutonomousSystem,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
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
		&m.BgpaasShared,
		&m.BgpaasIpv4MappedIpv6Nexthop,
		&m.BgpaasIPAddress,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func createBGPAsAServiceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["bgpaas_session_attributes"]; ok {
		results = append(results, "bgpaas_session_attributes = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["bgpaas_ip_address"]; ok {
		results = append(results, "bgpaas_ip_address = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListBGPAsAService(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.BGPAsAService, error) {
	result := models.MakeBGPAsAServiceSlice()
	whereQuery, values := createBGPAsAServiceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listBGPAsAServiceQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
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
