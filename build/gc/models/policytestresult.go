package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicytestresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicytestresultDud struct { 
    


    


    


    

}

// Policytestresult
type Policytestresult struct { 
    // Id - The ID of the policy being tested.
    Id string `json:"id"`


    // Name - The name of the policy being tested.
    Name string `json:"name"`


    // Result - The result of the evaluation against supplied test data.
    Result string `json:"result"`


    // PolicyConditionResults - The results of conditions, with their boolean result.
    PolicyConditionResults []Policyconditionresult `json:"policyConditionResults"`

}

// String returns a JSON representation of the model
func (o *Policytestresult) String() string {
    
    
    
     o.PolicyConditionResults = []Policyconditionresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policytestresult) MarshalJSON() ([]byte, error) {
    type Alias Policytestresult

    if PolicytestresultMarshalled {
        return []byte("{}"), nil
    }
    PolicytestresultMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Result string `json:"result"`
        
        PolicyConditionResults []Policyconditionresult `json:"policyConditionResults"`
        *Alias
    }{

        


        


        


        
        PolicyConditionResults: []Policyconditionresult{{}},
        

        Alias: (*Alias)(u),
    })
}

