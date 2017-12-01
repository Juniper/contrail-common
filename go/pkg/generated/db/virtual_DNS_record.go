package db

// virtual_DNS_record

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertVirtualDNSRecordQuery = "insert into `virtual_DNS_record` (`display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`record_name`,`record_class`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSRecordQuery = "update `virtual_DNS_record` set `display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`record_data` = ?,`record_type` = ?,`record_ttl_seconds` = ?,`record_mx_preference` = ?,`record_name` = ?,`record_class` = ?;"
const deleteVirtualDNSRecordQuery = "delete from `virtual_DNS_record` where uuid = ?"
const listVirtualDNSRecordQuery = "select `display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`record_name`,`record_class` from `virtual_DNS_record`"
const showVirtualDNSRecordQuery = "select `display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`record_name`,`record_class` from `virtual_DNS_record` where uuid = ?"

func CreateVirtualDNSRecord(tx *sql.Tx, model *models.VirtualDNSRecord) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.VirtualDNSRecordData.RecordData),
		string(model.VirtualDNSRecordData.RecordType),
		int(model.VirtualDNSRecordData.RecordTTLSeconds),
		int(model.VirtualDNSRecordData.RecordMXPreference),
		string(model.VirtualDNSRecordData.RecordName),
		string(model.VirtualDNSRecordData.RecordClass))
	return err
}

func scanVirtualDNSRecord(rows *sql.Rows) (*models.VirtualDNSRecord, error) {
	m := models.MakeVirtualDNSRecord()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.VirtualDNSRecordData.RecordData,
		&m.VirtualDNSRecordData.RecordType,
		&m.VirtualDNSRecordData.RecordTTLSeconds,
		&m.VirtualDNSRecordData.RecordMXPreference,
		&m.VirtualDNSRecordData.RecordName,
		&m.VirtualDNSRecordData.RecordClass); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func ListVirtualDNSRecord(tx *sql.Tx) ([]*models.VirtualDNSRecord, error) {
	result := models.MakeVirtualDNSRecordSlice()
	rows, err := tx.Query(listVirtualDNSRecordQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVirtualDNSRecord(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualDNSRecord(tx *sql.Tx, uuid string) (*models.VirtualDNSRecord, error) {
	rows, err := tx.Query(showVirtualDNSRecordQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualDNSRecord(rows)
	}
	return nil, nil
}

func UpdateVirtualDNSRecord(tx *sql.Tx, uuid string, model *models.VirtualDNSRecord) error {
	return nil
}

func DeleteVirtualDNSRecord(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualDNSRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
