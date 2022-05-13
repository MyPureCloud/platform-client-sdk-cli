package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserconversationsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserconversationsummaryDud struct { 
    


    


    


    


    


    


    


    

}

// Userconversationsummary
type Userconversationsummary struct { 
    // UserId
    UserId string `json:"userId"`


    // Call
    Call Mediasummary `json:"call"`


    // Callback
    Callback Mediasummary `json:"callback"`


    // Email
    Email Mediasummary `json:"email"`


    // Message
    Message Mediasummary `json:"message"`


    // Chat
    Chat Mediasummary `json:"chat"`


    // SocialExpression
    SocialExpression Mediasummary `json:"socialExpression"`


    // Video
    Video Mediasummary `json:"video"`

}

// String returns a JSON representation of the model
func (o *Userconversationsummary) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userconversationsummary) MarshalJSON() ([]byte, error) {
    type Alias Userconversationsummary

    if UserconversationsummaryMarshalled {
        return []byte("{}"), nil
    }
    UserconversationsummaryMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        Call Mediasummary `json:"call"`
        
        Callback Mediasummary `json:"callback"`
        
        Email Mediasummary `json:"email"`
        
        Message Mediasummary `json:"message"`
        
        Chat Mediasummary `json:"chat"`
        
        SocialExpression Mediasummary `json:"socialExpression"`
        
        Video Mediasummary `json:"video"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

