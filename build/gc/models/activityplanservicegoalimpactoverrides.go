package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanservicegoalimpactoverridesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanservicegoalimpactoverridesDud struct { 
    


    


    

}

// Activityplanservicegoalimpactoverrides
type Activityplanservicegoalimpactoverrides struct { 
    // AbandonRate - Abandon rate service goal override for the associated activity plan
    AbandonRate Activityplanabandonrateimpactoverride `json:"abandonRate"`


    // ServiceLevel - Service level goal override for the associated activity plan
    ServiceLevel Activityplanservicelevelimpactoverride `json:"serviceLevel"`


    // AverageSpeedOfAnswer - Average speed of answer service goal override for the associated activity plan
    AverageSpeedOfAnswer Activityplanasaimpactoverride `json:"averageSpeedOfAnswer"`

}

// String returns a JSON representation of the model
func (o *Activityplanservicegoalimpactoverrides) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanservicegoalimpactoverrides) MarshalJSON() ([]byte, error) {
    type Alias Activityplanservicegoalimpactoverrides

    if ActivityplanservicegoalimpactoverridesMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanservicegoalimpactoverridesMarshalled = true

    return json.Marshal(&struct {
        
        AbandonRate Activityplanabandonrateimpactoverride `json:"abandonRate"`
        
        ServiceLevel Activityplanservicelevelimpactoverride `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Activityplanasaimpactoverride `json:"averageSpeedOfAnswer"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

