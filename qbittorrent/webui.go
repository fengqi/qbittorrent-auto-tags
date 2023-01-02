package qbittorrent

import (
	"encoding/json"
	"fengqi/qbittorrent-auto-tags/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Login(c *config.Config) (*WebUI, error) {
	url := fmt.Sprintf("%s/api/v2/auth/login", c.Host)
	data := fmt.Sprintf("username=%s&password=%s", c.Username, c.Password)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &WebUI{
		Host:     c.Host,
		Username: c.Username,
		Password: c.Password,
		Cookie:   resp.Cookies(),
	}, nil
}

func (w *WebUI) Logout() {

}

func (w *WebUI) GetTorrentList() ([]*TorrentInfo, error) {
	url := fmt.Sprintf("%s/api/v2/torrents/info?filter=all&tag=&limit=%d", w.Host, 1000)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	for _, item := range w.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var torrentList []*TorrentInfo
	err = json.Unmarshal(bytes, &torrentList)
	if err != nil {
		panic(err)
	}

	return torrentList, nil
}

func (w *WebUI) AddTags(hashes, tag string) error {
	url := fmt.Sprintf("%s/api/v2/torrents/addTags", w.Host)
	data := fmt.Sprintf("hashes=%s&tags=%s", hashes, tag)

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, item := range w.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
