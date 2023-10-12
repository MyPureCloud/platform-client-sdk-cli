package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeattributionresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeattributionresultsresponseDud struct { 
    


    


    


    


    


    


    


    

}

// Outcomeattributionresultsresponse
type Outcomeattributionresultsresponse struct { 
    // OutcomeId - ID of Outcome.
    OutcomeId string `json:"outcomeId"`


    // Index - The index/position of the OutcomeAttribution in the original POST request.
    Index int `json:"index"`


    // ExternalContactId - The external contact ID of the customer who achieved the outcome.
    ExternalContactId string `json:"externalContactId"`


    // AssociatedValue - The total value associated with the customer's outcome.
    AssociatedValue float32 `json:"associatedValue"`


    // State - State of the Outcome Attribution entity.
    State string `json:"state"`


    // Message - Additional information on the state of the Outcome Attribution entity.
    Message string `json:"message"`


    // Touchpoints - List of interactions that led to this outcome being achieved.
    Touchpoints []Touchpointresponse `json:"touchpoints"`


    // CreatedDate - Date outcome was achieved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Outcomeattributionresultsresponse) String() string {
    
    
    
    
    
    
     o.Touchpoints = []Touchpointresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeattributionresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Outcomeattributionresultsresponse

    if OutcomeattributionresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    OutcomeattributionresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeId string `json:"outcomeId"`
        
        Index int `json:"index"`
        
        ExternalContactId string `json:"externalContactId"`
        
        AssociatedValue float32 `json:"associatedValue"`
        
        State string `json:"state"`
        
        Message string `json:"message"`
        
        Touchpoints []Touchpointresponse `json:"touchpoints"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        
        Touchpoints: []Touchpointresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

