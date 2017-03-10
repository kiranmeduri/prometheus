// TODO: package name, header

package merger

import (
	"time"

	"golang.org/x/net/context"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage/local"
	"github.com/prometheus/prometheus/storage/metric"
)

// Fanin is a local.Queryable that reads from multiple other Queryables.
type Fanin []promql.Queryable

// Querier implements local.Queryable.
func (f Fanin) Querier() (local.Querier, error) {
	queriers := make(faninQuerier, 0, len(f))
	for _, q := range f {
		querier, err := q.Querier()
		if err != nil {
			return nil, err
		}
		queriers = append(queriers, querier)
	}
	return queriers, nil
}

type faninQuerier []local.Querier

func (f faninQuerier) Close() error {
	for _, q := range f {
		if err := q.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (f faninQuerier) QueryRange(ctx context.Context, from, through model.Time, matchers ...*metric.LabelMatcher) ([]local.SeriesIterator, error) {
	return nil, nil
}

func (f faninQuerier) QueryInstant(ctx context.Context, ts model.Time, stalenessDelta time.Duration, matchers ...*metric.LabelMatcher) ([]local.SeriesIterator, error) {
	fpToIt := map[model.Fingerprint]local.SeriesIterator{}
	for _, q := range f {
		its, err := q.QueryInstant(ctx, ts, stalenessDelta, matchers...)
		if err != nil {
			return nil, err
		}
		mergeIterators(fpToIt, its)
	}
	its := make([]local.SeriesIterator, 0, len(fpToIt))
	for _, it := range fpToIt {
		its = append(its, it)
	}
	return nil, nil
}

func (f faninQuerier) MetricsForLabelMatchers(ctx context.Context, from, through model.Time, matcherSets ...metric.LabelMatchers) ([]metric.Metric, error) {
	return nil, nil
}

func (f faninQuerier) LastSampleForLabelMatchers(ctx context.Context, cutoff model.Time, matcherSets ...metric.LabelMatchers) (model.Vector, error) {
	return nil, nil
}

func (f faninQuerier) LabelValuesForLabelName(context.Context, model.LabelName) (model.LabelValues, error) {
	return nil, nil
}

type fanInIterator []local.SeriesIterator

func (f fanInIterator) ValueAtOrBeforeTime(t model.Time) model.SamplePair {
	latest := model.ZeroSamplePair
	for _, it := range f {
		v := it.ValueAtOrBeforeTime(t)
		if v.Timestamp.After(latest.Timestamp) {
			latest = v
		}
	}
	return latest
}

func (f fanInIterator) RangeValues(interval metric.Interval) []model.SamplePair {
	var values []model.SamplePair
	for _, it := range f {
		values = mergeSamples(values, it.RangeValues(interval))
	}
	return values
}

func (f fanInIterator) Metric() metric.Metric {
	return f[0].Metric()
}

func (f fanInIterator) Close() {
	for _, it := range f {
		it.Close()
	}
}

func mergeIterators(fpToIt map[model.Fingerprint]local.SeriesIterator, its []local.SeriesIterator) {
	// TODO
}

func mergeSamples(a, b []model.SamplePair) []model.SamplePair {
	result := make([]model.SamplePair, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i].Timestamp < b[j].Timestamp {
			result = append(result, a[i])
			i++
		} else if a[i].Timestamp > b[j].Timestamp {
			result = append(result, b[j])
			j++
		} else {
			result = append(result, a[i])
			i++
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}
