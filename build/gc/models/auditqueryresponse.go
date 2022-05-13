package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryresponseDud struct { }

// Auditqueryresponse
type Auditqueryresponse struct { }

// String returns a JSON representation of the model
func (o *Auditqueryresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryresponse

    if AuditqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

