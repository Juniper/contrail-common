package db
// discovery_service_assignment

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertDiscoveryServiceAssignmentQuery = "insert into `discovery_service_assignment` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDiscoveryServiceAssignmentQuery = "update `discovery_service_assignment` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?;"
const deleteDiscoveryServiceAssignmentQuery = "delete from `discovery_service_assignment`"
const selectDiscoveryServiceAssignmentQuery = "select `display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access` from `discovery_service_assignment`"

func CreateDiscoveryServiceAssignment(tx *sql.Tx, model *models.DiscoveryServiceAssignment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDiscoveryServiceAssignmentQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.UUID,
    model.FQName,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess)
    return err
}

func ListDiscoveryServiceAssignment(tx *sql.Tx) ([]*models.DiscoveryServiceAssignment, error) {
    result := models.MakeDiscoveryServiceAssignmentSlice()
    rows, err := tx.Query(selectDiscoveryServiceAssignmentQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeDiscoveryServiceAssignment()
            if err := rows.Scan(&m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.UUID,
                &m.FQName,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowDiscoveryServiceAssignment(db *sql.DB, id string, model *models.DiscoveryServiceAssignment) error {
    return nil
}

func UpdateDiscoveryServiceAssignment(db *sql.DB, id string, model *models.DiscoveryServiceAssignment) error {
    return nil
}

func DeleteDiscoveryServiceAssignment(db *sql.DB, id string) error {
    return nil
}