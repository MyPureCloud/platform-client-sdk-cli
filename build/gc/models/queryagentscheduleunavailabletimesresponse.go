package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentscheduleunavailabletimesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentscheduleunavailabletimesresponseDud struct { 
    


    

}

// Queryagentscheduleunavailabletimesresponse
type Queryagentscheduleunavailabletimesresponse struct { 
    // ConsideredInScheduleGeneration - Indicates whether the unavailability times were considered in schedule generation. Returns false when no schedule exists
    ConsideredInScheduleGeneration bool `json:"consideredInScheduleGeneration"`


    // ScheduleGenerationUnavailableTimes - List of the unavailable times used in schedule generation
    ScheduleGenerationUnavailableTimes []Agentscheduleunavailabletime `json:"scheduleGenerationUnavailableTimes"`

}

// String returns a JSON representation of the model
func (o *Queryagentscheduleunavailabletimesresponse) String() string {
    
     o.ScheduleGenerationUnavailableTimes = []Agentscheduleunavailabletime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentscheduleunavailabletimesresponse) MarshalJSON() ([]byte, error) {
    type Alias Queryagentscheduleunavailabletimesresponse

    if QueryagentscheduleunavailabletimesresponseMarshalled {
        return []byte("{}"), nil
    }
    QueryagentscheduleunavailabletimesresponseMarshalled = true

    return json.Marshal(&struct {
        
        ConsideredInScheduleGeneration bool `json:"consideredInScheduleGeneration"`
        
        ScheduleGenerationUnavailableTimes []Agentscheduleunavailabletime `json:"scheduleGenerationUnavailableTimes"`
        *Alias
    }{

        


        
        ScheduleGenerationUnavailableTimes: []Agentscheduleunavailabletime{{}},
        

        Alias: (*Alias)(u),
    })
}

