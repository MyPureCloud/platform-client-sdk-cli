package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AttributefilteritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AttributefilteritemDud struct { 
    


    


    

}

// Attributefilteritem
type Attributefilteritem struct { 
    // Id
    Id string `json:"id"`


    // Operator
    Operator string `json:"operator"`


    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Attributefilteritem) String() string {
    
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Attributefilteritem) MarshalJSON() ([]byte, error) {
    type Alias Attributefilteritem

    if AttributefilteritemMarshalled {
        return []byte("{}"), nil
    }
    AttributefilteritemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

