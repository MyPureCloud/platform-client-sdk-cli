package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingstickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingstickerDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Messagingsticker
type Messagingsticker struct { 
    


    // Name
    Name string `json:"name"`


    // ProviderStickerId - The sticker Id of the sticker, assigned by the sticker provider.
    ProviderStickerId int `json:"providerStickerId"`


    // ProviderPackageId - The package Id of the sticker, assigned by the sticker provider.
    ProviderPackageId int `json:"providerPackageId"`


    // PackageName - The package name of the sticker, assigned by the sticker provider.
    PackageName string `json:"packageName"`


    // MessengerType - The type of the messenger provider.
    MessengerType string `json:"messengerType"`


    // StickerType - The type of the sticker.
    StickerType string `json:"stickerType"`


    // ProviderVersion - The version of the sticker, assigned by the provider.
    ProviderVersion int `json:"providerVersion"`


    // UriLocation
    UriLocation string `json:"uriLocation"`


    

}

// String returns a JSON representation of the model
func (o *Messagingsticker) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingsticker) MarshalJSON() ([]byte, error) {
    type Alias Messagingsticker

    if MessagingstickerMarshalled {
        return []byte("{}"), nil
    }
    MessagingstickerMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ProviderStickerId int `json:"providerStickerId"`
        
        ProviderPackageId int `json:"providerPackageId"`
        
        PackageName string `json:"packageName"`
        
        MessengerType string `json:"messengerType"`
        
        StickerType string `json:"stickerType"`
        
        ProviderVersion int `json:"providerVersion"`
        
        UriLocation string `json:"uriLocation"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

