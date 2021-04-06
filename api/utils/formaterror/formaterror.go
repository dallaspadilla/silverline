package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "packetLoss") {
		return errors.New("packetLoss is wrong")
	}

	if strings.Contains(err, "minRoundTrip") {
		return errors.New("minRoundTrip is wrong")
	}

	if strings.Contains(err, "maxRoundTrip") {
		return errors.New("maxRoundTrip is wrong")
	}
	if strings.Contains(err, "avgRoundTrip") {
		return errors.New("avgRoundTrip is wrong")
	}
	return errors.New("Incorrect Details")
}
