package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastabandonrateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastabandonrateresponseDud struct { 
    

}

// Forecastabandonrateresponse
type Forecastabandonrateresponse struct { 
    // Percent - The target percent abandon rate goal
    Percent int `json:"percent"`

}

// String returns a JSON representation of the model
func (o *Forecastabandonrateresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastabandonrateresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastabandonrateresponse

    if ForecastabandonrateresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastabandonrateresponseMarshalled = true

    return json.Marshal(&struct { 
        Percent int `json:"percent"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

