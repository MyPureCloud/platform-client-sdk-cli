package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneycontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneycontextDud struct { 
    


    


    

}

// Journeycontext
type Journeycontext struct { 
    // Customer - A subset of the Journey System's customer data at a point-in-time (for external linkage and internal usage/context)
    Customer Journeycustomer `json:"customer"`


    // CustomerSession - A subset of the Journey System's tracked customer session data at a point-in-time (for external linkage and internal usage/context)
    CustomerSession Journeycustomersession `json:"customerSession"`


    // TriggeringAction - A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
    TriggeringAction Journeyaction `json:"triggeringAction"`

}

// String returns a JSON representation of the model
func (o *Journeycontext) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeycontext) MarshalJSON() ([]byte, error) {
    type Alias Journeycontext

    if JourneycontextMarshalled {
        return []byte("{}"), nil
    }
    JourneycontextMarshalled = true

    return json.Marshal(&struct { 
        Customer Journeycustomer `json:"customer"`
        
        CustomerSession Journeycustomersession `json:"customerSession"`
        
        TriggeringAction Journeyaction `json:"triggeringAction"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

