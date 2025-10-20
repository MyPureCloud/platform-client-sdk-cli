package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcontentofferMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcontentofferDud struct { 
    


    


    


    


    


    


    


    


    

}

// Patchcontentoffer
type Patchcontentoffer struct { 
    // ImageUrl - URL for image displayed to the customer when displaying content offer.
    ImageUrl string `json:"imageUrl"`


    // DisplayMode - The display mode of Genesys Widgets when displaying content offer.
    DisplayMode string `json:"displayMode"`


    // LayoutMode - The layout mode of the text shown to the user when displaying content offer.
    LayoutMode string `json:"layoutMode"`


    // Title - Title used in the header of the content offer.
    Title string `json:"title"`


    // Headline - Headline displayed above the body text of the content offer.
    Headline string `json:"headline"`


    // Body - Body text of the content offer.
    Body string `json:"body"`


    // CallToAction - Properties customizing the call to action button on the content offer.
    CallToAction Patchcalltoaction `json:"callToAction"`


    // Style - Properties customizing the styling of the content offer.
    Style Patchcontentofferstylingconfiguration `json:"style"`


    // ImageAltText - Image description text for accessibility compliance and assistive technology support.
    ImageAltText string `json:"imageAltText"`

}

// String returns a JSON representation of the model
func (o *Patchcontentoffer) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcontentoffer) MarshalJSON() ([]byte, error) {
    type Alias Patchcontentoffer

    if PatchcontentofferMarshalled {
        return []byte("{}"), nil
    }
    PatchcontentofferMarshalled = true

    return json.Marshal(&struct {
        
        ImageUrl string `json:"imageUrl"`
        
        DisplayMode string `json:"displayMode"`
        
        LayoutMode string `json:"layoutMode"`
        
        Title string `json:"title"`
        
        Headline string `json:"headline"`
        
        Body string `json:"body"`
        
        CallToAction Patchcalltoaction `json:"callToAction"`
        
        Style Patchcontentofferstylingconfiguration `json:"style"`
        
        ImageAltText string `json:"imageAltText"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

