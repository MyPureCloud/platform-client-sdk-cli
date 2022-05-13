package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseconfigDud struct { 
    


    


    


    

}

// Responseconfig - Defines response components of the Action Request.
type Responseconfig struct { 
    // TranslationMap - Map 'attribute name' and 'JSON path' pairs used to extract data from REST response.
    TranslationMap map[string]string `json:"translationMap"`


    // TranslationMapDefaults - Map 'attribute name' and 'default value' pairs used as fallback values if JSON path extraction fails for specified key.
    TranslationMapDefaults map[string]string `json:"translationMapDefaults"`


    // SuccessTemplate - Velocity template to build response to return from Action.
    SuccessTemplate string `json:"successTemplate"`


    // SuccessTemplateUri - URI to retrieve success template.
    SuccessTemplateUri string `json:"successTemplateUri"`

}

// String returns a JSON representation of the model
func (o *Responseconfig) String() string {
     o.TranslationMap = map[string]string{"": ""} 
     o.TranslationMapDefaults = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseconfig) MarshalJSON() ([]byte, error) {
    type Alias Responseconfig

    if ResponseconfigMarshalled {
        return []byte("{}"), nil
    }
    ResponseconfigMarshalled = true

    return json.Marshal(&struct {
        
        TranslationMap map[string]string `json:"translationMap"`
        
        TranslationMapDefaults map[string]string `json:"translationMapDefaults"`
        
        SuccessTemplate string `json:"successTemplate"`
        
        SuccessTemplateUri string `json:"successTemplateUri"`
        *Alias
    }{

        
        TranslationMap: map[string]string{"": ""},
        


        
        TranslationMapDefaults: map[string]string{"": ""},
        


        


        

        Alias: (*Alias)(u),
    })
}

