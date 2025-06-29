package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotpromptsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotpromptsegmentDud struct { 
    


    


    


    

}

// Textbotpromptsegment - Data for a single bot flow prompt segment.
type Textbotpromptsegment struct { 
    // Text - The text of this prompt segment.
    Text string `json:"text"`


    // VarType - The segment type which describes any semantics about the 'text' and also indicates which other field might include additional relevant info.
    VarType string `json:"type"`


    // Format - Additional details describing the segment’s contents, which the client should honour where possible.
    Format Format `json:"format"`


    // Content - Details to display Rich Media content. This is only populated when the segment 'type' is 'Rich Media'.
    Content []Conversationmessagecontent `json:"content"`

}

// String returns a JSON representation of the model
func (o *Textbotpromptsegment) String() string {
    
    
    
     o.Content = []Conversationmessagecontent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotpromptsegment) MarshalJSON() ([]byte, error) {
    type Alias Textbotpromptsegment

    if TextbotpromptsegmentMarshalled {
        return []byte("{}"), nil
    }
    TextbotpromptsegmentMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        VarType string `json:"type"`
        
        Format Format `json:"format"`
        
        Content []Conversationmessagecontent `json:"content"`
        *Alias
    }{

        


        


        


        
        Content: []Conversationmessagecontent{{}},
        

        Alias: (*Alias)(u),
    })
}

