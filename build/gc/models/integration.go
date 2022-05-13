package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    IntegrationType Integrationtype `json:"integrationType"`


    Notes string `json:"notes"`


    


    Config Integrationconfigurationinfo `json:"config"`


    ReportedState Integrationstatusinfo `json:"reportedState"`


    Attributes map[string]string `json:"attributes"`


    SelfUri string `json:"selfUri"`

}

// Integration - Details for an Integration
type Integration struct { 
    


    


    


    


    // IntendedState - Configured state of the integration.
    IntendedState string `json:"intendedState"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Integration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integration) MarshalJSON() ([]byte, error) {
    type Alias Integration

    if IntegrationMarshalled {
        return []byte("{}"), nil
    }
    IntegrationMarshalled = true

    return json.Marshal(&struct {
        
        IntendedState string `json:"intendedState"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

