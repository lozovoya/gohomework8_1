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
		// TODO: Add test cases.
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
