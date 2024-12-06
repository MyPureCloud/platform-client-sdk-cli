package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestsessionDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Knowledgeguestsession
type Knowledgeguestsession struct { 
    


    // App - The app where the session is started.
    App Knowledgeguestsessionapp `json:"app"`


    // CustomerId - An arbitrary ID for the customer starting the session. Used to track multiple sessions started by the same customer.
    CustomerId string `json:"customerId"`


    // PageUrl - URL of the page where the session is started.
    PageUrl string `json:"pageUrl"`


    // Contexts - The session contexts.
    Contexts []Knowledgeguestsessioncontext `json:"contexts"`


    // JourneySessionId - Journey session ID. Used to get the segments of the customer to filter search results.
    JourneySessionId string `json:"journeySessionId"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestsession) String() string {
    
    
    
     o.Contexts = []Knowledgeguestsessioncontext{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestsession) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestsession

    if KnowledgeguestsessionMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestsessionMarshalled = true

    return json.Marshal(&struct {
        
        App Knowledgeguestsessionapp `json:"app"`
        
        CustomerId string `json:"customerId"`
        
        PageUrl string `json:"pageUrl"`
        
        Contexts []Knowledgeguestsessioncontext `json:"contexts"`
        
        JourneySessionId string `json:"journeySessionId"`
        *Alias
    }{

        


        


        


        


        
        Contexts: []Knowledgeguestsessioncontext{{}},
        


        

        Alias: (*Alias)(u),
    })
}

