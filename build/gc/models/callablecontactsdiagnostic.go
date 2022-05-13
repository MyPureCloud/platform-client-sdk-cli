package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallablecontactsdiagnosticMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallablecontactsdiagnosticDud struct { 
    AttemptLimits Domainentityref `json:"attemptLimits"`


    DncLists []Domainentityref `json:"dncLists"`


    CallableTimeSet Domainentityref `json:"callableTimeSet"`


    RuleSets []Domainentityref `json:"ruleSets"`

}

// Callablecontactsdiagnostic
type Callablecontactsdiagnostic struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Callablecontactsdiagnostic) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callablecontactsdiagnostic) MarshalJSON() ([]byte, error) {
    type Alias Callablecontactsdiagnostic

    if CallablecontactsdiagnosticMarshalled {
        return []byte("{}"), nil
    }
    CallablecontactsdiagnosticMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

