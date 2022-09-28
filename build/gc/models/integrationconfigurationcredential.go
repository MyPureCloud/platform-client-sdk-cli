package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationconfigurationcredentialMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationconfigurationcredentialDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Integrationconfigurationcredential - Configuration credential for the integration
type Integrationconfigurationcredential struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Integrationconfigurationcredential) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationconfigurationcredential) MarshalJSON() ([]byte, error) {
    type Alias Integrationconfigurationcredential

    if IntegrationconfigurationcredentialMarshalled {
        return []byte("{}"), nil
    }
    IntegrationconfigurationcredentialMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

