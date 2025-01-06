package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatesessionquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatesessionquerypredicateDud struct { 
    


    

}

// Agentstatesessionquerypredicate
type Agentstatesessionquerypredicate struct { 
    // Dimension - Left hand side for dimension predicates
    Dimension string `json:"dimension"`


    // Value - Right hand side for dimension predicates
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Agentstatesessionquerypredicate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatesessionquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Agentstatesessionquerypredicate

    if AgentstatesessionquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    AgentstatesessionquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

