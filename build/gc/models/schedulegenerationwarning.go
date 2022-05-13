package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegenerationwarningMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegenerationwarningDud struct { 
    


    


    


    


    


    


    


    

}

// Schedulegenerationwarning - Schedule generation warning
type Schedulegenerationwarning struct { 
    // UserId - ID of the user in the warning
    UserId string `json:"userId"`


    // UserNotLicensed - Whether the user does not have the appropriate license to be scheduled
    UserNotLicensed bool `json:"userNotLicensed"`


    // UnableToMeetMaxDays - Whether the number of scheduled days exceeded the maximum days to schedule defined in the agent work plan
    UnableToMeetMaxDays bool `json:"unableToMeetMaxDays"`


    // UnableToScheduleRequiredDays - Days indicated as required to work in agent work plan where no viable shift was found to schedule
    UnableToScheduleRequiredDays []string `json:"unableToScheduleRequiredDays"`


    // UnableToMeetMinPaidForTheWeek - Whether the schedule did not meet the minimum paid time for the week defined in the agent work plan
    UnableToMeetMinPaidForTheWeek bool `json:"unableToMeetMinPaidForTheWeek"`


    // UnableToMeetMaxPaidForTheWeek - Whether the schedule exceeded the maximum paid time for the week defined in the agent work plan
    UnableToMeetMaxPaidForTheWeek bool `json:"unableToMeetMaxPaidForTheWeek"`


    // NoNeedDays - Days agent was scheduled but there was no need to meet. The scheduled days have no effect on service levels
    NoNeedDays []string `json:"noNeedDays"`


    // ShiftsTooCloseTogether - Whether the schedule did not meet the minimum time between shifts defined in the agent work plan
    ShiftsTooCloseTogether bool `json:"shiftsTooCloseTogether"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationwarning) String() string {
    
    
    
     o.UnableToScheduleRequiredDays = []string{""} 
    
    
     o.NoNeedDays = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegenerationwarning) MarshalJSON() ([]byte, error) {
    type Alias Schedulegenerationwarning

    if SchedulegenerationwarningMarshalled {
        return []byte("{}"), nil
    }
    SchedulegenerationwarningMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        UserNotLicensed bool `json:"userNotLicensed"`
        
        UnableToMeetMaxDays bool `json:"unableToMeetMaxDays"`
        
        UnableToScheduleRequiredDays []string `json:"unableToScheduleRequiredDays"`
        
        UnableToMeetMinPaidForTheWeek bool `json:"unableToMeetMinPaidForTheWeek"`
        
        UnableToMeetMaxPaidForTheWeek bool `json:"unableToMeetMaxPaidForTheWeek"`
        
        NoNeedDays []string `json:"noNeedDays"`
        
        ShiftsTooCloseTogether bool `json:"shiftsTooCloseTogether"`
        *Alias
    }{

        


        


        


        
        UnableToScheduleRequiredDays: []string{""},
        


        


        


        
        NoNeedDays: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

