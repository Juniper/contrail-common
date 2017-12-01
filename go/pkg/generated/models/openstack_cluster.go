package models

// OpenstackCluster

import "encoding/json"

type OpenstackCluster struct {
	AdminPassword                             string         `json:"admin_password"`
	PublicGateway                             string         `json:"public_gateway"`
	Perms2                                    *PermType2     `json:"perms2"`
	UUID                                      string         `json:"uuid"`
	Annotations                               *KeyValuePairs `json:"annotations"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members"`
	ExternalNetCidr                           string         `json:"external_net_cidr"`
	OpenstackWebui                            string         `json:"openstack_webui"`
	PublicIP                                  string         `json:"public_ip"`
	FQName                                    []string       `json:"fq_name"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time"`
	ProvisioningState                         string         `json:"provisioning_state"`
	DefaultJournalDrives                      string         `json:"default_journal_drives"`
	DefaultOsdDrives                          string         `json:"default_osd_drives"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end"`
	IDPerms                                   *IdPermsType   `json:"id_perms"`
	DisplayName                               string         `json:"display_name"`
	ProvisioningProgress                      int            `json:"provisioning_progress"`
	ContrailClusterID                         string         `json:"contrail_cluster_id"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start"`
	ProvisioningLog                           string         `json:"provisioning_log"`
}

func (model *OpenstackCluster) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeOpenstackCluster() *OpenstackCluster {
	return &OpenstackCluster{
		//TODO(nati): Apply default
		ExternalAllocationPoolStart:              "",
		ProvisioningLog:                          "",
		ContrailClusterID:                        "",
		DefaultCapacityDrives:                    "",
		DefaultStorageAccessBondInterfaceMembers: "",
		UUID:                                      "",
		Annotations:                               MakeKeyValuePairs(),
		ProvisioningProgressStage:                 "",
		AdminPassword:                             "",
		PublicGateway:                             "",
		Perms2:                                    MakePermType2(),
		PublicIP:                                  "",
		FQName:                                    []string{},
		ProvisioningStartTime:                     "",
		ProvisioningState:                         "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalNetCidr:                           "",
		OpenstackWebui:                            "",
		ExternalAllocationPoolEnd:                 "",
		IDPerms:                                   MakeIdPermsType(),
		DisplayName:                               "",
		ProvisioningProgress:                      0,
		DefaultJournalDrives:                      "",
		DefaultOsdDrives:                          "",
		DefaultPerformanceDrives:                  "",
	}
}

