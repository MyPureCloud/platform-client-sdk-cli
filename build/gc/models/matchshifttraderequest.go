package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MatchshifttraderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MatchshifttraderequestDud struct { 
    


    


    

}

// Matchshifttraderequest
type Matchshifttraderequest struct { 
    // ReceivingScheduleId - The ID of the schedule with which the shift trade is associated
    ReceivingScheduleId string `json:"receivingScheduleId"`


    // ReceivingShiftId - The ID of the shift the receiving user is giving up in trade, if applicable
    ReceivingShiftId string `json:"receivingShiftId"`


    // Metadata - Version metadata for the shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Matchshifttraderequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Matchshifttraderequest) MarshalJSON() ([]byte, error) {
    type Alias Matchshifttraderequest

    if MatchshifttraderequestMarshalled {
        return []byte("{}"), nil
    }
    MatchshifttraderequestMarshalled = true

    return json.Marshal(&struct { 
        ReceivingScheduleId string `json:"receivingScheduleId"`
        
        ReceivingShiftId string `json:"receivingShiftId"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

