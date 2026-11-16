package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastimportuploadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastimportuploadresponseDud struct { 
    


    

}

// Capacityplanforecastimportuploadresponse
type Capacityplanforecastimportuploadresponse struct { 
    // HourlyForecast - The upload URL, upload key, and required headers for the hourly forecast to be imported for the capacity plan
    HourlyForecast Capacityplanforecastimportupload `json:"hourlyForecast"`


    // DailyForecast - The upload URL, upload key, and required headers for the daily forecast to be imported for the capacity plan
    DailyForecast Capacityplanforecastimportupload `json:"dailyForecast"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastimportuploadresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastimportuploadresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastimportuploadresponse

    if CapacityplanforecastimportuploadresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastimportuploadresponseMarshalled = true

    return json.Marshal(&struct {
        
        HourlyForecast Capacityplanforecastimportupload `json:"hourlyForecast"`
        
        DailyForecast Capacityplanforecastimportupload `json:"dailyForecast"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

