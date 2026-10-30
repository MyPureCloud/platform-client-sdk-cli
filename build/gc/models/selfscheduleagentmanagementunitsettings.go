package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleagentmanagementunitsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleagentmanagementunitsettingsDud struct { 
    


    


    


    

}

// Selfscheduleagentmanagementunitsettings
type Selfscheduleagentmanagementunitsettings struct { 
    // Enabled - Whether agent self-scheduling is enabled for this management unit
    Enabled bool `json:"enabled"`


    // MinimumMinutesBeforeEarliestScheduleChange - The minimum number of minutes before the new activity start time to which the agent is allowed to self-schedule
    MinimumMinutesBeforeEarliestScheduleChange int `json:"minimumMinutesBeforeEarliestScheduleChange"`


    // MovableActivityCategories - Movable activity categories enabled for self-schedule
    MovableActivityCategories []string `json:"movableActivityCategories"`


    // MovableActivityCodeIds - Movable activity codes enabled for self-schedule
    MovableActivityCodeIds []string `json:"movableActivityCodeIds"`

}

// String returns a JSON representation of the model
func (o *Selfscheduleagentmanagementunitsettings) String() string {
    
    
     o.MovableActivityCategories = []string{""} 
     o.MovableActivityCodeIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleagentmanagementunitsettings) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleagentmanagementunitsettings

    if SelfscheduleagentmanagementunitsettingsMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleagentmanagementunitsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        MinimumMinutesBeforeEarliestScheduleChange int `json:"minimumMinutesBeforeEarliestScheduleChange"`
        
        MovableActivityCategories []string `json:"movableActivityCategories"`
        
        MovableActivityCodeIds []string `json:"movableActivityCodeIds"`
        *Alias
    }{

        


        


        
        MovableActivityCategories: []string{""},
        


        
        MovableActivityCodeIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

