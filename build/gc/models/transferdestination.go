package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransferdestinationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransferdestinationDud struct { 
    


    

}

// Transferdestination
type Transferdestination struct { 
    // UserId - The id of the user if the command destination is a user.
    UserId string `json:"userId"`


    // Address - The destination address if the command destination is an endpoint.
    Address string `json:"address"`

}

// String returns a JSON representation of the model
func (o *Transferdestination) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transferdestination) MarshalJSON() ([]byte, error) {
    type Alias Transferdestination

    if TransferdestinationMarshalled {
        return []byte("{}"), nil
    }
    TransferdestinationMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        Address string `json:"address"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

