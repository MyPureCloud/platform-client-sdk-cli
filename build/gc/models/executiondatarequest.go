package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExecutiondatarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExecutiondatarequestDud struct { 
    

}

// Executiondatarequest - Used to retrieve a set of executionData history by the respective ids
type Executiondatarequest struct { 
    // Ids - A list of ids to retrieve
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Executiondatarequest) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Executiondatarequest) MarshalJSON() ([]byte, error) {
    type Alias Executiondatarequest

    if ExecutiondatarequestMarshalled {
        return []byte("{}"), nil
    }
    ExecutiondatarequestMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

