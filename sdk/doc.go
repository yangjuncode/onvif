package sdk

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/juju/errors"
	"github.com/rs/zerolog"
)

var (
	// LoggerContext is the builder of a zerolog.Logger that is exposed to the application so that
	// options at the CLI might alter the formatting and the output of the logs.
	LoggerContext = zerolog.
			New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().Timestamp()

	// Logger is a zerolog logger, that can be safely used from any part of the application.
	// It gathers the format and the output.
	Logger = LoggerContext.Logger()
)

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	Logger.Debug().
		Str("msg", httpReply.Status).
		Int("status", httpReply.StatusCode).
		Str("action", tag).
		Msg("RPC")
	// TODO(jfsmig): extract the deadline from ctx.Deadline() and apply it on the reply reading
	if b, err := ioutil.ReadAll(httpReply.Body); err != nil {
		return errors.Annotate(err, "read")
	} else {
		// my case
		if tag == "Subscribe" ||
			tag == "CreatePullPointSubscription" ||
			tag == "GetStreamUri" ||
			tag == "GetSnapshotUri" ||
			tag == "GetEventProperties" ||
			tag == "GetServiceCapabilities" ||
			tag == "Unsubscribe" ||
			tag == "GetCapabilities" {
			fmt.Printf("body received for %s:%s\n", tag, b)
		}

		err = xml.Unmarshal(b, reply)
		return errors.Annotate(err, "decode")
	}
}