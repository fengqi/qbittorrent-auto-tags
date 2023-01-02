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
		return nil, err
	}

	for _, item := range w.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var torrentList []*TorrentInfo
	err = json.Unmarshal(bytes, &torrentList)

	return torrentList, nil
}

func (w *WebUI) GetTorrentTrackers(hash string) ([]*TorrentTracker, error) {
	url := fmt.Sprintf("%s/api/v2/torrents/trackers?hash=%s", w.Host, hash)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for _, item := range w.Cookie {
		req.AddCookie(item)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var trackerList []*TorrentTracker
	err = json.Unmarshal(bytes, &trackerList)
	if err != nil {
		return nil, err
	}

	i := 0
	for _, item := range trackerList {
		// tier >=0 的才有效
		if item.Tier >= 0 {
			trackerList[i] = item
			i++
		}
	}

	return trackerList[:i], nil
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
