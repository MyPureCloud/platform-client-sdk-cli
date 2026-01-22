package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsearchcriteriadtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsearchcriteriadtoDud struct { 
    


    


    


    


    


    


    

}

// Evaluationsearchcriteriadto
type Evaluationsearchcriteriadto struct { 
    // VarType - The type of query Operation to Perform
    VarType string `json:"type"`


    // Field - Field name to search against
    Field string `json:"field"`


    // EndValue - The end value of the range. This field is used for range search types. Date values must be in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
    EndValue string `json:"endValue"`


    // Values - A list of values for the search to match against. Only for Type: EXACT
    Values []string `json:"values"`


    // StartValue - The start value of the range. This field is used for range search types. Date values must be in the format yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
    StartValue string `json:"startValue"`


    // Value - A value for the search to match against
    Value string `json:"value"`


    // Operator - How to apply this search criteria against other criteria
    Operator string `json:"operator"`

}

// String returns a JSON representation of the model
func (o *Evaluationsearchcriteriadto) String() string {
    
    
    
     o.Values = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsearchcriteriadto) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsearchcriteriadto

    if EvaluationsearchcriteriadtoMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsearchcriteriadtoMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Field string `json:"field"`
        
        EndValue string `json:"endValue"`
        
        Values []string `json:"values"`
        
        StartValue string `json:"startValue"`
        
        Value string `json:"value"`
        
        Operator string `json:"operator"`
        *Alias
    }{

        


        


        


        
        Values: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

