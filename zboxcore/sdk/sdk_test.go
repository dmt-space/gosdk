package sdk

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSdk_GetAllocationMinLock(t *testing.T) {
	type fields struct {
		consensus     float32
		fullconsensus float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test_Result_Get_Consensus_Required_For_Ok",
			fields{
				consensus:     2,
				fullconsensus: 4,
			},
			"50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			// req := &Consensus{
			// 	consensus:     tt.fields.consensus,
			// 	fullconsensus: tt.fields.fullconsensus,
			// }
			sdkInitialized = true
			got, err  := GetAllocationMinLock(1,1,1,1,PriceRange{},PriceRange{},time.Duration(time.Hour))
			fmt.Println("-------",got,"==",err)
			assertions.Equal(err, true)
			assertions.Equal(got, tt.want)
		})
	}
}

func TestSdk_CreateAllocationForOwner(t *testing.T) {
	type fields struct {
		consensus     float32
		fullconsensus float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test_Result_Get_Consensus_Required_For_Ok",
			fields{
				consensus:     2,
				fullconsensus: 4,
			},
			"50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			// req := &Consensus{
			// 	consensus:     tt.fields.consensus,
			// 	fullconsensus: tt.fields.fullconsensus,
			// }
			sdkInitialized = true
			got, err  := CreateAllocationForOwner("","",1,1,1,1,PriceRange{},PriceRange{},time.Duration(time.Hour),1,[]string{})
			fmt.Println("-------",got,"==",err)
			assertions.Equal(err, true)
			assertions.Equal(got, tt.want)
		})
	}
}
