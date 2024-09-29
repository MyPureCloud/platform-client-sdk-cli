package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommunicationtranslationlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommunicationtranslationlistDud struct { 
    

}

// Communicationtranslationlist
type Communicationtranslationlist struct { 
    // Entities
    Entities []Communicationtranslation `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Communicationtranslationlist) String() string {
     o.Entities = []Communicationtranslation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Communicationtranslationlist) MarshalJSON() ([]byte, error) {
    type Alias Communicationtranslationlist

    if CommunicationtranslationlistMarshalled {
        return []byte("{}"), nil
    }
    CommunicationtranslationlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Communicationtranslation `json:"entities"`
        *Alias
    }{

        
        Entities: []Communicationtranslation{{}},
        

        Alias: (*Alias)(u),
    })
}

