package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastservicelevelresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastservicelevelresponseDud struct { 
    


    

}

// Forecastservicelevelresponse
type Forecastservicelevelresponse struct { 
    // Percent - The percent of calls to answer in the number of seconds defined
    Percent int `json:"percent"`


    // Seconds - The number of seconds to define for the percent of calls to be answered
    Seconds int `json:"seconds"`

}

// String returns a JSON representation of the model
func (o *Forecastservicelevelresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastservicelevelresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastservicelevelresponse

    if ForecastservicelevelresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastservicelevelresponseMarshalled = true

    return json.Marshal(&struct { 
        Percent int `json:"percent"`
        
        Seconds int `json:"seconds"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

