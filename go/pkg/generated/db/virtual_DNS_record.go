package db
// virtual_DNS_record

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualDNSRecordQuery = "insert into `virtual_DNS_record` (`record_name`,`record_class`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSRecordQuery = "update `virtual_DNS_record` set `record_name` = ?,`record_class` = ?,`record_data` = ?,`record_type` = ?,`record_ttl_seconds` = ?,`record_mx_preference` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteVirtualDNSRecordQuery = "delete from `virtual_DNS_record`"
const selectVirtualDNSRecordQuery = "select `record_name`,`record_class`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair` from `virtual_DNS_record`"

func CreateVirtualDNSRecord(tx *sql.Tx, model *models.VirtualDNSRecord) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSRecordQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.VirtualDNSRecordData.RecordName,
    model.VirtualDNSRecordData.RecordClass,
    model.VirtualDNSRecordData.RecordData,
    model.VirtualDNSRecordData.RecordType,
    model.VirtualDNSRecordData.RecordTTLSeconds,
    model.VirtualDNSRecordData.RecordMXPreference,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.UUID,
    model.FQName,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.DisplayName,
    model.Annotations.KeyValuePair)
    return err
}

func ListVirtualDNSRecord(tx *sql.Tx) ([]*models.VirtualDNSRecord, error) {
    result := models.MakeVirtualDNSRecordSlice()
    rows, err := tx.Query(selectVirtualDNSRecordQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualDNSRecord()
            if err := rows.Scan(&m.VirtualDNSRecordData.RecordName,
                &m.VirtualDNSRecordData.RecordClass,
                &m.VirtualDNSRecordData.RecordData,
                &m.VirtualDNSRecordData.RecordType,
                &m.VirtualDNSRecordData.RecordTTLSeconds,
                &m.VirtualDNSRecordData.RecordMXPreference,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.DisplayName,
                &m.Annotations.KeyValuePair); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowVirtualDNSRecord(db *sql.DB, id string, model *models.VirtualDNSRecord) error {
    return nil
}

func UpdateVirtualDNSRecord(db *sql.DB, id string, model *models.VirtualDNSRecord) error {
    return nil
}

func DeleteVirtualDNSRecord(db *sql.DB, id string) error {
    return nil
}