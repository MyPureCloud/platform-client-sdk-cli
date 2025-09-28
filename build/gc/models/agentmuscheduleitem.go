package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmuscheduleitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmuscheduleitemDud struct { 
    


    


    

}

// Agentmuscheduleitem
type Agentmuscheduleitem struct { 
    // Agent - The agent to whom this schedule applies. Note: selfUri will not be populated if retrieving result via downloadUrl
    Agent Userreference `json:"agent"`


    // Shifts - The shift definitions for this agent schedule
    Shifts []Agentmuscheduleshift `json:"shifts"`


    // FullDayTimeOffMarkerDates - The full day time off marker dates which apply to this agent schedule, interpreted in the time zone of the relevant business unit
    FullDayTimeOffMarkerDates []time.Time `json:"fullDayTimeOffMarkerDates"`

}

// String returns a JSON representation of the model
func (o *Agentmuscheduleitem) String() string {
    
     o.Shifts = []Agentmuscheduleshift{{}} 
     o.FullDayTimeOffMarkerDates = []time.Time{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmuscheduleitem) MarshalJSON() ([]byte, error) {
    type Alias Agentmuscheduleitem

    if AgentmuscheduleitemMarshalled {
        return []byte("{}"), nil
    }
    AgentmuscheduleitemMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        Shifts []Agentmuscheduleshift `json:"shifts"`
        
        FullDayTimeOffMarkerDates []time.Time `json:"fullDayTimeOffMarkerDates"`
        *Alias
    }{

        


        
        Shifts: []Agentmuscheduleshift{{}},
        


        
        FullDayTimeOffMarkerDates: []time.Time{{}},
        

        Alias: (*Alias)(u),
    })
}

