package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalyfilterrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalyfilterrequestDud struct { 
    


    

}

// Anomalyfilterrequest
type Anomalyfilterrequest struct { 
    // Interval - Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Filters - Expresses boolean logical predicates as well as dimensional filters
    Filters Anomalyqueryfilters `json:"filters"`

}

// String returns a JSON representation of the model
func (o *Anomalyfilterrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalyfilterrequest) MarshalJSON() ([]byte, error) {
    type Alias Anomalyfilterrequest

    if AnomalyfilterrequestMarshalled {
        return []byte("{}"), nil
    }
    AnomalyfilterrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Filters Anomalyqueryfilters `json:"filters"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

