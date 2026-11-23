package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypedirectionbyobjectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypedirectionbyobjectDud struct { 
    

}

// Capabilitiesdatatypedirectionbyobject
type Capabilitiesdatatypedirectionbyobject struct { 
    // Flow
    Flow Capabilitiesdatatypeinputoutputinfo `json:"flow"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypedirectionbyobject) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypedirectionbyobject) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypedirectionbyobject

    if CapabilitiesdatatypedirectionbyobjectMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypedirectionbyobjectMarshalled = true

    return json.Marshal(&struct {
        
        Flow Capabilitiesdatatypeinputoutputinfo `json:"flow"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

