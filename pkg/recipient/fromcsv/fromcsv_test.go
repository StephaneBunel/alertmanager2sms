package fromcsv_test

import (
	"testing"

	"github.com/StephaneBunel/alertmanager2sms/pkg/recipient/fromcsv"

	"github.com/spf13/viper"
)

func TestFromFile(t *testing.T) {
	conf := viper.New()
	conf.Set("filename", "../../../test/recipients.csv")

	recipRepo := fromcsv.New()
	err := recipRepo.Config(conf)
	if err != nil {
		t.Error(err)
		return
	}
	recipients := recipRepo.FindByName("bob")
	if len(recipients) < 1 {
		t.Error("Cannot find bob !!")
		return
	}
	phone0 := recipients[0].PhoneNumbers[0]
	if phone0 != "+123456789" {
		t.Error("bob phones[0] should be +123456789. Not", phone0)
		return
	}
}
