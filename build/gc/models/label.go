package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelDud struct { 
    


    


    Encoded bool `json:"encoded"`

}

// Label
type Label struct { 
    // Name - Name of the label
    Name string `json:"name"`


    // Value - Value of the label
    Value string `json:"value"`


    

}

// String returns a JSON representation of the model
func (o *Label) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Label) MarshalJSON() ([]byte, error) {
    type Alias Label

    if LabelMarshalled {
        return []byte("{}"), nil
    }
    LabelMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

