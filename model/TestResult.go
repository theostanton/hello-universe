package model

import (
	"theo.dev/hello-universe/model/PerformanceResult"
)

type TestResult struct {
	Id          int
	Performance PerformanceResult.PerformanceResult
}
