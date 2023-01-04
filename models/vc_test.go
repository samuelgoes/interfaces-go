package models

import (
	"reflect"
	"testing"
)

func TestJWTVerifiableCredential_GetCredentialSubject(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			if got := jvc.GetCredentialSubject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTVerifiableCredential_GetFormat(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			if got := jvc.GetFormat(); got != tt.want {
				t.Errorf("GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTVerifiableCredential_GetId(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			if got := jvc.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTVerifiableCredential_GetType(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			if got := jvc.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTVerifiableCredential_MarshalJSON(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			got, err := jvc.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTVerifiableCredential_ToString(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			jvc.ToString()
		})
	}
}

func TestJWTVerifiableCredential_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Id                   string
		Issuer               string
		Subject              string
		VerifiableCredential *LDPVerifiableCredential
		Format               string
	}
	type args struct {
		jsonData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jvc := JWTVerifiableCredential{
				Id:                   tt.fields.Id,
				Issuer:               tt.fields.Issuer,
				Subject:              tt.fields.Subject,
				VerifiableCredential: tt.fields.VerifiableCredential,
				Format:               tt.fields.Format,
			}
			if err := jvc.UnmarshalJSON(tt.args.jsonData); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLDPVerifiableCredential_GetCredentialSubject(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			if got := ldp.GetCredentialSubject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDPVerifiableCredential_GetFormat(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			if got := ldp.GetFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDPVerifiableCredential_GetId(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			if got := ldp.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDPVerifiableCredential_GetType(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			if got := ldp.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDPVerifiableCredential_MarshalJSON(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			got, err := ldp.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDPVerifiableCredential_ToString(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			ldp.ToString()
		})
	}
}

func TestLDPVerifiableCredential_UnmarshalJSON(t *testing.T) {
	type fields struct {
		CredentialSubject map[string]interface{}
		Id                string
		Type              []string
		Format            string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldp := &LDPVerifiableCredential{
				CredentialSubject: tt.fields.CredentialSubject,
				Id:                tt.fields.Id,
				Type:              tt.fields.Type,
				Format:            tt.fields.Format,
			}
			if err := ldp.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewJWTVerifiableCredential(t *testing.T) {
	tests := []struct {
		name string
		want JWTVerifiableCredential
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWTVerifiableCredential(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWTVerifiableCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLDPVerifiableCredential(t *testing.T) {
	tests := []struct {
		name string
		want LDPVerifiableCredential
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLDPVerifiableCredential(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLDPVerifiableCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}
