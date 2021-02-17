package model

import (
	"theo.dev/hello-universe/speedtester/model/PerformanceResult"
)

type TestResult struct {
	Id          int
	Performance PerformanceResult.PerformanceResult
}
