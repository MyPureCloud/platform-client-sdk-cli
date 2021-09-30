package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentfacetfilteritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentfacetfilteritemDud struct { 
    


    


    


    

}

// Contentfacetfilteritem
type Contentfacetfilteritem struct { 
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
func (o *Contentfacetfilteritem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentfacetfilteritem) MarshalJSON() ([]byte, error) {
    type Alias Contentfacetfilteritem

    if ContentfacetfilteritemMarshalled {
        return []byte("{}"), nil
    }
    ContentfacetfilteritemMarshalled = true

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

