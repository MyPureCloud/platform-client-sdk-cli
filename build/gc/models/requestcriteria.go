package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestcriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestcriteriaDud struct { 
    


    


    


    

}

// Requestcriteria
type Requestcriteria struct { 
    // Key - The criteria key.
    Key string `json:"key"`


    // Values - The criteria values.
    Values []string `json:"values"`


    // ShouldIgnoreCase - Should criteria be case insensitive.
    ShouldIgnoreCase bool `json:"shouldIgnoreCase"`


    // Operator - The comparison operator.
    Operator string `json:"operator"`

}

// String returns a JSON representation of the model
func (o *Requestcriteria) String() string {
    
     o.Values = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestcriteria) MarshalJSON() ([]byte, error) {
    type Alias Requestcriteria

    if RequestcriteriaMarshalled {
        return []byte("{}"), nil
    }
    RequestcriteriaMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Values []string `json:"values"`
        
        ShouldIgnoreCase bool `json:"shouldIgnoreCase"`
        
        Operator string `json:"operator"`
        *Alias
    }{

        


        
        Values: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

