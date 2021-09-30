package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QualityauditqueryexecutionstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QualityauditqueryexecutionstatusresponseDud struct { 
    


    


    


    


    


    

}

// Qualityauditqueryexecutionstatusresponse
type Qualityauditqueryexecutionstatusresponse struct { 
    // Id - Id of the audit query execution request.
    Id string `json:"id"`


    // State - Status of the audit query execution request.
    State string `json:"state"`


    // DateStart - Start date and time of the audit query execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // Interval - Interval for the audit query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Filters - Filters for the audit query.
    Filters []Qualityauditqueryfilter `json:"filters"`


    // Sort - Sort parameter for the audit query.
    Sort []Auditquerysort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Qualityauditqueryexecutionstatusresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Filters = []Qualityauditqueryfilter{{}} 
    
    
    
     o.Sort = []Auditquerysort{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Qualityauditqueryexecutionstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Qualityauditqueryexecutionstatusresponse

    if QualityauditqueryexecutionstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    QualityauditqueryexecutionstatusresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        State string `json:"state"`
        
        DateStart time.Time `json:"dateStart"`
        
        Interval string `json:"interval"`
        
        Filters []Qualityauditqueryfilter `json:"filters"`
        
        Sort []Auditquerysort `json:"sort"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Filters: []Qualityauditqueryfilter{{}},
        

        

        
        Sort: []Auditquerysort{{}},
        

        
        Alias: (*Alias)(u),
    })
}

