package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterdatahistoricaltweetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterdatahistoricaltweetresponseDud struct { 
    


    


    

}

// Twitterdatahistoricaltweetresponse
type Twitterdatahistoricaltweetresponse struct { 
    // MatchingTweets - The count of matching tweets using the searchTerms.
    MatchingTweets int `json:"matchingTweets"`


    // PreviousNumberOfDays - The number of days used for the query
    PreviousNumberOfDays int `json:"previousNumberOfDays"`


    // UsageStatistics - The tweet usage of the org
    UsageStatistics Tweetusage `json:"usageStatistics"`

}

// String returns a JSON representation of the model
func (o *Twitterdatahistoricaltweetresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterdatahistoricaltweetresponse) MarshalJSON() ([]byte, error) {
    type Alias Twitterdatahistoricaltweetresponse

    if TwitterdatahistoricaltweetresponseMarshalled {
        return []byte("{}"), nil
    }
    TwitterdatahistoricaltweetresponseMarshalled = true

    return json.Marshal(&struct {
        
        MatchingTweets int `json:"matchingTweets"`
        
        PreviousNumberOfDays int `json:"previousNumberOfDays"`
        
        UsageStatistics Tweetusage `json:"usageStatistics"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

