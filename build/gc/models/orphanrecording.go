package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrphanrecordingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrphanrecordingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Orphanrecording
type Orphanrecording struct { 
    


    // Name
    Name string `json:"name"`


    // CreatedTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedTime time.Time `json:"createdTime"`


    // RecoveredTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RecoveredTime time.Time `json:"recoveredTime"`


    // ProviderType
    ProviderType string `json:"providerType"`


    // MediaSizeBytes
    MediaSizeBytes int `json:"mediaSizeBytes"`


    // MediaType
    MediaType string `json:"mediaType"`


    // FileState
    FileState string `json:"fileState"`


    // ProviderEndpoint
    ProviderEndpoint Endpoint `json:"providerEndpoint"`


    // Recording
    Recording Recording `json:"recording"`


    // OrphanStatus - The status of the orphaned recording's conversation.
    OrphanStatus string `json:"orphanStatus"`


    // SourceOrphaningId - An identifier used during recovery operations by the supplying hybrid platform to track back and determine which interaction this recording is associated with
    SourceOrphaningId string `json:"sourceOrphaningId"`


    

}

// String returns a JSON representation of the model
func (o *Orphanrecording) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Orphanrecording) MarshalJSON() ([]byte, error) {
    type Alias Orphanrecording

    if OrphanrecordingMarshalled {
        return []byte("{}"), nil
    }
    OrphanrecordingMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        CreatedTime time.Time `json:"createdTime"`
        
        RecoveredTime time.Time `json:"recoveredTime"`
        
        ProviderType string `json:"providerType"`
        
        MediaSizeBytes int `json:"mediaSizeBytes"`
        
        MediaType string `json:"mediaType"`
        
        FileState string `json:"fileState"`
        
        ProviderEndpoint Endpoint `json:"providerEndpoint"`
        
        Recording Recording `json:"recording"`
        
        OrphanStatus string `json:"orphanStatus"`
        
        SourceOrphaningId string `json:"sourceOrphaningId"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

