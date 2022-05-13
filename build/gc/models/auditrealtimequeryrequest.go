package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditrealtimequeryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditrealtimequeryrequestDud struct { 
    


    


    


    


    


    

}

// Auditrealtimequeryrequest
type Auditrealtimequeryrequest struct { 
    // Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // ServiceName - Name of the service to query audits for.
    ServiceName string `json:"serviceName"`


    // Filters - Additional filters for the query.
    Filters []Auditqueryfilter `json:"filters"`


    // Sort - Sort parameter for the query.
    Sort []Auditquerysort `json:"sort"`


    // PageNumber - Page number
    PageNumber int `json:"pageNumber"`


    // PageSize - Page size
    PageSize int `json:"pageSize"`

}

// String returns a JSON representation of the model
func (o *Auditrealtimequeryrequest) String() string {
    
    
     o.Filters = []Auditqueryfilter{{}} 
     o.Sort = []Auditquerysort{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditrealtimequeryrequest) MarshalJSON() ([]byte, error) {
    type Alias Auditrealtimequeryrequest

    if AuditrealtimequeryrequestMarshalled {
        return []byte("{}"), nil
    }
    AuditrealtimequeryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        ServiceName string `json:"serviceName"`
        
        Filters []Auditqueryfilter `json:"filters"`
        
        Sort []Auditquerysort `json:"sort"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        *Alias
    }{

        


        


        
        Filters: []Auditqueryfilter{{}},
        


        
        Sort: []Auditquerysort{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

