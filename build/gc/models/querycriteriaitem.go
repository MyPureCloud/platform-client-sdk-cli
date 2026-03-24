package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerycriteriaitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerycriteriaitemDud struct { 
    


    


    

}

// Querycriteriaitem - A singular item used to query libraries
type Querycriteriaitem struct { 
    // Key - The key to filter on
    Key string `json:"key"`


    // Operator - The operator for comparison
    Operator string `json:"operator"`


    // Value - The target value to match
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Querycriteriaitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querycriteriaitem) MarshalJSON() ([]byte, error) {
    type Alias Querycriteriaitem

    if QuerycriteriaitemMarshalled {
        return []byte("{}"), nil
    }
    QuerycriteriaitemMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

