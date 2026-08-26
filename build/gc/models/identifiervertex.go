package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentifiervertexMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentifiervertexDud struct { 
    VertexId string `json:"vertexId"`


    NormalizedType string `json:"normalizedType"`


    NormalizedValue string `json:"normalizedValue"`

}

// Identifiervertex
type Identifiervertex struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Identifiervertex) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identifiervertex) MarshalJSON() ([]byte, error) {
    type Alias Identifiervertex

    if IdentifiervertexMarshalled {
        return []byte("{}"), nil
    }
    IdentifiervertexMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

