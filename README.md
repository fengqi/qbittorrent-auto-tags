# qbittorrent-auto-tags

根据`tracker`地址给种子设置标签，如 `https://tracker.baidu.com` 则设置标签为 `baidu.com`你也可以在 `config.json` 里配置别名。

# 怎么使用
根据平台下载对应的release文件，复制 example.config.json 为 config.json，修改 qbittorrent 的配置信息以及自定义标签，然后放到定时任务定期运行。

> 目前提供：Linux、Mac、Windows的二进制文件，其他平台可自行编译

# 注释事项

1. 运行模式为定时任务，非守护进程，每次处理1000条没有标签的种子（打算做成可配置），处理完会自动退出。

2. 部分种子调用 qbittorrent 接口获取不到trakcer地址，所以无法处理。

3. 如果需要自定义标签，请修改 config.json 文件的 sites 段，已经设置的标签请先在 webui 删除，等待重新处理。
