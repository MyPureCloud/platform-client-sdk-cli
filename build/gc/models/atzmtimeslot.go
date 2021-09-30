package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AtzmtimeslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AtzmtimeslotDud struct { 
    


    

}

// Atzmtimeslot
type Atzmtimeslot struct { 
    // EarliestCallableTime - The earliest time to dial a contact. Valid format is HH:mm
    EarliestCallableTime string `json:"earliestCallableTime"`


    // LatestCallableTime - The latest time to dial a contact. Valid format is HH:mm
    LatestCallableTime string `json:"latestCallableTime"`

}

// String returns a JSON representation of the model
func (o *Atzmtimeslot) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Atzmtimeslot) MarshalJSON() ([]byte, error) {
    type Alias Atzmtimeslot

    if AtzmtimeslotMarshalled {
        return []byte("{}"), nil
    }
    AtzmtimeslotMarshalled = true

    return json.Marshal(&struct { 
        EarliestCallableTime string `json:"earliestCallableTime"`
        
        LatestCallableTime string `json:"latestCallableTime"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

