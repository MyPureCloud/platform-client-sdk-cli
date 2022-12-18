package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaregionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaregionsDud struct { 
    AwsHomeRegion string `json:"awsHomeRegion"`


    AwsCoreRegions []string `json:"awsCoreRegions"`


    AwsSatelliteRegions []string `json:"awsSatelliteRegions"`

}

// Mediaregions
type Mediaregions struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Mediaregions) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaregions) MarshalJSON() ([]byte, error) {
    type Alias Mediaregions

    if MediaregionsMarshalled {
        return []byte("{}"), nil
    }
    MediaregionsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

