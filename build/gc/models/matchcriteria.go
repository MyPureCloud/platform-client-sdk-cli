package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MatchcriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MatchcriteriaDud struct { 
    


    


    


    

}

// Matchcriteria - Defines a simple matching condition
type Matchcriteria struct { 
    // JsonPath - The Goessner json path of the field to match
    JsonPath string `json:"jsonPath"`


    // Operator - The type of operation to perform for matching check
    Operator string `json:"operator"`


    // Value - The value to match on. Only one of value and values can be included
    Value interface{} `json:"value"`


    // Values - The list of values to match on. Only one of value and values can be included
    Values []map[string]interface{} `json:"values"`

}

// String returns a JSON representation of the model
func (o *Matchcriteria) String() string {
    
    
     o.Value = Interface{} 
     o.Values = []map[string]interface{}{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Matchcriteria) MarshalJSON() ([]byte, error) {
    type Alias Matchcriteria

    if MatchcriteriaMarshalled {
        return []byte("{}"), nil
    }
    MatchcriteriaMarshalled = true

    return json.Marshal(&struct {
        
        JsonPath string `json:"jsonPath"`
        
        Operator string `json:"operator"`
        
        Value interface{} `json:"value"`
        
        Values []map[string]interface{} `json:"values"`
        *Alias
    }{

        


        


        
        Value: Interface{},
        


        
        Values: []map[string]interface{}{{}},
        

        Alias: (*Alias)(u),
    })
}

