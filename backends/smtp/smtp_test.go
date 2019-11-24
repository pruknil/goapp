package smtp

import "testing"

func TestSendMail(t *testing.T) {
	type args struct {
		target  string
		body    string
		subject string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case1", args: args{
			target:  "pruknil@gmail.com",
			body:    "Hahaha test",
			subject: "test email",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.target, tt.args.body, tt.args.subject); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
