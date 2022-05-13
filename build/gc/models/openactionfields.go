package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenactionfieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenactionfieldsDud struct { 
    


    

}

// Openactionfields
type Openactionfields struct { 
    // OpenAction - The specific type of the open action.
    OpenAction Domainentityref `json:"openAction"`


    // ConfigurationFields - Custom fields defined in the schema referenced by the open action type selected.
    ConfigurationFields map[string]interface{} `json:"configurationFields"`

}

// String returns a JSON representation of the model
func (o *Openactionfields) String() string {
    
     o.ConfigurationFields = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openactionfields) MarshalJSON() ([]byte, error) {
    type Alias Openactionfields

    if OpenactionfieldsMarshalled {
        return []byte("{}"), nil
    }
    OpenactionfieldsMarshalled = true

    return json.Marshal(&struct {
        
        OpenAction Domainentityref `json:"openAction"`
        
        ConfigurationFields map[string]interface{} `json:"configurationFields"`
        *Alias
    }{

        


        
        ConfigurationFields: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

