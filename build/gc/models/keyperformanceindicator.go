package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KeyperformanceindicatorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KeyperformanceindicatorDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    OptimizationType string `json:"optimizationType"`


    ProblemType string `json:"problemType"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Description string `json:"description"`


    KpiType string `json:"kpiType"`


    Source string `json:"source"`


    WrapUpCodeConfig Wrapupcodeconfig `json:"wrapUpCodeConfig"`


    OutcomeConfig Outcomeconfig `json:"outcomeConfig"`


    Status string `json:"status"`


    KpiGroup string `json:"kpiGroup"`


    Queues []string `json:"queues"`


    SelfUri string `json:"selfUri"`

}

// Keyperformanceindicator
type Keyperformanceindicator struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Keyperformanceindicator) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Keyperformanceindicator) MarshalJSON() ([]byte, error) {
    type Alias Keyperformanceindicator

    if KeyperformanceindicatorMarshalled {
        return []byte("{}"), nil
    }
    KeyperformanceindicatorMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

