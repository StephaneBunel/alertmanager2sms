package fromcsv

import (
	"testing"

	// "github.com/StephaneBunel/alertmanager2sms/pkg/recipient/handler/fromcsv"

	"github.com/spf13/viper"
)

func TestFromFile(t *testing.T) {
	conf := viper.New()
	conf.Set("filename", "../../../../test/recipient.handler.fromcsv/recipients.csv")

	recipRepo := New()
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
		t.Error("bob phones[0] should be +123456789. Got:", phone0)
		return
	}
}
