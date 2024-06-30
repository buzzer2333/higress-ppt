## 编译
plugins/wasm-go/extensions/ai-cache

```
tinygo build -o main.wasm -scheduler=none -target=wasi -gc=custom -tags="custommalloc nottinygc_finalizer" ./
```

## 镜像
docker login --username=小学生反弹 registry.cn-shanghai.aliyuncs.com

docker buildx build --platform linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/ai-cache/wasm:v1.0 . --push


## higress 平台配置 镜像地址

## 验证
```
curl 'http://localhost:8080/api/openai/v1/chat/completions' \
  -H 'Accept: application/json, text/event-stream' \
  -H 'Content-Type: application/json' \
  --data-raw '{"model":"qwen-long","frequency_penalty":0,"max_tokens":800,"stream":false,"messages":[{"role":"user","content":"higress项目主仓库的github地址是什么"}],"presence_penalty":0,"temperature":0.7,"top_p":0.95}'
```

## 提交线上测评

```
zip -vr plugin.zip address.txt plugin
```
```
zip -vr plugin.zip address.txt plugins
```