package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsuploadschemaDud struct { 
    

}

// Decisionmetricsuploadschema
type Decisionmetricsuploadschema struct { 
    // UserMetrics - Decision metrics to be uploaded
    UserMetrics []Decisionmetricsuploaddata `json:"userMetrics"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsuploadschema) String() string {
     o.UserMetrics = []Decisionmetricsuploaddata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsuploadschema

    if DecisionmetricsuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsuploadschemaMarshalled = true

    return json.Marshal(&struct {
        
        UserMetrics []Decisionmetricsuploaddata `json:"userMetrics"`
        *Alias
    }{

        
        UserMetrics: []Decisionmetricsuploaddata{{}},
        

        Alias: (*Alias)(u),
    })
}

