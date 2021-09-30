package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopyattachmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopyattachmentsrequestDud struct { 
    


    

}

// Copyattachmentsrequest
type Copyattachmentsrequest struct { 
    // SourceMessage - A reference to the email message within the current conversation that owns the attachments to be copied
    SourceMessage Domainentityref `json:"sourceMessage"`


    // Attachments - A list of attachments that will be copied from the source message to the current draft
    Attachments []Attachment `json:"attachments"`

}

// String returns a JSON representation of the model
func (o *Copyattachmentsrequest) String() string {
    
    
    
    
    
    
     o.Attachments = []Attachment{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copyattachmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Copyattachmentsrequest

    if CopyattachmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    CopyattachmentsrequestMarshalled = true

    return json.Marshal(&struct { 
        SourceMessage Domainentityref `json:"sourceMessage"`
        
        Attachments []Attachment `json:"attachments"`
        
        *Alias
    }{
        

        

        

        
        Attachments: []Attachment{{}},
        

        
        Alias: (*Alias)(u),
    })
}

