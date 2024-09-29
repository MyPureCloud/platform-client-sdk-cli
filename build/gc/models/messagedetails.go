package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagedetailsDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Messagedetails
type Messagedetails struct { 
    // MessageId - UUID identifying the message media.
    MessageId string `json:"messageId"`


    // MessageURI - A URI for this message entity.
    MessageURI string `json:"messageURI"`


    // MessageStatus - Indicates the delivery status of the message.
    MessageStatus string `json:"messageStatus"`


    // MessageSegmentCount - The message segment count, greater than 1 if the message content was split into multiple parts for this message type, e.g. SMS character limits.
    MessageSegmentCount int `json:"messageSegmentCount"`


    // MessageTime - The time when the message was sent or received. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    MessageTime time.Time `json:"messageTime"`


    // Media - The media (images, files, etc) associated with this message, if any
    Media []Messagemedia `json:"media"`


    // Stickers - One or more stickers associated with this message, if any
    Stickers []Messagesticker `json:"stickers"`


    // MessageMetadata - Information that describes the content of the message, if any
    MessageMetadata Conversationmessagemetadata `json:"messageMetadata"`


    // SocialVisibility - For social media messages, the visibility of the message in the originating social platform
    SocialVisibility string `json:"socialVisibility"`


    // ErrorInfo - Provider specific error information for a communication.
    ErrorInfo Errorbody `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Messagedetails) String() string {
    
    
    
    
    
     o.Media = []Messagemedia{{}} 
     o.Stickers = []Messagesticker{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagedetails) MarshalJSON() ([]byte, error) {
    type Alias Messagedetails

    if MessagedetailsMarshalled {
        return []byte("{}"), nil
    }
    MessagedetailsMarshalled = true

    return json.Marshal(&struct {
        
        MessageId string `json:"messageId"`
        
        MessageURI string `json:"messageURI"`
        
        MessageStatus string `json:"messageStatus"`
        
        MessageSegmentCount int `json:"messageSegmentCount"`
        
        MessageTime time.Time `json:"messageTime"`
        
        Media []Messagemedia `json:"media"`
        
        Stickers []Messagesticker `json:"stickers"`
        
        MessageMetadata Conversationmessagemetadata `json:"messageMetadata"`
        
        SocialVisibility string `json:"socialVisibility"`
        
        ErrorInfo Errorbody `json:"errorInfo"`
        *Alias
    }{

        


        


        


        


        


        
        Media: []Messagemedia{{}},
        


        
        Stickers: []Messagesticker{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

