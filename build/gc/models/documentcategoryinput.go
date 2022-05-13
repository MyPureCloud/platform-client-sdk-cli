package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentcategoryinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentcategoryinputDud struct { 
    

}

// Documentcategoryinput
type Documentcategoryinput struct { 
    // Id - KnowledgeBase Category ID
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Documentcategoryinput) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentcategoryinput) MarshalJSON() ([]byte, error) {
    type Alias Documentcategoryinput

    if DocumentcategoryinputMarshalled {
        return []byte("{}"), nil
    }
    DocumentcategoryinputMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

