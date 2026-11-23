package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzdefaultsearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzdefaultsearchresponseDud struct { 
    


    


    

}

// Authzdefaultsearchresponse
type Authzdefaultsearchresponse struct { 
    // Schema - The schema for default search criteria.
    Schema Domainentityref `json:"schema"`


    // ItemType - The itemType for default search criteria.
    ItemType string `json:"itemType"`


    // AttributeName - The attributeName for default search criteria.
    AttributeName string `json:"attributeName"`

}

// String returns a JSON representation of the model
func (o *Authzdefaultsearchresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzdefaultsearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Authzdefaultsearchresponse

    if AuthzdefaultsearchresponseMarshalled {
        return []byte("{}"), nil
    }
    AuthzdefaultsearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Schema Domainentityref `json:"schema"`
        
        ItemType string `json:"itemType"`
        
        AttributeName string `json:"attributeName"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

