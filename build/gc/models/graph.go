package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GraphMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GraphDud struct { 
    Vertices Graphvertices `json:"vertices"`


    Edges []Graphedge `json:"edges"`

}

// Graph
type Graph struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Graph) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Graph) MarshalJSON() ([]byte, error) {
    type Alias Graph

    if GraphMarshalled {
        return []byte("{}"), nil
    }
    GraphMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

