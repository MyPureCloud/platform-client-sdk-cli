package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateshifttradejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateshifttradejobrequestDud struct { 
    


    


    


    


    

}

// Updateshifttradejobrequest
type Updateshifttradejobrequest struct { 
    // WeekDate - The start week date of this shift in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // Target - Target of the shift trade, if applicable
    Target Valuewrappershifttradetargetrequestitem `json:"target"`


    // ExpirationDate - When this shift trade will expire. Date time is represented as an ISO-8601 string
    ExpirationDate Valuewrapperdate `json:"expirationDate"`


    // AcceptableIntervals - Time frames when the initiating user is willing to accept a shift in exchange. Setting the enclosed list to empty will make this a one sided trade request.
    AcceptableIntervals Listwrapperrequireddaterange `json:"acceptableIntervals"`


    // Metadata - Version metadata for the shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateshifttradejobrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateshifttradejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateshifttradejobrequest

    if UpdateshifttradejobrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateshifttradejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        Target Valuewrappershifttradetargetrequestitem `json:"target"`
        
        ExpirationDate Valuewrapperdate `json:"expirationDate"`
        
        AcceptableIntervals Listwrapperrequireddaterange `json:"acceptableIntervals"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

