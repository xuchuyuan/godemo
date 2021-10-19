package protocol

import (
	"log"
	"strconv"
	"strings"
)

const (
	StatusReply    = '+'
	ErrorReply     = '-'
	IntegerReply   = ':'
	BulkReply      = '$'
	MultiBulkReply = '*'

	OkReply   = "OK"
	PongReply = "PONG"
)

func GetReply(reply []byte) (interface{}, error) {
	replyType := reply[0]
	switch replyType {
	case StatusReply:
		return doStatusReply(reply[1:])
	case ErrorReply:
		return doStatusReply(reply[1:])
	case IntegerReply:
		return doIntegerReply(reply[1:])
	case BulkReply:
		return doBulkReply(reply[1:])
	case MultiBulkReply:
		return doMultiBulkReply(reply[1:])
	default:
		return nil, nil
	}
}

/*正确 | 错误 返回*/
func doStatusReply(reply []byte) (string, error) {
	return string(reply), nil
}

/*整数返回*/
func doIntegerReply(reply []byte) (int, error) {
	pos := getFlagPos('\r', reply)
	result, err := strconv.Atoi(string(reply[:pos]))
	if err != nil {
		return 0, err
	}

	return result, nil
}

/*获取内容*/
func doBulkReply(reply []byte) (interface{}, error) {
	pos := getFlagPos('\r', reply)
	vlen, err := strconv.Atoi(string(reply[:pos]))
	if err != nil {
		return nil, err
	}

	start := pos + 2
	end := start + vlen

	return string(reply[start:end]), nil
}

/*批量获取多个值*/
func doMultiBulkReply(reply []byte) (interface{}, error) {
	replyStrs := strings.Split(string(reply), "\r\n")
	loop, err := strconv.Atoi(replyStrs[0])
	if err != nil {
		return nil, err
	}

	result := []string{}
	log.Printf("replyStrs=%v, loop=%v", len(replyStrs[7:8]), loop)
	for i := 2; i < len(replyStrs); i++ {
		if i%2 == 0 {
			result = append(result, replyStrs[i])
		}
	}
	return result, nil
}

/*寻找位置*/
func getFlagPos(flag byte, reply []byte) int {
	pos := 0
	for _, v := range reply {
		if v == flag {
			break
		}
		pos++
	}

	return pos
}
