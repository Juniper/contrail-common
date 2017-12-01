package db
// contrail_cluster

import (
  	"database/sql"
    "github.com/Juniper/contrail-common/go/pkg/generated/models"
)

const insertContrailClusterQuery = "insert into `contrail_cluster` (`contrail_webui`,`default_vrouter_bond_interface_members`,`statistics_ttl`,`flow_ttl`,`fq_name`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`config_audit_ttl`,`data_ttl`,`default_gateway`,`uuid`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`default_vrouter_bond_interface`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailClusterQuery = "update `contrail_cluster` set `contrail_webui` = ?,`default_vrouter_bond_interface_members` = ?,`statistics_ttl` = ?,`flow_ttl` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`config_audit_ttl` = ?,`data_ttl` = ?,`default_gateway` = ?,`uuid` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`default_vrouter_bond_interface` = ?;"
const deleteContrailClusterQuery = "delete from `contrail_cluster`"
const selectContrailClusterQuery = "select `contrail_webui`,`default_vrouter_bond_interface_members`,`statistics_ttl`,`flow_ttl`,`fq_name`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`config_audit_ttl`,`data_ttl`,`default_gateway`,`uuid`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`default_vrouter_bond_interface` from `contrail_cluster`"

func CreateContrailCluster(tx *sql.Tx, model *models.ContrailCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailClusterQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(model.ContrailWebui,
    model.DefaultVrouterBondInterfaceMembers,
    model.StatisticsTTL,
    model.FlowTTL,
    model.FQName,
    model.DisplayName,
    model.Annotations.KeyValuePair,
    model.Perms2.Share,
    model.Perms2.Owner,
    model.Perms2.OwnerAccess,
    model.Perms2.GlobalAccess,
    model.ConfigAuditTTL,
    model.DataTTL,
    model.DefaultGateway,
    model.UUID,
    model.IDPerms.Permissions.Owner,
    model.IDPerms.Permissions.OwnerAccess,
    model.IDPerms.Permissions.OtherAccess,
    model.IDPerms.Permissions.Group,
    model.IDPerms.Permissions.GroupAccess,
    model.IDPerms.Enable,
    model.IDPerms.Description,
    model.IDPerms.Created,
    model.IDPerms.Creator,
    model.IDPerms.UserVisible,
    model.IDPerms.LastModified,
    model.DefaultVrouterBondInterface)
    return err
}

func ListContrailCluster(tx *sql.Tx) ([]*models.ContrailCluster, error) {
    result := models.MakeContrailClusterSlice()
    rows, err := tx.Query(selectContrailClusterQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
            m := models.MakeContrailCluster()
            if err := rows.Scan(&m.ContrailWebui,
                &m.DefaultVrouterBondInterfaceMembers,
                &m.StatisticsTTL,
                &m.FlowTTL,
                &m.FQName,
                &m.DisplayName,
                &m.Annotations.KeyValuePair,
                &m.Perms2.Share,
                &m.Perms2.Owner,
                &m.Perms2.OwnerAccess,
                &m.Perms2.GlobalAccess,
                &m.ConfigAuditTTL,
                &m.DataTTL,
                &m.DefaultGateway,
                &m.UUID,
                &m.IDPerms.Permissions.Owner,
                &m.IDPerms.Permissions.OwnerAccess,
                &m.IDPerms.Permissions.OtherAccess,
                &m.IDPerms.Permissions.Group,
                &m.IDPerms.Permissions.GroupAccess,
                &m.IDPerms.Enable,
                &m.IDPerms.Description,
                &m.IDPerms.Created,
                &m.IDPerms.Creator,
                &m.IDPerms.UserVisible,
                &m.IDPerms.LastModified,
                &m.DefaultVrouterBondInterface); err != nil {
                    return nil, err
            }
            result = append(result, m)
    }
    if err := rows.Err(); err != nil {
            return nil, err
    }
    return result, nil
}

func ShowContrailCluster(db *sql.DB, id string, model *models.ContrailCluster) error {
    return nil
}

func UpdateContrailCluster(db *sql.DB, id string, model *models.ContrailCluster) error {
    return nil
}

func DeleteContrailCluster(db *sql.DB, id string) error {
    return nil
}