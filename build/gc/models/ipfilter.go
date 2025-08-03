package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IpfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IpfilterDud struct { }

// Ipfilter - Configuration for filtering tracking based on IP addresses.
type Ipfilter struct { }

// String returns a JSON representation of the model
func (o *Ipfilter) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ipfilter) MarshalJSON() ([]byte, error) {
    type Alias Ipfilter

    if IpfilterMarshalled {
        return []byte("{}"), nil
    }
    IpfilterMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

