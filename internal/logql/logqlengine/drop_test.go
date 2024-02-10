package logqlengine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/tdakkota/docker-logql/internal/logql"
)

func TestDropLabels(t *testing.T) {
	tests := []struct {
		input        map[logql.Label]pcommon.Value
		drop         []logql.Label
		matchers     []logql.LabelMatcher
		expectLabels map[logql.Label]pcommon.Value
	}{
		{
			map[logql.Label]pcommon.Value{},
			[]logql.Label{`foo`},
			nil,
			map[logql.Label]pcommon.Value{},
		},
		{
			map[logql.Label]pcommon.Value{
				`foo`: pcommon.NewValueStr(`1`),
				`bar`: pcommon.NewValueStr(`2`),
				`baz`: pcommon.NewValueStr(`3`),
			},
			[]logql.Label{`foo`},
			nil,
			map[logql.Label]pcommon.Value{
				`bar`: pcommon.NewValueStr(`2`),
				`baz`: pcommon.NewValueStr(`3`),
			},
		},
		{
			map[logql.Label]pcommon.Value{
				`bar`: pcommon.NewValueStr(`2`),
			},
			[]logql.Label{`foo`},
			nil,
			map[logql.Label]pcommon.Value{
				`bar`: pcommon.NewValueStr(`2`),
			},
		},

		// Label matcher tests.
		{
			map[logql.Label]pcommon.Value{
				`foo`: pcommon.NewValueStr(`1`),
				`bar`: pcommon.NewValueStr(`2`),
				`baz`: pcommon.NewValueStr(`3`),
			},
			[]logql.Label{`foo`},
			[]logql.LabelMatcher{
				{Label: `bar`, Op: logql.OpEq, Value: `2`},
			},
			map[logql.Label]pcommon.Value{
				`baz`: pcommon.NewValueStr(`3`),
			},
		},
		{
			map[logql.Label]pcommon.Value{
				`foo`: pcommon.NewValueStr(`1`),
				`bar`: pcommon.NewValueStr(`2`),
				`baz`: pcommon.NewValueStr(`3`),
			},
			[]logql.Label{`foo`},
			[]logql.LabelMatcher{
				{Label: `bar`, Op: logql.OpNotEq, Value: `2`},
			},
			map[logql.Label]pcommon.Value{
				`bar`: pcommon.NewValueStr(`2`),
				`baz`: pcommon.NewValueStr(`3`),
			},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			e, err := buildDropLabels(&logql.DropLabelsExpr{
				Labels:   tt.drop,
				Matchers: tt.matchers,
			})
			require.NoError(t, err)

			set := newLabelSet()
			set.labels = tt.input
			newLine, ok := e.Process(0, ``, set)
			// Ensure that processor does not change the line.
			require.Equal(t, ``, newLine)
			require.True(t, ok)

			for k, expect := range tt.expectLabels {
				got, ok := set.Get(k)
				require.Truef(t, ok, "key %q", k)
				require.Equal(t, expect, got)
			}
		})
	}
}
