package models

// SNMPCredentials

import "encoding/json"

type SNMPCredentials struct {
	V3SecurityName           string `json:"v3_security_name"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol"`
	V3PrivacyPassword        string `json:"v3_privacy_password"`
	V3AuthenticationPassword string `json:"v3_authentication_password"`
	V3EngineTime             int    `json:"v3_engine_time"`
	V3EngineBoots            int    `json:"v3_engine_boots"`
	V3EngineID               string `json:"v3_engine_id"`
	LocalPort                int    `json:"local_port"`
	V3SecurityEngineID       string `json:"v3_security_engine_id"`
	Timeout                  int    `json:"timeout"`
	Version                  int    `json:"version"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol"`
	Retries                  int    `json:"retries"`
	V3SecurityLevel          string `json:"v3_security_level"`
	V3Context                string `json:"v3_context"`
	V2Community              string `json:"v2_community"`
	V3ContextEngineID        string `json:"v3_context_engine_id"`
}

func (model *SNMPCredentials) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeSNMPCredentials() *SNMPCredentials {
	return &SNMPCredentials{
		//TODO(nati): Apply default
		V3SecurityLevel:          "",
		V3Context:                "",
		V2Community:              "",
		V3ContextEngineID:        "",
		Version:                  0,
		V3PrivacyProtocol:        "",
		Retries:                  0,
		V3PrivacyPassword:        "",
		V3SecurityName:           "",
		V3AuthenticationProtocol: "",
		V3EngineBoots:            0,
		V3AuthenticationPassword: "",
		V3EngineTime:             0,
		V3SecurityEngineID:       "",
		Timeout:                  0,
		V3EngineID:               "",
		LocalPort:                0,
	}
}

func InterfaceToSNMPCredentials(iData interface{}) *SNMPCredentials {
	data := iData.(map[string]interface{})
	return &SNMPCredentials{
		V3AuthenticationPassword: data["v3_authentication_password"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3AuthenticationPassword","GoType":"string"}
		V3EngineTime: data["v3_engine_time"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3EngineTime","GoType":"int"}
		V3EngineBoots: data["v3_engine_boots"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3EngineBoots","GoType":"int"}
		V3EngineID: data["v3_engine_id"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3EngineID","GoType":"string"}
		LocalPort: data["local_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LocalPort","GoType":"int"}
		V3SecurityEngineID: data["v3_security_engine_id"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3SecurityEngineID","GoType":"string"}
		Timeout: data["timeout"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Timeout","GoType":"int"}
		V3PrivacyProtocol: data["v3_privacy_protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3PrivacyProtocol","GoType":"string"}
		Retries: data["retries"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Retries","GoType":"int"}
		V3SecurityLevel: data["v3_security_level"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3SecurityLevel","GoType":"string"}
		V3Context: data["v3_context"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3Context","GoType":"string"}
		V2Community: data["v2_community"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V2Community","GoType":"string"}
		V3ContextEngineID: data["v3_context_engine_id"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3ContextEngineID","GoType":"string"}
		Version: data["version"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Version","GoType":"int"}
		V3SecurityName: data["v3_security_name"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3SecurityName","GoType":"string"}
		V3AuthenticationProtocol: data["v3_authentication_protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3AuthenticationProtocol","GoType":"string"}
		V3PrivacyPassword: data["v3_privacy_password"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3PrivacyPassword","GoType":"string"}

	}
}

func InterfaceToSNMPCredentialsSlice(data interface{}) []*SNMPCredentials {
	list := data.([]interface{})
	result := MakeSNMPCredentialsSlice()
	for _, item := range list {
		result = append(result, InterfaceToSNMPCredentials(item))
	}
	return result
}

func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}
