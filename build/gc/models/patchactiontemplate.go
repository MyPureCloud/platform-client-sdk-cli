package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchactiontemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchactiontemplateDud struct { 
    


    


    


    


    

}

// Patchactiontemplate
type Patchactiontemplate struct { 
    // Name - Name of the action template.
    Name string `json:"name"`


    // Description - Description of the action template's functionality.
    Description string `json:"description"`


    // MediaType - Media type of action described by the action template.
    MediaType string `json:"mediaType"`


    // State - Whether the action template is currently active, inactive or deleted.
    State string `json:"state"`


    // ContentOffer - Properties used to configure an action of type content offer
    ContentOffer Patchcontentoffer `json:"contentOffer"`

}

// String returns a JSON representation of the model
func (o *Patchactiontemplate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchactiontemplate) MarshalJSON() ([]byte, error) {
    type Alias Patchactiontemplate

    if PatchactiontemplateMarshalled {
        return []byte("{}"), nil
    }
    PatchactiontemplateMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        MediaType string `json:"mediaType"`
        
        State string `json:"state"`
        
        ContentOffer Patchcontentoffer `json:"contentOffer"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

