package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulehistorydroppedchangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulehistorydroppedchangeDud struct { 
    


    


    


    

}

// Buagentschedulehistorydroppedchange
type Buagentschedulehistorydroppedchange struct { 
    // Metadata - The metadata of the change, including who and when the change was made
    Metadata Buagentschedulehistorychangemetadata `json:"metadata"`


    // ShiftIds - The IDs of deleted shifts
    ShiftIds []string `json:"shiftIds"`


    // FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
    FullDayTimeOffMarkerDates []time.Time `json:"fullDayTimeOffMarkerDates"`


    // Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
    Deletes Buagentschedulehistorydeletedchange `json:"deletes"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydroppedchange) String() string {
    
    
    
    
    
    
     o.ShiftIds = []string{""} 
    
    
    
     o.FullDayTimeOffMarkerDates = []time.Time{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulehistorydroppedchange) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulehistorydroppedchange

    if BuagentschedulehistorydroppedchangeMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulehistorydroppedchangeMarshalled = true

    return json.Marshal(&struct { 
        Metadata Buagentschedulehistorychangemetadata `json:"metadata"`
        
        ShiftIds []string `json:"shiftIds"`
        
        FullDayTimeOffMarkerDates []time.Time `json:"fullDayTimeOffMarkerDates"`
        
        Deletes Buagentschedulehistorydeletedchange `json:"deletes"`
        
        *Alias
    }{
        

        

        

        
        ShiftIds: []string{""},
        

        

        
        FullDayTimeOffMarkerDates: []time.Time{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

