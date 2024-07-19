package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailmessageDud struct { 
    Id string `json:"id"`


    Conversation Conversation `json:"conversation"`


    


    AudioRecordingDurationSeconds int `json:"audioRecordingDurationSeconds"`


    AudioRecordingSizeBytes int `json:"audioRecordingSizeBytes"`


    Transcription string `json:"transcription"`


    CreatedDate time.Time `json:"createdDate"`


    ModifiedDate time.Time `json:"modifiedDate"`


    DeletedDate time.Time `json:"deletedDate"`


    CallerAddress string `json:"callerAddress"`


    CallerName string `json:"callerName"`


    CallerUser User `json:"callerUser"`


    


    


    User User `json:"user"`


    Group Group `json:"group"`


    Queue Queue `json:"queue"`


    CopiedFrom Voicemailcopyrecord `json:"copiedFrom"`


    CopiedTo []Voicemailcopyrecord `json:"copiedTo"`


    


    SelfUri string `json:"selfUri"`

}

// Voicemailmessage
type Voicemailmessage struct { 
    


    


    // Read - Whether the voicemail message is marked as read
    Read bool `json:"read"`


    


    


    


    


    


    


    


    


    


    // Deleted - Whether the voicemail message has been marked as deleted
    Deleted bool `json:"deleted"`


    // Note - An optional note
    Note string `json:"note"`


    


    


    


    


    


    // DeleteRetentionPolicy - The retention policy for this voicemail when deleted is set to true
    DeleteRetentionPolicy Voicemailretentionpolicy `json:"deleteRetentionPolicy"`


    

}

// String returns a JSON representation of the model
func (o *Voicemailmessage) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailmessage) MarshalJSON() ([]byte, error) {
    type Alias Voicemailmessage

    if VoicemailmessageMarshalled {
        return []byte("{}"), nil
    }
    VoicemailmessageMarshalled = true

    return json.Marshal(&struct {
        
        Read bool `json:"read"`
        
        Deleted bool `json:"deleted"`
        
        Note string `json:"note"`
        
        DeleteRetentionPolicy Voicemailretentionpolicy `json:"deleteRetentionPolicy"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

