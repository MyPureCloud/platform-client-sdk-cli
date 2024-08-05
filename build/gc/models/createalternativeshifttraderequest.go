package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatealternativeshifttraderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatealternativeshifttraderequestDud struct { 
    


    


    


    


    

}

// Createalternativeshifttraderequest
type Createalternativeshifttraderequest struct { 
    // JobId - The ID of this alternative shift job
    JobId string `json:"jobId"`


    // DropShiftReferenceKeys - A list of offered shift reference keys an agent wants to drop
    DropShiftReferenceKeys []string `json:"dropShiftReferenceKeys"`


    // PickupShiftReferenceKeys - A list of offered shift reference keys an agent wants to pick up
    PickupShiftReferenceKeys []string `json:"pickupShiftReferenceKeys"`


    // AlternativeShiftTradeGranularity - The granularity of alternative shifts to be traded
    AlternativeShiftTradeGranularity string `json:"alternativeShiftTradeGranularity"`


    // ExpirationDate - The date when the trade will expire in ISO-8601 format. The trade cannot be approved after expiration
    ExpirationDate time.Time `json:"expirationDate"`

}

// String returns a JSON representation of the model
func (o *Createalternativeshifttraderequest) String() string {
    
     o.DropShiftReferenceKeys = []string{""} 
     o.PickupShiftReferenceKeys = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createalternativeshifttraderequest) MarshalJSON() ([]byte, error) {
    type Alias Createalternativeshifttraderequest

    if CreatealternativeshifttraderequestMarshalled {
        return []byte("{}"), nil
    }
    CreatealternativeshifttraderequestMarshalled = true

    return json.Marshal(&struct {
        
        JobId string `json:"jobId"`
        
        DropShiftReferenceKeys []string `json:"dropShiftReferenceKeys"`
        
        PickupShiftReferenceKeys []string `json:"pickupShiftReferenceKeys"`
        
        AlternativeShiftTradeGranularity string `json:"alternativeShiftTradeGranularity"`
        
        ExpirationDate time.Time `json:"expirationDate"`
        *Alias
    }{

        


        
        DropShiftReferenceKeys: []string{""},
        


        
        PickupShiftReferenceKeys: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

