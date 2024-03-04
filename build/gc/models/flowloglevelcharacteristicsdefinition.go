package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowloglevelcharacteristicsdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowloglevelcharacteristicsdefinitionDud struct { 
    Id string `json:"id"`


    MinimumLevel string `json:"minimumLevel"`


    DependsOn Flowcharacteristics `json:"dependsOn"`

}

// Flowloglevelcharacteristicsdefinition - Defines a characteristic that can be captured by data providers
type Flowloglevelcharacteristicsdefinition struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Flowloglevelcharacteristicsdefinition) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowloglevelcharacteristicsdefinition) MarshalJSON() ([]byte, error) {
    type Alias Flowloglevelcharacteristicsdefinition

    if FlowloglevelcharacteristicsdefinitionMarshalled {
        return []byte("{}"), nil
    }
    FlowloglevelcharacteristicsdefinitionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

