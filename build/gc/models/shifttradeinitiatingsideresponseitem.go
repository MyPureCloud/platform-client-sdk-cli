package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradeinitiatingsideresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradeinitiatingsideresponseitemDud struct { 
    


    


    


    

}

// Shifttradeinitiatingsideresponseitem
type Shifttradeinitiatingsideresponseitem struct { 
    // User - The user who initiated this trade
    User Userreference `json:"user"`


    // ManagementUnit - The management unit of the user who initiated this trade
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Schedule - Associated schedule information for the initiating user
    Schedule Schedulereferencewithbusinessunit `json:"schedule"`


    // Shift - The shift offered for trade by the initiating user
    Shift Shifttradeshiftresponseitem `json:"shift"`

}

// String returns a JSON representation of the model
func (o *Shifttradeinitiatingsideresponseitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradeinitiatingsideresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradeinitiatingsideresponseitem

    if ShifttradeinitiatingsideresponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradeinitiatingsideresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Schedule Schedulereferencewithbusinessunit `json:"schedule"`
        
        Shift Shifttradeshiftresponseitem `json:"shift"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

