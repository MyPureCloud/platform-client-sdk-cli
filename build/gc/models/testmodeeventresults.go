package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestmodeeventresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestmodeeventresultsDud struct { 
    


    

}

// Testmodeeventresults - Information about event test mode execution
type Testmodeeventresults struct { 
    // SchemaValidation - Information about the validation of the schema of the event body passed in to test mode
    SchemaValidation Testschemaoperation `json:"schemaValidation"`


    // TriggerMatchValidation - Information about matched and unmatched triggers
    TriggerMatchValidation Testmatcheseventoperation `json:"triggerMatchValidation"`

}

// String returns a JSON representation of the model
func (o *Testmodeeventresults) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testmodeeventresults) MarshalJSON() ([]byte, error) {
    type Alias Testmodeeventresults

    if TestmodeeventresultsMarshalled {
        return []byte("{}"), nil
    }
    TestmodeeventresultsMarshalled = true

    return json.Marshal(&struct {
        
        SchemaValidation Testschemaoperation `json:"schemaValidation"`
        
        TriggerMatchValidation Testmatcheseventoperation `json:"triggerMatchValidation"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

