package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReschedulingoptionsrunresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReschedulingoptionsrunresponseDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Reschedulingoptionsrunresponse
type Reschedulingoptionsrunresponse struct { 
    // ExistingSchedule - The existing schedule to which this reschedule run applies
    ExistingSchedule Buschedulereference `json:"existingSchedule"`


    // StartDate - The start date of the period to reschedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date of the period to reschedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // ManagementUnits - Per-management unit rescheduling options
    ManagementUnits []Reschedulingmanagementunitresponse `json:"managementUnits"`


    // AgentCount - The number of agents to be considered in the reschedule
    AgentCount int `json:"agentCount"`


    // ActivityCodeIds - The IDs of the activity codes being considered for reschedule
    ActivityCodeIds []string `json:"activityCodeIds"`


    // DoNotChangeWeeklyPaidTime - Whether weekly paid time is allowed to be changed
    DoNotChangeWeeklyPaidTime bool `json:"doNotChangeWeeklyPaidTime"`


    // DoNotChangeDailyPaidTime - Whether daily paid time is allowed to be changed
    DoNotChangeDailyPaidTime bool `json:"doNotChangeDailyPaidTime"`


    // DoNotChangeShiftStartTimes - Whether shift start times are allowed to be changed
    DoNotChangeShiftStartTimes bool `json:"doNotChangeShiftStartTimes"`


    // DoNotChangeManuallyEditedShifts - Whether manually edited shifts are allowed to be changed
    DoNotChangeManuallyEditedShifts bool `json:"doNotChangeManuallyEditedShifts"`

}

// String returns a JSON representation of the model
func (o *Reschedulingoptionsrunresponse) String() string {
    
    
    
     o.ManagementUnits = []Reschedulingmanagementunitresponse{{}} 
    
     o.ActivityCodeIds = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reschedulingoptionsrunresponse) MarshalJSON() ([]byte, error) {
    type Alias Reschedulingoptionsrunresponse

    if ReschedulingoptionsrunresponseMarshalled {
        return []byte("{}"), nil
    }
    ReschedulingoptionsrunresponseMarshalled = true

    return json.Marshal(&struct {
        
        ExistingSchedule Buschedulereference `json:"existingSchedule"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        ManagementUnits []Reschedulingmanagementunitresponse `json:"managementUnits"`
        
        AgentCount int `json:"agentCount"`
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        DoNotChangeWeeklyPaidTime bool `json:"doNotChangeWeeklyPaidTime"`
        
        DoNotChangeDailyPaidTime bool `json:"doNotChangeDailyPaidTime"`
        
        DoNotChangeShiftStartTimes bool `json:"doNotChangeShiftStartTimes"`
        
        DoNotChangeManuallyEditedShifts bool `json:"doNotChangeManuallyEditedShifts"`
        *Alias
    }{

        


        


        


        
        ManagementUnits: []Reschedulingmanagementunitresponse{{}},
        


        


        
        ActivityCodeIds: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

