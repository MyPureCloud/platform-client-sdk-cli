package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcentersettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcentersettingsDud struct { 
    


    


    


    


    


    


    

}

// Supportcentersettings - Settings concerning support center
type Supportcentersettings struct { 
    // Enabled - Whether or not support center is enabled
    Enabled bool `json:"enabled"`


    // KnowledgeBase - The knowledge base for support center
    KnowledgeBase Addressableentityref `json:"knowledgeBase"`


    // CustomMessages - Customizable display texts for support center
    CustomMessages []Supportcentercustommessage `json:"customMessages"`


    // RouterType - Router type for support center
    RouterType string `json:"routerType"`


    // Screens - Available screens for the support center with its modules
    Screens []Supportcenterscreen `json:"screens"`


    // EnabledCategories - Enabled article categories for support center
    EnabledCategories []Addressableentityref `json:"enabledCategories"`


    // StyleSetting - Style attributes for support center
    StyleSetting Supportcenterstylesetting `json:"styleSetting"`

}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
    
    
     o.CustomMessages = []Supportcentercustommessage{{}} 
    
     o.Screens = []Supportcenterscreen{{}} 
     o.EnabledCategories = []Addressableentityref{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcentersettings) MarshalJSON() ([]byte, error) {
    type Alias Supportcentersettings

    if SupportcentersettingsMarshalled {
        return []byte("{}"), nil
    }
    SupportcentersettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        KnowledgeBase Addressableentityref `json:"knowledgeBase"`
        
        CustomMessages []Supportcentercustommessage `json:"customMessages"`
        
        RouterType string `json:"routerType"`
        
        Screens []Supportcenterscreen `json:"screens"`
        
        EnabledCategories []Addressableentityref `json:"enabledCategories"`
        
        StyleSetting Supportcenterstylesetting `json:"styleSetting"`
        *Alias
    }{

        


        


        
        CustomMessages: []Supportcentercustommessage{{}},
        


        


        
        Screens: []Supportcenterscreen{{}},
        


        
        EnabledCategories: []Addressableentityref{{}},
        


        

        Alias: (*Alias)(u),
    })
}

