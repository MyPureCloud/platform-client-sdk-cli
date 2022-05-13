package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulehistoryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulehistoryresponseDud struct { 
    


    


    


    

}

// Buagentschedulehistoryresponse
type Buagentschedulehistoryresponse struct { 
    // PriorPublishedSchedules - The list of previously published schedules
    PriorPublishedSchedules []Buschedulereference `json:"priorPublishedSchedules"`


    // BasePublishedSchedule - The originally published agent schedules
    BasePublishedSchedule Buagentschedulehistorychange `json:"basePublishedSchedule"`


    // DroppedChanges - The changes dropped from the schedule history. This will happen if the schedule history is too large
    DroppedChanges []Buagentschedulehistorydroppedchange `json:"droppedChanges"`


    // Changes - The list of changes for the schedule history
    Changes []Buagentschedulehistorychange `json:"changes"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistoryresponse) String() string {
     o.PriorPublishedSchedules = []Buschedulereference{{}} 
    
     o.DroppedChanges = []Buagentschedulehistorydroppedchange{{}} 
     o.Changes = []Buagentschedulehistorychange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulehistoryresponse) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulehistoryresponse

    if BuagentschedulehistoryresponseMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulehistoryresponseMarshalled = true

    return json.Marshal(&struct {
        
        PriorPublishedSchedules []Buschedulereference `json:"priorPublishedSchedules"`
        
        BasePublishedSchedule Buagentschedulehistorychange `json:"basePublishedSchedule"`
        
        DroppedChanges []Buagentschedulehistorydroppedchange `json:"droppedChanges"`
        
        Changes []Buagentschedulehistorychange `json:"changes"`
        *Alias
    }{

        
        PriorPublishedSchedules: []Buschedulereference{{}},
        


        


        
        DroppedChanges: []Buagentschedulehistorydroppedchange{{}},
        


        
        Changes: []Buagentschedulehistorychange{{}},
        

        Alias: (*Alias)(u),
    })
}

