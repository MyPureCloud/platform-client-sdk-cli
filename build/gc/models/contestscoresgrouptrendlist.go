package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoresgrouptrendlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoresgrouptrendlistDud struct { 
    

}

// Contestscoresgrouptrendlist
type Contestscoresgrouptrendlist struct { 
    // Entities
    Entities []Contestscoresgrouptrend `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Contestscoresgrouptrendlist) String() string {
     o.Entities = []Contestscoresgrouptrend{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoresgrouptrendlist) MarshalJSON() ([]byte, error) {
    type Alias Contestscoresgrouptrendlist

    if ContestscoresgrouptrendlistMarshalled {
        return []byte("{}"), nil
    }
    ContestscoresgrouptrendlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contestscoresgrouptrend `json:"entities"`
        *Alias
    }{

        
        Entities: []Contestscoresgrouptrend{{}},
        

        Alias: (*Alias)(u),
    })
}

