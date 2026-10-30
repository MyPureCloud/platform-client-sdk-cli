package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfschedulemanagementunitsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfschedulemanagementunitsettingsrequestDud struct { 
    


    


    


    


    


    

}

// Selfschedulemanagementunitsettingsrequest
type Selfschedulemanagementunitsettingsrequest struct { 
    // Enabled - Whether agent self-scheduling is enabled for this management unit
    Enabled bool `json:"enabled"`


    // CheckWorkPlanConstraints - Whether self-scheduling is further restricted by work plan constraints
    CheckWorkPlanConstraints bool `json:"checkWorkPlanConstraints"`


    // MinimumMinutesBeforeEarliestScheduleChange - The minimum number of minutes before the new activity start time to which the agent is allowed to self-schedule
    MinimumMinutesBeforeEarliestScheduleChange int `json:"minimumMinutesBeforeEarliestScheduleChange"`


    // MovableActivityCategories - Movable activity categories enabled for self-scheduling
    MovableActivityCategories Setwrapperactivitycategory `json:"movableActivityCategories"`


    // MovableActivityCodeIds - Movable activity codes enabled for self-scheduling
    MovableActivityCodeIds Setwrapperstring `json:"movableActivityCodeIds"`


    // Metadata - Version metadata for this management unit's self-scheduling settings
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Selfschedulemanagementunitsettingsrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfschedulemanagementunitsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Selfschedulemanagementunitsettingsrequest

    if SelfschedulemanagementunitsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    SelfschedulemanagementunitsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        CheckWorkPlanConstraints bool `json:"checkWorkPlanConstraints"`
        
        MinimumMinutesBeforeEarliestScheduleChange int `json:"minimumMinutesBeforeEarliestScheduleChange"`
        
        MovableActivityCategories Setwrapperactivitycategory `json:"movableActivityCategories"`
        
        MovableActivityCodeIds Setwrapperstring `json:"movableActivityCodeIds"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

