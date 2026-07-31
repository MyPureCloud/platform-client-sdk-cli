package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsupdateerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsupdateerrorDud struct { 
    


    

}

// Decisionmetricsupdateerror
type Decisionmetricsupdateerror struct { 
    // User - The user for the decision metrics row where errors were found
    User Userreference `json:"user"`


    // Errors - Errors found during the update process
    Errors []string `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsupdateerror) String() string {
    
     o.Errors = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsupdateerror) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsupdateerror

    if DecisionmetricsupdateerrorMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsupdateerrorMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Errors []string `json:"errors"`
        *Alias
    }{

        


        
        Errors: []string{""},
        

        Alias: (*Alias)(u),
    })
}

