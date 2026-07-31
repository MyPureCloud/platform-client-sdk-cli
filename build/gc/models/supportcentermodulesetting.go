package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcentermodulesettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcentermodulesettingDud struct { 
    


    


    


    

}

// Supportcentermodulesetting
type Supportcentermodulesetting struct { 
    // VarType - Screen module type. Required if support center is enabled
    VarType string `json:"type"`


    // Enabled - Whether or not knowledge portal (previously support center) screen module is enabled. If not provided, the module is treated as hidden
    Enabled bool `json:"enabled"`


    // CompactCategoryModuleTemplate - Compact category module template
    CompactCategoryModuleTemplate Supportcentercompactcategorymoduletemplate `json:"compactCategoryModuleTemplate"`


    // DetailedCategoryModuleTemplate - Detailed category module template
    DetailedCategoryModuleTemplate Supportcenterdetailedcategorymoduletemplate `json:"detailedCategoryModuleTemplate"`

}

// String returns a JSON representation of the model
func (o *Supportcentermodulesetting) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcentermodulesetting) MarshalJSON() ([]byte, error) {
    type Alias Supportcentermodulesetting

    if SupportcentermodulesettingMarshalled {
        return []byte("{}"), nil
    }
    SupportcentermodulesettingMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Enabled bool `json:"enabled"`
        
        CompactCategoryModuleTemplate Supportcentercompactcategorymoduletemplate `json:"compactCategoryModuleTemplate"`
        
        DetailedCategoryModuleTemplate Supportcenterdetailedcategorymoduletemplate `json:"detailedCategoryModuleTemplate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

