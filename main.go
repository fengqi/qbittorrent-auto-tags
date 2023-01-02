package main

import (
	"fengqi/qbittorrent-auto-tags/config"
	"fengqi/qbittorrent-auto-tags/qbittorrent"
	"fmt"
)

func main() {
	c, err := config.LoadConfig("./config.json")
	if err != nil {
		fmt.Printf("[ERR] load config err: %v\n", err)
		return
	}

	webui, err := qbittorrent.Login(c)
	if err != nil {
		fmt.Printf("[ERR] login to qbittorrent err %v\n", err)
		return
	}

	torrentsList, err := webui.GetTorrentList()
	if err != nil {
		fmt.Printf("[ERR] get torrent list err %v\n", err)
		return
	}

	for _, i := range torrentsList {
		if i.Tracker == "" {
			trackerList, err := webui.GetTorrentTrackers(i.Hash)
			if err == nil && len(trackerList) > 0 {
				i.Tracker = trackerList[0].Url // 暂时默认用第一个
			}
		}

		tag, err := i.GetTrackerTag()
		if err != nil {
			fmt.Printf("[ERR] get %s tag err: %v\n", i.Name, err)
			continue
		}

		if custom, ok := c.Sites[tag]; ok {
			tag = custom
		}

		err = webui.AddTags(i.Hash, tag)
		if err != nil {
			fmt.Printf("[ERR] add tag %s to %s err: %v\n", tag, i.Name, err)
			continue
		}

		fmt.Printf("[INFO] add tag %s to %s\n", tag, i.Name)
	}
}
