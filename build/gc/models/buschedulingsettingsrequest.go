package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulingsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulingsettingsrequestDud struct { 
    


    


    


    

}

// Buschedulingsettingsrequest
type Buschedulingsettingsrequest struct { 
    // MessageSeverities - Schedule generation message severity configuration
    MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`


    // SyncTimeOffProperties - Synchronize set of time off properties from scheduled activities to time off requests when the schedule is published.
    SyncTimeOffProperties Setwrappersynctimeoffproperty `json:"syncTimeOffProperties"`


    // ServiceGoalImpact - Configures the max percent increase and decrease of service goals for this business unit
    ServiceGoalImpact Wfmservicegoalimpactsettings `json:"serviceGoalImpact"`


    // AllowWorkPlanPerMinuteGranularity - Indicates whether or not per minute granularity for scheduling will be enabled for this business unit
    AllowWorkPlanPerMinuteGranularity bool `json:"allowWorkPlanPerMinuteGranularity"`

}

// String returns a JSON representation of the model
func (o *Buschedulingsettingsrequest) String() string {
     o.MessageSeverities = []Schedulermessagetypeseverity{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulingsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Buschedulingsettingsrequest

    if BuschedulingsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    BuschedulingsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`
        
        SyncTimeOffProperties Setwrappersynctimeoffproperty `json:"syncTimeOffProperties"`
        
        ServiceGoalImpact Wfmservicegoalimpactsettings `json:"serviceGoalImpact"`
        
        AllowWorkPlanPerMinuteGranularity bool `json:"allowWorkPlanPerMinuteGranularity"`
        *Alias
    }{

        
        MessageSeverities: []Schedulermessagetypeseverity{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

