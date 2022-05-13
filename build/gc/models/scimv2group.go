package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2groupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2groupDud struct { 
    Id string `json:"id"`


    Schemas []string `json:"schemas"`


    DisplayName string `json:"displayName"`


    


    


    Meta Scimmetadata `json:"meta"`

}

// Scimv2group - Defines a SCIM group.
type Scimv2group struct { 
    


    


    


    // ExternalId - The external ID of the group. Set by the provisioning client. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readWrite\".
    ExternalId string `json:"externalId"`


    // Members - The list of members in the group.
    Members []Scimv2memberreference `json:"members"`


    

}

// String returns a JSON representation of the model
func (o *Scimv2group) String() string {
    
     o.Members = []Scimv2memberreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2group) MarshalJSON() ([]byte, error) {
    type Alias Scimv2group

    if Scimv2groupMarshalled {
        return []byte("{}"), nil
    }
    Scimv2groupMarshalled = true

    return json.Marshal(&struct {
        
        ExternalId string `json:"externalId"`
        
        Members []Scimv2memberreference `json:"members"`
        *Alias
    }{

        


        


        


        


        
        Members: []Scimv2memberreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

