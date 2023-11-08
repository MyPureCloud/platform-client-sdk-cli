package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActioneventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActioneventrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Actioneventrequest
type Actioneventrequest struct { 
    


    // SessionId - UUID of the customer session for this action.
    SessionId string `json:"sessionId"`


    // ActionId - UUID for the action, as returned by the Ping endpoint when the action was qualified.
    ActionId string `json:"actionId"`


    // ActionState - State the action is transitioning to.
    ActionState string `json:"actionState"`


    // ErrorCode - Client defined error code (when state transitions to errored)
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - Message of the error returned when the action fails (when state transitions to errored)
    ErrorMessage string `json:"errorMessage"`


    

}

// String returns a JSON representation of the model
func (o *Actioneventrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actioneventrequest) MarshalJSON() ([]byte, error) {
    type Alias Actioneventrequest

    if ActioneventrequestMarshalled {
        return []byte("{}"), nil
    }
    ActioneventrequestMarshalled = true

    return json.Marshal(&struct {
        
        SessionId string `json:"sessionId"`
        
        ActionId string `json:"actionId"`
        
        ActionState string `json:"actionState"`
        
        ErrorCode string `json:"errorCode"`
        
        ErrorMessage string `json:"errorMessage"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

