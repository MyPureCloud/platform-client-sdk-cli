package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionsettingsrequestDud struct { 
    


    

}

// Anomalydetectionsettingsrequest
type Anomalydetectionsettingsrequest struct { 
    // TimeZone - The time zone to be used for both display and time series data evaluation of anomalies, in IANA format
    TimeZone string `json:"timeZone"`


    // StartDayOfWeek - The start day of the week that will be used for weekly aggregations
    StartDayOfWeek string `json:"startDayOfWeek"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectionsettingsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionsettingsrequest

    if AnomalydetectionsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        TimeZone string `json:"timeZone"`
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

