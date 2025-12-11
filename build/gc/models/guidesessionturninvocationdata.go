package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturninvocationdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturninvocationdataDud struct { 
    


    


    

}

// Guidesessionturninvocationdata
type Guidesessionturninvocationdata struct { 
    // Group - The action group of the invocation data.
    Group string `json:"group"`


    // ActionName - The action name of the invocation data.
    ActionName string `json:"actionName"`


    // Output - The output of the invocation data.
    Output string `json:"output"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturninvocationdata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturninvocationdata) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturninvocationdata

    if GuidesessionturninvocationdataMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturninvocationdataMarshalled = true

    return json.Marshal(&struct {
        
        Group string `json:"group"`
        
        ActionName string `json:"actionName"`
        
        Output string `json:"output"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

