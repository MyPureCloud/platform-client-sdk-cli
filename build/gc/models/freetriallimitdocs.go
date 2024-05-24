package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FreetriallimitdocsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FreetriallimitdocsDud struct { 
    

}

// Freetriallimitdocs
type Freetriallimitdocs struct { 
    // Namespaces
    Namespaces []Freetrialnamespace `json:"namespaces"`

}

// String returns a JSON representation of the model
func (o *Freetriallimitdocs) String() string {
     o.Namespaces = []Freetrialnamespace{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Freetriallimitdocs) MarshalJSON() ([]byte, error) {
    type Alias Freetriallimitdocs

    if FreetriallimitdocsMarshalled {
        return []byte("{}"), nil
    }
    FreetriallimitdocsMarshalled = true

    return json.Marshal(&struct {
        
        Namespaces []Freetrialnamespace `json:"namespaces"`
        *Alias
    }{

        
        Namespaces: []Freetrialnamespace{{}},
        

        Alias: (*Alias)(u),
    })
}

