package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransfersfulltimeequivalentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransfersfulltimeequivalentDud struct { 
    


    

}

// Transfersfulltimeequivalent
type Transfersfulltimeequivalent struct { 
    // DestinationStaffingGroupId - The target staffing group that will receive the full time equivalent when agents are transferred from one staffing group to another
    DestinationStaffingGroupId string `json:"destinationStaffingGroupId"`


    // TransferType - The duration of the transfer full time equivalent from one staffing group to another
    TransferType string `json:"transferType"`

}

// String returns a JSON representation of the model
func (o *Transfersfulltimeequivalent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transfersfulltimeequivalent) MarshalJSON() ([]byte, error) {
    type Alias Transfersfulltimeequivalent

    if TransfersfulltimeequivalentMarshalled {
        return []byte("{}"), nil
    }
    TransfersfulltimeequivalentMarshalled = true

    return json.Marshal(&struct {
        
        DestinationStaffingGroupId string `json:"destinationStaffingGroupId"`
        
        TransferType string `json:"transferType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

