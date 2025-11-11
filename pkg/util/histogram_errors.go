// SPDX-License-Identifier: AGPL-3.0-only

package util

import (
	"errors"
	"fmt"

	"github.com/prometheus/prometheus/model/histogram"

	"github.com/grafana/mimir/pkg/util/globalerror"
)

func ConvertHistogramErrorToGlobalError(histErr histogram.Error) (globalerror.ID, error) {
	switch {
	case errors.Is(histErr, histogram.ErrHistogramCountMismatch):
		return globalerror.NativeHistogramCountMismatch, nil
	case errors.Is(histErr, histogram.ErrHistogramCountNotBigEnough):
		return globalerror.NativeHistogramCountNotBigEnough, nil
	case errors.Is(histErr, histogram.ErrHistogramNegativeBucketCount):
		return globalerror.NativeHistogramNegativeBucketCount, nil
	case errors.Is(histErr, histogram.ErrHistogramSpanNegativeOffset):
		return globalerror.NativeHistogramSpanNegativeOffset, nil
	case errors.Is(histErr, histogram.ErrHistogramSpansBucketsMismatch):
		return globalerror.NativeHistogramSpansBucketsMismatch, nil
	case errors.Is(histErr, histogram.ErrHistogramCustomBucketsMismatch):
		return globalerror.NativeHistogramCustomBucketsMismatch, nil
	case errors.Is(histErr, histogram.ErrHistogramCustomBucketsInvalid):
		return globalerror.NativeHistogramCustomBucketsInvalid, nil
	case errors.Is(histErr, histogram.ErrHistogramCustomBucketsInfinite):
		return globalerror.NativeHistogramCustomBucketsInfinite, nil
	default:
		return "", fmt.Errorf("unknown histogram error: %w", histErr)
	}
}
