package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperschedulesetrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperschedulesetrequestDud struct { 
    

}

// Listwrapperschedulesetrequest
type Listwrapperschedulesetrequest struct { 
    // Values
    Values []Schedulesetrequest `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperschedulesetrequest) String() string {
     o.Values = []Schedulesetrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperschedulesetrequest) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperschedulesetrequest

    if ListwrapperschedulesetrequestMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperschedulesetrequestMarshalled = true

    return json.Marshal(&struct {
        
        Values []Schedulesetrequest `json:"values"`
        *Alias
    }{

        
        Values: []Schedulesetrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

