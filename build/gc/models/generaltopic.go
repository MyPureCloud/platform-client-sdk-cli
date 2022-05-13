package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneraltopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneraltopicDud struct { 
    

}

// Generaltopic
type Generaltopic struct { 
    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Generaltopic) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generaltopic) MarshalJSON() ([]byte, error) {
    type Alias Generaltopic

    if GeneraltopicMarshalled {
        return []byte("{}"), nil
    }
    GeneraltopicMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

