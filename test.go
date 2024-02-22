package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestX(t *testing.T) {

	type Y struct {
		A string
		B string
		C *string
	}

	type Xsub1 struct {
		YType *Y
	}

	type Xsub struct {
		Xsub1
	}

	type X struct {
		Xsub
		NumExport int
		//numUnExport int
		CreateAt time.Time
		UpdateAt time.Time
	}

	yVal1C := "test1"
	//yVal2C := "test1"

	yVal1 := Y{"test", "test1", &yVal1C}
	//yVal2 := Y{"test", "test2", &yVal2C}

	num1 := X{
		Xsub: Xsub{
			Xsub1{
				YType: &yVal1,
			},
		},
		NumExport: 100,
		//numUnExport: -1,
		CreateAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdateAt: time.Now(),
	}
	num2 := X{}

	opts := []cmp.Option{
		//cmpopts.IgnoreUnexported(X{}),
		cmpopts.IgnoreFields(X{}, "CreateAt", "UpdateAt"),
		//cmpopts.IgnoreFields(Y{}, "B"),
	}

	if diff := cmp.Diff(num1, num2, opts...); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}
