package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturninvocationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturninvocationresponseDud struct { 
    


    


    


    

}

// Guidesessionturninvocationresponse
type Guidesessionturninvocationresponse struct { 
    // Id - The action ID of the invocation data.
    Id string `json:"id"`


    // Group - The action group of the invocation data.
    Group string `json:"group"`


    // ActionName - The action name of the invocation data.
    ActionName string `json:"actionName"`


    // Parameters - The parameters of the invocation response.
    Parameters []Guidesessionturninvocationparameters `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturninvocationresponse) String() string {
    
    
    
     o.Parameters = []Guidesessionturninvocationparameters{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturninvocationresponse) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturninvocationresponse

    if GuidesessionturninvocationresponseMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturninvocationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Group string `json:"group"`
        
        ActionName string `json:"actionName"`
        
        Parameters []Guidesessionturninvocationparameters `json:"parameters"`
        *Alias
    }{

        


        


        


        
        Parameters: []Guidesessionturninvocationparameters{{}},
        

        Alias: (*Alias)(u),
    })
}

