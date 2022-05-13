package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Createintegrationrequest - Details for an Integration
type Createintegrationrequest struct { 
    


    // Name - The name of the integration, used to distinguish this integration from others of the same type.
    Name string `json:"name"`


    // IntegrationType - Type of the integration to create.
    IntegrationType Integrationtype `json:"integrationType"`


    

}

// String returns a JSON representation of the model
func (o *Createintegrationrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Createintegrationrequest

    if CreateintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        IntegrationType Integrationtype `json:"integrationType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

