package queuing

import "testing"

func TestTicketIDGenerator(t *testing.T) {
	type args struct {
		prefix string
	}
	var tests []struct {
		name string
		args args
		want Ticket
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TicketIDGenerator(tt.args.prefix); got != tt.want {
				t.Errorf("TicketIDGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
