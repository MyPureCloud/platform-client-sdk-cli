package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Recording
type Recording struct { 
    


    // Name
    Name string `json:"name"`


    // ConversationId
    ConversationId string `json:"conversationId"`


    // Path
    Path string `json:"path"`


    // StartTime - The start time of the recording. Null when there is no playable media.
    StartTime string `json:"startTime"`


    // EndTime - The end time of the recording. Null when there is no playable media.
    EndTime string `json:"endTime"`


    // Media - The media type of the recording. This could be audio, chat, messaging, email, or screen.
    Media string `json:"media"`


    // MediaSubtype - The media subtype of the recording.
    MediaSubtype string `json:"mediaSubtype"`


    // MediaSubject - The media subject of the recording.
    MediaSubject string `json:"mediaSubject"`


    // Annotations - Annotations that belong to the recording.
    Annotations []Annotation `json:"annotations"`


    // Transcript - Represents a chat transcript
    Transcript []Chatmessage `json:"transcript"`


    // EmailTranscript - Represents an email transcript
    EmailTranscript []Recordingemailmessage `json:"emailTranscript"`


    // MessagingTranscript - Represents a messaging transcript
    MessagingTranscript []Recordingmessagingmessage `json:"messagingTranscript"`


    // FileState - Represents the current file state for a recording. Examples: Uploading, Archived, etc
    FileState string `json:"fileState"`


    // RestoreExpirationTime - The amount of time a restored recording will remain restored before being archived again. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RestoreExpirationTime time.Time `json:"restoreExpirationTime"`


    // MediaUris - The different mediaUris for the recording. Null when there is no playable media.
    MediaUris map[string]Mediaresult `json:"mediaUris"`


    // EstimatedTranscodeTimeMs
    EstimatedTranscodeTimeMs int `json:"estimatedTranscodeTimeMs"`


    // ActualTranscodeTimeMs
    ActualTranscodeTimeMs int `json:"actualTranscodeTimeMs"`


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


    // OutputDurationMs - Duration of transcoded media in milliseconds
    OutputDurationMs int `json:"outputDurationMs"`


    // OutputSizeInBytes - Size of transcoded media in bytes. 0 if there is no transcoded media.
    OutputSizeInBytes int `json:"outputSizeInBytes"`


    // MaxAllowedRestorationsForOrg - How many archive restorations the organization is allowed to have. Deprecated - Always returns 10000 since the restoration limit is no longer enforced.
    MaxAllowedRestorationsForOrg int `json:"maxAllowedRestorationsForOrg"`


    // RemainingRestorationsAllowedForOrg - The remaining archive restorations the organization has. Deprecated - Always returns 10000 since the restoration limit is no longer enforced.
    RemainingRestorationsAllowedForOrg int `json:"remainingRestorationsAllowedForOrg"`


    // SessionId - The session id represents an external resource id, such as email, call, chat, etc
    SessionId string `json:"sessionId"`


    // Users - The users participating in the conversation
    Users []User `json:"users"`


    // RecordingFileRole - Role of the file recording. It can be either customer_experience or adhoc.
    RecordingFileRole string `json:"recordingFileRole"`


    // RecordingErrorStatus - Status of a recording that cannot be returned because of an error
    RecordingErrorStatus string `json:"recordingErrorStatus"`


    // OriginalRecordingStartTime - The start time of the full recording, before any segment access restrictions are applied. Null when there is no playable media. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    OriginalRecordingStartTime time.Time `json:"originalRecordingStartTime"`


    // CreationTime - The creation time of the recording. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreationTime time.Time `json:"creationTime"`


    

}

// String returns a JSON representation of the model
func (o *Recording) String() string {
    
    
    
    
    
    
    
    
     o.Annotations = []Annotation{{}} 
     o.Transcript = []Chatmessage{{}} 
     o.EmailTranscript = []Recordingemailmessage{{}} 
     o.MessagingTranscript = []Recordingmessagingmessage{{}} 
    
    
     o.MediaUris = map[string]Mediaresult{"": {}} 
    
    
    
    
    
    
    
    
    
    
    
    
     o.Users = []User{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recording) MarshalJSON() ([]byte, error) {
    type Alias Recording

    if RecordingMarshalled {
        return []byte("{}"), nil
    }
    RecordingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ConversationId string `json:"conversationId"`
        
        Path string `json:"path"`
        
        StartTime string `json:"startTime"`
        
        EndTime string `json:"endTime"`
        
        Media string `json:"media"`
        
        MediaSubtype string `json:"mediaSubtype"`
        
        MediaSubject string `json:"mediaSubject"`
        
        Annotations []Annotation `json:"annotations"`
        
        Transcript []Chatmessage `json:"transcript"`
        
        EmailTranscript []Recordingemailmessage `json:"emailTranscript"`
        
        MessagingTranscript []Recordingmessagingmessage `json:"messagingTranscript"`
        
        FileState string `json:"fileState"`
        
        RestoreExpirationTime time.Time `json:"restoreExpirationTime"`
        
        MediaUris map[string]Mediaresult `json:"mediaUris"`
        
        EstimatedTranscodeTimeMs int `json:"estimatedTranscodeTimeMs"`
        
        ActualTranscodeTimeMs int `json:"actualTranscodeTimeMs"`
        
        ArchiveDate time.Time `json:"archiveDate"`
        
        ArchiveMedium string `json:"archiveMedium"`
        
        DeleteDate time.Time `json:"deleteDate"`
        
        ExportDate time.Time `json:"exportDate"`
        
        ExportedDate time.Time `json:"exportedDate"`
        
        OutputDurationMs int `json:"outputDurationMs"`
        
        OutputSizeInBytes int `json:"outputSizeInBytes"`
        
        MaxAllowedRestorationsForOrg int `json:"maxAllowedRestorationsForOrg"`
        
        RemainingRestorationsAllowedForOrg int `json:"remainingRestorationsAllowedForOrg"`
        
        SessionId string `json:"sessionId"`
        
        Users []User `json:"users"`
        
        RecordingFileRole string `json:"recordingFileRole"`
        
        RecordingErrorStatus string `json:"recordingErrorStatus"`
        
        OriginalRecordingStartTime time.Time `json:"originalRecordingStartTime"`
        
        CreationTime time.Time `json:"creationTime"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Annotations: []Annotation{{}},
        


        
        Transcript: []Chatmessage{{}},
        


        
        EmailTranscript: []Recordingemailmessage{{}},
        


        
        MessagingTranscript: []Recordingmessagingmessage{{}},
        


        


        


        
        MediaUris: map[string]Mediaresult{"": {}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Users: []User{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

