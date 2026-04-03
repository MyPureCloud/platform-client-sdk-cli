package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateshifttradestatejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateshifttradestatejobrequestDud struct { 
    


    


    

}

// Updateshifttradestatejobrequest
type Updateshifttradestatejobrequest struct { 
    // WeekDate - The start week date of this shift in the business unit time zone (yyyy-MM-dd format). Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // State - The new state to set on the shift trade
    State string `json:"state"`


    // Metadata - Version metadata for the shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateshifttradestatejobrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateshifttradestatejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateshifttradestatejobrequest

    if UpdateshifttradestatejobrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateshifttradestatejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate time.Time `json:"weekDate"`
        
        State string `json:"state"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

