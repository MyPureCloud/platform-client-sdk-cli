package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TweetusageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TweetusageDud struct { 
    


    


    

}

// Tweetusage
type Tweetusage struct { 
    // IngestionLimit - Ingestion limit
    IngestionLimit int `json:"ingestionLimit"`


    // TweetCount - The number of tweets consumed
    TweetCount int `json:"tweetCount"`


    // DateStart - The start of the usage period for the currently consumed tweets. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`

}

// String returns a JSON representation of the model
func (o *Tweetusage) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Tweetusage) MarshalJSON() ([]byte, error) {
    type Alias Tweetusage

    if TweetusageMarshalled {
        return []byte("{}"), nil
    }
    TweetusageMarshalled = true

    return json.Marshal(&struct {
        
        IngestionLimit int `json:"ingestionLimit"`
        
        TweetCount int `json:"tweetCount"`
        
        DateStart time.Time `json:"dateStart"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

