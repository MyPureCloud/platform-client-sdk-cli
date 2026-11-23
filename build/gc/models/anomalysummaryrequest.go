package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalysummaryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalysummaryrequestDud struct { 
    


    


    

}

// Anomalysummaryrequest
type Anomalysummaryrequest struct { 
    // Interval - Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Filters - Expresses boolean logical predicates as well as dimensional filters
    Filters Anomalyqueryfilters `json:"filters"`


    // Settings - The summary settings which include the granularity for time window grouping
    Settings Anomalysummarysettings `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Anomalysummaryrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalysummaryrequest) MarshalJSON() ([]byte, error) {
    type Alias Anomalysummaryrequest

    if AnomalysummaryrequestMarshalled {
        return []byte("{}"), nil
    }
    AnomalysummaryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Filters Anomalyqueryfilters `json:"filters"`
        
        Settings Anomalysummarysettings `json:"settings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

