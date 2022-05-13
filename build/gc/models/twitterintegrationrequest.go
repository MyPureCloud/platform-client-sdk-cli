package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Twitterintegrationrequest
type Twitterintegrationrequest struct { 
    


    // Name - The name of the Twitter Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // AccessTokenKey - The Access Token Key from Twitter messenger
    AccessTokenKey string `json:"accessTokenKey"`


    // AccessTokenSecret - The Access Token Secret from Twitter messenger
    AccessTokenSecret string `json:"accessTokenSecret"`


    // ConsumerKey - The Consumer Key from Twitter messenger
    ConsumerKey string `json:"consumerKey"`


    // ConsumerSecret - The Consumer Secret from Twitter messenger
    ConsumerSecret string `json:"consumerSecret"`


    // Tier - The type of twitter account to be used for the integration
    Tier string `json:"tier"`


    // EnvName - The Twitter environment name, e.g.: env-beta (required for premium tier)
    EnvName string `json:"envName"`


    

}

// String returns a JSON representation of the model
func (o *Twitterintegrationrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Twitterintegrationrequest

    if TwitterintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    TwitterintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        AccessTokenKey string `json:"accessTokenKey"`
        
        AccessTokenSecret string `json:"accessTokenSecret"`
        
        ConsumerKey string `json:"consumerKey"`
        
        ConsumerSecret string `json:"consumerSecret"`
        
        Tier string `json:"tier"`
        
        EnvName string `json:"envName"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

