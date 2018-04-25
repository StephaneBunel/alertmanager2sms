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

func (rh *csvRepositoryHandle) ReadFromFile(filename string) error {
	csvFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return rh.ReadFromString(string(csvFile))
}

func (rh *csvRepositoryHandle) ReadFromString(body string) error {
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
			err := rh.Add(recip)
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
func (rh *csvRepositoryHandle) Config(conf *viper.Viper) error {
	rh.linkedRecipient = list.New()
	rh.conf = conf
	if filename := conf.GetString("filename"); filename != "" {
		return rh.ReadFromFile(filename)
	}
	return errors.New("you must give a filename to use CSV with recipients")
}

func (rh *csvRepositoryHandle) Info() domain.RecipientRepositoryHandlerInfo {
	return domain.RecipientRepositoryHandlerInfo{
		Name:    "csv",
		Authors: "Stéphane Bunel",
		Version: "0.1",
	}
}

func (rh *csvRepositoryHandle) Add(recip *domain.Recipient) error {
	rh.linkedRecipient.PushBack(recip)
	rlog.Debugf("Add recipient: %v\n", recip)
	return nil
}

func (rh *csvRepositoryHandle) FindByName(findName string) domain.RecipientList {
	results := make(domain.RecipientList, 0)
	for e := rh.linkedRecipient.Front(); e != nil; e = e.Next() {
		r := e.Value.(*domain.Recipient)
		if strings.ToLower(r.Name) == strings.ToLower(findName) {
			results = append(results, r)
		}
	}
	return results
}
