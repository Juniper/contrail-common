package db
// qos_queue

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertQosQueueQuery = "insert into `qos_queue` (`qos_queue_identifier`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`max_bandwidth`,`min_bandwidth`,`key_value_pair`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosQueueQuery = "update `qos_queue` set `qos_queue_identifier` = ?,`display_name` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`max_bandwidth` = ?,`min_bandwidth` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?;"
const deleteQosQueueQuery = "delete from `qos_queue`"
const selectQosQueueQuery = "select `qos_queue_identifier`,`display_name`,`owner_access`,`global_access`,`share`,`owner`,`max_bandwidth`,`min_bandwidth`,`key_value_pair`,`uuid`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator` from `qos_queue`"

func CreateQosQueue(tx *sql.Tx, model *models.QosQueue) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosQueueQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.QosQueueIdentifier,
    model.DisplayName,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.MaxBandwidth,
    model.MinBandwidth,
    model.Annotations.KeyValuePair,
    model.UUID,
    model.FQName,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator)
    return err
}

func ListQosQueue(tx *sql.Tx) ([]*models.QosQueue, error) {
    result := models.MakeQosQueueSlice()
    rows, err := tx.Query(selectQosQueueQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeQosQueue()
            if err := rows.Scan(&m.QosQueueIdentifier,
                &m.DisplayName,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.MaxBandwidth,
                &m.MinBandwidth,
                &m.Annotations.KeyValuePair,
                &m.UUID,
                &m.FQName,
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
                &m.IDPerms.Creator); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowQosQueue(db *sql.DB, id string, model *models.QosQueue) error {
    return nil
}

func UpdateQosQueue(db *sql.DB, id string, model *models.QosQueue) error {
    return nil
}

func DeleteQosQueue(db *sql.DB, id string) error {
    return nil
}