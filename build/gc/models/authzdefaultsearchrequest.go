package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzdefaultsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzdefaultsearchrequestDud struct { 
    


    

}

// Authzdefaultsearchrequest
type Authzdefaultsearchrequest struct { 
    // SchemaName - The schema name for the custom user attribute
    SchemaName string `json:"schemaName"`


    // AttributeName - The attribute name for the custom user attribute
    AttributeName string `json:"attributeName"`

}

// String returns a JSON representation of the model
func (o *Authzdefaultsearchrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzdefaultsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Authzdefaultsearchrequest

    if AuthzdefaultsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    AuthzdefaultsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SchemaName string `json:"schemaName"`
        
        AttributeName string `json:"attributeName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

