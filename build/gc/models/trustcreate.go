package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustcreateDud struct { 
    


    


    


    


    

}

// Trustcreate
type Trustcreate struct { 
    // PairingId - The pairing Id created by the trustee. This is required to prove that the trustee agrees to the relationship.  Not required when creating a default pairing with Customer Care.
    PairingId string `json:"pairingId"`


    // Enabled - If disabled no trustee user will have access, even if they were previously added.
    Enabled bool `json:"enabled"`


    // Users - The list of users and their roles to which access will be granted. The users are from the trustee and the roles are from the trustor. If no users are specified, at least one group is required.
    Users []Trustmembercreate `json:"users"`


    // Groups - The list of groups and their roles to which access will be granted. The groups are from the trustee and the roles are from the trustor. If no groups are specified, at least one user is required.
    Groups []Trustmembercreate `json:"groups"`


    // DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpired time.Time `json:"dateExpired"`

}

// String returns a JSON representation of the model
func (o *Trustcreate) String() string {
    
    
     o.Users = []Trustmembercreate{{}} 
     o.Groups = []Trustmembercreate{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustcreate) MarshalJSON() ([]byte, error) {
    type Alias Trustcreate

    if TrustcreateMarshalled {
        return []byte("{}"), nil
    }
    TrustcreateMarshalled = true

    return json.Marshal(&struct {
        
        PairingId string `json:"pairingId"`
        
        Enabled bool `json:"enabled"`
        
        Users []Trustmembercreate `json:"users"`
        
        Groups []Trustmembercreate `json:"groups"`
        
        DateExpired time.Time `json:"dateExpired"`
        *Alias
    }{

        


        


        
        Users: []Trustmembercreate{{}},
        


        
        Groups: []Trustmembercreate{{}},
        


        

        Alias: (*Alias)(u),
    })
}

