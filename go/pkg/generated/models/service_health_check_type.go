package models

// ServiceHealthCheckType

import "encoding/json"

type ServiceHealthCheckType struct {
	HealthCheckType HealthCheckType         `json:"health_check_type"`
	HTTPMethod      string                  `json:"http_method"`
	Timeout         int                     `json:"timeout"`
	MonitorType     HealthCheckProtocolType `json:"monitor_type"`
	DelayUsecs      int                     `json:"delayUsecs"`
	TimeoutUsecs    int                     `json:"timeoutUsecs"`
	Delay           int                     `json:"delay"`
	URLPath         string                  `json:"url_path"`
	Enabled         bool                    `json:"enabled"`
	ExpectedCodes   string                  `json:"expected_codes"`
	MaxRetries      int                     `json:"max_retries"`
}

func (model *ServiceHealthCheckType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeServiceHealthCheckType() *ServiceHealthCheckType {
	return &ServiceHealthCheckType{
		//TODO(nati): Apply default
		Delay:           0,
		HealthCheckType: MakeHealthCheckType(),
		HTTPMethod:      "",
		Timeout:         0,
		MonitorType:     MakeHealthCheckProtocolType(),
		DelayUsecs:      0,
		TimeoutUsecs:    0,
		MaxRetries:      0,
		URLPath:         "",
		Enabled:         false,
		ExpectedCodes:   "",
	}
}

func InterfaceToServiceHealthCheckType(iData interface{}) *ServiceHealthCheckType {
	data := iData.(map[string]interface{})
	return &ServiceHealthCheckType{
		Timeout: data["timeout"].(int),

		//{"Title":"","Description":"Time in seconds to wait for response","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Timeout","GoType":"int"}
		MonitorType: InterfaceToHealthCheckProtocolType(data["monitor_type"]),

		//{"Title":"","Description":"Protocol used to monitor health, currently only HTTP, ICMP(ping), and BFD are supported","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["PING","HTTP","BFD"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/HealthCheckProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"MonitorType","GoType":"HealthCheckProtocolType"}
		DelayUsecs: data["delayUsecs"].(int),

		//{"Title":"","Description":"Time in micro seconds at which health check is repeated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DelayUsecs","GoType":"int"}
		TimeoutUsecs: data["timeoutUsecs"].(int),

		//{"Title":"","Description":"Time in micro seconds to wait for response","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TimeoutUsecs","GoType":"int"}
		Delay: data["delay"].(int),

		//{"Title":"","Description":"Time in seconds at which health check is repeated","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Delay","GoType":"int"}
		HealthCheckType: InterfaceToHealthCheckType(data["health_check_type"]),

		//{"Title":"","Description":"Health check type, currently only link-local, end-to-end and segment are supported","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["link-local","end-to-end","segment"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/HealthCheckType","CollectionType":"","Column":"","Item":null,"GoName":"HealthCheckType","GoType":"HealthCheckType"}
		HTTPMethod: data["http_method"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, type of http method used like GET, PUT, POST etc","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"HTTPMethod","GoType":"string"}
		Enabled: data["enabled"].(bool),

		//{"Title":"","Description":"Administratively enable or disable this health check.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Enabled","GoType":"bool"}
		ExpectedCodes: data["expected_codes"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, expected return code for HTTP operations like 200 ok.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ExpectedCodes","GoType":"string"}
		MaxRetries: data["max_retries"].(int),

		//{"Title":"","Description":"Number of failures before declaring health bad","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MaxRetries","GoType":"int"}
		URLPath: data["url_path"].(string),

		//{"Title":"","Description":"In case monitor protocol is HTTP, URL to be used. In case of ICMP, ip address","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"URLPath","GoType":"string"}

	}
}

func InterfaceToServiceHealthCheckTypeSlice(data interface{}) []*ServiceHealthCheckType {
	list := data.([]interface{})
	result := MakeServiceHealthCheckTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceHealthCheckType(item))
	}
	return result
}

func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}
