package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueanomalyexportjobtemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueanomalyexportjobtemplateDud struct { 
    

}

// Queueanomalyexportjobtemplate
type Queueanomalyexportjobtemplate struct { 
    // Entities
    Entities []Queueanomalyexport `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Queueanomalyexportjobtemplate) String() string {
     o.Entities = []Queueanomalyexport{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueanomalyexportjobtemplate) MarshalJSON() ([]byte, error) {
    type Alias Queueanomalyexportjobtemplate

    if QueueanomalyexportjobtemplateMarshalled {
        return []byte("{}"), nil
    }
    QueueanomalyexportjobtemplateMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Queueanomalyexport `json:"entities"`
        *Alias
    }{

        
        Entities: []Queueanomalyexport{{}},
        

        Alias: (*Alias)(u),
    })
}

