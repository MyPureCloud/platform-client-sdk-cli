package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmoduleinformsteprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmoduleinformsteprequestDud struct { 
    


    


    


    


    


    

}

// Learningmoduleinformsteprequest - Learning module inform steps request
type Learningmoduleinformsteprequest struct { 
    // VarType - The learning module inform step type
    VarType string `json:"type"`


    // Name - The name of the inform step or content
    Name string `json:"name"`


    // Value - The value for inform step
    Value string `json:"value"`


    // SharingUri - The sharing uri for Content type inform step
    SharingUri string `json:"sharingUri"`


    // ContentType - The document type for Content type Inform step
    ContentType string `json:"contentType"`


    // Order - The order of inform step in a learning module
    Order int `json:"order"`

}

// String returns a JSON representation of the model
func (o *Learningmoduleinformsteprequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmoduleinformsteprequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmoduleinformsteprequest

    if LearningmoduleinformsteprequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmoduleinformsteprequestMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Name string `json:"name"`
        
        Value string `json:"value"`
        
        SharingUri string `json:"sharingUri"`
        
        ContentType string `json:"contentType"`
        
        Order int `json:"order"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

