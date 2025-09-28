package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FeaturesupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FeaturesupportDud struct { 
    Feature string `json:"feature"`


    SupportLevel string `json:"supportLevel"`


    Details string `json:"details"`

}

// Featuresupport
type Featuresupport struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Featuresupport) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Featuresupport) MarshalJSON() ([]byte, error) {
    type Alias Featuresupport

    if FeaturesupportMarshalled {
        return []byte("{}"), nil
    }
    FeaturesupportMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

