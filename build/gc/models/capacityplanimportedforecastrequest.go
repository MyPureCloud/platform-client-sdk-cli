package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanimportedforecastrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanimportedforecastrequestDud struct { 
    


    

}

// Capacityplanimportedforecastrequest
type Capacityplanimportedforecastrequest struct { 
    // HourlyForecastUploadKey - The uploadKey returned in the hourlyForecast field of the capacity plan forecast upload URL response
    HourlyForecastUploadKey string `json:"hourlyForecastUploadKey"`


    // DailyForecastUploadKey - The uploadKey returned in the dailyForecast field of the capacity plan forecast upload URL response
    DailyForecastUploadKey string `json:"dailyForecastUploadKey"`

}

// String returns a JSON representation of the model
func (o *Capacityplanimportedforecastrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanimportedforecastrequest) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanimportedforecastrequest

    if CapacityplanimportedforecastrequestMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanimportedforecastrequestMarshalled = true

    return json.Marshal(&struct {
        
        HourlyForecastUploadKey string `json:"hourlyForecastUploadKey"`
        
        DailyForecastUploadKey string `json:"dailyForecastUploadKey"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

