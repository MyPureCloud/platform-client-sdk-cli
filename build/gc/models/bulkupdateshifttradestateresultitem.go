package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateshifttradestateresultitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateshifttradestateresultitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Bulkupdateshifttradestateresultitem
type Bulkupdateshifttradestateresultitem struct { 
    


    // State - The state of the shift trade after the update request is processed
    State string `json:"state"`


    // ReviewedBy - The user who reviewed the request, if applicable
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The date the request was reviewed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReviewedDate time.Time `json:"reviewedDate"`


    // FailureReason - The reason the update failed, if applicable
    FailureReason string `json:"failureReason"`


    // Metadata - Version metadata for the shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresultitem) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateshifttradestateresultitem) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateshifttradestateresultitem

    if BulkupdateshifttradestateresultitemMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateshifttradestateresultitemMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        
        FailureReason string `json:"failureReason"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

