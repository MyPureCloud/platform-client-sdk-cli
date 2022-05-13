package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QmauditqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QmauditqueryrequestDud struct { 
    


    


    

}

// Qmauditqueryrequest
type Qmauditqueryrequest struct { 
    // Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Filters - List of filters for the query.
    Filters []Qualityauditqueryfilter `json:"filters"`


    // Sort - Sort parameter for the query.
    Sort []Auditquerysort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Qmauditqueryrequest) String() string {
    
     o.Filters = []Qualityauditqueryfilter{{}} 
     o.Sort = []Auditquerysort{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Qmauditqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Qmauditqueryrequest

    if QmauditqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    QmauditqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
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

