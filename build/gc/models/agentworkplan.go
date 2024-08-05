package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentworkplan
type Agentworkplan struct { 
    


    // Name
    Name string `json:"name"`


    // ConstrainWeeklyPaidTime - Whether the weekly paid time constraint is enabled for this work plan
    ConstrainWeeklyPaidTime bool `json:"constrainWeeklyPaidTime"`


    // FlexibleWeeklyPaidTime - Whether the weekly paid time constraint is flexible for this work plan
    FlexibleWeeklyPaidTime bool `json:"flexibleWeeklyPaidTime"`


    // WeeklyExactPaidMinutes - Exact weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == false
    WeeklyExactPaidMinutes int `json:"weeklyExactPaidMinutes"`


    // WeeklyMinimumPaidMinutes - Minimum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
    WeeklyMinimumPaidMinutes int `json:"weeklyMinimumPaidMinutes"`


    // WeeklyMaximumPaidMinutes - Maximum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
    WeeklyMaximumPaidMinutes int `json:"weeklyMaximumPaidMinutes"`


    // OptionalDays - Optional days to schedule for this work plan
    OptionalDays Setwrapperdayofweek `json:"optionalDays"`


    // Shifts - Shifts in this work plan
    Shifts []Agentworkplanshift `json:"shifts"`


    

}

// String returns a JSON representation of the model
func (o *Agentworkplan) String() string {
    
    
    
    
    
    
    
     o.Shifts = []Agentworkplanshift{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplan) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplan

    if AgentworkplanMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ConstrainWeeklyPaidTime bool `json:"constrainWeeklyPaidTime"`
        
        FlexibleWeeklyPaidTime bool `json:"flexibleWeeklyPaidTime"`
        
        WeeklyExactPaidMinutes int `json:"weeklyExactPaidMinutes"`
        
        WeeklyMinimumPaidMinutes int `json:"weeklyMinimumPaidMinutes"`
        
        WeeklyMaximumPaidMinutes int `json:"weeklyMaximumPaidMinutes"`
        
        OptionalDays Setwrapperdayofweek `json:"optionalDays"`
        
        Shifts []Agentworkplanshift `json:"shifts"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Shifts: []Agentworkplanshift{{}},
        


        

        Alias: (*Alias)(u),
    })
}

