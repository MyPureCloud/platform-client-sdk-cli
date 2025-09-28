package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagecontentDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Conversationmessagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Conversationmessagecontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // Location - Location content.
    Location Conversationcontentlocation `json:"location"`


    // Attachment - Attachment content.
    Attachment Conversationcontentattachment `json:"attachment"`


    // QuickReply - Quick reply content.
    QuickReply Conversationcontentquickreply `json:"quickReply"`


    // ButtonResponse - Button response content.
    ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`


    // Template - Template notification content.
    Template Conversationcontentnotificationtemplate `json:"template"`


    // Story - Ephemeral story content.
    Story Conversationcontentstory `json:"story"`


    // Card - Card content
    Card Conversationcontentcard `json:"card"`


    // Carousel - Carousel content
    Carousel Conversationcontentcarousel `json:"carousel"`


    // Text - Text content.
    Text Conversationcontenttext `json:"text"`


    // QuickReplyV2 - Quick reply V2 content.
    QuickReplyV2 Conversationcontentquickreplyv2 `json:"quickReplyV2"`


    // Reactions - A set of reactions to a message.
    Reactions []Conversationcontentreaction `json:"reactions"`


    // DatePicker - Date picker content.
    DatePicker Conversationcontentdatepicker `json:"datePicker"`


    // InteractiveApplication - InteractiveApplication content.
    InteractiveApplication Conversationcontentinteractiveapplication `json:"interactiveApplication"`


    // ListPicker - List picker content.
    ListPicker Conversationcontentlistpicker `json:"listPicker"`


    // PaymentRequest - Payment request content.
    PaymentRequest Conversationcontentpaymentrequest `json:"paymentRequest"`


    // PaymentResponse - Payment response content.
    PaymentResponse Conversationcontentpaymentresponse `json:"paymentResponse"`


    // Push - Push content.
    Push Conversationcontentpush `json:"push"`


    // Form - Form content.
    Form Conversationcontentform `json:"form"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.Reactions = []Conversationcontentreaction{{}} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagecontent

    if ConversationmessagecontentMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        Location Conversationcontentlocation `json:"location"`
        
        Attachment Conversationcontentattachment `json:"attachment"`
        
        QuickReply Conversationcontentquickreply `json:"quickReply"`
        
        ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`
        
        Template Conversationcontentnotificationtemplate `json:"template"`
        
        Story Conversationcontentstory `json:"story"`
        
        Card Conversationcontentcard `json:"card"`
        
        Carousel Conversationcontentcarousel `json:"carousel"`
        
        Text Conversationcontenttext `json:"text"`
        
        QuickReplyV2 Conversationcontentquickreplyv2 `json:"quickReplyV2"`
        
        Reactions []Conversationcontentreaction `json:"reactions"`
        
        DatePicker Conversationcontentdatepicker `json:"datePicker"`
        
        InteractiveApplication Conversationcontentinteractiveapplication `json:"interactiveApplication"`
        
        ListPicker Conversationcontentlistpicker `json:"listPicker"`
        
        PaymentRequest Conversationcontentpaymentrequest `json:"paymentRequest"`
        
        PaymentResponse Conversationcontentpaymentresponse `json:"paymentResponse"`
        
        Push Conversationcontentpush `json:"push"`
        
        Form Conversationcontentform `json:"form"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        Reactions: []Conversationcontentreaction{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

