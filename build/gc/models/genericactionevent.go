package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenericactioneventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenericactioneventDud struct { 
    


    


    


    

}

// Genericactionevent
type Genericactionevent struct { 
    // Action - The action that triggered the event.
    Action Genericeventaction `json:"action"`


    // ActionMap - The action map that triggered the action.
    ActionMap Actioneventactionmap `json:"actionMap"`


    // ErrorCode - Code of the error returned when the action fails.
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - Message of the error returned when the action fails.
    ErrorMessage string `json:"errorMessage"`

}

// String returns a JSON representation of the model
func (o *Genericactionevent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Genericactionevent) MarshalJSON() ([]byte, error) {
    type Alias Genericactionevent

    if GenericactioneventMarshalled {
        return []byte("{}"), nil
    }
    GenericactioneventMarshalled = true

    return json.Marshal(&struct {
        
        Action Genericeventaction `json:"action"`
        
        ActionMap Actioneventactionmap `json:"actionMap"`
        
        ErrorCode string `json:"errorCode"`
        
        ErrorMessage string `json:"errorMessage"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

