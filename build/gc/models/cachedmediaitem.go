package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CachedmediaitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CachedmediaitemDud struct { 
    Id string `json:"id"`


    Url string `json:"url"`


    DownloadUrl string `json:"downloadUrl"`


    MediaType string `json:"mediaType"`


    ContentLengthBytes int `json:"contentLengthBytes"`


    DateCreated time.Time `json:"dateCreated"`


    DateExpires time.Time `json:"dateExpires"`


    SelfUri string `json:"selfUri"`

}

// Cachedmediaitem - Defines an external media that has been ingested and cached by Genesys Cloud for conversation messaging
type Cachedmediaitem struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Cachedmediaitem) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cachedmediaitem) MarshalJSON() ([]byte, error) {
    type Alias Cachedmediaitem

    if CachedmediaitemMarshalled {
        return []byte("{}"), nil
    }
    CachedmediaitemMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

