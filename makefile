GO ?= go

# Custom configuration | 独立配置
# Service name | 项目名称
SERVICE=usercenter
# Service name in specific style | 项目经过style格式化的名称
SERVICE_STYLE=usercenter
# The project version, if you don't use git, you should set it manually | 项目版本，如果不使用git请手动设置
VERSION=$(shell git describe --tags --always)
# The project file name style | 项目文件命名风格
PROJECT_STYLE=go_zero
# Whether to use i18n | 是否启用 i18n
PROJECT_I18N=true

LDFLAGS := -s -w

# Swagger type, support yml,json | Swagger 文件类型，支持yml,json
SWAGGER_TYPE=json

.PHONY: help
help: # Show help | 显示帮助
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: test
test: # Run Test for the project | 运行项目测试
	go test -v --cover .

.PHONY: install
install: # Install the necessary tools | 安装必要的工具
	@echo "检测 swagger 是否已安装"
	make check_swagger

check_swagger:
	@command -v swagger > /dev/null || (echo "Swagger is not installed. Installing..." && make install_swagger)
	@echo "swagger 已安装"

install_swagger:
	$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: gen-rpc
gen-rpc:  #Generate RPC files from proto | 生成 RPC 的代码
	goctl rpc protoc ./$(SERVICE_STYLE).proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style=$(PROJECT_STYLE) -m
	@echo "Generate RPC codes successfully"

.PHONY: gen-entity
gen-entity: # Create new entity | 创建 新的 模型实例
	go run -mod=mod entgo.io/ent/cmd/ent new --target ent/schema ${entity}

.PHONY: gen-ent
gen-ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema --feature  sql/execquery,intercept,privacy,entql,dialect

.PHONY: db
db: # Show database table | 查看 数据库 表
	atlas schema inspect \
    -u "ent://ent/schema" \
    --dev-url "sqlite://file?mode=memory&_fk=1" \
    --format '{{ sql . "  " }}'

.PHONY: migrate
migrate: # Migrate db to mysql | 迁移数据库
	atlas migrate apply \
    --dir "file://ent/migrate/migrations" \
    --url "mysql://root:123456@localhost:3306/gbill"

.PHONY: docker
docker: # Build docker by Dockerfile | 构建 docker 镜像 | make docker DOCKER_REMOTE=xxxxxx
	docker build -f Dockerfile -t $(DOCKER_REMOTE)/$(SERVICE_STYLE)-rpc:${VERSION} .

.PHONY: run
run: # Run by docker-compose | 在 docker 中运行
	docker-compose down
	docker-compose up -d
