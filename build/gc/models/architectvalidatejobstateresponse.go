package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectvalidatejobstateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectvalidatejobstateresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Architectvalidatejobstateresponse
type Architectvalidatejobstateresponse struct { 
    


    // Status - Status of the Architect Validate Job
    Status string `json:"status"`


    // Command - The command executed by the Architect Job
    Command string `json:"command"`


    // Messages - Warnings and Errors messages of the Architect Job
    Messages []Architectjobmessage `json:"messages"`


    

}

// String returns a JSON representation of the model
func (o *Architectvalidatejobstateresponse) String() string {
    
    
     o.Messages = []Architectjobmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectvalidatejobstateresponse) MarshalJSON() ([]byte, error) {
    type Alias Architectvalidatejobstateresponse

    if ArchitectvalidatejobstateresponseMarshalled {
        return []byte("{}"), nil
    }
    ArchitectvalidatejobstateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Command string `json:"command"`
        
        Messages []Architectjobmessage `json:"messages"`
        *Alias
    }{

        


        


        


        
        Messages: []Architectjobmessage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

