package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesactionDud struct { 
    


    


    


    Limits Capabilitiesactionlimitinfo `json:"limits"`

}

// Capabilitiesaction
type Capabilitiesaction struct { 
    // FlowTypes
    FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`


    // CategoryInfo
    CategoryInfo Capabilitiescategoryinfo `json:"categoryInfo"`


    // FlowTypesDeprecated
    FlowTypesDeprecated map[string]Capabilitiesavailability `json:"flowTypesDeprecated"`


    

}

// String returns a JSON representation of the model
func (o *Capabilitiesaction) String() string {
     o.FlowTypes = map[string]Capabilitiesavailability{"": {}} 
    
     o.FlowTypesDeprecated = map[string]Capabilitiesavailability{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesaction) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesaction

    if CapabilitiesactionMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesactionMarshalled = true

    return json.Marshal(&struct {
        
        FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`
        
        CategoryInfo Capabilitiescategoryinfo `json:"categoryInfo"`
        
        FlowTypesDeprecated map[string]Capabilitiesavailability `json:"flowTypesDeprecated"`
        *Alias
    }{

        
        FlowTypes: map[string]Capabilitiesavailability{"": {}},
        


        


        
        FlowTypesDeprecated: map[string]Capabilitiesavailability{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

