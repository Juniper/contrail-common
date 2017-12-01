package models

// FirewallRule

import "encoding/json"

type FirewallRule struct {
	MatchTagTypes *FirewallRuleMatchTagsTypeIdList `json:"match_tag_types"`
	MatchTags     *FirewallRuleMatchTagsType       `json:"match_tags"`
	Perms2        *PermType2                       `json:"perms2"`
	Endpoint2     *FirewallRuleEndpointType        `json:"endpoint_2"`
	Service       *FirewallServiceType             `json:"service"`
	Direction     FirewallRuleDirectionType        `json:"direction"`
	UUID          string                           `json:"uuid"`
	FQName        []string                         `json:"fq_name"`
	IDPerms       *IdPermsType                     `json:"id_perms"`
	DisplayName   string                           `json:"display_name"`
	Annotations   *KeyValuePairs                   `json:"annotations"`
	Endpoint1     *FirewallRuleEndpointType        `json:"endpoint_1"`
	ActionList    *ActionListType                  `json:"action_list"`
}

func (model *FirewallRule) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallRule() *FirewallRule {
	return &FirewallRule{
		//TODO(nati): Apply default
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Endpoint1:     MakeFirewallRuleEndpointType(),
		ActionList:    MakeActionListType(),
		Direction:     MakeFirewallRuleDirectionType(),
		UUID:          "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		Endpoint2:     MakeFirewallRuleEndpointType(),
		Service:       MakeFirewallServiceType(),
		MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
		MatchTags:     MakeFirewallRuleMatchTagsType(),
		Perms2:        MakePermType2(),
	}
}

