package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributeslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributeslistingDud struct { 
    

}

// Customattributeslisting
type Customattributeslisting struct { 
    // Results
    Results []Customattributes `json:"results"`

}

// String returns a JSON representation of the model
func (o *Customattributeslisting) String() string {
     o.Results = []Customattributes{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributeslisting) MarshalJSON() ([]byte, error) {
    type Alias Customattributeslisting

    if CustomattributeslistingMarshalled {
        return []byte("{}"), nil
    }
    CustomattributeslistingMarshalled = true

    return json.Marshal(&struct {
        
        Results []Customattributes `json:"results"`
        *Alias
    }{

        
        Results: []Customattributes{{}},
        

        Alias: (*Alias)(u),
    })
}

