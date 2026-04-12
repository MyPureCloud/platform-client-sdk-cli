package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnflowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnflowDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Reportingturnflow
type Reportingturnflow struct { 
    


    // VarType - The type of the flow
    VarType string `json:"type"`


    // Name - The name of the flow
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Reportingturnflow) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnflow) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnflow

    if ReportingturnflowMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnflowMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

