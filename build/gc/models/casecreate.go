package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasecreateDud struct { 
    


    


    


    


    


    


    


    

}

// Casecreate
type Casecreate struct { 
    // CaseplanId - The ID of the caseplan to create the case from.
    CaseplanId string `json:"caseplanId"`


    // OwnerId - The ID of the owner of the case.
    OwnerId string `json:"ownerId"`


    // Summary - Overview information for the Case. Valid length between 3 and 512 characters.
    Summary string `json:"summary"`


    // ExternalContactId - The ID of the External Contact associated with the Case.
    ExternalContactId string `json:"externalContactId"`


    // ConversationId - The ID of conversation associated with the Case.
    ConversationId string `json:"conversationId"`


    // WorkitemId - The ID of the workitem associated with the Case.
    WorkitemId string `json:"workitemId"`


    // TtlSeconds - The epoch timestamp in seconds specifying the time-to-live for the lifetime of the Case. Can not be greater than 365 days from the current time.
    TtlSeconds int `json:"ttlSeconds"`


    // Intake - The intake data for the Case. Maximum of 10 intake objects allowed.
    Intake []Intake `json:"intake"`

}

// String returns a JSON representation of the model
func (o *Casecreate) String() string {
    
    
    
    
    
    
    
     o.Intake = []Intake{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casecreate) MarshalJSON() ([]byte, error) {
    type Alias Casecreate

    if CasecreateMarshalled {
        return []byte("{}"), nil
    }
    CasecreateMarshalled = true

    return json.Marshal(&struct {
        
        CaseplanId string `json:"caseplanId"`
        
        OwnerId string `json:"ownerId"`
        
        Summary string `json:"summary"`
        
        ExternalContactId string `json:"externalContactId"`
        
        ConversationId string `json:"conversationId"`
        
        WorkitemId string `json:"workitemId"`
        
        TtlSeconds int `json:"ttlSeconds"`
        
        Intake []Intake `json:"intake"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Intake: []Intake{{}},
        

        Alias: (*Alias)(u),
    })
}

