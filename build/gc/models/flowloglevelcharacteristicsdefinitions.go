package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowloglevelcharacteristicsdefinitionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowloglevelcharacteristicsdefinitionsDud struct { 
    

}

// Flowloglevelcharacteristicsdefinitions - A set of available characteristic definitions
type Flowloglevelcharacteristicsdefinitions struct { 
    // Characteristics
    Characteristics []Flowloglevelcharacteristicsdefinition `json:"characteristics"`

}

// String returns a JSON representation of the model
func (o *Flowloglevelcharacteristicsdefinitions) String() string {
     o.Characteristics = []Flowloglevelcharacteristicsdefinition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowloglevelcharacteristicsdefinitions) MarshalJSON() ([]byte, error) {
    type Alias Flowloglevelcharacteristicsdefinitions

    if FlowloglevelcharacteristicsdefinitionsMarshalled {
        return []byte("{}"), nil
    }
    FlowloglevelcharacteristicsdefinitionsMarshalled = true

    return json.Marshal(&struct {
        
        Characteristics []Flowloglevelcharacteristicsdefinition `json:"characteristics"`
        *Alias
    }{

        
        Characteristics: []Flowloglevelcharacteristicsdefinition{{}},
        

        Alias: (*Alias)(u),
    })
}

