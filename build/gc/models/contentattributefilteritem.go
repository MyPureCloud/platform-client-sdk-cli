package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentattributefilteritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentattributefilteritemDud struct { 
    


    


    

}

// Contentattributefilteritem
type Contentattributefilteritem struct { 
    // Id
    Id string `json:"id"`


    // Operator
    Operator string `json:"operator"`


    // Values
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Contentattributefilteritem) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentattributefilteritem) MarshalJSON() ([]byte, error) {
    type Alias Contentattributefilteritem

    if ContentattributefilteritemMarshalled {
        return []byte("{}"), nil
    }
    ContentattributefilteritemMarshalled = true

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

