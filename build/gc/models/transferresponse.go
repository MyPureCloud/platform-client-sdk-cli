package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransferresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransferresponseDud struct { 
    


    


    


    


    


    


    

}

// Transferresponse
type Transferresponse struct { 
    // Id - The id of the command.
    Id string `json:"id"`


    // State - The state of the command.
    State string `json:"state"`


    // DateIssued - The date/time that this command was issued. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateIssued time.Time `json:"dateIssued"`


    // Initiator - The initiator of the command.
    Initiator Transferinitiator `json:"initiator"`


    // ModifiedBy - The user or entity that modified the command.
    ModifiedBy Transferresponsemodifiedby `json:"modifiedBy"`


    // Destination - The destination of the command.
    Destination Transferdestination `json:"destination"`


    // TransferType - The type of transfer to perform.
    TransferType string `json:"transferType"`

}

// String returns a JSON representation of the model
func (o *Transferresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transferresponse) MarshalJSON() ([]byte, error) {
    type Alias Transferresponse

    if TransferresponseMarshalled {
        return []byte("{}"), nil
    }
    TransferresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        State string `json:"state"`
        
        DateIssued time.Time `json:"dateIssued"`
        
        Initiator Transferinitiator `json:"initiator"`
        
        ModifiedBy Transferresponsemodifiedby `json:"modifiedBy"`
        
        Destination Transferdestination `json:"destination"`
        
        TransferType string `json:"transferType"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

