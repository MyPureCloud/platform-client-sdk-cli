package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnactionDud struct { 
    


    


    


    

}

// Reportingturnaction
type Reportingturnaction struct { 
    // ActionId - The ID of the action in the bot flow.
    ActionId string `json:"actionId"`


    // ActionName - The name of the action in the bot flow.
    ActionName string `json:"actionName"`


    // ActionNumber - The number of the action in the bot flow.
    ActionNumber int `json:"actionNumber"`


    // ActionType
    ActionType string `json:"actionType"`

}

// String returns a JSON representation of the model
func (o *Reportingturnaction) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnaction) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnaction

    if ReportingturnactionMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnactionMarshalled = true

    return json.Marshal(&struct {
        
        ActionId string `json:"actionId"`
        
        ActionName string `json:"actionName"`
        
        ActionNumber int `json:"actionNumber"`
        
        ActionType string `json:"actionType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

