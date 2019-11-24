package smtp

import "testing"

func TestMySmtp_BuildMail(t *testing.T) {
	type fields struct {
		ISmtp ISmtp
		umail string
		upw   string
		host  string
	}
	type args struct {
		target  string
		body    string
		subject string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Case1", fields: struct {
			ISmtp ISmtp
			umail string
			upw   string
			host  string
		}{ISmtp: &MockSmtp{}, umail: "p_nilsuriyakon@hotmail.com", upw: "Aoom1346", host: "smtp.office365.com:587"}, args: struct {
			target  string
			body    string
			subject string
		}{target: "pruknil@gmail.com", body: "Test Body", subject: "Title naja"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MailService{
				ISmtp: tt.fields.ISmtp,
				Config: Config{
					from:     tt.fields.umail,
					password: tt.fields.upw,
					host:     tt.fields.host,
				},
			}
			if err := s.BuildMail(tt.args.target, tt.args.body, tt.args.subject); (err != nil) != tt.wantErr {
				t.Errorf("BuildMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
