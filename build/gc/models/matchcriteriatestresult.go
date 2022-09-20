package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MatchcriteriatestresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MatchcriteriatestresultDud struct { 
    


    


    


    


    


    


    

}

// Matchcriteriatestresult - Results of a matching expression
type Matchcriteriatestresult struct { 
    // JsonPath - The Goessner json path of the field to match
    JsonPath string `json:"jsonPath"`


    // Operator - The type of operation to perform for matching check
    Operator string `json:"operator"`


    // Value - The value to match on. Only one of value and values can be included
    Value interface{} `json:"value"`


    // Values - The list of values to match on. Only one of value and values can be included
    Values []map[string]interface{} `json:"values"`


    // GeneratedJsonPathCondition - The generated json path condition
    GeneratedJsonPathCondition string `json:"generatedJsonPathCondition"`


    // Match - Did the generated json path condition match
    Match bool `json:"match"`


    // JsonPathExtraction - The json paths and their values that were compared
    JsonPathExtraction []Matchtestresult `json:"jsonPathExtraction"`

}

// String returns a JSON representation of the model
func (o *Matchcriteriatestresult) String() string {
    
    
     o.Value = Interface{} 
     o.Values = []map[string]interface{}{{}} 
    
    
     o.JsonPathExtraction = []Matchtestresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Matchcriteriatestresult) MarshalJSON() ([]byte, error) {
    type Alias Matchcriteriatestresult

    if MatchcriteriatestresultMarshalled {
        return []byte("{}"), nil
    }
    MatchcriteriatestresultMarshalled = true

    return json.Marshal(&struct {
        
        JsonPath string `json:"jsonPath"`
        
        Operator string `json:"operator"`
        
        Value interface{} `json:"value"`
        
        Values []map[string]interface{} `json:"values"`
        
        GeneratedJsonPathCondition string `json:"generatedJsonPathCondition"`
        
        Match bool `json:"match"`
        
        JsonPathExtraction []Matchtestresult `json:"jsonPathExtraction"`
        *Alias
    }{

        


        


        
        Value: Interface{},
        


        
        Values: []map[string]interface{}{{}},
        


        


        


        
        JsonPathExtraction: []Matchtestresult{{}},
        

        Alias: (*Alias)(u),
    })
}

