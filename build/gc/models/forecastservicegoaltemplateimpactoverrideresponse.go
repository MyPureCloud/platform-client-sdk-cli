package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastservicegoaltemplateimpactoverrideresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastservicegoaltemplateimpactoverrideresponseDud struct { 
    


    


    

}

// Forecastservicegoaltemplateimpactoverrideresponse
type Forecastservicegoaltemplateimpactoverrideresponse struct { 
    // ServiceLevel - Allowed service level percent increase and decrease; undefined if the goal is not enabled
    ServiceLevel Wfmservicegoalimpact `json:"serviceLevel"`


    // AverageSpeedOfAnswer - Allowed average speed of answer percent increase and decrease; undefined if the goal is not enabled
    AverageSpeedOfAnswer Wfmservicegoalimpact `json:"averageSpeedOfAnswer"`


    // AbandonRate - Allowed abandon rate percent increase and decrease; undefined if the goal is not enabled
    AbandonRate Wfmservicegoalimpact `json:"abandonRate"`

}

// String returns a JSON representation of the model
func (o *Forecastservicegoaltemplateimpactoverrideresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastservicegoaltemplateimpactoverrideresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastservicegoaltemplateimpactoverrideresponse

    if ForecastservicegoaltemplateimpactoverrideresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastservicegoaltemplateimpactoverrideresponseMarshalled = true

    return json.Marshal(&struct {
        
        ServiceLevel Wfmservicegoalimpact `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Wfmservicegoalimpact `json:"averageSpeedOfAnswer"`
        
        AbandonRate Wfmservicegoalimpact `json:"abandonRate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

