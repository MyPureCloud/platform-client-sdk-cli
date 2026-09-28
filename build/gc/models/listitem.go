package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListitemDud struct { 
    


    


    


    

}

// Listitem
type Listitem struct { 
    // Value - The value returned when this item is selected.
    Value string `json:"value"`


    // Synonyms - Alternative phrases that should match this value. Used only with Exact match type.
    Synonyms []string `json:"synonyms"`


    // Active - Whether this list item is active and available for selection.
    Active bool `json:"active"`


    // Description - Description of this value for semantic matching. Used only with Semantic match type.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Listitem) String() string {
    
     o.Synonyms = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listitem) MarshalJSON() ([]byte, error) {
    type Alias Listitem

    if ListitemMarshalled {
        return []byte("{}"), nil
    }
    ListitemMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        
        Synonyms []string `json:"synonyms"`
        
        Active bool `json:"active"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        
        Synonyms: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

