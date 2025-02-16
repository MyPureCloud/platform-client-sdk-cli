package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsernextactivityreminderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsernextactivityreminderDud struct { 
    ActivityCategory string `json:"activityCategory"`


    StartDate time.Time `json:"startDate"`

}

// Usernextactivityreminder
type Usernextactivityreminder struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Usernextactivityreminder) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usernextactivityreminder) MarshalJSON() ([]byte, error) {
    type Alias Usernextactivityreminder

    if UsernextactivityreminderMarshalled {
        return []byte("{}"), nil
    }
    UsernextactivityreminderMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

