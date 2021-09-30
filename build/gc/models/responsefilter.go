package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsefilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsefilterDud struct { 
    


    


    

}

// Responsefilter - Used to filter response queries
type Responsefilter struct { 
    // Name - Field to filter on. Allowed values are 'name' and 'libraryId.
    Name string `json:"name"`


    // Operator - Filter operation: IN, EQUALS, NOTEQUALS.
    Operator string `json:"operator"`


    // Values - Values to filter on.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Responsefilter) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsefilter) MarshalJSON() ([]byte, error) {
    type Alias Responsefilter

    if ResponsefilterMarshalled {
        return []byte("{}"), nil
    }
    ResponsefilterMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        
        *Alias
    }{
        

        

        

        

        

        
        Values: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

