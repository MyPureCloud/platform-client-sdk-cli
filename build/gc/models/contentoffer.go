package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentofferMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentofferDud struct { 
    


    


    


    


    


    


    


    

}

// Contentoffer
type Contentoffer struct { 
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
    CallToAction Calltoaction `json:"callToAction"`


    // Style - Properties customizing the styling of the content offer.
    Style Contentofferstylingconfiguration `json:"style"`

}

// String returns a JSON representation of the model
func (o *Contentoffer) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentoffer) MarshalJSON() ([]byte, error) {
    type Alias Contentoffer

    if ContentofferMarshalled {
        return []byte("{}"), nil
    }
    ContentofferMarshalled = true

    return json.Marshal(&struct { 
        ImageUrl string `json:"imageUrl"`
        
        DisplayMode string `json:"displayMode"`
        
        LayoutMode string `json:"layoutMode"`
        
        Title string `json:"title"`
        
        Headline string `json:"headline"`
        
        Body string `json:"body"`
        
        CallToAction Calltoaction `json:"callToAction"`
        
        Style Contentofferstylingconfiguration `json:"style"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

