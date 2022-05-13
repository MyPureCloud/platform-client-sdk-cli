package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntradayperformancepredictiondataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntradayperformancepredictiondataDud struct { 
    


    


    

}

// Intradayperformancepredictiondata
type Intradayperformancepredictiondata struct { 
    // ServiceLevelPercent - Percentage of interactions that meets service level target as defined in the matching service goal templates
    ServiceLevelPercent float64 `json:"serviceLevelPercent"`


    // AverageSpeedOfAnswerSeconds - Predicted average time in seconds it takes to answer an interaction once the interaction becomes available to be routed
    AverageSpeedOfAnswerSeconds float64 `json:"averageSpeedOfAnswerSeconds"`


    // OccupancyPercent - Percentage of on-queue time for all agents in this group that are occupied handling interactions
    OccupancyPercent float64 `json:"occupancyPercent"`

}

// String returns a JSON representation of the model
func (o *Intradayperformancepredictiondata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intradayperformancepredictiondata) MarshalJSON() ([]byte, error) {
    type Alias Intradayperformancepredictiondata

    if IntradayperformancepredictiondataMarshalled {
        return []byte("{}"), nil
    }
    IntradayperformancepredictiondataMarshalled = true

    return json.Marshal(&struct {
        
        ServiceLevelPercent float64 `json:"serviceLevelPercent"`
        
        AverageSpeedOfAnswerSeconds float64 `json:"averageSpeedOfAnswerSeconds"`
        
        OccupancyPercent float64 `json:"occupancyPercent"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

