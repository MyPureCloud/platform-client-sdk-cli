package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UrlconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UrlconditionDud struct { 
    


    

}

// Urlcondition
type Urlcondition struct { 
    // Values - The URL condition value.
    Values []string `json:"values"`


    // Operator - The comparison operator.
    Operator string `json:"operator"`

}

// String returns a JSON representation of the model
func (o *Urlcondition) String() string {
     o.Values = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Urlcondition) MarshalJSON() ([]byte, error) {
    type Alias Urlcondition

    if UrlconditionMarshalled {
        return []byte("{}"), nil
    }
    UrlconditionMarshalled = true

    return json.Marshal(&struct {
        
        Values []string `json:"values"`
        
        Operator string `json:"operator"`
        *Alias
    }{

        
        Values: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

