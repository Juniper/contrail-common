package db

// customer_attachment

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

const insertCustomerAttachmentQuery = "insert into `customer_attachment` (`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateCustomerAttachmentQuery = "update `customer_attachment` set `global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteCustomerAttachmentQuery = "delete from `customer_attachment` where uuid = ?"
const listCustomerAttachmentQuery = "select `global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair` from `customer_attachment`"
const showCustomerAttachmentQuery = "select `global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair` from `customer_attachment` where uuid = ?"

func CreateCustomerAttachment(tx *sql.Tx, model *models.CustomerAttachment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertCustomerAttachmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.GlobalAccess),
		util.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		util.MustJSON(model.FQName),
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
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		util.MustJSON(model.Annotations.KeyValuePair))
	return err
}

func scanCustomerAttachment(rows *sql.Rows) (*models.CustomerAttachment, error) {
	m := models.MakeCustomerAttachment()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func ListCustomerAttachment(tx *sql.Tx) ([]*models.CustomerAttachment, error) {
	result := models.MakeCustomerAttachmentSlice()
	rows, err := tx.Query(listCustomerAttachmentQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanCustomerAttachment(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowCustomerAttachment(tx *sql.Tx, uuid string) (*models.CustomerAttachment, error) {
	rows, err := tx.Query(showCustomerAttachmentQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanCustomerAttachment(rows)
	}
	return nil, nil
}

func UpdateCustomerAttachment(tx *sql.Tx, uuid string, model *models.CustomerAttachment) error {
	return nil
}

func DeleteCustomerAttachment(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteCustomerAttachmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
