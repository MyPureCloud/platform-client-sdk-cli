package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypecharacteristicmodelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypecharacteristicmodelDud struct { 
    


    


    

}

// Capabilitiesdatatypecharacteristicmodel
type Capabilitiesdatatypecharacteristicmodel struct { 
    // Passthrough
    Passthrough Capabilitiesfuncopoverloadidmodel `json:"passthrough"`


    // Properties
    Properties []string `json:"properties"`


    // Property
    Property string `json:"property"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypecharacteristicmodel) String() string {
    
     o.Properties = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypecharacteristicmodel) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypecharacteristicmodel

    if CapabilitiesdatatypecharacteristicmodelMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypecharacteristicmodelMarshalled = true

    return json.Marshal(&struct {
        
        Passthrough Capabilitiesfuncopoverloadidmodel `json:"passthrough"`
        
        Properties []string `json:"properties"`
        
        Property string `json:"property"`
        *Alias
    }{

        


        
        Properties: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