func InterfaceToOpenstackCluster(iData interface{}) *OpenstackCluster {
	data := iData.(map[string]interface{})
	return &OpenstackCluster{
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"Title":"Provisioning Progress","Description":"","SQL":"int","Default":0,"Operation":"","Presence":"","Type":"integer","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress","Item":null,"GoName":"ProvisioningProgress","GoType":"int"}
		DefaultJournalDrives: data["default_journal_drives"].(string),

		//{"Title":"Journal Drives  for Storage Node Role","Description":"SSD Drives to use for journals","SQL":"varchar(255)","Default":"sdf","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_journal_drives","Item":null,"GoName":"DefaultJournalDrives","GoType":"string"}
		DefaultOsdDrives: data["default_osd_drives"].(string),

		//{"Title":"Stoage Drives for Storage Node Role","Description":"Drives to use for cloud storage","SQL":"varchar(255)","Default":"sdc,sdd,sde","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_osd_drives","Item":null,"GoName":"DefaultOsdDrives","GoType":"string"}
		DefaultPerformanceDrives: data["default_performance_drives"].(string),

		//{"Title":"Default Performance Drive  for Controller Node Role","Description":"Drives for performance oriented application such as journaling  for Controller Node Role","SQL":"varchar(255)","Default":"sdf","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_performance_drives","Item":null,"GoName":"DefaultPerformanceDrives","GoType":"string"}
		ExternalAllocationPoolEnd: data["external_allocation_pool_end"].(string),

		//{"Title":"External Allocation pool end","Description":"End of the allocation pool range","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_allocation_pool_end","Item":null,"GoName":"ExternalAllocationPoolEnd","GoType":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}
		ContrailClusterID: data["contrail_cluster_id"].(string),

		//{"Title":"Contrail Cluster ID","Description":"contrial cluster ID","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"contrail_cluster_id","Item":null,"GoName":"ContrailClusterID","GoType":"string"}
		DefaultCapacityDrives: data["default_capacity_drives"].(string),

		//{"Title":"Default Capacity Drives  for Controller Node Role","Description":"Drives for capacity oriented applications such as logging for Controller Node Role","SQL":"varchar(255)","Default":"sdc,sdd,sde","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_capacity_drives","Item":null,"GoName":"DefaultCapacityDrives","GoType":"string"}
		DefaultStorageAccessBondInterfaceMembers: data["default_storage_access_bond_interface_members"].(string),

		//{"Title":"Default Storage Access  Bond Interface Members","Description":"Storage Management  Bond Interface Members","SQL":"varchar(255)","Default":"ens8f0,ens8f1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_storage_access_bond_interface_members","Item":null,"GoName":"DefaultStorageAccessBondInterfaceMembers","GoType":"string"}
		ExternalAllocationPoolStart: data["external_allocation_pool_start"].(string),

		//{"Title":"External Allocation pool start","Description":"Start of the allocation pool range","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_allocation_pool_start","Item":null,"GoName":"ExternalAllocationPoolStart","GoType":"string"}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"Title":"Provisioning Log","Description":"","SQL":"text","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_log","Item":null,"GoName":"ProvisioningLog","GoType":"string"}
		AdminPassword: data["admin_password"].(string),

		//{"Title":"Admin Password","Description":"Password for admin openstack account","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"admin_password","Item":null,"GoName":"AdminPassword","GoType":"string"}
		PublicGateway: data["public_gateway"].(string),

		//{"Title":"Public Gateway","Description":"Gateway for public VIP","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"public_gateway","Item":null,"GoName":"PublicGateway","GoType":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"Title":"Provisioning Progress Stage","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress_stage","Item":null,"GoName":"ProvisioningProgressStage","GoType":"string"}
		ProvisioningState: data["provisioning_state"].(string),

		//{"Title":"Provisioning Status","Description":"","SQL":"varchar(255)","Default":"CREATED","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_state","Item":null,"GoName":"ProvisioningState","GoType":"string"}
		DefaultStorageBackendBondInterfaceMembers: data["default_storage_backend_bond_interface_members"].(string),

		//{"Title":"Default Storage Backend Bond Interface Members","Description":"Storage Backend Bond Interface Members","SQL":"varchar(255)","Default":"ens9f0,ens9f1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_storage_backend_bond_interface_members","Item":null,"GoName":"DefaultStorageBackendBondInterfaceMembers","GoType":"string"}
		ExternalNetCidr: data["external_net_cidr"].(string),

		//{"Title":"External Network CIDR","Description":"Subnet to use for external network","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_net_cidr","Item":null,"GoName":"ExternalNetCidr","GoType":"string"}
		OpenstackWebui: data["openstack_webui"].(string),

		//{"Title":"OpenStack WebUI","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"openstack_webui","Item":null,"GoName":"OpenstackWebui","GoType":"string"}
		PublicIP: data["public_ip"].(string),

		//{"Title":"Public IP","Description":"Public Virtual IP (VIP)","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"public_ip","Item":null,"GoName":"PublicIP","GoType":"string"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"Title":"Time provisioning started","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_start_time","Item":null,"GoName":"ProvisioningStartTime","GoType":"string"}

	}
}

func InterfaceToOpenstackClusterSlice(data interface{}) []*OpenstackCluster {
	list := data.([]interface{})
	result := MakeOpenstackClusterSlice()
	for _, item := range list {
		result = append(result, InterfaceToOpenstackCluster(item))
	}
	return result
}

func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
