package models

// PolicyRuleType

import "encoding/json"

type PolicyRuleType struct {
	ActionList   *ActionListType `json:"action_list"`
	RuleUUID     string          `json:"rule_uuid"`
	Application  []string        `json:"application"`
	LastModified string          `json:"last_modified"`
	SRCAddresses []*AddressType  `json:"src_addresses"`
	RuleSequence *SequenceType   `json:"rule_sequence"`
	Protocol     string          `json:"protocol"`
	DSTAddresses []*AddressType  `json:"dst_addresses"`
	Created      string          `json:"created"`
	DSTPorts     []*PortType     `json:"dst_ports"`
	Ethertype    EtherType       `json:"ethertype"`
	SRCPorts     []*PortType     `json:"src_ports"`
	Direction    DirectionType   `json:"direction"`
}

func (model *PolicyRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePolicyRuleType() *PolicyRuleType {
	return &PolicyRuleType{
		//TODO(nati): Apply default
		Direction: MakeDirectionType(),

		DSTAddresses: MakeAddressTypeSlice(),

		Created: "",

		DSTPorts: MakePortTypeSlice(),

		Ethertype: MakeEtherType(),

		SRCPorts: MakePortTypeSlice(),

		Protocol:     "",
		ActionList:   MakeActionListType(),
		RuleUUID:     "",
		Application:  []string{},
		LastModified: "",

		SRCAddresses: MakeAddressTypeSlice(),

		RuleSequence: MakeSequenceType(),
	}
}

func InterfaceToPolicyRuleType(iData interface{}) *PolicyRuleType {
	data := iData.(map[string]interface{})
	return &PolicyRuleType{

		DSTPorts: InterfaceToPortTypeSlice(data["dst_ports"]),

		//{"Title":"","Description":"Range of destination  port for layer 4 protocol","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPorts","GoType":"PortType"},"GoName":"DSTPorts","GoType":"[]*PortType"}
		Ethertype: InterfaceToEtherType(data["ethertype"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["IPv4","IPv6"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/EtherType","CollectionType":"","Column":"","Item":null,"GoName":"Ethertype","GoType":"EtherType"}

		SRCPorts: InterfaceToPortTypeSlice(data["src_ports"]),

		//{"Title":"","Description":"Range of source port for layer 4 protocol","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPorts","GoType":"PortType"},"GoName":"SRCPorts","GoType":"[]*PortType"}
		Direction: InterfaceToDirectionType(data["direction"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["\u003e","\u003c\u003e"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DirectionType","CollectionType":"","Column":"","Item":null,"GoName":"Direction","GoType":"DirectionType"}

		DSTAddresses: InterfaceToAddressTypeSlice(data["dst_addresses"]),

		//{"Title":"","Description":"Destination ip matching criteria","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"DSTAddresses","GoType":"AddressType"},"GoName":"DSTAddresses","GoType":"[]*AddressType"}
		Created: data["created"].(string),

		//{"Title":"","Description":"timestamp when security group rule object gets created","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Created","GoType":"string"}
		Application: data["application"].([]string),

		//{"Title":"","Description":"Optionally application can be specified instead of protocol and port. not currently implemented","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Application","GoType":"string"},"GoName":"Application","GoType":"[]string"}
		LastModified: data["last_modified"].(string),

		//{"Title":"","Description":"timestamp when security group rule object gets updated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LastModified","GoType":"string"}

		SRCAddresses: InterfaceToAddressTypeSlice(data["src_addresses"]),

		//{"Title":"","Description":"Source ip matching criteria","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"SRCAddresses","GoType":"AddressType"},"GoName":"SRCAddresses","GoType":"[]*AddressType"}
		RuleSequence: InterfaceToSequenceType(data["rule_sequence"]),

		//{"Title":"","Description":"Deprecated, Will be removed because rules themselves are already an ordered list","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"major":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Major","GoType":"int"},"minor":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Minor","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SequenceType","CollectionType":"","Column":"","Item":null,"GoName":"RuleSequence","GoType":"SequenceType"}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"Layer 4 protocol in ip packet","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"}
		ActionList: InterfaceToActionListType(data["action_list"]),

		//{"Title":"","Description":"Actions to be performed if packets match condition","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"alert":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Alert","GoType":"bool"},"apply_service":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ApplyService","GoType":"string"},"GoName":"ApplyService","GoType":"[]string"},"assign_routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AssignRoutingInstance","GoType":"string"},"gateway_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GatewayName","GoType":"string"},"log":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Log","GoType":"bool"},"mirror_to":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string"},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string"},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string"},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string"},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool"},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType"},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool"},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType"},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string"},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType"},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string"},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType"},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType"},"qos_action":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"QosAction","GoType":"string"},"simple_action":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["deny","pass"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SimpleActionType","CollectionType":"","Column":"","Item":null,"GoName":"SimpleAction","GoType":"SimpleActionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ActionListType","CollectionType":"","Column":"","Item":null,"GoName":"ActionList","GoType":"ActionListType"}
		RuleUUID: data["rule_uuid"].(string),

		//{"Title":"","Description":"Rule UUID is identifier used in flow records to identify rule","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleUUID","GoType":"string"}

	}
}

func InterfaceToPolicyRuleTypeSlice(data interface{}) []*PolicyRuleType {
	list := data.([]interface{})
	result := MakePolicyRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPolicyRuleType(item))
	}
	return result
}

func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
