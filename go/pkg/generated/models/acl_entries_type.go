package models
// AclEntriesType



import "encoding/json"

type AclEntriesType struct {

    Dynamic bool `json:"dynamic"`
    ACLRule []*AclRuleType `json:"acl_rule"`
}

func (model *AclEntriesType) String() string{
    b, _ := json.Marshal(model)
    return string(b)
}

func MakeAclEntriesType() *AclEntriesType{
    return &AclEntriesType{
    //TODO(nati): Apply default
    Dynamic: false,
        
            // {"Title":"","Description":"For Internal use only","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"action_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"alert":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Alert","GoType":"bool"},"apply_service":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ApplyService","GoType":"string"},"GoName":"ApplyService","GoType":"[]string"},"assign_routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AssignRoutingInstance","GoType":"string"},"gateway_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GatewayName","GoType":"string"},"log":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Log","GoType":"bool"},"mirror_to":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string"},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string"},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string"},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string"},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool"},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType"},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool"},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType"},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string"},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType"},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string"},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType"},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType"},"qos_action":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"QosAction","GoType":"string"},"simple_action":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["deny","pass"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SimpleActionType","CollectionType":"","Column":"","Item":null,"GoName":"SimpleAction","GoType":"SimpleActionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ActionListType","CollectionType":"","Column":"","Item":null,"GoName":"ActionList","GoType":"ActionListType"},"direction":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["\u003e","\u003c\u003e"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DirectionType","CollectionType":"","Column":"","Item":null,"GoName":"Direction","GoType":"DirectionType"},"match_condition":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"dst_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"DSTAddress","GoType":"AddressType"},"dst_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"PortType"},"ethertype":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["IPv4","IPv6"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/EtherType","CollectionType":"","Column":"","Item":null,"GoName":"Ethertype","GoType":"EtherType"},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"},"src_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"SRCAddress","GoType":"AddressType"},"src_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MatchConditionType","CollectionType":"","Column":"","Item":null,"GoName":"MatchCondition","GoType":"MatchConditionType"},"rule_uuid":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleUUID","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AclRuleType","CollectionType":"","Column":"","Item":null,"GoName":"ACLRule","GoType":"AclRuleType"},"GoName":"ACLRule","GoType":"[]*AclRuleType"}
            
                ACLRule:  MakeAclRuleTypeSlice(),
            
        
    }
}

func InterfaceToAclEntriesType(iData interface{}) *AclEntriesType {
    data := iData.(map[string]interface{})
    return &AclEntriesType{
    Dynamic: data["dynamic"].(bool),
        
        //{"Title":"","Description":"For Internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Dynamic","GoType":"bool"}
        
            
                ACLRule:  InterfaceToAclRuleTypeSlice(data["acl_rule"]),
            
        
        //{"Title":"","Description":"For Internal use only","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"action_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"alert":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Alert","GoType":"bool"},"apply_service":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ApplyService","GoType":"string"},"GoName":"ApplyService","GoType":"[]string"},"assign_routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AssignRoutingInstance","GoType":"string"},"gateway_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GatewayName","GoType":"string"},"log":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Log","GoType":"bool"},"mirror_to":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string"},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string"},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string"},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string"},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool"},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType"},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool"},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType"},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string"},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType"},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string"},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType"},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType"},"qos_action":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"QosAction","GoType":"string"},"simple_action":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["deny","pass"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SimpleActionType","CollectionType":"","Column":"","Item":null,"GoName":"SimpleAction","GoType":"SimpleActionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ActionListType","CollectionType":"","Column":"","Item":null,"GoName":"ActionList","GoType":"ActionListType"},"direction":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["\u003e","\u003c\u003e"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DirectionType","CollectionType":"","Column":"","Item":null,"GoName":"Direction","GoType":"DirectionType"},"match_condition":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"dst_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"DSTAddress","GoType":"AddressType"},"dst_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"PortType"},"ethertype":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["IPv4","IPv6"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/EtherType","CollectionType":"","Column":"","Item":null,"GoName":"Ethertype","GoType":"EtherType"},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"},"src_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"SRCAddress","GoType":"AddressType"},"src_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MatchConditionType","CollectionType":"","Column":"","Item":null,"GoName":"MatchCondition","GoType":"MatchConditionType"},"rule_uuid":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleUUID","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AclRuleType","CollectionType":"","Column":"","Item":null,"GoName":"ACLRule","GoType":"AclRuleType"},"GoName":"ACLRule","GoType":"[]*AclRuleType"}
        
    }
}


func InterfaceToAclEntriesTypeSlice(data interface{}) []*AclEntriesType {
    list := data.([]interface{})
    result := MakeAclEntriesTypeSlice()
    for _, item := range list {
        result = append(result, InterfaceToAclEntriesType(item))
    }
    return result
}

func MakeAclEntriesTypeSlice() []*AclEntriesType {
    return []*AclEntriesType{}
}
