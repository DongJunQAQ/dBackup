# 定义项目名称
BINARY_NAME := dBackup
# 定义编译输出目录
BUILD_DIR := ./bin
# 定义源代码入口
MAIN_PACKAGE := ./main.go
# 定义man手册输出目录
DOC_DIR := ./doc
# 定义默认行为，执行make时等同于执行make build
.DEFAULT_GOAL := build

# 根据不同的平台设置不同的变量
ifeq ($(OS),Windows_NT)  # 如果当前系统为Windows_NT的话执行一下操作
    BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME).exe  # 获取二进制文件的完整路径
    RMDIR := powershell -Command "Remove-Item -Path $(BUILD_DIR) -Recurse -Force"  # 删除编译输出目录(bin目录)
else
    BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)
    RMDIR := rm -rf $(BUILD_DIR) $(DOC_DIR)
endif

## build: Compile Project
build:
	@echo "Start compiling this project on the $(OS) platform..."
	go build -o $(BINARY_PATH) $(MAIN_PACKAGE)  # 在Makefile的规则中，要执行的shell命令必须以一个制表符（Tab）开头，而不能是空格
	@echo "Compilation completed: $(BINARY_PATH)"

## clean: Clean Artifacts
clean:
	@echo "Cleaning compilation artifacts..."
	$(RMDIR)
	@echo "Cleanup completed"

## doc: Generate man Pages In Linux Systems
doc:
	@echo "Generate man page file..."
	./$(BINARY_PATH)/$(MAIN_PACKAGE) man
	\cp doc/* /usr/share/man/man1/
	@echo "Generate completed"

## help: Print Help Info
help: Makefile
	@echo -e "\nUsage: make <TARGETS> \n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"