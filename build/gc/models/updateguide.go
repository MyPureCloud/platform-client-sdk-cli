package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateguideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateguideDud struct { 
    

}

// Updateguide - Request body for updating a guide
type Updateguide struct { 
    // Name - The name of the guide.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Updateguide) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateguide) MarshalJSON() ([]byte, error) {
    type Alias Updateguide

    if UpdateguideMarshalled {
        return []byte("{}"), nil
    }
    UpdateguideMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

