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

// Supportcentersettings - Settings concerning knowledge portal (previously support center)
type Supportcentersettings struct { 
    // Enabled - Whether or not knowledge portal (previously support center) is enabled
    Enabled bool `json:"enabled"`


    // KnowledgeBase - The knowledge base for knowledge portal (previously support center)
    KnowledgeBase Addressableentityref `json:"knowledgeBase"`


    // CustomMessages - Customizable display texts for knowledge portal (previously support center)
    CustomMessages []Supportcentercustommessage `json:"customMessages"`


    // RouterType - Router type for knowledge portal (previously support center)
    RouterType string `json:"routerType"`


    // Screens - Available screens for the knowledge portal (previously support center) with its modules
    Screens []Supportcenterscreen `json:"screens"`


    // EnabledCategories - Featured categories for knowledge portal (previously support center) home screen
    EnabledCategories []Supportcentercategory `json:"enabledCategories"`


    // StyleSetting - Style attributes for knowledge portal (previously support center)
    StyleSetting Supportcenterstylesetting `json:"styleSetting"`


    // Feedback - Customer feedback settings
    Feedback Supportcenterfeedbacksettings `json:"feedback"`

}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
    
    
     o.CustomMessages = []Supportcentercustommessage{{}} 
    
     o.Screens = []Supportcenterscreen{{}} 
     o.EnabledCategories = []Supportcentercategory{{}} 
    
    

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
        
        EnabledCategories []Supportcentercategory `json:"enabledCategories"`
        
        StyleSetting Supportcenterstylesetting `json:"styleSetting"`
        
        Feedback Supportcenterfeedbacksettings `json:"feedback"`
        *Alias
    }{

        


        


        
        CustomMessages: []Supportcentercustommessage{{}},
        


        


        
        Screens: []Supportcenterscreen{{}},
        


        
        EnabledCategories: []Supportcentercategory{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

