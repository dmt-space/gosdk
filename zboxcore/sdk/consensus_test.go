package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsensus_isConsensusMin(t *testing.T) {
	type fields struct {
		consensus       float32
		consensusThresh float32
		fullconsensus   float32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			"Test_Is_Consensus_Min_True",
			fields{
				consensus:       2,
				consensusThresh: 50,
				fullconsensus:   4,
			},
			true,
		},
		{
			"Test_Is_Consensus_Min_False",
			fields{
				consensus:       1,
				consensusThresh: 50,
				fullconsensus:   4,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &Consensus{
				consensus:       tt.fields.consensus,
				consensusThresh: tt.fields.consensusThresh,
				fullconsensus:   tt.fields.fullconsensus,
			}
			got := req.isConsensusMin()
			assertion := assert.New(t)
			var check = assertion.False
			if tt.want {
				check = assertion.True
			}
			check(got)
		})
	}
}

func TestConsensus_isConsensusOk(t *testing.T) {
	type fields struct {
		consensus       float32
		consensusThresh float32
		fullconsensus   float32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			"Test_Is_Consensus_OK_True",
			fields{
				consensus:       3,
				consensusThresh: 50,
				fullconsensus:   4,
			},
			true,
		},
		{
			"Test_Is_Consensus_OK_False",
			fields{
				consensus:       2,
				consensusThresh: 50,
				fullconsensus:   4,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &Consensus{
				consensus:       tt.fields.consensus,
				consensusThresh: tt.fields.consensusThresh,
				fullconsensus:   tt.fields.fullconsensus,
			}
			got := req.isConsensusOk()
			assertion := assert.New(t)
			var check = assertion.False
			if tt.want {
				check = assertion.True
			}
			check(got)
		})
	}
}

func TestConsensus_getConsensusRequiredForOk(t *testing.T) {
	type fields struct {
		consensusThresh float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			"Test_Result_Get_Consensus_Required_For_Ok",
			fields{
				consensusThresh: 50,
			},
			60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			req := &Consensus{
				consensusThresh: tt.fields.consensusThresh,
			}
			got := req.getConsensusRequiredForOk()
			assertions.Equal(got, tt.want)
		})
	}
}

func TestConsensus_getConsensusRate(t *testing.T) {
	type fields struct {
		consensus     float32
		fullconsensus float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			"Test_Result_Get_Consensus_Required_For_Ok",
			fields{
				consensus:     2,
				fullconsensus: 4,
			},
			50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			req := &Consensus{
				consensus:     tt.fields.consensus,
				fullconsensus: tt.fields.fullconsensus,
			}
			got := req.getConsensusRate()
			assertions.Equal(got, tt.want)
		})
	}
}