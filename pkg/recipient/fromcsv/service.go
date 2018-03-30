package fromcsv

import (
	"container/list"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

func (rr *csvRecipientRepository) ReadFromFile(filename string) error {
	csvFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return rr.ReadFromString(string(csvFile))
}

func (rr *csvRecipientRepository) ReadFromString(body string) error {
	r := csv.NewReader(strings.NewReader(body))
	r.Comment = '#'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) >= 2 {
			recip := domain.NewRecipient()
			recip.Name = record[0]
			recip.PhoneNumbers = strings.Split(record[1], ":")
			err := rr.Add(recip)
			if err != nil {
				return err
			}
			rlog.Debug(recip)
		} else {
			rlog.Warn("Invalide CSV record:", record)
		}
	}
	return nil
}

// -- begin of domain.IRecipientRepository interface
func (rr *csvRecipientRepository) Config(conf *viper.Viper) error {
	rr.linkedRecipient = list.New()
	rr.conf = conf
	if filename := conf.GetString("filename"); filename != "" {
		return rr.ReadFromFile(filename)
	}
	return errors.New("You must give a filename to use CSV with recipients.")
}

func (rr *csvRecipientRepository) Info() domain.RecipientRepositoryInfo {
	return domain.RecipientRepositoryInfo{
		Name: "csv",
	}
}

func (rr *csvRecipientRepository) Add(recip *domain.Recipient) error {
	rr.linkedRecipient.PushBack(recip)
	rlog.Debugf("Add recipient: %v\n", recip)
	return nil
}

func (rr *csvRecipientRepository) FindByName(findName string) domain.RecipientList {
	results := make(domain.RecipientList, 0)
	for e := rr.linkedRecipient.Front(); e != nil; e = e.Next() {
		r := e.Value.(*domain.Recipient)
		if strings.ToLower(r.Name) == strings.ToLower(findName) {
			results = append(results, r)
		}
	}
	return results
}
