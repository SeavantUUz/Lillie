package tool

import (
    "github.com/SeavantUUz/Lillie/protocol"
    "strconv"
)

func RequestKey(request *protocol.Request) string {
    return "request:" + strconv.FormatInt(int64(request.Operation), 10)
}
