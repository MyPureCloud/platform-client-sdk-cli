package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateunavailabletimesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateunavailabletimesrequestDud struct { 
    

}

// Updateunavailabletimesrequest
type Updateunavailabletimesrequest struct { 
    // Entities - A list of unavailable time spans to update
    Entities []Updateunavailabletime `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Updateunavailabletimesrequest) String() string {
     o.Entities = []Updateunavailabletime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateunavailabletimesrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateunavailabletimesrequest

    if UpdateunavailabletimesrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateunavailabletimesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Updateunavailabletime `json:"entities"`
        *Alias
    }{

        
        Entities: []Updateunavailabletime{{}},
        

        Alias: (*Alias)(u),
    })
}

