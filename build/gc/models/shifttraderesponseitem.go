package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttraderesponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttraderesponseitemDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Shifttraderesponseitem
type Shifttraderesponseitem struct { 
    // Id - The ID of this shift trade
    Id string `json:"id"`


    // State - The state of this shift trade
    State string `json:"state"`


    // ExpirationDate - When this shift trade will expire. Date time is represented as an ISO-8601 string
    ExpirationDate time.Time `json:"expirationDate"`


    // Initiating - Details about the initiating user involved in this shift trade
    Initiating Shifttradeinitiatingsideresponseitem `json:"initiating"`


    // Receiving - Details about the receiving user involved in this shift trade
    Receiving Shifttradereceivingsideresponseitem `json:"receiving"`


    // AcceptableIntervals - Time frames when the initiating user is willing to accept trades. Empty means giving up the shift
    AcceptableIntervals []Requireddaterange `json:"acceptableIntervals"`


    // OneSided - Whether this is a one-sided shift trade (e.g. the initiating user is not asking for a shift in return)
    OneSided bool `json:"oneSided"`


    // Target - The user to whom the shift trade request was sent in a direct trade, or the user with whom a shift trade was Matched
    Target Shifttradetargetresponseitem `json:"target"`


    // ReviewedBy - The admin who approved or denied this shift trade
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The timestamp of when the trade request was reviewed by an admin in ISO-8601 format
    ReviewedDate time.Time `json:"reviewedDate"`


    // Metadata - Version metadata for this shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Shifttraderesponseitem) String() string {
    
    
    
    
    
     o.AcceptableIntervals = []Requireddaterange{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttraderesponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttraderesponseitem

    if ShifttraderesponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttraderesponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        State string `json:"state"`
        
        ExpirationDate time.Time `json:"expirationDate"`
        
        Initiating Shifttradeinitiatingsideresponseitem `json:"initiating"`
        
        Receiving Shifttradereceivingsideresponseitem `json:"receiving"`
        
        AcceptableIntervals []Requireddaterange `json:"acceptableIntervals"`
        
        OneSided bool `json:"oneSided"`
        
        Target Shifttradetargetresponseitem `json:"target"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        
        AcceptableIntervals: []Requireddaterange{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

