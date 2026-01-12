package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BureschedulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BureschedulerequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Bureschedulerequest
type Bureschedulerequest struct { 
    // StartDate - The start of the range to reschedule.  Defaults to the beginning of the schedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end of the range to reschedule.  Defaults the the end of the schedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // AgentIds - The IDs of the agents to consider for rescheduling.  Omit to consider all agents in the specified management units.Agents not in the specified management units will be ignored
    AgentIds []string `json:"agentIds"`


    // ActivityCodeIds - The IDs of the activity codes to consider for rescheduling.  Omit to consider all activity codes
    ActivityCodeIds []string `json:"activityCodeIds"`


    // ManagementUnitIds - The IDs of the management units to reschedule
    ManagementUnitIds []string `json:"managementUnitIds"`


    // DoNotChangeWeeklyPaidTime - Instructs the scheduler whether it is allowed to change weekly paid time
    DoNotChangeWeeklyPaidTime bool `json:"doNotChangeWeeklyPaidTime"`


    // DoNotChangeDailyPaidTime - Instructs the scheduler whether it is allowed to change daily paid time
    DoNotChangeDailyPaidTime bool `json:"doNotChangeDailyPaidTime"`


    // DoNotChangeShiftStartTimes - Instructs the scheduler whether it is allowed to change shift start times
    DoNotChangeShiftStartTimes bool `json:"doNotChangeShiftStartTimes"`


    // DoNotChangeManuallyEditedShifts - Instructs the scheduler whether it is allowed to change manually edited shifts
    DoNotChangeManuallyEditedShifts bool `json:"doNotChangeManuallyEditedShifts"`


    // ActivitySmoothingType - Overrides the default BU level activity smoothing type for this reschedule run
    ActivitySmoothingType string `json:"activitySmoothingType"`


    // InduceScheduleVariability - Overrides the default BU level induce schedule variability setting for this reschedule run
    InduceScheduleVariability bool `json:"induceScheduleVariability"`


    // UseUnavailableTimesSnapshot - Whether to use original unavailable times from schedule generation or latest saved unavailable times for this reschedule run
    UseUnavailableTimesSnapshot bool `json:"useUnavailableTimesSnapshot"`

}

// String returns a JSON representation of the model
func (o *Bureschedulerequest) String() string {
    
    
     o.AgentIds = []string{""} 
     o.ActivityCodeIds = []string{""} 
     o.ManagementUnitIds = []string{""} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bureschedulerequest) MarshalJSON() ([]byte, error) {
    type Alias Bureschedulerequest

    if BureschedulerequestMarshalled {
        return []byte("{}"), nil
    }
    BureschedulerequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        AgentIds []string `json:"agentIds"`
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        DoNotChangeWeeklyPaidTime bool `json:"doNotChangeWeeklyPaidTime"`
        
        DoNotChangeDailyPaidTime bool `json:"doNotChangeDailyPaidTime"`
        
        DoNotChangeShiftStartTimes bool `json:"doNotChangeShiftStartTimes"`
        
        DoNotChangeManuallyEditedShifts bool `json:"doNotChangeManuallyEditedShifts"`
        
        ActivitySmoothingType string `json:"activitySmoothingType"`
        
        InduceScheduleVariability bool `json:"induceScheduleVariability"`
        
        UseUnavailableTimesSnapshot bool `json:"useUnavailableTimesSnapshot"`
        *Alias
    }{

        


        


        
        AgentIds: []string{""},
        


        
        ActivityCodeIds: []string{""},
        


        
        ManagementUnitIds: []string{""},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

