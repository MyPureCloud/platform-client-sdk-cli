package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseravailabletimesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseravailabletimesDud struct { 
    User Userreference `json:"user"`


    AvailableTimes []Availabletime `json:"availableTimes"`

}

// Useravailabletimes
type Useravailabletimes struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Useravailabletimes) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useravailabletimes) MarshalJSON() ([]byte, error) {
    type Alias Useravailabletimes

    if UseravailabletimesMarshalled {
        return []byte("{}"), nil
    }
    UseravailabletimesMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

