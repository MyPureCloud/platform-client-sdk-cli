package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulehistorydeletedchangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulehistorydeletedchangeDud struct { 
    


    


    

}

// Buagentschedulehistorydeletedchange
type Buagentschedulehistorydeletedchange struct { 
    // ShiftIds - The IDs of deleted shifts
    ShiftIds []string `json:"shiftIds"`


    // FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
    FullDayTimeOffMarkerDates []time.Time `json:"fullDayTimeOffMarkerDates"`


    // AgentSchedule - Whether the entire agent schedule was deleted
    AgentSchedule bool `json:"agentSchedule"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydeletedchange) String() string {
    
    
     o.ShiftIds = []string{""} 
    
    
    
     o.FullDayTimeOffMarkerDates = []time.Time{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulehistorydeletedchange) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulehistorydeletedchange

    if BuagentschedulehistorydeletedchangeMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulehistorydeletedchangeMarshalled = true

    return json.Marshal(&struct { 
        ShiftIds []string `json:"shiftIds"`
        
        FullDayTimeOffMarkerDates []time.Time `json:"fullDayTimeOffMarkerDates"`
        
        AgentSchedule bool `json:"agentSchedule"`
        
        *Alias
    }{
        

        
        ShiftIds: []string{""},
        

        

        
        FullDayTimeOffMarkerDates: []time.Time{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

