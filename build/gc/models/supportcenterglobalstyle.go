package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterglobalstyleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterglobalstyleDud struct { 
    


    


    


    


    


    

}

// Supportcenterglobalstyle
type Supportcenterglobalstyle struct { 
    // BackgroundColor - Global background color, in hexadecimal format, eg #ffffff
    BackgroundColor string `json:"backgroundColor"`


    // PrimaryColor - Global primary color, in hexadecimal format, eg #ffffff
    PrimaryColor string `json:"primaryColor"`


    // PrimaryColorDark - Global dark primary color, in hexadecimal format, eg #ffffff
    PrimaryColorDark string `json:"primaryColorDark"`


    // PrimaryColorLight - Global light primary color, in hexadecimal format, eg #ffffff
    PrimaryColorLight string `json:"primaryColorLight"`


    // TextColor - Global text color, in hexadecimal format, eg #ffffff
    TextColor string `json:"textColor"`


    // FontFamily - Global font family
    FontFamily string `json:"fontFamily"`

}

// String returns a JSON representation of the model
func (o *Supportcenterglobalstyle) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterglobalstyle) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterglobalstyle

    if SupportcenterglobalstyleMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterglobalstyleMarshalled = true

    return json.Marshal(&struct { 
        BackgroundColor string `json:"backgroundColor"`
        
        PrimaryColor string `json:"primaryColor"`
        
        PrimaryColorDark string `json:"primaryColorDark"`
        
        PrimaryColorLight string `json:"primaryColorLight"`
        
        TextColor string `json:"textColor"`
        
        FontFamily string `json:"fontFamily"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

