package helpers

import (
	"Restobook/delivery/common"
	"fmt"
	"strconv"
)

func NumberToDayConverter(openDay, closeDay []string) (string, string, error) {
	openStr := ""
	closeStr := ""
	for j := 0; j < len(openDay); j++ {
		for k := 0; k < len(common.Daytoint); k++ {
			if openDay[j] == strconv.Itoa(common.Daytoint[k].No) {
				openStr += fmt.Sprintf("%v,", common.Daytoint[k].Day)
			}
		}
	}
	for l := 0; l < len(closeDay); l++ {
		for m := 0; m < len(common.Daytoint); m++ {
			if closeDay[l] == strconv.Itoa(common.Daytoint[m].No) {
				closeStr += fmt.Sprintf("%v,", common.Daytoint[m].Day)
			}
		}
	}
	return openStr, closeStr, nil
}
