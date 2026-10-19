package cli

import "testing"

func TestIsCmdWithPassword(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		want bool
	}{
		{
			name: "password command",
			cmd:  `CREATE USER test WITH PASSWORD "secret"`,
			want: true,
		},
		{
			name: "pwd command",
			cmd:  `ALTER USER test WITH PWD "secret"`,
			want: true,
		},
		{
			name: "case insensitive",
			cmd:  `ALTER USER test WITH Password "secret"`,
			want: true,
		},
		{
			name: "ordinary command",
			cmd:  "SHOW SPACES;",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCmdWithPassword(tt.cmd); got != tt.want {
				t.Errorf("isCmdWithPassword(%q) = %v, want %v", tt.cmd, got, tt.want)
			}
		})
	}
}
