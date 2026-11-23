package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypeDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Capabilitiesdatatype
type Capabilitiesdatatype struct { 
    // FlowTypes
    FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`


    // AllowCollection
    AllowCollection bool `json:"allowCollection"`


    // Characteristics
    Characteristics Capabilitiesdatatypecharacteristicsmodel `json:"characteristics"`


    // Direction
    Direction Capabilitiesdatatypedirectionbyobject `json:"direction"`


    // Facets
    Facets Capabilitiesdatatypefacetsmodel `json:"facets"`


    // Id
    Id string `json:"id"`


    // IsAbstract
    IsAbstract bool `json:"isAbstract"`


    // IsAnAnyType
    IsAnAnyType Capabilitiesavailability `json:"isAnAnyType"`


    // IsComplex
    IsComplex bool `json:"isComplex"`


    // IsEnum
    IsEnum bool `json:"isEnum"`


    // IsEnumContainer
    IsEnumContainer bool `json:"isEnumContainer"`


    // Name
    Name string `json:"name"`


    // Properties
    Properties []Capabilitiesdatatypeproperty `json:"properties"`


    // UserCreatable
    UserCreatable Capabilitiesavailability `json:"userCreatable"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatype) String() string {
     o.FlowTypes = map[string]Capabilitiesavailability{"": {}} 
    
    
    
    
    
    
    
    
    
    
    
     o.Properties = []Capabilitiesdatatypeproperty{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatype) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatype

    if CapabilitiesdatatypeMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypeMarshalled = true

    return json.Marshal(&struct {
        
        FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`
        
        AllowCollection bool `json:"allowCollection"`
        
        Characteristics Capabilitiesdatatypecharacteristicsmodel `json:"characteristics"`
        
        Direction Capabilitiesdatatypedirectionbyobject `json:"direction"`
        
        Facets Capabilitiesdatatypefacetsmodel `json:"facets"`
        
        Id string `json:"id"`
        
        IsAbstract bool `json:"isAbstract"`
        
        IsAnAnyType Capabilitiesavailability `json:"isAnAnyType"`
        
        IsComplex bool `json:"isComplex"`
        
        IsEnum bool `json:"isEnum"`
        
        IsEnumContainer bool `json:"isEnumContainer"`
        
        Name string `json:"name"`
        
        Properties []Capabilitiesdatatypeproperty `json:"properties"`
        
        UserCreatable Capabilitiesavailability `json:"userCreatable"`
        *Alias
    }{

        
        FlowTypes: map[string]Capabilitiesavailability{"": {}},
        


        


        


        


        


        


        


        


        


        


        


        


        
        Properties: []Capabilitiesdatatypeproperty{{}},
        


        

        Alias: (*Alias)(u),
    })
}

