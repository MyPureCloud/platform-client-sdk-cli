package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalysummaryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalysummaryresponseDud struct { 
    


    


    

}

// Anomalysummaryresponse
type Anomalysummaryresponse struct { 
    // Interval - Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // TimeZone - The time zone identifier for the organization, in IANA format
    TimeZone string `json:"timeZone"`


    // Summaries - The summarized anomaly count for each time window
    Summaries []Anomalysummary `json:"summaries"`

}

// String returns a JSON representation of the model
func (o *Anomalysummaryresponse) String() string {
    
    
     o.Summaries = []Anomalysummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalysummaryresponse) MarshalJSON() ([]byte, error) {
    type Alias Anomalysummaryresponse

    if AnomalysummaryresponseMarshalled {
        return []byte("{}"), nil
    }
    AnomalysummaryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        TimeZone string `json:"timeZone"`
        
        Summaries []Anomalysummary `json:"summaries"`
        *Alias
    }{

        


        


        
        Summaries: []Anomalysummary{{}},
        

        Alias: (*Alias)(u),
    })
}

