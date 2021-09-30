package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastservicegoaltemplateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastservicegoaltemplateresponseDud struct { 
    


    


    

}

// Forecastservicegoaltemplateresponse
type Forecastservicegoaltemplateresponse struct { 
    // ServiceLevel - The service level goal for this forecast
    ServiceLevel Forecastservicelevelresponse `json:"serviceLevel"`


    // AverageSpeedOfAnswer - The average speed of answer goal for this forecast
    AverageSpeedOfAnswer Forecastaveragespeedofanswerresponse `json:"averageSpeedOfAnswer"`


    // AbandonRate - The abandon rate goal for this forecast
    AbandonRate Forecastabandonrateresponse `json:"abandonRate"`

}

// String returns a JSON representation of the model
func (o *Forecastservicegoaltemplateresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastservicegoaltemplateresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastservicegoaltemplateresponse

    if ForecastservicegoaltemplateresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastservicegoaltemplateresponseMarshalled = true

    return json.Marshal(&struct { 
        ServiceLevel Forecastservicelevelresponse `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Forecastaveragespeedofanswerresponse `json:"averageSpeedOfAnswer"`
        
        AbandonRate Forecastabandonrateresponse `json:"abandonRate"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

