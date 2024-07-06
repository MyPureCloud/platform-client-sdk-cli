package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreaterecognitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreaterecognitionDud struct { 
    


    


    


    


    


    

}

// Createrecognition
type Createrecognition struct { 
    // RecipientId - The recipient of the recognition
    RecipientId string `json:"recipientId"`


    // VarType - The type of the recognition
    VarType string `json:"type"`


    // Title - The title of the recognition. Max length of 100 characters (optional)
    Title string `json:"title"`


    // Note - The note of the recognition. Max length of 800 characters (optional)
    Note string `json:"note"`


    // ContextType - The context type (optional)
    ContextType string `json:"contextType"`


    // ContextId - The context id (optional)
    ContextId string `json:"contextId"`

}

// String returns a JSON representation of the model
func (o *Createrecognition) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createrecognition) MarshalJSON() ([]byte, error) {
    type Alias Createrecognition

    if CreaterecognitionMarshalled {
        return []byte("{}"), nil
    }
    CreaterecognitionMarshalled = true

    return json.Marshal(&struct {
        
        RecipientId string `json:"recipientId"`
        
        VarType string `json:"type"`
        
        Title string `json:"title"`
        
        Note string `json:"note"`
        
        ContextType string `json:"contextType"`
        
        ContextId string `json:"contextId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

