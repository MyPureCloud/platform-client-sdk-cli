package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecognitionbaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecognitionbaseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Recognitionbase
type Recognitionbase struct { 
    


    // Recipient - The recipient of the recognition
    Recipient Userreference `json:"recipient"`


    // CreatedBy - The creator of the recognition
    CreatedBy Userreference `json:"createdBy"`


    // DateCreated - The creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // VarType - The type of recognition
    VarType string `json:"type"`


    // Title - The recognition title
    Title string `json:"title"`


    // Note - The recognition note
    Note string `json:"note"`


    // ContextType - The context type (optional)
    ContextType string `json:"contextType"`


    // ContextId - The context id (optional)
    ContextId string `json:"contextId"`


    

}

// String returns a JSON representation of the model
func (o *Recognitionbase) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recognitionbase) MarshalJSON() ([]byte, error) {
    type Alias Recognitionbase

    if RecognitionbaseMarshalled {
        return []byte("{}"), nil
    }
    RecognitionbaseMarshalled = true

    return json.Marshal(&struct {
        
        Recipient Userreference `json:"recipient"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
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

