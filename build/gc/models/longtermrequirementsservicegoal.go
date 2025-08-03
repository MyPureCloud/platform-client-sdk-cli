package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermrequirementsservicegoalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermrequirementsservicegoalDud struct { 
    


    


    


    


    


    


    

}

// Longtermrequirementsservicegoal
type Longtermrequirementsservicegoal struct { 
    // UseAsaTarget - Toggle for target average speed of answer from service goals
    UseAsaTarget bool `json:"useAsaTarget"`


    // TargetAsaSec - Value of target average speed of answer from service goals
    TargetAsaSec float64 `json:"targetAsaSec"`


    // UseServiceLevelTarget - Toggle for service level objective from service goals
    UseServiceLevelTarget bool `json:"useServiceLevelTarget"`


    // ServiceLevelObjectiveSeconds - Value of service level objective seconds from service goals
    ServiceLevelObjectiveSeconds float64 `json:"serviceLevelObjectiveSeconds"`


    // ServiceLevelGoalPercent - Value of service level objective percent from service goals
    ServiceLevelGoalPercent float64 `json:"serviceLevelGoalPercent"`


    // UseAbandonRateGoal - Toggle for abandon rate from service goals
    UseAbandonRateGoal bool `json:"useAbandonRateGoal"`


    // AbandonRateGoalPercent - Value of abandon rate objective
    AbandonRateGoalPercent float64 `json:"abandonRateGoalPercent"`

}

// String returns a JSON representation of the model
func (o *Longtermrequirementsservicegoal) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermrequirementsservicegoal) MarshalJSON() ([]byte, error) {
    type Alias Longtermrequirementsservicegoal

    if LongtermrequirementsservicegoalMarshalled {
        return []byte("{}"), nil
    }
    LongtermrequirementsservicegoalMarshalled = true

    return json.Marshal(&struct {
        
        UseAsaTarget bool `json:"useAsaTarget"`
        
        TargetAsaSec float64 `json:"targetAsaSec"`
        
        UseServiceLevelTarget bool `json:"useServiceLevelTarget"`
        
        ServiceLevelObjectiveSeconds float64 `json:"serviceLevelObjectiveSeconds"`
        
        ServiceLevelGoalPercent float64 `json:"serviceLevelGoalPercent"`
        
        UseAbandonRateGoal bool `json:"useAbandonRateGoal"`
        
        AbandonRateGoalPercent float64 `json:"abandonRateGoalPercent"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

