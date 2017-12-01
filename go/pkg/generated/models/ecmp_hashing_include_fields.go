package models

// EcmpHashingIncludeFields

import "encoding/json"

type EcmpHashingIncludeFields struct {
	DestinationPort   bool `json:"destination_port"`
	DestinationIP     bool `json:"destination_ip"`
	IPProtocol        bool `json:"ip_protocol"`
	SourceIP          bool `json:"source_ip"`
	HashingConfigured bool `json:"hashing_configured"`
	SourcePort        bool `json:"source_port"`
}

func (model *EcmpHashingIncludeFields) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeEcmpHashingIncludeFields() *EcmpHashingIncludeFields {
	return &EcmpHashingIncludeFields{
		//TODO(nati): Apply default
		HashingConfigured: false,
		SourcePort:        false,
		DestinationPort:   false,
		DestinationIP:     false,
		IPProtocol:        false,
		SourceIP:          false,
	}
}

func InterfaceToEcmpHashingIncludeFields(iData interface{}) *EcmpHashingIncludeFields {
	data := iData.(map[string]interface{})
	return &EcmpHashingIncludeFields{
		DestinationIP: data["destination_ip"].(bool),

		//{"Title":"","Description":"When false, do not use destination ip in the ECMP hash calculation","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DestinationIP","GoType":"bool"}
		IPProtocol: data["ip_protocol"].(bool),

		//{"Title":"","Description":"When false, do not use ip protocol in the ECMP hash calculation","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPProtocol","GoType":"bool"}
		SourceIP: data["source_ip"].(bool),

		//{"Title":"","Description":"When false, do not use source ip in the ECMP hash calculation","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SourceIP","GoType":"bool"}
		HashingConfigured: data["hashing_configured"].(bool),

		//{"Title":"","Description":"When True, Packet header fields used in calculating ECMP hash is decided by following flags","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"HashingConfigured","GoType":"bool"}
		SourcePort: data["source_port"].(bool),

		//{"Title":"","Description":"When false, do not use source port in the ECMP hash calculation","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SourcePort","GoType":"bool"}
		DestinationPort: data["destination_port"].(bool),

		//{"Title":"","Description":"When false, do not use destination port in the ECMP hash calculation","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DestinationPort","GoType":"bool"}

	}
}

func InterfaceToEcmpHashingIncludeFieldsSlice(data interface{}) []*EcmpHashingIncludeFields {
	list := data.([]interface{})
	result := MakeEcmpHashingIncludeFieldsSlice()
	for _, item := range list {
		result = append(result, InterfaceToEcmpHashingIncludeFields(item))
	}
	return result
}

func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
	return []*EcmpHashingIncludeFields{}
}
