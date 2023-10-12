package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CriteriaitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CriteriaitemDud struct { 
    


    


    

}

// Criteriaitem - A singular criteria used to query executionData.
type Criteriaitem struct { 
    // Key - The id of the criteria to be checked.
    Key string `json:"key"`


    // Operator - The operator used to check on the criteria id.
    Operator string `json:"operator"`


    // Value - The target value used to query on.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Criteriaitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Criteriaitem) MarshalJSON() ([]byte, error) {
    type Alias Criteriaitem

    if CriteriaitemMarshalled {
        return []byte("{}"), nil
    }
    CriteriaitemMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

