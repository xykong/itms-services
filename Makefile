
.PHONY : itms-services prepare

itms-services : prepare
	@go-bindata -o ./generator/bindata.go -pkg generator -nometadata template
	@go build


prepare :
	@command -v go-bindata >/dev/null 2>&1 || { \
  		echo >&2 "Build require go-bindata but it's not installed. Installing..."; \
  		go get -u github.com/go-bindata/go-bindata/...; }
