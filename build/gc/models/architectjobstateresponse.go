package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectjobstateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectjobstateresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Architectjobstateresponse
type Architectjobstateresponse struct { 
    


    // Flow - Flow created from the Architect Job
    Flow Addressableentityref `json:"flow"`


    // Status - Status of the Architect Job
    Status string `json:"status"`


    // Command - The command executed by the Architect Job
    Command string `json:"command"`


    // Messages - Warnings and Errors messages of the Architect Job
    Messages []Architectjobmessage `json:"messages"`


    

}

// String returns a JSON representation of the model
func (o *Architectjobstateresponse) String() string {
    
    
    
     o.Messages = []Architectjobmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectjobstateresponse) MarshalJSON() ([]byte, error) {
    type Alias Architectjobstateresponse

    if ArchitectjobstateresponseMarshalled {
        return []byte("{}"), nil
    }
    ArchitectjobstateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Flow Addressableentityref `json:"flow"`
        
        Status string `json:"status"`
        
        Command string `json:"command"`
        
        Messages []Architectjobmessage `json:"messages"`
        *Alias
    }{

        


        


        


        


        
        Messages: []Architectjobmessage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

