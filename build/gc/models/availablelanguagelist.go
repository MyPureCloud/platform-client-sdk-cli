package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailablelanguagelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailablelanguagelistDud struct { 
    

}

// Availablelanguagelist
type Availablelanguagelist struct { 
    // Languages
    Languages []string `json:"languages"`

}

// String returns a JSON representation of the model
func (o *Availablelanguagelist) String() string {
     o.Languages = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availablelanguagelist) MarshalJSON() ([]byte, error) {
    type Alias Availablelanguagelist

    if AvailablelanguagelistMarshalled {
        return []byte("{}"), nil
    }
    AvailablelanguagelistMarshalled = true

    return json.Marshal(&struct {
        
        Languages []string `json:"languages"`
        *Alias
    }{

        
        Languages: []string{""},
        

        Alias: (*Alias)(u),
    })
}

