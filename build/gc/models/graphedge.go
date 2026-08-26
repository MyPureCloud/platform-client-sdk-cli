package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GraphedgeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GraphedgeDud struct { 
    From string `json:"from"`


    To string `json:"to"`

}

// Graphedge
type Graphedge struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Graphedge) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Graphedge) MarshalJSON() ([]byte, error) {
    type Alias Graphedge

    if GraphedgeMarshalled {
        return []byte("{}"), nil
    }
    GraphedgeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

