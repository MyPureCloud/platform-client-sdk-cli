package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeinterfaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeinterfaceDud struct { 
    


    


    


    


    


    


    


    

}

// Edgeinterface
type Edgeinterface struct { 
    // VarType
    VarType string `json:"type"`


    // IpAddress
    IpAddress string `json:"ipAddress"`


    // Name
    Name string `json:"name"`


    // MacAddress
    MacAddress string `json:"macAddress"`


    // IfName
    IfName string `json:"ifName"`


    // Endpoints
    Endpoints []Domainentityref `json:"endpoints"`


    // LineTypes
    LineTypes []string `json:"lineTypes"`


    // AddressFamilyId
    AddressFamilyId string `json:"addressFamilyId"`

}

// String returns a JSON representation of the model
func (o *Edgeinterface) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Endpoints = []Domainentityref{{}} 
    
    
    
     o.LineTypes = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeinterface) MarshalJSON() ([]byte, error) {
    type Alias Edgeinterface

    if EdgeinterfaceMarshalled {
        return []byte("{}"), nil
    }
    EdgeinterfaceMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        IpAddress string `json:"ipAddress"`
        
        Name string `json:"name"`
        
        MacAddress string `json:"macAddress"`
        
        IfName string `json:"ifName"`
        
        Endpoints []Domainentityref `json:"endpoints"`
        
        LineTypes []string `json:"lineTypes"`
        
        AddressFamilyId string `json:"addressFamilyId"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Endpoints: []Domainentityref{{}},
        

        

        
        LineTypes: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

