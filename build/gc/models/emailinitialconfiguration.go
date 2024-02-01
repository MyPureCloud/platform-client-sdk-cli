package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailinitialconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailinitialconfigurationDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Emailinitialconfiguration
type Emailinitialconfiguration struct { 
    // To - An email address that this email is to.
    To string `json:"to"`


    // From - An email address that this email is from.
    From string `json:"from"`


    // Cc - An email addresses that this email is carbon copied to.
    Cc []string `json:"cc"`


    // Bcc - An email addresses that this email is blind carbon copied to.
    Bcc []string `json:"bcc"`


    // Subject - The subject for this email.
    Subject string `json:"subject"`


    // PreviousEmailId - UUID identifying the most recent previous email communication ID from the same participant on this email conversation. Will be null if this is a new participant.
    PreviousEmailId string `json:"previousEmailId"`


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
func (o *Emailinitialconfiguration) String() string {
    
    
     o.Cc = []string{""} 
     o.Bcc = []string{""} 
    
    
    
    
    
    
     o.AdditionalInfo = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailinitialconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Emailinitialconfiguration

    if EmailinitialconfigurationMarshalled {
        return []byte("{}"), nil
    }
    EmailinitialconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        To string `json:"to"`
        
        From string `json:"from"`
        
        Cc []string `json:"cc"`
        
        Bcc []string `json:"bcc"`
        
        Subject string `json:"subject"`
        
        PreviousEmailId string `json:"previousEmailId"`
        
        Held bool `json:"held"`
        
        Alerting bool `json:"alerting"`
        
        Inbound bool `json:"inbound"`
        
        InvitedBy string `json:"invitedBy"`
        
        AdditionalInfo map[string]string `json:"additionalInfo"`
        *Alias
    }{

        


        


        
        Cc: []string{""},
        


        
        Bcc: []string{""},
        


        


        


        


        


        


        


        
        AdditionalInfo: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

