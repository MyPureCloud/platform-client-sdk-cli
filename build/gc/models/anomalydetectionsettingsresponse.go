package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionsettingsresponseDud struct { 
    


    


    

}

// Anomalydetectionsettingsresponse
type Anomalydetectionsettingsresponse struct { 
    // TimeZone - The time zone to be used for both display and time series data evaluation of anomalies, in IANA format
    TimeZone string `json:"timeZone"`


    // StartDayOfWeek - The start day of the week that will be used for weekly aggregations
    StartDayOfWeek string `json:"startDayOfWeek"`


    // DateModified - Date and time when the anomaly detection configuration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectionsettingsresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionsettingsresponse

    if AnomalydetectionsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        TimeZone string `json:"timeZone"`
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

