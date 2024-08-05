package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshifttradebulkupdatetemplateitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshifttradebulkupdatetemplateitemDud struct { 
    


    


    


    


    


    

}

// Alternativeshifttradebulkupdatetemplateitem
type Alternativeshifttradebulkupdatetemplateitem struct { 
    // TradeId - The ID of this alternative shift trade
    TradeId string `json:"tradeId"`


    // State - The current state of this alternative shift trade request
    State string `json:"state"`


    // FailureReason - The reason the update failed, if applicable
    FailureReason string `json:"failureReason"`


    // AdminDateReviewed - The timestamp of when the trade request was manually reviewed by an admin in ISO-8601 format
    AdminDateReviewed time.Time `json:"adminDateReviewed"`


    // AdminReviewedBy - The admin who manually reviewed this alternative shift trade after system denial
    AdminReviewedBy Userreference `json:"adminReviewedBy"`


    // Metadata - Version metadata for this alternative shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Alternativeshifttradebulkupdatetemplateitem) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshifttradebulkupdatetemplateitem) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshifttradebulkupdatetemplateitem

    if AlternativeshifttradebulkupdatetemplateitemMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshifttradebulkupdatetemplateitemMarshalled = true

    return json.Marshal(&struct {
        
        TradeId string `json:"tradeId"`
        
        State string `json:"state"`
        
        FailureReason string `json:"failureReason"`
        
        AdminDateReviewed time.Time `json:"adminDateReviewed"`
        
        AdminReviewedBy Userreference `json:"adminReviewedBy"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

