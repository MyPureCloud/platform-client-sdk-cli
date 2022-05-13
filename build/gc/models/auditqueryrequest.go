package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryrequestDud struct { 
    


    


    


    

}

// Auditqueryrequest
type Auditqueryrequest struct { 
    // Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // ServiceName - Name of the service to query audits for.
    ServiceName string `json:"serviceName"`


    // Filters - Additional filters for the query.
    Filters []Auditqueryfilter `json:"filters"`


    // Sort - Sort parameter for the query.
    Sort []Auditquerysort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Auditqueryrequest) String() string {
    
    
     o.Filters = []Auditqueryfilter{{}} 
     o.Sort = []Auditquerysort{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryrequest

    if AuditqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        ServiceName string `json:"serviceName"`
        
        Filters []Auditqueryfilter `json:"filters"`
        
        Sort []Auditquerysort `json:"sort"`
        *Alias
    }{

        


        


        
        Filters: []Auditqueryfilter{{}},
        


        
        Sort: []Auditquerysort{{}},
        

        Alias: (*Alias)(u),
    })
}

