package main

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func Test_handlers_validateData(t *testing.T) {
	v1 := VerifyRequest{
		Name: "Boby",
	}
	v2 := VerifyRequest{
		Name:   "Martin",
		Family: "Fowler",
	}

	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		r VerifyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{

		{
			name: "should return err",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				r: v1,
			},
			wantErr: true,
		},
		{
			name: "no error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				r: v2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlers{
				validator: tt.fields.validator,
			}
			if err := h.validateData(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
