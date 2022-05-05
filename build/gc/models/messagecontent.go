package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagecontentDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Messagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Messagecontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // Location - Location content.
    Location Contentlocation `json:"location"`


    // Attachment - Attachment content.
    Attachment Contentattachment `json:"attachment"`


    // QuickReply - Quick reply content.
    QuickReply Contentquickreply `json:"quickReply"`


    // ButtonResponse - Button response content.
    ButtonResponse Contentbuttonresponse `json:"buttonResponse"`


    // Generic - Generic content (Deprecated).
    Generic Contentgeneric `json:"generic"`


    // List - List content (Deprecated).
    List Contentlist `json:"list"`


    // Template - Template notification content.
    Template Contentnotificationtemplate `json:"template"`


    // Reactions - A set of reactions to a message.
    Reactions []Contentreaction `json:"reactions"`


    // Mention - Mention content.
    Mention Messagingrecipient `json:"mention"`


    // Postback - Structured message postback (Deprecated).
    Postback Contentpostback `json:"postback"`


    // Story - Ephemeral story content.
    Story Contentstory `json:"story"`


    // Card - Card content
    Card Contentcard `json:"card"`


    // Carousel - Carousel content
    Carousel Contentcarousel `json:"carousel"`

}

// String returns a JSON representation of the model
func (o *Messagecontent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Reactions = []Contentreaction{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagecontent) MarshalJSON() ([]byte, error) {
    type Alias Messagecontent

    if MessagecontentMarshalled {
        return []byte("{}"), nil
    }
    MessagecontentMarshalled = true

    return json.Marshal(&struct { 
        ContentType string `json:"contentType"`
        
        Location Contentlocation `json:"location"`
        
        Attachment Contentattachment `json:"attachment"`
        
        QuickReply Contentquickreply `json:"quickReply"`
        
        ButtonResponse Contentbuttonresponse `json:"buttonResponse"`
        
        Generic Contentgeneric `json:"generic"`
        
        List Contentlist `json:"list"`
        
        Template Contentnotificationtemplate `json:"template"`
        
        Reactions []Contentreaction `json:"reactions"`
        
        Mention Messagingrecipient `json:"mention"`
        
        Postback Contentpostback `json:"postback"`
        
        Story Contentstory `json:"story"`
        
        Card Contentcard `json:"card"`
        
        Carousel Contentcarousel `json:"carousel"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Reactions: []Contentreaction{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

