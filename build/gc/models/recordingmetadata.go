package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingmetadataDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Recordingmetadata
type Recordingmetadata struct { 
    


    // Name
    Name string `json:"name"`


    // ConversationId
    ConversationId string `json:"conversationId"`


    // Path
    Path string `json:"path"`


    // StartTime - The start time of the recording for screen recordings. Null for other types.
    StartTime string `json:"startTime"`


    // EndTime
    EndTime string `json:"endTime"`


    // Media - The type of media that the recording is. At the moment that could be audio, chat, email, or message.
    Media string `json:"media"`


    // Annotations - Annotations that belong to the recording. Populated when recording filestate is AVAILABLE.
    Annotations []Annotation `json:"annotations"`


    // FileState - Represents the current file state for a recording. Examples: Uploading, Archived, etc
    FileState string `json:"fileState"`


    // RestoreExpirationTime - The amount of time a restored recording will remain restored before being archived again. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RestoreExpirationTime time.Time `json:"restoreExpirationTime"`


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


    // MaxAllowedRestorationsForOrg - How many archive restorations the organization is allowed to have.
    MaxAllowedRestorationsForOrg int `json:"maxAllowedRestorationsForOrg"`


    // RemainingRestorationsAllowedForOrg - The remaining archive restorations the organization has.
    RemainingRestorationsAllowedForOrg int `json:"remainingRestorationsAllowedForOrg"`


    // SessionId - The session id represents an external resource id, such as email, call, chat, etc
    SessionId string `json:"sessionId"`


    

}

// String returns a JSON representation of the model
func (o *Recordingmetadata) String() string {
    
    
    
    
    
    
     o.Annotations = []Annotation{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingmetadata) MarshalJSON() ([]byte, error) {
    type Alias Recordingmetadata

    if RecordingmetadataMarshalled {
        return []byte("{}"), nil
    }
    RecordingmetadataMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ConversationId string `json:"conversationId"`
        
        Path string `json:"path"`
        
        StartTime string `json:"startTime"`
        
        EndTime string `json:"endTime"`
        
        Media string `json:"media"`
        
        Annotations []Annotation `json:"annotations"`
        
        FileState string `json:"fileState"`
        
        RestoreExpirationTime time.Time `json:"restoreExpirationTime"`
        
        ArchiveDate time.Time `json:"archiveDate"`
        
        ArchiveMedium string `json:"archiveMedium"`
        
        DeleteDate time.Time `json:"deleteDate"`
        
        ExportDate time.Time `json:"exportDate"`
        
        ExportedDate time.Time `json:"exportedDate"`
        
        MaxAllowedRestorationsForOrg int `json:"maxAllowedRestorationsForOrg"`
        
        RemainingRestorationsAllowedForOrg int `json:"remainingRestorationsAllowedForOrg"`
        
        SessionId string `json:"sessionId"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Annotations: []Annotation{{}},
        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

