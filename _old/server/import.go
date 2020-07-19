package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ImportServer(configuration, settings, event, eventRules, entrylist, bop, assistRules io.Reader) error {
	server := new(ServerSettings)

	if err := importFile(configuration, &server.Configuration); err != nil {
		return err
	}

	if err := importFile(settings, &server.Settings); err != nil {
		return err
	}

	if err := importFile(event, &server.Event); err != nil {
		return err
	}

	if err := importFile(eventRules, &server.EventRules); err != nil {
		return err
	}

	if err := importFile(entrylist, &server.Entrylist); err != nil {
		return err
	}

	if err := importFile(bop, &server.Bop); err != nil {
		return err
	}

	if err := importFile(assistRules, &server.AssistRules); err != nil {
		return err
	}

	if err := SaveServerSettings(server); err != nil {
		return err
	}

	return nil
}

func importFile(reader io.Reader, config interface{}) error {
	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return err
	}

	return nil
}
