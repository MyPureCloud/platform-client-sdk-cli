package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GraphverticesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GraphverticesDud struct { 
    Contacts []Contactvertex `json:"contacts"`


    Identifiers []Identifiervertex `json:"identifiers"`

}

// Graphvertices
type Graphvertices struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Graphvertices) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Graphvertices) MarshalJSON() ([]byte, error) {
    type Alias Graphvertices

    if GraphverticesMarshalled {
        return []byte("{}"), nil
    }
    GraphverticesMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

