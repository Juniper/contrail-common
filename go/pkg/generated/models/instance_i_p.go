package models

// InstanceIP

import "encoding/json"

type InstanceIP struct {
	ServiceHealthCheckIP  bool                `json:"service_health_check_ip"`
	SecondaryIPTrackingIP *SubnetType         `json:"secondary_ip_tracking_ip"`
	InstanceIPAddress     IpAddressType       `json:"instance_ip_address"`
	InstanceIPLocalIP     bool                `json:"instance_ip_local_ip"`
	InstanceIPSecondary   bool                `json:"instance_ip_secondary"`
	Perms2                *PermType2          `json:"perms2"`
	Annotations           *KeyValuePairs      `json:"annotations"`
	InstanceIPMode        AddressMode         `json:"instance_ip_mode"`
	ServiceInstanceIP     bool                `json:"service_instance_ip"`
	IDPerms               *IdPermsType        `json:"id_perms"`
	DisplayName           string              `json:"display_name"`
	SubnetUUID            string              `json:"subnet_uuid"`
	InstanceIPFamily      IpAddressFamilyType `json:"instance_ip_family"`
	UUID                  string              `json:"uuid"`
	FQName                []string            `json:"fq_name"`
}

func (model *InstanceIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeInstanceIP() *InstanceIP {
	return &InstanceIP{
		//TODO(nati): Apply default
		ServiceInstanceIP:     false,
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		InstanceIPMode:        MakeAddressMode(),
		InstanceIPFamily:      MakeIpAddressFamilyType(),
		UUID:                  "",
		FQName:                []string{},
		SubnetUUID:            "",
		SecondaryIPTrackingIP: MakeSubnetType(),
		ServiceHealthCheckIP:  false,
		InstanceIPLocalIP:     false,
		InstanceIPSecondary:   false,
		Perms2:                MakePermType2(),
		InstanceIPAddress:     MakeIpAddressType(),
	}
}

func InterfaceToInstanceIP(iData interface{}) *InstanceIP {
	data := iData.(map[string]interface{})
	return &InstanceIP{
		SubnetUUID: data["subnet_uuid"].(string),

		//{"Title":"","Description":"This instance ip was allocated from this Subnet(UUID).","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"subnet_uuid","Item":null,"GoName":"SubnetUUID","GoType":"string"}
		InstanceIPFamily: InterfaceToIpAddressFamilyType(data["instance_ip_family"]),

		//{"Title":"","Description":"Ip address family for instance ip, IPv4(v4) or IPv6(v6).","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["v4","v6"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressFamilyType","CollectionType":"","Column":"instance_ip_family","Item":null,"GoName":"InstanceIPFamily","GoType":"IpAddressFamilyType"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		ServiceHealthCheckIP: data["service_health_check_ip"].(bool),

		//{"Title":"","Description":"This instance ip is used as service health check source ip","SQL":"bool","Default":false,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"service_health_check_ip","Item":null,"GoName":"ServiceHealthCheckIP","GoType":"bool"}
		SecondaryIPTrackingIP: InterfaceToSubnetType(data["secondary_ip_tracking_ip"]),

		//{"Title":"","Description":"When this instance ip is secondary ip, it can track activeness of another ip.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix_len","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SecondaryIPTrackingIP","GoType":"SubnetType"}
		InstanceIPAddress: InterfaceToIpAddressType(data["instance_ip_address"]),

		//{"Title":"","Description":"Ip address value for instance ip.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"instance_ip_address","Item":null,"GoName":"InstanceIPAddress","GoType":"IpAddressType"}
		InstanceIPLocalIP: data["instance_ip_local_ip"].(bool),

		//{"Title":"","Description":"This instance ip is local to compute and will not be exported to other nodes.","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"instance_ip_local_ip","Item":null,"GoName":"InstanceIPLocalIP","GoType":"bool"}
		InstanceIPSecondary: data["instance_ip_secondary"].(bool),

		//{"Title":"","Description":"This instance ip is secondary ip of the interface.","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"instance_ip_secondary","Item":null,"GoName":"InstanceIPSecondary","GoType":"bool"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		InstanceIPMode: InterfaceToAddressMode(data["instance_ip_mode"]),

		//{"Title":"","Description":"Ip address HA mode in case this instance ip is used in more than one interface, active-Active or active-Standby.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"instance_ip_mode","Item":null,"GoName":"InstanceIPMode","GoType":"AddressMode"}
		ServiceInstanceIP: data["service_instance_ip"].(bool),

		//{"Title":"","Description":"This instance ip is used as service chain next hop","SQL":"bool","Default":false,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"service_instance_ip","Item":null,"GoName":"ServiceInstanceIP","GoType":"bool"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"permissions_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"permissions_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}

	}
}

func InterfaceToInstanceIPSlice(data interface{}) []*InstanceIP {
	list := data.([]interface{})
	result := MakeInstanceIPSlice()
	for _, item := range list {
		result = append(result, InterfaceToInstanceIP(item))
	}
	return result
}

func MakeInstanceIPSlice() []*InstanceIP {
	return []*InstanceIP{}
}
