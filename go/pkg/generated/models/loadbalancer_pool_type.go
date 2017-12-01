package models

// LoadbalancerPoolType

import "encoding/json"

type LoadbalancerPoolType struct {
	LoadbalancerMethod    LoadbalancerMethodType   `json:"loadbalancer_method"`
	Status                string                   `json:"status"`
	Protocol              LoadbalancerProtocolType `json:"protocol"`
	SubnetID              UuidStringType           `json:"subnet_id"`
	SessionPersistence    SessionPersistenceType   `json:"session_persistence"`
	AdminState            bool                     `json:"admin_state"`
	PersistenceCookieName string                   `json:"persistence_cookie_name"`
	StatusDescription     string                   `json:"status_description"`
}

func (model *LoadbalancerPoolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeLoadbalancerPoolType() *LoadbalancerPoolType {
	return &LoadbalancerPoolType{
		//TODO(nati): Apply default
		LoadbalancerMethod:    MakeLoadbalancerMethodType(),
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		SessionPersistence:    MakeSessionPersistenceType(),
		AdminState:            false,
		PersistenceCookieName: "",
		StatusDescription:     "",
	}
}

func InterfaceToLoadbalancerPoolType(iData interface{}) *LoadbalancerPoolType {
	data := iData.(map[string]interface{})
	return &LoadbalancerPoolType{
		StatusDescription: data["status_description"].(string),

		//{"Title":"","Description":"Operating status description for this loadbalancer pool.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StatusDescription","GoType":"string"}
		LoadbalancerMethod: InterfaceToLoadbalancerMethodType(data["loadbalancer_method"]),

		//{"Title":"","Description":"Load balancing method ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["ROUND_ROBIN","LEAST_CONNECTIONS","SOURCE_IP"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LoadbalancerMethodType","CollectionType":"","Column":"","Item":null,"GoName":"LoadbalancerMethod","GoType":"LoadbalancerMethodType"}
		Status: data["status"].(string),

		//{"Title":"","Description":"Operating status for this loadbalancer pool.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Status","GoType":"string"}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"Title":"","Description":"IP protocol string like http, https or tcp.","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LoadbalancerProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"LoadbalancerProtocolType"}
		SubnetID: InterfaceToUuidStringType(data["subnet_id"]),

		//{"Title":"","Description":"UUID of the subnet from where the members of the pool are reachable.","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UuidStringType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetID","GoType":"UuidStringType"}
		SessionPersistence: InterfaceToSessionPersistenceType(data["session_persistence"]),

		//{"Title":"","Description":"Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SessionPersistenceType","CollectionType":"","Column":"","Item":null,"GoName":"SessionPersistence","GoType":"SessionPersistenceType"}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"Administrative up or down","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool"}
		PersistenceCookieName: data["persistence_cookie_name"].(string),

		//{"Title":"","Description":"To Be Added","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"PersistenceCookieName","GoType":"string"}

	}
}

func InterfaceToLoadbalancerPoolTypeSlice(data interface{}) []*LoadbalancerPoolType {
	list := data.([]interface{})
	result := MakeLoadbalancerPoolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerPoolType(item))
	}
	return result
}

func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
	return []*LoadbalancerPoolType{}
}
