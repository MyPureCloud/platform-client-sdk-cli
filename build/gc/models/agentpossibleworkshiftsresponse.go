package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentpossibleworkshiftsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentpossibleworkshiftsresponseDud struct { 
    


    


    


    


    

}

// Agentpossibleworkshiftsresponse
type Agentpossibleworkshiftsresponse struct { 
    // WeekStartDate - Start date of requested effective work plan. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekStartDate time.Time `json:"weekStartDate"`


    // Pattern - Each element is the ID of an effective work plan for a specific week
    Pattern []int `json:"pattern"`


    // WeeklyPossibleWorkShifts - Each element is a weekly effective work plan that can be used for multiple weeks
    WeeklyPossibleWorkShifts []Possibleworkshiftsforweek `json:"weeklyPossibleWorkShifts"`


    // SchedulerIntervalLengthMinutes - Number of minutes in each interval in the intervalScheduleProbabilities
    SchedulerIntervalLengthMinutes int `json:"schedulerIntervalLengthMinutes"`


    // TimeZone - The time zone of the business unit
    TimeZone string `json:"timeZone"`

}

// String returns a JSON representation of the model
func (o *Agentpossibleworkshiftsresponse) String() string {
    
     o.Pattern = []int{0} 
     o.WeeklyPossibleWorkShifts = []Possibleworkshiftsforweek{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentpossibleworkshiftsresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentpossibleworkshiftsresponse

    if AgentpossibleworkshiftsresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentpossibleworkshiftsresponseMarshalled = true

    return json.Marshal(&struct {
        
        WeekStartDate time.Time `json:"weekStartDate"`
        
        Pattern []int `json:"pattern"`
        
        WeeklyPossibleWorkShifts []Possibleworkshiftsforweek `json:"weeklyPossibleWorkShifts"`
        
        SchedulerIntervalLengthMinutes int `json:"schedulerIntervalLengthMinutes"`
        
        TimeZone string `json:"timeZone"`
        *Alias
    }{

        


        
        Pattern: []int{0},
        


        
        WeeklyPossibleWorkShifts: []Possibleworkshiftsforweek{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

