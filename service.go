package vkn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type srv struct {
	token string
	cnf   Config
}

// New creates a new VKN service instance.
func New(cnf Config) Vkn {
	return &srv{
		cnf: cnf,
	}
}

func (s *srv) login() error {
	data := url.Values{}
	data.Set("assoscmd", "anologin")
	data.Set("rtype", "json")
	data.Set("userid", s.cnf.Username)
	data.Set("sifre", s.cnf.Password)
	data.Set("sifre2", s.cnf.Password)
	data.Set("parola", "1")
	client := &http.Client{}
	r, _ := http.NewRequest("POST", "https://earsivportal.efatura.gov.tr/earsiv-services/assos-login", bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result loginResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	if result.Token == "" {
		return fmt.Errorf("Token not found")
	}
	s.token = result.Token
	return nil
}

func (s *srv) checkLogin() error {
	if s.token == "" {
		return s.login()
	}
	return nil
}

func (s *srv) GetRecipient(vkn string) (*GetRecipientResponse, error) {
	err := s.checkLogin()
	if err != nil {
		return nil, err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	data.Set("cmd", "SICIL_VEYA_MERNISTEN_BILGILERI_GETIR")
	data.Set("callid", id.String())
	data.Set("pageName", "RG_BASITFATURA")
	data.Set("token", s.token)
	data.Set("jp", `{"vknTcknn":"`+vkn+`"}`)
	client := &http.Client{}
	r, _ := http.NewRequest("POST", "https://earsivportal.efatura.gov.tr/earsiv-services/dispatch", bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *GetRecipientResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
