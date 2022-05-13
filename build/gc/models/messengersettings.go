package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessengersettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessengersettingsDud struct { 
    


    


    


    


    


    

}

// Messengersettings - Settings concerning messenger
type Messengersettings struct { 
    // Enabled - Whether or not messenger is enabled
    Enabled bool `json:"enabled"`


    // Styles - The style settings for messenger
    Styles Messengerstyles `json:"styles"`


    // LauncherButton - The launcher button settings for messenger
    LauncherButton Launcherbuttonsettings `json:"launcherButton"`


    // FileUpload - The file upload settings for messenger
    FileUpload Fileuploadsettings `json:"fileUpload"`


    // Apps - The apps embedded in the messenger
    Apps Messengerapps `json:"apps"`


    // HomeScreen - The homescreen settings for messenger
    HomeScreen Messengerhomescreen `json:"homeScreen"`

}

// String returns a JSON representation of the model
func (o *Messengersettings) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messengersettings) MarshalJSON() ([]byte, error) {
    type Alias Messengersettings

    if MessengersettingsMarshalled {
        return []byte("{}"), nil
    }
    MessengersettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Styles Messengerstyles `json:"styles"`
        
        LauncherButton Launcherbuttonsettings `json:"launcherButton"`
        
        FileUpload Fileuploadsettings `json:"fileUpload"`
        
        Apps Messengerapps `json:"apps"`
        
        HomeScreen Messengerhomescreen `json:"homeScreen"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

