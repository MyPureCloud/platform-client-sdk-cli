package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RmsassetaddressablerefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RmsassetaddressablerefDud struct { 
    


    


    


    

}

// Rmsassetaddressableref
type Rmsassetaddressableref struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // AssetUsage - Usage of the asset. Null for inline images, 'FileAttachments' for file attachments
    AssetUsage string `json:"assetUsage"`


    // ContentLocation - Content location URL for the asset
    ContentLocation string `json:"contentLocation"`

}

// String returns a JSON representation of the model
func (o *Rmsassetaddressableref) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rmsassetaddressableref) MarshalJSON() ([]byte, error) {
    type Alias Rmsassetaddressableref

    if RmsassetaddressablerefMarshalled {
        return []byte("{}"), nil
    }
    RmsassetaddressablerefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        AssetUsage string `json:"assetUsage"`
        
        ContentLocation string `json:"contentLocation"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

