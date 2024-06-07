package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditrealtimerelatedrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditrealtimerelatedrequestDud struct { 
    


    


    

}

// Auditrealtimerelatedrequest
type Auditrealtimerelatedrequest struct { 
    // AuditId - The id of the audit of which related audits will be retrieved.
    AuditId string `json:"auditId"`


    // TrustorOrgId - The id of the trustor org to which the audit belongs. Used when searching for audits performed by a trustee user within a trustor org.
    TrustorOrgId string `json:"trustorOrgId"`


    // Sort - Sort parameter for the query.
    Sort []Auditquerysort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Auditrealtimerelatedrequest) String() string {
    
    
     o.Sort = []Auditquerysort{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditrealtimerelatedrequest) MarshalJSON() ([]byte, error) {
    type Alias Auditrealtimerelatedrequest

    if AuditrealtimerelatedrequestMarshalled {
        return []byte("{}"), nil
    }
    AuditrealtimerelatedrequestMarshalled = true

    return json.Marshal(&struct {
        
        AuditId string `json:"auditId"`
        
        TrustorOrgId string `json:"trustorOrgId"`
        
        Sort []Auditquerysort `json:"sort"`
        *Alias
    }{

        


        


        
        Sort: []Auditquerysort{{}},
        

        Alias: (*Alias)(u),
    })
}

