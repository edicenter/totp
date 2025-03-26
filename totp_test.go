package main

import "testing"

type Data struct {
	unixTime int64
	totp     string
	secret   string
}

func TestTOTP(t *testing.T) {

	data := []Data{}
	data = append(data, Data{unixTime: 1742912054, totp: "077949", secret: "JBSWY3DPEHPK3PXP"})
	data = append(data, Data{unixTime: 1742911803, totp: "313515", secret: "JBSWY3DPEHPK3PXP"})

	for _, d := range data {
		if gotTotp, err := GetTOTP(d.secret, d.unixTime); gotTotp != d.totp || err != nil {
			t.Errorf("Got: %v; Expected: %v; Testdata: %+v; Error: %v", gotTotp, d.totp, d, err)
		}
	}

}
