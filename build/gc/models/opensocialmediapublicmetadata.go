package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediapublicmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediapublicmetadataDud struct { 
    


    


    


    


    


    

}

// Opensocialmediapublicmetadata - Information about a public message.
type Opensocialmediapublicmetadata struct { 
    // RootId - The id of the root public message.
    RootId string `json:"rootId"`


    // ReplyToId - The id of the message this public message is replying to.
    ReplyToId string `json:"replyToId"`


    // Source - The source of the public message. Useful when there could be more than location. Channel specific, e.g., for Facebook it's a source page.
    Source string `json:"source"`


    // Url - The URL of the social post on the native platform.
    Url string `json:"url"`


    // MentionIds - IDs of accounts referenced by name within the message text or caption (for example '@brandpage' in a post body or comment). A 'mention' here uses the same convention as Meta (Facebook, Instagram) and LinkedIn, where it denotes a textual reference to an account. Distinct from tagIds, which represent explicit associations with the message. Null or absent when no accounts are referenced in the text.
    MentionIds []string `json:"mentionIds"`


    // TagIds - IDs of accounts attached to the message itself, independent of the text (for example a person tagged in an Instagram photo so their profile is linked from the image). A 'tag' here uses the same convention as Meta (Facebook, Instagram) and LinkedIn, where it denotes an explicit association with content rather than a textual reference. Distinct from mentionIds, which represent in-text/caption references. Null or absent when no accounts are tagged on the message.
    TagIds []string `json:"tagIds"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediapublicmetadata) String() string {
    
    
    
    
     o.MentionIds = []string{""} 
     o.TagIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediapublicmetadata) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediapublicmetadata

    if OpensocialmediapublicmetadataMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediapublicmetadataMarshalled = true

    return json.Marshal(&struct {
        
        RootId string `json:"rootId"`
        
        ReplyToId string `json:"replyToId"`
        
        Source string `json:"source"`
        
        Url string `json:"url"`
        
        MentionIds []string `json:"mentionIds"`
        
        TagIds []string `json:"tagIds"`
        *Alias
    }{

        


        


        


        


        
        MentionIds: []string{""},
        


        
        TagIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

