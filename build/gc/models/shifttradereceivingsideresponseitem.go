package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradereceivingsideresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradereceivingsideresponseitemDud struct { 
    


    

}

// Shifttradereceivingsideresponseitem
type Shifttradereceivingsideresponseitem struct { 
    // User - The receiving user that was matched in a shift trade
    User Userreference `json:"user"`


    // Shift - The shift being traded by the receiving user, or null if picking up a shift in a one-sided trade
    Shift Shifttradeshiftresponseitem `json:"shift"`

}

// String returns a JSON representation of the model
func (o *Shifttradereceivingsideresponseitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradereceivingsideresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradereceivingsideresponseitem

    if ShifttradereceivingsideresponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradereceivingsideresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Shift Shifttradeshiftresponseitem `json:"shift"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

