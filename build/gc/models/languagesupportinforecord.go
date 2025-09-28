package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LanguagesupportinforecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LanguagesupportinforecordDud struct { 
    Language string `json:"language"`


    FeatureSupport []Featuresupport `json:"featureSupport"`

}

// Languagesupportinforecord
type Languagesupportinforecord struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Languagesupportinforecord) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Languagesupportinforecord) MarshalJSON() ([]byte, error) {
    type Alias Languagesupportinforecord

    if LanguagesupportinforecordMarshalled {
        return []byte("{}"), nil
    }
    LanguagesupportinforecordMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

