package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"reflect"
	"testing"
)

func TestNewAccAddressData(t *testing.T) {
	type args struct {
		value sdkTypes.AccAddress
	}
	var tests []struct {
		name string
		args args
		want types.Data
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccAddressData(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccAddressData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadAccAddressData(t *testing.T) {
	type args struct {
		dataString string
	}
	var tests []struct {
		name    string
		args    args
		want    types.Data
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAccAddressData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAccAddressData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAccAddressData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressDataFromInterface(t *testing.T) {
	type args struct {
		data types.Data
	}
	var tests []struct {
		name    string
		args    args
		want    accAddressData
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accAddressDataFromInterface(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("accAddressDataFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accAddressDataFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_AsAccAddressData(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name    string
		fields  fields
		want    sdkTypes.AccAddress
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			got, err := accAddressData.AsAccAddressData()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsAccAddressData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsAccAddressData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_AsDec(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name    string
		fields  fields
		want    sdkTypes.Dec
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			got, err := accAddressData.AsDec()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsDec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsDec() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_AsHeight(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name    string
		fields  fields
		want    types.Height
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			got, err := accAddressData.AsHeight()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsHeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsHeight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_AsID(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name    string
		fields  fields
		want    types.ID
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			got, err := accAddressData.AsID()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_AsString(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			got, err := accAddressData.AsString()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_Equal(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	type args struct {
		data types.Data
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			if got := accAddressData.Equal(tt.args.data); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name   string
		fields fields
		want   types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			if got := accAddressData.GenerateHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_Get(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name   string
		fields fields
		want   interface{}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			if got := accAddressData.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_GetTypeID(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name   string
		fields fields
		want   types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			if got := accAddressData.GetTypeID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_String(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			if got := accAddressData.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressData_ZeroValue(t *testing.T) {
	type fields struct {
		Value sdkTypes.AccAddress
	}
	var tests []struct {
		name   string
		fields fields
		want   types.Data
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			if got := accAddressData.ZeroValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
