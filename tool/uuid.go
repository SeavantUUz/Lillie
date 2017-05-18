package tool

import (
    "crypto/rand"
    "fmt"
)

func UUID() (uuid string) {
    b := make([]byte, 8)
    if _, err := rand.Read(b); err != nil {
        return "00000000"
    }
    uuid = fmt.Sprintf("%X", b)
    return
}
