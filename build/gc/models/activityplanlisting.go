package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanlistingDud struct { 
    

}

// Activityplanlisting
type Activityplanlisting struct { 
    // Entities
    Entities []Activityplanlistitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Activityplanlisting) String() string {
     o.Entities = []Activityplanlistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanlisting) MarshalJSON() ([]byte, error) {
    type Alias Activityplanlisting

    if ActivityplanlistingMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Activityplanlistitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Activityplanlistitem{{}},
        

        Alias: (*Alias)(u),
    })
}

