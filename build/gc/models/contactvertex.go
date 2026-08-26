package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactvertexMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactvertexDud struct { 
    VertexId string `json:"vertexId"`


    Contact Externalcontact `json:"contact"`

}

// Contactvertex
type Contactvertex struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Contactvertex) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactvertex) MarshalJSON() ([]byte, error) {
    type Alias Contactvertex

    if ContactvertexMarshalled {
        return []byte("{}"), nil
    }
    ContactvertexMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

