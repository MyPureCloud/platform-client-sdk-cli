package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MatchshifttradejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MatchshifttradejobrequestDud struct { 
    


    


    


    

}

// Matchshifttradejobrequest
type Matchshifttradejobrequest struct { 
    // WeekDate - The start week date of the initiating shift in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // ReceivingSchedule - Associated schedule information for the receiving user
    ReceivingSchedule Receivingschedulelookup `json:"receivingSchedule"`


    // ReceivingShiftId - The ID of the shift the receiving user is giving up in trade, if applicable
    ReceivingShiftId string `json:"receivingShiftId"`


    // Metadata - Version metadata for the shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Matchshifttradejobrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Matchshifttradejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Matchshifttradejobrequest

    if MatchshifttradejobrequestMarshalled {
        return []byte("{}"), nil
    }
    MatchshifttradejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        ReceivingSchedule Receivingschedulelookup `json:"receivingSchedule"`
        
        ReceivingShiftId string `json:"receivingShiftId"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

