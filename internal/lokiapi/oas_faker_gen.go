// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"fmt"
)

// SetFake set fake values.
func (s *Error) SetFake() {
	var unwrapped string
	{
		unwrapped = "string"
	}
	*s = Error(unwrapped)
}

// SetFake set fake values.
func (s *FPoint) SetFake() {
	{
		{
			s.T = float64(0)
		}
	}
	{
		{
			s.V = "string"
		}
	}
}

// SetFake set fake values.
func (s *IndexStats) SetFake() {
	{
		{
			s.Streams = int(0)
		}
	}
	{
		{
			s.Chunks = int(0)
		}
	}
	{
		{
			s.Entries = int(0)
		}
	}
	{
		{
			s.Bytes = int(0)
		}
	}
}

// SetFake set fake values.
func (s *LabelSet) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *Labels) SetFake() {
	{
		{
			s.Data = nil
			for i := 0; i < 0; i++ {
				var elem string
				{
					elem = "string"
				}
				s.Data = append(s.Data, elem)
			}
		}
	}
	{
		{
			s.Status = "string"
		}
	}
}

// SetFake set fake values.
func (s *LogEntry) SetFake() {
	{
		{
			s.T = uint64(0)
		}
	}
	{
		{
			s.V = "string"
		}
	}
}

// SetFake set fake values.
func (s *Maps) SetFake() {
	{
		{
			s.Data = nil
			for i := 0; i < 0; i++ {
				var elem MapsDataItem
				{
					elem.SetFake()
				}
				s.Data = append(s.Data, elem)
			}
		}
	}
	{
		{
			s.Status = "string"
		}
	}
}

// SetFake set fake values.
func (s *MapsDataItem) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *Matrix) SetFake() {
	var unwrapped []Series
	{
		unwrapped = nil
		for i := 0; i < 0; i++ {
			var elem Series
			{
				elem.SetFake()
			}
			unwrapped = append(unwrapped, elem)
		}
	}
	*s = Matrix(unwrapped)
}

// SetFake set fake values.
func (s *MatrixResult) SetFake() {
	{
		{
			s.Result.SetFake()
		}
	}
	{
		{ // Keep pointer nil to prevent infinite recursion.
			s.Stats = nil
		}
	}
}

// SetFake set fake values.
func (s *OptLabelSet) SetFake() {
	var elem LabelSet
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *Push) SetFake() {
	{
		{
			s.Streams = nil
			for i := 0; i < 0; i++ {
				var elem Stream
				{
					elem.SetFake()
				}
				s.Streams = append(s.Streams, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *QueryResponse) SetFake() {
	{
		{
			s.Status = "string"
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *QueryResponseData) SetFake() {
	var variant MatrixResult

	{
		variant.SetFake()
	}
	s.SetMatrixResult(variant)
}

// SetFake set fake values.
func (s *Sample) SetFake() {
	{
		{
			s.Metric.SetFake()
		}
	}
	{
		{
			s.Value.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ScalarResult) SetFake() {
	{
		{
			s.Result.SetFake()
		}
	}
	{
		{ // Keep pointer nil to prevent infinite recursion.
			s.Stats = nil
		}
	}
}

// SetFake set fake values.
func (s *Series) SetFake() {
	{
		{
			s.Metric.SetFake()
		}
	}
	{
		{
			s.Values = nil
			for i := 0; i < 0; i++ {
				var elem FPoint
				{
					elem.SetFake()
				}
				s.Values = append(s.Values, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *Stats) SetFake() {
}

// SetFake set fake values.
func (s *Stream) SetFake() {
	{
		{
			s.Stream.SetFake()
		}
	}
	{
		{
			s.Values = nil
			for i := 0; i < 0; i++ {
				var elem LogEntry
				{
					elem.SetFake()
				}
				s.Values = append(s.Values, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *Streams) SetFake() {
	var unwrapped []Stream
	{
		unwrapped = nil
		for i := 0; i < 0; i++ {
			var elem Stream
			{
				elem.SetFake()
			}
			unwrapped = append(unwrapped, elem)
		}
	}
	*s = Streams(unwrapped)
}

// SetFake set fake values.
func (s *StreamsResult) SetFake() {
	{
		{
			s.Result.SetFake()
		}
	}
	{
		{ // Keep pointer nil to prevent infinite recursion.
			s.Stats = nil
		}
	}
}

// SetFake set fake values.
func (s *Values) SetFake() {
	{
		{
			s.Data = nil
			for i := 0; i < 0; i++ {
				var elem string
				{
					elem = "string"
				}
				s.Data = append(s.Data, elem)
			}
		}
	}
	{
		{
			s.Status = "string"
		}
	}
}

// SetFake set fake values.
func (s *Vector) SetFake() {
	var unwrapped []Sample
	{
		unwrapped = nil
		for i := 0; i < 0; i++ {
			var elem Sample
			{
				elem.SetFake()
			}
			unwrapped = append(unwrapped, elem)
		}
	}
	*s = Vector(unwrapped)
}

// SetFake set fake values.
func (s *VectorResult) SetFake() {
	{
		{
			s.Result.SetFake()
		}
	}
	{
		{ // Keep pointer nil to prevent infinite recursion.
			s.Stats = nil
		}
	}
}
