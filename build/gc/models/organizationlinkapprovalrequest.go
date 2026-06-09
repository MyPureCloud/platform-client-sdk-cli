package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationlinkapprovalrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationlinkapprovalrequestDud struct { 
    

}

// Organizationlinkapprovalrequest
type Organizationlinkapprovalrequest struct { 
    // Approval - Value for approving or rejecting an organization link, true is approved, false is rejected
    Approval bool `json:"approval"`

}

// String returns a JSON representation of the model
func (o *Organizationlinkapprovalrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationlinkapprovalrequest) MarshalJSON() ([]byte, error) {
    type Alias Organizationlinkapprovalrequest

    if OrganizationlinkapprovalrequestMarshalled {
        return []byte("{}"), nil
    }
    OrganizationlinkapprovalrequestMarshalled = true

    return json.Marshal(&struct {
        
        Approval bool `json:"approval"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

