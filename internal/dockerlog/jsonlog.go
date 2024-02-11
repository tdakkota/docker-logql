package dockerlog

import (
	"io"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/tdakkota/docker-logql/internal/iterators"
	"github.com/tdakkota/docker-logql/internal/logstorage"
	"github.com/tdakkota/docker-logql/internal/otelstorage"
)

// ParseFile parses Docker JSON log file.
func ParseFile(f io.ReadCloser, resource otelstorage.Attrs) (iterators.Iterator[logstorage.Record], error) {
	return &jsonIter{
		dec:      jx.Decode(f, 4096),
		closer:   f,
		err:      nil,
		resource: resource,
	}, nil
}

// jsonIter iterates over JSON log file made by docker.
type jsonIter struct {
	dec    *jx.Decoder
	closer io.Closer
	err    error

	resource otelstorage.Attrs
}

var _ iterators.Iterator[logstorage.Record] = (*jsonIter)(nil)

// Next returns true, if there is element and fills t.
func (i *jsonIter) Next(r *logstorage.Record) (ok bool) {
	// Reset record.
	*r = logstorage.Record{
		Attrs:         otelstorage.Attrs(pcommon.NewMap()),
		ResourceAttrs: i.resource,
	}

	ok, i.err = i.parseNext(r)
	return ok
}

func (i *jsonIter) parseNext(r *logstorage.Record) (bool, error) {
	if tt := i.dec.Next(); tt == jx.Invalid {
		return false, nil
	}
	if err := i.dec.ObjBytes(func(d *jx.Decoder, key []byte) (err error) {
		switch string(key) {
		case "log":
			r.Body, err = d.Str()
			if err != nil {
				return errors.Wrap(err, "read log message")
			}
			return nil
		case "stream":
			stream, err := d.Str()
			if err != nil {
				return errors.Wrap(err, "read stream")
			}

			m := r.Attrs.AsMap()
			m.PutStr("stream", stream)
			return nil
		case "time":
			t, err := d.Str()
			if err != nil {
				return errors.Wrap(err, "read timestamp")
			}

			ts, err := time.Parse(time.RFC3339Nano, t)
			if err != nil {
				return errors.Wrap(err, "parse timestamp")
			}

			r.Timestamp = otelstorage.NewTimestampFromTime(ts)
			r.ObservedTimestamp = otelstorage.NewTimestampFromTime(ts)
			return nil
		default:
			return d.Skip()
		}
	}); err != nil {
		return false, errors.Wrap(err, "read json")
	}
	return true, nil
}

// Err returns an error caused during iteration, if any.
func (i *jsonIter) Err() error {
	return i.err
}

// Close closes iterator.
func (i *jsonIter) Close() error {
	return i.closer.Close()
}
