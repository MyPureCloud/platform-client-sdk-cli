package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalorganizationtrustorlinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalorganizationtrustorlinkDud struct { 
    Id string `json:"id"`


    


    


    


    ExternalOrganizationUri string `json:"externalOrganizationUri"`


    SelfUri string `json:"selfUri"`

}

// Externalorganizationtrustorlink
type Externalorganizationtrustorlink struct { 
    


    // ExternalOrganizationId - The id of a PureCloud External Organization entity in the External Contacts system that will be used to represent the trustor org
    ExternalOrganizationId string `json:"externalOrganizationId"`


    // TrustorOrgId - The id of a PureCloud organization that has granted trust to this PureCloud organization
    TrustorOrgId string `json:"trustorOrgId"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    


    

}

// String returns a JSON representation of the model
func (o *Externalorganizationtrustorlink) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalorganizationtrustorlink) MarshalJSON() ([]byte, error) {
    type Alias Externalorganizationtrustorlink

    if ExternalorganizationtrustorlinkMarshalled {
        return []byte("{}"), nil
    }
    ExternalorganizationtrustorlinkMarshalled = true

    return json.Marshal(&struct {
        
        ExternalOrganizationId string `json:"externalOrganizationId"`
        
        TrustorOrgId string `json:"trustorOrgId"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

