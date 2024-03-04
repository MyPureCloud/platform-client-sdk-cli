package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessaginginitialconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessaginginitialconfigurationDud struct { 
    


    


    


    


    


    


    


    

}

// Messaginginitialconfiguration
type Messaginginitialconfiguration struct { 
    // ToAddress - Address for the participant on receiving side of the message conversation. If the address is a phone number, E.164 format is recommended.
    ToAddress string `json:"toAddress"`


    // FromAddress - Address for the participant on the sending side of the message conversation. If the address is a phone number, E.164 format is recommended.
    FromAddress string `json:"fromAddress"`


    // MessageType - The type of message platform from which the message originated.
    MessageType string `json:"messageType"`


    // Held - Indicates that this communication's initial state is held.
    Held bool `json:"held"`


    // Alerting - Indicates that this communication's initial state is alerting. If false, the communication started in a connected state.
    Alerting bool `json:"alerting"`


    // Inbound - Indicates the direction of this communication with respect to the contact center. `true` means the communication is INBOUND. `false` means the communication is OUTBOUND.
    Inbound bool `json:"inbound"`


    // InvitedBy - The id of the communication (the \"peer\") that \"invited\" this communication, if this occurred.
    InvitedBy string `json:"invitedBy"`


    // AdditionalInfo - Additional metadata about this session which should be recorded by the platform but which will not be indexed or searchable. Primarily for diagnostic value. Any information that needs to be accessible through other components like Analytics should be moved to dedicated fields.
    AdditionalInfo map[string]string `json:"additionalInfo"`

}

// String returns a JSON representation of the model
func (o *Messaginginitialconfiguration) String() string {
    
    
    
    
    
    
    
     o.AdditionalInfo = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messaginginitialconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Messaginginitialconfiguration

    if MessaginginitialconfigurationMarshalled {
        return []byte("{}"), nil
    }
    MessaginginitialconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        ToAddress string `json:"toAddress"`
        
        FromAddress string `json:"fromAddress"`
        
        MessageType string `json:"messageType"`
        
        Held bool `json:"held"`
        
        Alerting bool `json:"alerting"`
        
        Inbound bool `json:"inbound"`
        
        InvitedBy string `json:"invitedBy"`
        
        AdditionalInfo map[string]string `json:"additionalInfo"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        AdditionalInfo: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

