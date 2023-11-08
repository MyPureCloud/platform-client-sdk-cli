package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenactionpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenactionpropertiesDud struct { 
    


    

}

// Openactionproperties
type Openactionproperties struct { 
    // OpenActionName - The specific type of the open action.
    OpenActionName string `json:"openActionName"`


    // ConfigurationFields - Custom fields defined in the schema referenced by the open action type selected.
    ConfigurationFields map[string]interface{} `json:"configurationFields"`

}

// String returns a JSON representation of the model
func (o *Openactionproperties) String() string {
    
     o.ConfigurationFields = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openactionproperties) MarshalJSON() ([]byte, error) {
    type Alias Openactionproperties

    if OpenactionpropertiesMarshalled {
        return []byte("{}"), nil
    }
    OpenactionpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        OpenActionName string `json:"openActionName"`
        
        ConfigurationFields map[string]interface{} `json:"configurationFields"`
        *Alias
    }{

        


        
        ConfigurationFields: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

