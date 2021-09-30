package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulehistorychangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulehistorychangeDud struct { 
    


    


    


    

}

// Buagentschedulehistorychange
type Buagentschedulehistorychange struct { 
    // Metadata - The metadata of the change, including who and when the change was made
    Metadata Buagentschedulehistorychangemetadata `json:"metadata"`


    // Shifts - The list of changed shifts
    Shifts []Buagentscheduleshift `json:"shifts"`


    // FullDayTimeOffMarkers - The list of changed full day time off markers
    FullDayTimeOffMarkers []Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`


    // Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
    Deletes Buagentschedulehistorydeletedchange `json:"deletes"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychange) String() string {
    
    
    
    
    
    
     o.Shifts = []Buagentscheduleshift{{}} 
    
    
    
     o.FullDayTimeOffMarkers = []Bufulldaytimeoffmarker{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulehistorychange) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulehistorychange

    if BuagentschedulehistorychangeMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulehistorychangeMarshalled = true

    return json.Marshal(&struct { 
        Metadata Buagentschedulehistorychangemetadata `json:"metadata"`
        
        Shifts []Buagentscheduleshift `json:"shifts"`
        
        FullDayTimeOffMarkers []Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`
        
        Deletes Buagentschedulehistorydeletedchange `json:"deletes"`
        
        *Alias
    }{
        

        

        

        
        Shifts: []Buagentscheduleshift{{}},
        

        

        
        FullDayTimeOffMarkers: []Bufulldaytimeoffmarker{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

