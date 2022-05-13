package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentofferstylingconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentofferstylingconfigurationDud struct { 
    


    


    


    


    


    


    

}

// Contentofferstylingconfiguration
type Contentofferstylingconfiguration struct { 
    // Position - Properties for customizing the positioning of the content offer.
    Position Contentpositionproperties `json:"position"`


    // Offer - Properties for customizing the appearance of the content offer.
    Offer Contentofferstyleproperties `json:"offer"`


    // CloseButton - Properties for customizing the appearance of the close button.
    CloseButton Closebuttonstyleproperties `json:"closeButton"`


    // CtaButton - Properties for customizing the appearance of the CTA button.
    CtaButton Ctabuttonstyleproperties `json:"ctaButton"`


    // Title - Properties for customizing the appearance of the title text.
    Title Textstyleproperties `json:"title"`


    // Headline - Properties for customizing the appearance of the headline text.
    Headline Textstyleproperties `json:"headline"`


    // Body - Properties for customizing the appearance of the body text.
    Body Textstyleproperties `json:"body"`

}

// String returns a JSON representation of the model
func (o *Contentofferstylingconfiguration) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Contentofferstylingconfiguration

    if ContentofferstylingconfigurationMarshalled {
        return []byte("{}"), nil
    }
    ContentofferstylingconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Position Contentpositionproperties `json:"position"`
        
        Offer Contentofferstyleproperties `json:"offer"`
        
        CloseButton Closebuttonstyleproperties `json:"closeButton"`
        
        CtaButton Ctabuttonstyleproperties `json:"ctaButton"`
        
        Title Textstyleproperties `json:"title"`
        
        Headline Textstyleproperties `json:"headline"`
        
        Body Textstyleproperties `json:"body"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

