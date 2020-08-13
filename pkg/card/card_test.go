package card

import "testing"

func TestMapRowToTransaction(t *testing.T) {
	type args struct {
		row []string
	}
	tests := []struct {
		name       string
		args       args
		wantAmount int64
		wantOwner  int
	}{
		{
			name:       "normal",
			args:       args{row: []string{"151801", "6040", "4"}},
			wantAmount: 151801,
			wantOwner:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAmount, gotOwner := MapRowToTransaction(tt.args.row)
			if gotAmount != tt.wantAmount {
				t.Errorf("MapRowToTransaction() gotAmount = %v, want %v", gotAmount, tt.wantAmount)
			}
			if gotOwner != tt.wantOwner {
				t.Errorf("MapRowToTransaction() gotOwner = %v, want %v", gotOwner, tt.wantOwner)
			}
		})
	}
}
