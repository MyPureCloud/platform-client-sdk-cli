package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingcontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingcontentDud struct { 
    ContentType string `json:"contentType"`


    Attachment Webmessagingattachment `json:"attachment"`


    


    


    


    


    


    


    

}

// Webmessagingcontent - Message content element.
type Webmessagingcontent struct { 
    


    


    // QuickReply - Quick reply content.
    QuickReply Webmessagingquickreply `json:"quickReply"`


    // ButtonResponse - Button response content.
    ButtonResponse Webmessagingbuttonresponse `json:"buttonResponse"`


    // Generic - Generic content (Deprecated).
    Generic Webmessaginggeneric `json:"generic"`


    // Card - Card content
    Card Contentcard `json:"card"`


    // Carousel - Carousel content
    Carousel Contentcarousel `json:"carousel"`


    // DatePicker - DatePicker content
    DatePicker Contentdatepicker `json:"datePicker"`


    // ListPicker - ListPicker content
    ListPicker Conversationcontentlistpicker `json:"listPicker"`

}

// String returns a JSON representation of the model
func (o *Webmessagingcontent) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingcontent) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingcontent

    if WebmessagingcontentMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingcontentMarshalled = true

    return json.Marshal(&struct {
        
        QuickReply Webmessagingquickreply `json:"quickReply"`
        
        ButtonResponse Webmessagingbuttonresponse `json:"buttonResponse"`
        
        Generic Webmessaginggeneric `json:"generic"`
        
        Card Contentcard `json:"card"`
        
        Carousel Contentcarousel `json:"carousel"`
        
        DatePicker Contentdatepicker `json:"datePicker"`
        
        ListPicker Conversationcontentlistpicker `json:"listPicker"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

