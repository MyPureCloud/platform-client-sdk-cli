package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastimportuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastimportuploadrequestDud struct { 
    


    

}

// Capacityplanforecastimportuploadrequest
type Capacityplanforecastimportuploadrequest struct { 
    // HourlyForecast - Content length request details for the hourly forecast to be imported for the capacity plan
    HourlyForecast Uploadurlrequestbody `json:"hourlyForecast"`


    // DailyForecast - Content length request details for the daily forecast to be imported for the capacity plan
    DailyForecast Uploadurlrequestbody `json:"dailyForecast"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastimportuploadrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastimportuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastimportuploadrequest

    if CapacityplanforecastimportuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastimportuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        HourlyForecast Uploadurlrequestbody `json:"hourlyForecast"`
        
        DailyForecast Uploadurlrequestbody `json:"dailyForecast"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

