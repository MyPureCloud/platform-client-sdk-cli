package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcontentofferstylingconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcontentofferstylingconfigurationDud struct { 
    


    


    


    


    


    


    

}

// Patchcontentofferstylingconfiguration
type Patchcontentofferstylingconfiguration struct { 
    // Position - Properties for customizing the positioning of the content offer.
    Position Patchcontentpositionproperties `json:"position"`


    // Offer - Properties for customizing the appearance of the content offer.
    Offer Patchcontentofferstyleproperties `json:"offer"`


    // CloseButton - Properties for customizing the appearance of the close button.
    CloseButton Patchclosebuttonstyleproperties `json:"closeButton"`


    // CtaButton - Properties for customizing the appearance of the CTA button.
    CtaButton Patchctabuttonstyleproperties `json:"ctaButton"`


    // Title - Properties for customizing the appearance of the title text.
    Title Patchtextstyleproperties `json:"title"`


    // Headline - Properties for customizing the appearance of the headline text.
    Headline Patchtextstyleproperties `json:"headline"`


    // Body - Properties for customizing the appearance of the body text.
    Body Patchtextstyleproperties `json:"body"`

}

// String returns a JSON representation of the model
func (o *Patchcontentofferstylingconfiguration) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcontentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Patchcontentofferstylingconfiguration

    if PatchcontentofferstylingconfigurationMarshalled {
        return []byte("{}"), nil
    }
    PatchcontentofferstylingconfigurationMarshalled = true

    return json.Marshal(&struct { 
        Position Patchcontentpositionproperties `json:"position"`
        
        Offer Patchcontentofferstyleproperties `json:"offer"`
        
        CloseButton Patchclosebuttonstyleproperties `json:"closeButton"`
        
        CtaButton Patchctabuttonstyleproperties `json:"ctaButton"`
        
        Title Patchtextstyleproperties `json:"title"`
        
        Headline Patchtextstyleproperties `json:"headline"`
        
        Body Patchtextstyleproperties `json:"body"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

