package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrencelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrencelistingDud struct { 
    


    

}

// Activityplanoccurrencelisting
type Activityplanoccurrencelisting struct { 
    // Entities
    Entities []Activityplanoccurrencelistitem `json:"entities"`


    // Schedules - A list of schedules to which the sessions are attached. Reference this by index to resolve the schedule for a particular session
    Schedules []Buschedulesummary `json:"schedules"`

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencelisting) String() string {
     o.Entities = []Activityplanoccurrencelistitem{{}} 
     o.Schedules = []Buschedulesummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrencelisting) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrencelisting

    if ActivityplanoccurrencelistingMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrencelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Activityplanoccurrencelistitem `json:"entities"`
        
        Schedules []Buschedulesummary `json:"schedules"`
        *Alias
    }{

        
        Entities: []Activityplanoccurrencelistitem{{}},
        


        
        Schedules: []Buschedulesummary{{}},
        

        Alias: (*Alias)(u),
    })
}

