package encoding

func TestCsv2Strings(t *testing.T) {
	type args struct {
		csvSource []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CSV with commas",
			args: args{
				csvSource: []byte("Source,English\nhello,world\n\"hel,lo\",\"wor,ld\"\n"),
			},
			want:    "\"hello\" = \"world\";\n\"hel,lo\" = \"wor,ld\";",
			wantErr: false,
		},
		{
			name: "Empty CSV",
			args: args{
				csvSource: []byte("Source,English"),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Csv2Strings(tt.args.csvSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("Csv2Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Csv2Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXml2Csv(t *testing.T) {
	type args struct {
		xmlData []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Single line",
			args: args{
				xmlData: []byte("<xml><string name=\"hello\">world</string></xml>"),
			},
			want:    []byte("key,translation\nhello,world"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Xml2Csv(tt.args.xmlData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Xml2Csv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Xml2Csv() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
