package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemfilterDud struct { 
    


    


    


    

}

// Workitemfilter
type Workitemfilter struct { 
    // Name - Attribute name.
    Name string `json:"name"`


    // VarType - Attribute type.
    VarType string `json:"type"`


    // Operator - Filter operator.
    Operator string `json:"operator"`


    // Values - List of values to be used in the filter.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Workitemfilter) String() string {
    
    
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemfilter) MarshalJSON() ([]byte, error) {
    type Alias Workitemfilter

    if WorkitemfilterMarshalled {
        return []byte("{}"), nil
    }
    WorkitemfilterMarshalled = true

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

