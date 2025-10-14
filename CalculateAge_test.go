package utilities

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	type args struct {
		birthDate     time.Time
		referenceDate []time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantAge int
	}{
		{
			name: "Ater birthday",
			args: args{
				birthDate: time.Date(2003, time.March, 17, 0, 0, 0, 0, time.UTC),
				referenceDate: []time.Time{
					time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantAge: 22,
		},
		{
			name: "Before birthday",
			args: args{
				birthDate: time.Date(2003, time.March, 17, 0, 0, 0, 0, time.UTC),
				referenceDate: []time.Time{
					time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantAge: 21,
		},
		{
			name: "At birthday",
			args: args{
				birthDate: time.Date(2003, time.March, 17, 0, 0, 0, 0, time.UTC),
				referenceDate: []time.Time{
					time.Date(2025, time.March, 17, 0, 0, 0, 0, time.UTC),
				},
			},
			wantAge: 22,
		},
		{
			name: "Totally before",
			args: args{
				birthDate: time.Date(2003, time.March, 17, 0, 0, 0, 0, time.UTC),
				referenceDate: []time.Time{
					time.Date(2000, time.March, 17, 0, 0, 0, 0, time.UTC),
				},
			},
			wantAge: 0,
		},
		{
			name: "Ater birthday - now",
			args: args{
				birthDate:     time.Now().AddDate(-22, 0, -1),
				referenceDate: []time.Time{},
			},
			wantAge: 22,
		},
		{
			name: "Before birthday - now",
			args: args{
				birthDate:     time.Now().AddDate(-22, 0, 1),
				referenceDate: []time.Time{},
			},
			wantAge: 21,
		},
		{
			name: "At birthday - now",
			args: args{
				birthDate:     time.Now().AddDate(-22, 0, 0),
				referenceDate: []time.Time{},
			},
			wantAge: 22,
		},
		{
			name: "Totally before - now",
			args: args{
				birthDate:     time.Date(9003, time.March, 17, 0, 0, 0, 0, time.UTC),
				referenceDate: []time.Time{},
			},
			wantAge: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAge := CalculateAge(tt.args.birthDate, tt.args.referenceDate...); gotAge != tt.wantAge {
				t.Errorf("CalculateAge() = %v, want %v", gotAge, tt.wantAge)
			}
		})
	}
}
