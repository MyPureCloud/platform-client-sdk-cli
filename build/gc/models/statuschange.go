package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StatuschangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StatuschangeDud struct { 
    DateStatusChanged time.Time `json:"dateStatusChanged"`


    Status string `json:"status"`


    PreviousStatus string `json:"previousStatus"`


    Message string `json:"message"`


    ChangedBy string `json:"changedBy"`


    RejectReason string `json:"rejectReason"`

}

// Statuschange
type Statuschange struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Statuschange) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Statuschange) MarshalJSON() ([]byte, error) {
    type Alias Statuschange

    if StatuschangeMarshalled {
        return []byte("{}"), nil
    }
    StatuschangeMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

