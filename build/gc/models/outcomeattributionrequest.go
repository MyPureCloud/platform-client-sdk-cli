package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeattributionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeattributionrequestDud struct { 
    


    


    


    


    

}

// Outcomeattributionrequest
type Outcomeattributionrequest struct { 
    // OutcomeId - ID of Outcome.
    OutcomeId string `json:"outcomeId"`


    // ExternalContactId - The external contact ID of the customer who achieved the outcome.
    ExternalContactId string `json:"externalContactId"`


    // AssociatedValue - The total value associated with the customer's outcome.
    AssociatedValue float32 `json:"associatedValue"`


    // Touchpoints - List of interactions that led to this outcome being achieved.
    Touchpoints []Touchpoint `json:"touchpoints"`


    // CreatedDate - Date outcome was achieved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Outcomeattributionrequest) String() string {
    
    
    
     o.Touchpoints = []Touchpoint{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeattributionrequest) MarshalJSON() ([]byte, error) {
    type Alias Outcomeattributionrequest

    if OutcomeattributionrequestMarshalled {
        return []byte("{}"), nil
    }
    OutcomeattributionrequestMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeId string `json:"outcomeId"`
        
        ExternalContactId string `json:"externalContactId"`
        
        AssociatedValue float32 `json:"associatedValue"`
        
        Touchpoints []Touchpoint `json:"touchpoints"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        
        Touchpoints: []Touchpoint{{}},
        


        

        Alias: (*Alias)(u),
    })
}

