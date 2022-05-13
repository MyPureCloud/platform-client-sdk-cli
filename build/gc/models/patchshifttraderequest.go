package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchshifttraderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchshifttraderequestDud struct { 
    


    


    


    

}

// Patchshifttraderequest
type Patchshifttraderequest struct { 
    // ReceivingUserId - Update the ID of the receiving user to direct the request at a specific user, or set the wrapped id to null to open up a trade to be matched by any user.
    ReceivingUserId Valuewrapperstring `json:"receivingUserId"`


    // Expiration - Update the expiration time for this shift trade.
    Expiration Valuewrapperdate `json:"expiration"`


    // AcceptableIntervals - Update the acceptable intervals the initiating user is willing to accept in trade. Setting the enclosed list to empty will make this a one sided trade request
    AcceptableIntervals Listwrapperinterval `json:"acceptableIntervals"`


    // Metadata - Version metadata
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Patchshifttraderequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchshifttraderequest) MarshalJSON() ([]byte, error) {
    type Alias Patchshifttraderequest

    if PatchshifttraderequestMarshalled {
        return []byte("{}"), nil
    }
    PatchshifttraderequestMarshalled = true

    return json.Marshal(&struct {
        
        ReceivingUserId Valuewrapperstring `json:"receivingUserId"`
        
        Expiration Valuewrapperdate `json:"expiration"`
        
        AcceptableIntervals Listwrapperinterval `json:"acceptableIntervals"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

