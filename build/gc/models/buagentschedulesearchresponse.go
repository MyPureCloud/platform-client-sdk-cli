package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulesearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulesearchresponseDud struct { 
    


    


    

}

// Buagentschedulesearchresponse
type Buagentschedulesearchresponse struct { 
    // User - The user to whom this agent schedule applies
    User Userreference `json:"user"`


    // Shifts - The shift definitions for this agent schedule
    Shifts []Buagentscheduleshift `json:"shifts"`


    // FullDayTimeOffMarkers - Full day time off markers which apply to this agent schedule
    FullDayTimeOffMarkers []Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulesearchresponse) String() string {
    
     o.Shifts = []Buagentscheduleshift{{}} 
     o.FullDayTimeOffMarkers = []Bufulldaytimeoffmarker{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulesearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulesearchresponse

    if BuagentschedulesearchresponseMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulesearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Shifts []Buagentscheduleshift `json:"shifts"`
        
        FullDayTimeOffMarkers []Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`
        *Alias
    }{

        


        
        Shifts: []Buagentscheduleshift{{}},
        


        
        FullDayTimeOffMarkers: []Bufulldaytimeoffmarker{{}},
        

        Alias: (*Alias)(u),
    })
}

