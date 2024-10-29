package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelsfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelsfilterDud struct { 
    

}

// Labelsfilter
type Labelsfilter struct { 
    // Entities - A list of labels to filter by. Articles matching any of the specified labels can be accessed.
    Entities []Labelentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Labelsfilter) String() string {
     o.Entities = []Labelentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelsfilter) MarshalJSON() ([]byte, error) {
    type Alias Labelsfilter

    if LabelsfilterMarshalled {
        return []byte("{}"), nil
    }
    LabelsfilterMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Labelentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Labelentity{{}},
        

        Alias: (*Alias)(u),
    })
}

