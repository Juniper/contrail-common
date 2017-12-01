package db
// virtual_DNS

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertVirtualDNSQuery = "insert into `virtual_DNS` (`next_virtual_DNS`,`dynamic_records_from_client`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`external_visible`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSQuery = "update `virtual_DNS` set `next_virtual_DNS` = ?,`dynamic_records_from_client` = ?,`reverse_resolution` = ?,`default_ttl_seconds` = ?,`record_order` = ?,`floating_ip_record` = ?,`domain_name` = ?,`external_visible` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteVirtualDNSQuery = "delete from `virtual_DNS`"
const selectVirtualDNSQuery = "select `next_virtual_DNS`,`dynamic_records_from_client`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`external_visible`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid` from `virtual_DNS`"

func CreateVirtualDNS(tx *sql.Tx, model *models.VirtualDNS) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.VirtualDNSData.NextVirtualDNS,
    model.VirtualDNSData.DynamicRecordsFromClient,
    model.VirtualDNSData.ReverseResolution,
    model.VirtualDNSData.DefaultTTLSeconds,
    model.VirtualDNSData.RecordOrder,
    model.VirtualDNSData.FloatingIPRecord,
    model.VirtualDNSData.DomainName,
    model.VirtualDNSData.ExternalVisible,
    model.FQName,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID)
    return err
}

func ListVirtualDNS(tx *sql.Tx) ([]*models.VirtualDNS, error) {
    result := models.MakeVirtualDNSSlice()
    rows, err := tx.Query(selectVirtualDNSQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeVirtualDNS()
            if err := rows.Scan(&m.VirtualDNSData.NextVirtualDNS,
                &m.VirtualDNSData.DynamicRecordsFromClient,
                &m.VirtualDNSData.ReverseResolution,
                &m.VirtualDNSData.DefaultTTLSeconds,
                &m.VirtualDNSData.RecordOrder,
                &m.VirtualDNSData.FloatingIPRecord,
                &m.VirtualDNSData.DomainName,
                &m.VirtualDNSData.ExternalVisible,
                &m.FQName,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowVirtualDNS(db *sql.DB, id string, model *models.VirtualDNS) error {
    return nil
}

func UpdateVirtualDNS(db *sql.DB, id string, model *models.VirtualDNS) error {
    return nil
}

func DeleteVirtualDNS(db *sql.DB, id string) error {
    return nil
}