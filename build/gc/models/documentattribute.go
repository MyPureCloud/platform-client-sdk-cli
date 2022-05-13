package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentattributeDud struct { 
    


    

}

// Documentattribute
type Documentattribute struct { 
    // Attribute
    Attribute Attribute `json:"attribute"`


    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Documentattribute) String() string {
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentattribute) MarshalJSON() ([]byte, error) {
    type Alias Documentattribute

    if DocumentattributeMarshalled {
        return []byte("{}"), nil
    }
    DocumentattributeMarshalled = true

    return json.Marshal(&struct {
        
        Attribute Attribute `json:"attribute"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

