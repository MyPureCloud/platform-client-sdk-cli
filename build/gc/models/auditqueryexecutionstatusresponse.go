package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryexecutionstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryexecutionstatusresponseDud struct { 
    


    


    


    


    


    


    

}

// Auditqueryexecutionstatusresponse
type Auditqueryexecutionstatusresponse struct { 
    // Id - Id of the audit query execution request.
    Id string `json:"id"`


    // State - Status of the audit query execution request.
    State string `json:"state"`


    // StartDate - Start date and time of the audit query execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // Interval - Interval for the audit query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // ServiceName - Service name for the audit query.
    ServiceName string `json:"serviceName"`


    // Filters - Filters for the audit query.
    Filters []Auditqueryfilter `json:"filters"`


    // Sort - Sort parameter for the audit query.
    Sort []Auditquerysort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionstatusresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Filters = []Auditqueryfilter{{}} 
    
    
    
     o.Sort = []Auditquerysort{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryexecutionstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryexecutionstatusresponse

    if AuditqueryexecutionstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryexecutionstatusresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        State string `json:"state"`
        
        StartDate time.Time `json:"startDate"`
        
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

