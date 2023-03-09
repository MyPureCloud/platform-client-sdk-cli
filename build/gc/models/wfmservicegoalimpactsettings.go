package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmservicegoalimpactsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmservicegoalimpactsettingsDud struct { 
    


    


    

}

// Wfmservicegoalimpactsettings
type Wfmservicegoalimpactsettings struct { 
    // ServiceLevel - Allowed service level percent increase and decrease
    ServiceLevel Wfmservicegoalimpact `json:"serviceLevel"`


    // AverageSpeedOfAnswer - Allowed average speed of answer percent increase and decrease
    AverageSpeedOfAnswer Wfmservicegoalimpact `json:"averageSpeedOfAnswer"`


    // AbandonRate - Allowed abandon rate percent increase and decrease
    AbandonRate Wfmservicegoalimpact `json:"abandonRate"`

}

// String returns a JSON representation of the model
func (o *Wfmservicegoalimpactsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmservicegoalimpactsettings) MarshalJSON() ([]byte, error) {
    type Alias Wfmservicegoalimpactsettings

    if WfmservicegoalimpactsettingsMarshalled {
        return []byte("{}"), nil
    }
    WfmservicegoalimpactsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ServiceLevel Wfmservicegoalimpact `json:"serviceLevel"`
        
        AverageSpeedOfAnswer Wfmservicegoalimpact `json:"averageSpeedOfAnswer"`
        
        AbandonRate Wfmservicegoalimpact `json:"abandonRate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

