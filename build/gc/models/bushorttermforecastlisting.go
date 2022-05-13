package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BushorttermforecastlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BushorttermforecastlistingDud struct { 
    

}

// Bushorttermforecastlisting
type Bushorttermforecastlisting struct { 
    // Entities
    Entities []Bushorttermforecastlistitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastlisting) String() string {
     o.Entities = []Bushorttermforecastlistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bushorttermforecastlisting) MarshalJSON() ([]byte, error) {
    type Alias Bushorttermforecastlisting

    if BushorttermforecastlistingMarshalled {
        return []byte("{}"), nil
    }
    BushorttermforecastlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Bushorttermforecastlistitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Bushorttermforecastlistitem{{}},
        

        Alias: (*Alias)(u),
    })
}

