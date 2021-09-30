package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FreeseatingconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FreeseatingconfigurationDud struct { 
    


    

}

// Freeseatingconfiguration
type Freeseatingconfiguration struct { 
    // FreeSeatingState - The FreeSeatingState for FreeSeatingConfiguration. Can be ON, OFF, or PARTIAL. ON meaning disassociate the user after the ttl expires, OFF meaning never disassociate the user, and PARTIAL meaning only disassociate when a user explicitly clicks log out.
    FreeSeatingState string `json:"freeSeatingState"`


    // TtlMinutes - The amount of time in minutes until an offline user is disassociated from their station
    TtlMinutes int `json:"ttlMinutes"`

}

// String returns a JSON representation of the model
func (o *Freeseatingconfiguration) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Freeseatingconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Freeseatingconfiguration

    if FreeseatingconfigurationMarshalled {
        return []byte("{}"), nil
    }
    FreeseatingconfigurationMarshalled = true

    return json.Marshal(&struct { 
        FreeSeatingState string `json:"freeSeatingState"`
        
        TtlMinutes int `json:"ttlMinutes"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