func InterfaceToFirewallRule(iData interface{}) *FirewallRule {
	data := iData.(map[string]interface{})
	return &FirewallRule{
		Direction: InterfaceToFirewallRuleDirectionType(data["direction"]),

		//{"Title":"","Description":"Direction in the rule","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":["\u003c","\u003e","\u003c\u003e"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallRuleDirectionType","CollectionType":"","Column":"direction","Item":null,"GoName":"Direction","GoType":"FirewallRuleDirectionType"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}
		Endpoint1: InterfaceToFirewallRuleEndpointType(data["endpoint_1"]),

		//{"Title":"","Description":"match condition for traffic source","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"object","Permission":null,"Properties":{"address_group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"address_group","Item":null,"GoName":"AddressGroup","GoType":"string"},"any":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"exclusive","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"any","Item":null,"GoName":"Any","GoType":"bool"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix_len","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"tag_ids":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"tag_ids","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagIds","GoType":"int"},"GoName":"TagIds","GoType":"[]int"},"tags":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"tags","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tags","GoType":"string"},"GoName":"Tags","GoType":"[]string"},"virtual_network":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"virtual_network","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallRuleEndpointType","CollectionType":"","Column":"","Item":null,"GoName":"Endpoint1","GoType":"FirewallRuleEndpointType"}
		ActionList: InterfaceToActionListType(data["action_list"]),

		//{"Title":"","Description":"Actions to be performed if packets match condition","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"object","Permission":null,"Properties":{"alert":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"alert","Item":null,"GoName":"Alert","GoType":"bool"},"apply_service":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"apply_service","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ApplyService","GoType":"string"},"GoName":"ApplyService","GoType":"[]string"},"assign_routing_instance":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"assign_routing_instance","Item":null,"GoName":"AssignRoutingInstance","GoType":"string"},"gateway_name":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gateway_name","Item":null,"GoName":"GatewayName","GoType":"string"},"log":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"log","Item":null,"GoName":"Log","GoType":"bool"},"mirror_to":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"analyzer_ip_address","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string"},"analyzer_mac_address":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"analyzer_mac_address","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string"},"analyzer_name":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"analyzer_name","Item":null,"GoName":"AnalyzerName","GoType":"string"},"encapsulation":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"encapsulation","Item":null,"GoName":"Encapsulation","GoType":"string"},"juniper_header":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"juniper_header","Item":null,"GoName":"JuniperHeader","GoType":"bool"},"nh_mode":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"nh_mode","Item":null,"GoName":"NHMode","GoType":"NHModeType"},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"nic_assisted_mirroring","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool"},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"nic_assisted_mirroring_vlan","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType"},"routing_instance":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"routing_instance","Item":null,"GoName":"RoutingInstance","GoType":"string"},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"vni","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType"},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vtep_dst_ip_address","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string"},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vtep_dst_mac_address","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType"},"udp_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"udp_port","Item":null,"GoName":"UDPPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType"},"qos_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"qos_action","Item":null,"GoName":"QosAction","GoType":"string"},"simple_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["deny","pass"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SimpleActionType","CollectionType":"","Column":"simple_action","Item":null,"GoName":"SimpleAction","GoType":"SimpleActionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ActionListType","CollectionType":"","Column":"","Item":null,"GoName":"ActionList","GoType":"ActionListType"}
		MatchTagTypes: InterfaceToFirewallRuleMatchTagsTypeIdList(data["match_tag_types"]),

		//{"Title":"","Description":"matching tags ids for source and destination endpoints","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"object","Permission":null,"Properties":{"tag_type":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"tag_type","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagType","GoType":"int"},"GoName":"TagType","GoType":"[]int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallRuleMatchTagsTypeIdList","CollectionType":"","Column":"","Item":null,"GoName":"MatchTagTypes","GoType":"FirewallRuleMatchTagsTypeIdList"}
		MatchTags: InterfaceToFirewallRuleMatchTagsType(data["match_tags"]),

		//{"Title":"","Description":"matching tags for source and destination endpoints","SQL":"text","Default":null,"Operation":"","Presence":"required","Type":"object","Permission":null,"Properties":{"tag_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagList","GoType":"string"},"GoName":"TagList","GoType":"[]string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallRuleMatchTagsType","CollectionType":"list","Column":"match_tags","Item":null,"GoName":"MatchTags","GoType":"FirewallRuleMatchTagsType"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		Endpoint2: InterfaceToFirewallRuleEndpointType(data["endpoint_2"]),

		//{"Title":"","Description":"match condition for traffic destination","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"object","Permission":null,"Properties":{"address_group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_address_group","Item":null,"GoName":"AddressGroup","GoType":"string"},"any":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"exclusive","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_any","Item":null,"GoName":"Any","GoType":"bool"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_subnet_ip_prefix","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_subnet_ip_prefix_len","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"tag_ids":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_tag_ids","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagIds","GoType":"int"},"GoName":"TagIds","GoType":"[]int"},"tags":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_tags","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tags","GoType":"string"},"GoName":"Tags","GoType":"[]string"},"virtual_network":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"endpoint_2_virtual_network","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallRuleEndpointType","CollectionType":"","Column":"","Item":null,"GoName":"Endpoint2","GoType":"FirewallRuleEndpointType"}
		Service: InterfaceToFirewallServiceType(data["service"]),

		//{"Title":"","Description":"Service (port, protocol) for packets match condition","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"dst_ports":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"dst_ports_end_port","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"dst_ports_start_port","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPorts","GoType":"PortType"},"protocol":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"protocol","Item":null,"GoName":"Protocol","GoType":"string"},"protocol_id":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"protocol_id","Item":null,"GoName":"ProtocolID","GoType":"int"},"src_ports":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"end_port","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"start_port","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPorts","GoType":"PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallServiceType","CollectionType":"","Column":"","Item":null,"GoName":"Service","GoType":"FirewallServiceType"}

	}
}

func InterfaceToFirewallRuleSlice(data interface{}) []*FirewallRule {
	list := data.([]interface{})
	result := MakeFirewallRuleSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRule(item))
	}
	return result
}

func MakeFirewallRuleSlice() []*FirewallRule {
	return []*FirewallRule{}
}
