package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestmoderesultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestmoderesultsDud struct { 
    


    


    


    

}

// Testmoderesults - Information about trigger test mode execution
type Testmoderesults struct { 
    // SchemaValidation - Information about the validation of the schema of the event body passed in to test mode
    SchemaValidation Testschemaoperation `json:"schemaValidation"`


    // TargetValidation - Information about the validation of the trigger target
    TargetValidation Testtargetoperation `json:"targetValidation"`


    // JsonPathValidation - Information about the json path matching criteria
    JsonPathValidation Testmatchesoperation `json:"jsonPathValidation"`


    // TriggerMatches - Whether the trigger would have matched on the provided event body
    TriggerMatches bool `json:"triggerMatches"`

}

// String returns a JSON representation of the model
func (o *Testmoderesults) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testmoderesults) MarshalJSON() ([]byte, error) {
    type Alias Testmoderesults

    if TestmoderesultsMarshalled {
        return []byte("{}"), nil
    }
    TestmoderesultsMarshalled = true

    return json.Marshal(&struct {
        
        SchemaValidation Testschemaoperation `json:"schemaValidation"`
        
        TargetValidation Testtargetoperation `json:"targetValidation"`
        
        JsonPathValidation Testmatchesoperation `json:"jsonPathValidation"`
        
        TriggerMatches bool `json:"triggerMatches"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

