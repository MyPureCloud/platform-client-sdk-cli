package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingjobsqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingjobsqueryDud struct { 
    


    


    


    


    

}

// Recordingjobsquery
type Recordingjobsquery struct { 
    // Action - Operation to perform bulk task
    Action string `json:"action"`


    // ActionDate - The date when the action will be performed. If the operation will cause the delete date of a recording to be older than the export date, the export date will be adjusted to the delete date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ActionDate time.Time `json:"actionDate"`


    // IntegrationId - IntegrationId to Access AWS S3 bucket for bulk recording exports. This field is required and used only for EXPORT action.
    IntegrationId string `json:"integrationId"`


    // IncludeScreenRecordings - Include Screen recordings for export action, default value = true 
    IncludeScreenRecordings bool `json:"includeScreenRecordings"`


    // ConversationQuery - Conversation Query. Note: After the recording is created, it might take up to 48 hours for the recording to be included in the submitted job query.
    ConversationQuery Asyncconversationquery `json:"conversationQuery"`

}

// String returns a JSON representation of the model
func (o *Recordingjobsquery) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingjobsquery) MarshalJSON() ([]byte, error) {
    type Alias Recordingjobsquery

    if RecordingjobsqueryMarshalled {
        return []byte("{}"), nil
    }
    RecordingjobsqueryMarshalled = true

    return json.Marshal(&struct { 
        Action string `json:"action"`
        
        ActionDate time.Time `json:"actionDate"`
        
        IntegrationId string `json:"integrationId"`
        
        IncludeScreenRecordings bool `json:"includeScreenRecordings"`
        
        ConversationQuery Asyncconversationquery `json:"conversationQuery"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}
