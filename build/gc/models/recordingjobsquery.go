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
    // Action - Operation to perform bulk task. The date when the action will be performed can either be specified as an absolute date for all recordings with the actionDate/screenRecordingActionDate parameters, or as the number of days after each recording's creation time with the actionAge/screenRecordingActionAge parameters. If the operation will cause the delete date of a recording to be older than the export date, the export date will be adjusted to the delete date.
    Action string `json:"action"`


    // ActionDate - The date when the action will be performed. If screenRecordingActionDate is also provided, this value is only used for non-screen recordings. Otherwise this value is used for all recordings. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ActionDate time.Time `json:"actionDate"`


    // IntegrationId - IntegrationId to Access AWS S3 bucket for bulk recording exports. This field is required and used only for EXPORT action.
    IntegrationId string `json:"integrationId"`


    // IncludeScreenRecordings - Whether to include Screen recordings for the action, default value = true 
    IncludeScreenRecordings bool `json:"includeScreenRecordings"`


    // ConversationQuery - Conversation Query. Note: After the recording is created, it might take up to 48 hours for the recording to be included in the submitted job query.  This result depends on the analytics data lake job completion. See also: https://developer.genesys.cloud/analyticsdatamanagement/analytics/jobs/conversation-details-job#data-availability
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

