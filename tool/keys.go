package tool

import (
    "strconv"
)

func RequestKey(userId uint64) string {
    return "request:" + strconv.FormatInt(userId, 10)
}

func InboxKey(userId uint64) string {
    return "inbox:" + strconv.FormatInt(userId, 10)
}
