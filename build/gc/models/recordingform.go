package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingformDud struct { 
    


    


    


    


    


    


    

}

// Recordingform
type Recordingform struct { 
    // Introduction - The introduction component, used to give an intro into what the form entails.
    Introduction Recordingintroduction `json:"introduction"`


    // FormPages - Form pages.
    FormPages []Recordingformpage `json:"formPages"`


    // ReceivedMessage - Defines the initial prompt message structure containing title and subtitle fields that are displayed to the end user when a form requires completion.
    ReceivedMessage Receivedreplymessage `json:"receivedMessage"`


    // ReplyMessage - The reply message after the user has filled out the form received.
    ReplyMessage Receivedreplymessage `json:"replyMessage"`


    // Response - Content of the payload included in the Form response.
    Response []Recordingformresponsecomponent `json:"response"`


    // OriginatingMessageId - Reference to the id of the original message.
    OriginatingMessageId string `json:"originatingMessageId"`


    // CannedResponseId - The id of the canned response which was used to create the form.
    CannedResponseId string `json:"cannedResponseId"`

}

// String returns a JSON representation of the model
func (o *Recordingform) String() string {
    
     o.FormPages = []Recordingformpage{{}} 
    
    
     o.Response = []Recordingformresponsecomponent{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingform) MarshalJSON() ([]byte, error) {
    type Alias Recordingform

    if RecordingformMarshalled {
        return []byte("{}"), nil
    }
    RecordingformMarshalled = true

    return json.Marshal(&struct {
        
        Introduction Recordingintroduction `json:"introduction"`
        
        FormPages []Recordingformpage `json:"formPages"`
        
        ReceivedMessage Receivedreplymessage `json:"receivedMessage"`
        
        ReplyMessage Receivedreplymessage `json:"replyMessage"`
        
        Response []Recordingformresponsecomponent `json:"response"`
        
        OriginatingMessageId string `json:"originatingMessageId"`
        
        CannedResponseId string `json:"cannedResponseId"`
        *Alias
    }{

        


        
        FormPages: []Recordingformpage{{}},
        


        


        


        
        Response: []Recordingformresponsecomponent{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

