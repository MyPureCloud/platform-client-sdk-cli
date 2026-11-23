package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesfuncopoverloadidmodelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesfuncopoverloadidmodelDud struct { 
    


    

}

// Capabilitiesfuncopoverloadidmodel
type Capabilitiesfuncopoverloadidmodel struct { 
    // FuncOp
    FuncOp string `json:"funcOp"`


    // OverloadId
    OverloadId string `json:"overloadId"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesfuncopoverloadidmodel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesfuncopoverloadidmodel) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesfuncopoverloadidmodel

    if CapabilitiesfuncopoverloadidmodelMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesfuncopoverloadidmodelMarshalled = true

    return json.Marshal(&struct {
        
        FuncOp string `json:"funcOp"`
        
        OverloadId string `json:"overloadId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

