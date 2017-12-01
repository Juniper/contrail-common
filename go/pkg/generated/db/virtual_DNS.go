package db

// virtual_DNS

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVirtualDNSQuery = "insert into `virtual_DNS` (`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`external_visible`,`next_virtual_DNS`,`dynamic_records_from_client`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSQuery = "update `virtual_DNS` set `fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`reverse_resolution` = ?,`default_ttl_seconds` = ?,`record_order` = ?,`floating_ip_record` = ?,`domain_name` = ?,`external_visible` = ?,`next_virtual_DNS` = ?,`dynamic_records_from_client` = ?;"
const deleteVirtualDNSQuery = "delete from `virtual_DNS` where uuid = ?"
const listVirtualDNSQuery = "select `fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`external_visible`,`next_virtual_DNS`,`dynamic_records_from_client` from `virtual_DNS`"
const showVirtualDNSQuery = "select `fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`external_visible`,`next_virtual_DNS`,`dynamic_records_from_client` from `virtual_DNS` where uuid = ?"

func CreateVirtualDNS(tx *sql.Tx, model *models.VirtualDNS) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(util.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		bool(model.VirtualDNSData.ReverseResolution),
		int(model.VirtualDNSData.DefaultTTLSeconds),
		string(model.VirtualDNSData.RecordOrder),
		string(model.VirtualDNSData.FloatingIPRecord),
		string(model.VirtualDNSData.DomainName),
		bool(model.VirtualDNSData.ExternalVisible),
		string(model.VirtualDNSData.NextVirtualDNS),
		bool(model.VirtualDNSData.DynamicRecordsFromClient))
	return err
}

func scanVirtualDNS(rows *sql.Rows) (*models.VirtualDNS, error) {
	m := models.MakeVirtualDNS()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&m.VirtualDNSData.ReverseResolution,
		&m.VirtualDNSData.DefaultTTLSeconds,
		&m.VirtualDNSData.RecordOrder,
		&m.VirtualDNSData.FloatingIPRecord,
		&m.VirtualDNSData.DomainName,
		&m.VirtualDNSData.ExternalVisible,
		&m.VirtualDNSData.NextVirtualDNS,
		&m.VirtualDNSData.DynamicRecordsFromClient); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func ListVirtualDNS(tx *sql.Tx) ([]*models.VirtualDNS, error) {
	result := models.MakeVirtualDNSSlice()
	rows, err := tx.Query(listVirtualDNSQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualDNS(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualDNS(tx *sql.Tx, uuid string) (*models.VirtualDNS, error) {
	rows, err := tx.Query(showVirtualDNSQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualDNS(rows)
	}
	return nil, nil
}

func UpdateVirtualDNS(tx *sql.Tx, uuid string, model *models.VirtualDNS) error {
	return nil
}

func DeleteVirtualDNS(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualDNSQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
