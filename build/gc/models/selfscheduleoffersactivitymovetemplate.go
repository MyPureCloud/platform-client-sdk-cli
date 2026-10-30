package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleoffersactivitymovetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleoffersactivitymovetemplateDud struct { 
    


    


    


    


    


    


    

}

// Selfscheduleoffersactivitymovetemplate
type Selfscheduleoffersactivitymovetemplate struct { 
    // Schedule - The schedule associated with the shift the agent wants to edit.
    Schedule Buschedulereference `json:"schedule"`


    // ShiftId - The ID of the shift to which the self-schedule activity move is associated
    ShiftId string `json:"shiftId"`


    // ReferenceKey - The unique identifier for the offered shift edit that was used for this shift edit
    ReferenceKey string `json:"referenceKey"`


    // StartDate - The start date/time for the shift in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The duration of this shift in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // ActivityCode - The activity related to this shift
    ActivityCode Activitycodereference `json:"activityCode"`


    // AlternateStartDates - The alternative start dates for this shift
    AlternateStartDates []time.Time `json:"alternateStartDates"`

}

// String returns a JSON representation of the model
func (o *Selfscheduleoffersactivitymovetemplate) String() string {
    
    
    
    
    
    
     o.AlternateStartDates = []time.Time{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleoffersactivitymovetemplate) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleoffersactivitymovetemplate

    if SelfscheduleoffersactivitymovetemplateMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleoffersactivitymovetemplateMarshalled = true

    return json.Marshal(&struct {
        
        Schedule Buschedulereference `json:"schedule"`
        
        ShiftId string `json:"shiftId"`
        
        ReferenceKey string `json:"referenceKey"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        ActivityCode Activitycodereference `json:"activityCode"`
        
        AlternateStartDates []time.Time `json:"alternateStartDates"`
        *Alias
    }{

        


        


        


        


        


        


        
        AlternateStartDates: []time.Time{{}},
        

        Alias: (*Alias)(u),
    })
}

