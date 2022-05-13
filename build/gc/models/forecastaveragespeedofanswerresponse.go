package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastaveragespeedofanswerresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastaveragespeedofanswerresponseDud struct { 
    

}

// Forecastaveragespeedofanswerresponse
type Forecastaveragespeedofanswerresponse struct { 
    // Seconds - the average speed of answer goal in seconds
    Seconds int `json:"seconds"`

}

// String returns a JSON representation of the model
func (o *Forecastaveragespeedofanswerresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastaveragespeedofanswerresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastaveragespeedofanswerresponse

    if ForecastaveragespeedofanswerresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastaveragespeedofanswerresponseMarshalled = true

    return json.Marshal(&struct {
        
        Seconds int `json:"seconds"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

