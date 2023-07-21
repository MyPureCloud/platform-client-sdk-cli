package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentsjourneycontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentsjourneycontextDud struct { 
    


    


    

}

// Webdeploymentsjourneycontext
type Webdeploymentsjourneycontext struct { 
    // JourneyAction - A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
    JourneyAction Journeyaction `json:"journeyAction"`


    // Customer - Journey customer information. Used for linking the authenticated customer with the journey. 
    Customer Journeycustomer `json:"customer"`


    // CustomerSession - Contains the Journey System's customer session details.
    CustomerSession Journeycustomersession `json:"customerSession"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsjourneycontext) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentsjourneycontext) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentsjourneycontext

    if WebdeploymentsjourneycontextMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentsjourneycontextMarshalled = true

    return json.Marshal(&struct {
        
        JourneyAction Journeyaction `json:"journeyAction"`
        
        Customer Journeycustomer `json:"customer"`
        
        CustomerSession Journeycustomersession `json:"customerSession"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

