package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingtemplatecarouselMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingtemplatecarouselDud struct { 
    

}

// Recordingtemplatecarousel
type Recordingtemplatecarousel struct { 
    // Cards - An array of template card objects.
    Cards []Recordingtemplatecard `json:"cards"`

}

// String returns a JSON representation of the model
func (o *Recordingtemplatecarousel) String() string {
     o.Cards = []Recordingtemplatecard{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingtemplatecarousel) MarshalJSON() ([]byte, error) {
    type Alias Recordingtemplatecarousel

    if RecordingtemplatecarouselMarshalled {
        return []byte("{}"), nil
    }
    RecordingtemplatecarouselMarshalled = true

    return json.Marshal(&struct {
        
        Cards []Recordingtemplatecard `json:"cards"`
        *Alias
    }{

        
        Cards: []Recordingtemplatecard{{}},
        

        Alias: (*Alias)(u),
    })
}

