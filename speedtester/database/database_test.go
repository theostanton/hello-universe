package database

import (
	"testing"
	"theo.dev/hello-universe/speedtester/model"
	"theo.dev/hello-universe/speedtester/model/PerformanceResult"
)

type args struct {
	id int
}

type testCase struct {
	name       string
	args       args
	wantResult PerformanceResult.PerformanceResult
	wantErr    bool
}

func TestSelect(t *testing.T) {
	testCases := []testCase{

		{
			name: "Failed",
			args: args{
				id: 1,
			},
			wantErr: true,
		},

		{
			name: "Success",
			args: args{
				id: 2,
			},
			wantResult: "AVERAGE",
			wantErr:    false,
		},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Select(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Select() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		testResult model.TestResult
	}
	tests := []struct {
		name    string
		args    args
		wantId  int
		wantErr bool
	}{
		{
			name: "Successful insert",
			args: args{
				testResult: model.TestResult{
					Performance: "AVERAGE",
				},
			},
			wantId:  2,
			wantErr: false,
		},
		{
			name: "Unsuccessful insert - duplicate primary key",
			args: args{
				testResult: model.TestResult{
					Performance: "AVERAGE",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := Insert(tt.args.testResult)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Insert() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
