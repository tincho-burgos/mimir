// SPDX-License-Identifier: AGPL-3.0-only

package util

import (
	"errors"
	"fmt"

	"github.com/prometheus/prometheus/model/histogram"

	"github.com/grafana/mimir/pkg/util/globalerror"
)

func ConvertHistogramErrorToGlobalError(err error) (globalerror.ID, error) {
	switch {
	case errors.Is(err, histogram.ErrHistogramCountMismatch):
		return globalerror.NativeHistogramCountMismatch, nil
	case errors.Is(err, histogram.ErrHistogramCountNotBigEnough):
		return globalerror.NativeHistogramCountNotBigEnough, nil
	case errors.Is(err, histogram.ErrHistogramNegativeBucketCount):
		return globalerror.NativeHistogramNegativeBucketCount, nil
	case errors.Is(err, histogram.ErrHistogramSpanNegativeOffset):
		return globalerror.NativeHistogramSpanNegativeOffset, nil
	case errors.Is(err, histogram.ErrHistogramSpansBucketsMismatch):
		return globalerror.NativeHistogramSpansBucketsMismatch, nil
	case errors.Is(err, histogram.ErrHistogramCustomBucketsMismatch):
		return globalerror.NativeHistogramCustomBucketsMismatch, nil
	case errors.Is(err, histogram.ErrHistogramCustomBucketsInvalid):
		return globalerror.NativeHistogramCustomBucketsInvalid, nil
	case errors.Is(err, histogram.ErrHistogramCustomBucketsInfinite):
		return globalerror.NativeHistogramCustomBucketsInfinite, nil
	default:
		return "", fmt.Errorf("unknown histogram error: %w", err)
	}
}
