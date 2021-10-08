package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingbuttonresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingbuttonresponseDud struct { 
    


    


    


    

}

// Webmessagingbuttonresponse - Button response object representing the click of a structured message button, such as a quick reply.
type Webmessagingbuttonresponse struct { 
    // Id - An ID assigned to the button response (Deprecated).
    Id string `json:"id"`


    // VarType - Describes the button that resulted in the Button Response.
    VarType string `json:"type"`


    // Text - The response text from the button click.
    Text string `json:"text"`


    // Payload - The response payload associated with the clicked button.
    Payload string `json:"payload"`

}

// String returns a JSON representation of the model
func (o *Webmessagingbuttonresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingbuttonresponse) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingbuttonresponse

    if WebmessagingbuttonresponseMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingbuttonresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

