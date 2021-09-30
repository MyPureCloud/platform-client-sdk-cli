package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingoptionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingoptionsrequestDud struct { 
    

}

// Schedulingoptionsrequest
type Schedulingoptionsrequest struct { 
    // NoForecastOptions - Schedule generation options to apply if no forecast is supplied
    NoForecastOptions Schedulingnoforecastoptionsrequest `json:"noForecastOptions"`

}

// String returns a JSON representation of the model
func (o *Schedulingoptionsrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingoptionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Schedulingoptionsrequest

    if SchedulingoptionsrequestMarshalled {
        return []byte("{}"), nil
    }
    SchedulingoptionsrequestMarshalled = true

    return json.Marshal(&struct { 
        NoForecastOptions Schedulingnoforecastoptionsrequest `json:"noForecastOptions"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

