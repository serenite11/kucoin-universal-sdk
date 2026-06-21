package util

import (
	"fmt"
	"sort"
	"strings"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/interfaces"
)

const EMPTY_ARGS_STR = "EMPTY_ARGS"

type SubInfo struct {
	Prefix   string
	Args     []string
	Callback interfaces.WebSocketMessageCallback
}

func (info *SubInfo) ToId() string {
	sort.Strings(info.Args)
	argsStr := EMPTY_ARGS_STR
	if len(info.Args) > 0 {
		argsStr = strings.Join(info.Args, ",")
	}
	return fmt.Sprintf("%s@@%s", info.Prefix, argsStr)
}

func (info *SubInfo) FromId(id string) error {
	parts := strings.Split(id, "@@")
	if len(parts) != 2 {
		return fmt.Errorf("SubInfo.FromId: invalid id format: %s", id)
	}

	info.Prefix = parts[0]
	if parts[1] != EMPTY_ARGS_STR {
		info.Args = strings.Split(parts[1], ",")
	} else {
		info.Args = []string{}
	}
	return nil
}

func (info *SubInfo) Topics() []string {
	if info.Args == nil || len(info.Args) == 0 {
		return []string{info.Prefix}
	}

	ret := make([]string, 0, len(info.Args))
	for _, i := range info.Args {
		ret = append(ret, fmt.Sprintf("%s:%s", info.Prefix, i))
	}
	return ret
}

func (info *SubInfo) SubTopic() string {
	if info.Args == nil || len(info.Args) == 0 {
		return info.Prefix
	}
	return fmt.Sprintf("%s:%s", info.Prefix, strings.Join(info.Args, ","))
}
