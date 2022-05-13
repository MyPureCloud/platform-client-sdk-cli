package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserstateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserstateDud struct { 
    


    


    


    StateChangeDate time.Time `json:"stateChangeDate"`

}

// Userstate
type Userstate struct { 
    // State - User's current state.
    State string `json:"state"`


    // Version - Version of this user.
    Version int `json:"version"`


    // StateChangeReason - Reason for a change in the user's state.
    StateChangeReason string `json:"stateChangeReason"`


    

}

// String returns a JSON representation of the model
func (o *Userstate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userstate) MarshalJSON() ([]byte, error) {
    type Alias Userstate

    if UserstateMarshalled {
        return []byte("{}"), nil
    }
    UserstateMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Version int `json:"version"`
        
        StateChangeReason string `json:"stateChangeReason"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

