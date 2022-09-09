package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingretentionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingretentionDud struct { 
    


    


    


    


    


    


    


    

}

// Recordingretention
type Recordingretention struct { 
    // ConversationId
    ConversationId string `json:"conversationId"`


    // RecordingId
    RecordingId string `json:"recordingId"`


    // ArchiveDate - The date the recording will be archived. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ArchiveDate time.Time `json:"archiveDate"`


    // ArchiveMedium - The type of archive medium used. Example: CloudArchive
    ArchiveMedium string `json:"archiveMedium"`


    // DeleteDate - The date the recording will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DeleteDate time.Time `json:"deleteDate"`


    // ExportDate - The date the recording will be exported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExportDate time.Time `json:"exportDate"`


    // ExportedDate - The date the recording was exported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExportedDate time.Time `json:"exportedDate"`


    // CreationTime - The creation time of the recording. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreationTime time.Time `json:"creationTime"`

}

// String returns a JSON representation of the model
func (o *Recordingretention) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingretention) MarshalJSON() ([]byte, error) {
    type Alias Recordingretention

    if RecordingretentionMarshalled {
        return []byte("{}"), nil
    }
    RecordingretentionMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        RecordingId string `json:"recordingId"`
        
        ArchiveDate time.Time `json:"archiveDate"`
        
        ArchiveMedium string `json:"archiveMedium"`
        
        DeleteDate time.Time `json:"deleteDate"`
        
        ExportDate time.Time `json:"exportDate"`
        
        ExportedDate time.Time `json:"exportedDate"`
        
        CreationTime time.Time `json:"creationTime"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

