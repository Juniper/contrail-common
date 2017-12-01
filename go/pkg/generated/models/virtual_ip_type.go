package models
// VirtualIpType



import "encoding/json"

type VirtualIpType struct {

    Address IpAddressType `json:"address"`
    ProtocolPort int `json:"protocol_port"`
    Status string `json:"status"`
    Protocol LoadbalancerProtocolType `json:"protocol"`
    PersistenceType SessionPersistenceType `json:"persistence_type"`
    ConnectionLimit int `json:"connection_limit"`
    AdminState bool `json:"admin_state"`
    StatusDescription string `json:"status_description"`
    SubnetID UuidStringType `json:"subnet_id"`
    PersistenceCookieName string `json:"persistence_cookie_name"`
}

func (model *VirtualIpType) String() string{
    b, _ := json.Marshal(model)
    return string(b)
}

func MakeVirtualIpType() *VirtualIpType{
    return &VirtualIpType{
    //TODO(nati): Apply default
    PersistenceCookieName: "",
        ConnectionLimit: 0,
        AdminState: false,
        StatusDescription: "",
        SubnetID: MakeUuidStringType(),
        PersistenceType: MakeSessionPersistenceType(),
        Address: MakeIpAddressType(),
        ProtocolPort: 0,
        Status: "",
        Protocol: MakeLoadbalancerProtocolType(),
        
    }
}

func InterfaceToVirtualIpType(iData interface{}) *VirtualIpType {
    data := iData.(map[string]interface{})
    return &VirtualIpType{
    Status: data["status"].(string),
        
        //{"Title":"","Description":"Operating status for this virtual ip.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Status","GoType":"string"}
        Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),
        
        //{"Title":"","Description":"IP protocol string like http, https or tcp.","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LoadbalancerProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"LoadbalancerProtocolType"}
        PersistenceType: InterfaceToSessionPersistenceType(data["persistence_type"]),
        
        //{"Title":"","Description":"Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SessionPersistenceType","CollectionType":"","Column":"","Item":null,"GoName":"PersistenceType","GoType":"SessionPersistenceType"}
        Address: InterfaceToIpAddressType(data["address"]),
        
        //{"Title":"","Description":"IP address automatically allocated by system.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"Address","GoType":"IpAddressType"}
        ProtocolPort: data["protocol_port"].(int),
        
        //{"Title":"","Description":"Layer 4 protocol destination port.","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolPort","GoType":"int"}
        StatusDescription: data["status_description"].(string),
        
        //{"Title":"","Description":"Operating status description this virtual ip.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StatusDescription","GoType":"string"}
        SubnetID: InterfaceToUuidStringType(data["subnet_id"]),
        
        //{"Title":"","Description":"UUID of subnet in which to allocate the Virtual IP.","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UuidStringType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetID","GoType":"UuidStringType"}
        PersistenceCookieName: data["persistence_cookie_name"].(string),
        
        //{"Title":"","Description":"Set this string if the relation of client and server(pool member) need to persist.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"PersistenceCookieName","GoType":"string"}
        ConnectionLimit: data["connection_limit"].(int),
        
        //{"Title":"","Description":"Maximum number of concurrent connections","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ConnectionLimit","GoType":"int"}
        AdminState: data["admin_state"].(bool),
        
        //{"Title":"","Description":"Administrative up or down.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool"}
        
    }
}


func InterfaceToVirtualIpTypeSlice(data interface{}) []*VirtualIpType {
    list := data.([]interface{})
    result := MakeVirtualIpTypeSlice()
    for _, item := range list {
        result = append(result, InterfaceToVirtualIpType(item))
    }
    return result
}

func MakeVirtualIpTypeSlice() []*VirtualIpType {
    return []*VirtualIpType{}
}
