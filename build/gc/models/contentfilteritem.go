package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentfilteritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentfilteritemDud struct { 
    


    


    


    

}

// Contentfilteritem
type Contentfilteritem struct { 
    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // Operator
    Operator string `json:"operator"`


    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Contentfilteritem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentfilteritem) MarshalJSON() ([]byte, error) {
    type Alias Contentfilteritem

    if ContentfilteritemMarshalled {
        return []byte("{}"), nil
    }
    ContentfilteritemMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Values: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

