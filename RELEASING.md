# 发布规则与流程

本文档记录 `itms-services` 项目的版本号规则和发布流程。

## 版本号规则

版本格式：`v1.0.<commit-count>`

- **主版本号** （`1`）：固定为 `1`，代表首个正式版本系列。
- **次版本号** （`0`）：固定为 `0`，重大架构变更时手动递增。
- **修订号** （`<commit-count>`）：等于当前仓库的 **总 commit 数** ，自动递增，每次提交后唯一。

```bash
# 查询当前应使用的修订号
git rev-list --count HEAD
```

**示例：**

| commit 数 | 版本号 |
|---|---|
| 11 | `v1.0.11` |
| 12 | `v1.0.12` |
| 20 | `v1.0.20` |

## 发布流程

### 第一步：确认代码就绪

```bash
# 确保工作区干净、所有变更已提交
git status

# 构建验证
go build ./...
go vet ./...
```

### 第二步：获取版本号

```bash
# 获取当前 commit 数作为修订号
git rev-list --count HEAD
# 例如输出 12，则版本号为 v1.0.12
```

### 第三步：打 Tag 并推送

```bash
# 将 <N> 替换为实际 commit 数
git tag -a v1.0.<N> -m "v1.0.<N> - <本次发布的简要说明>"

# 推送 tag 到远程，自动触发 GitHub Actions release workflow
git push origin v1.0.<N>
```

### 第四步：等待 CI/CD 自动完成

推送 tag 后，GitHub Actions 的 **Release** workflow（`.github/workflows/release.yml`）
会自动运行 GoReleaser，构建并发布以下平台的二进制包：

| 平台 | 文件格式 |
|---|---|
| macOS Intel（darwin/amd64）| `.tar.gz` |
| macOS Apple Silicon（darwin/arm64）| `.tar.gz` |
| Linux amd64 | `.tar.gz` |
| Linux arm64 | `.tar.gz` |
| Windows amd64 | `.zip` |
| Windows arm64 | `.zip` |
| 校验文件 | `checksums.txt` |

### 第五步：确认发布结果

```bash
# 查看 workflow 运行状态
gh run list --limit 3

# 查看发布详情
gh release view v1.0.<N>
```

## 快速参考

```bash
# 一键获取下一个版本号
echo "v1.0.$(git rev-list --count HEAD)"

# 完整发布三步命令（将 N 替换为实际数值）
N=$(git rev-list --count HEAD)
git tag -a "v1.0.${N}" -m "v1.0.${N}"
git push origin "v1.0.${N}"
```

## 注意事项

- **不要** 在未提交所有变更的情况下打 tag，否则版本号与代码状态不一致。
- **不要** 手动在 GitHub 上创建 Release，统一通过推送 tag 触发自动化流程。
- 如需删除错误的 tag：

```bash
# 删除本地 tag
git tag -d v1.0.<N>

# 删除远程 tag 和 release
git push origin :refs/tags/v1.0.<N>
gh release delete v1.0.<N> --yes
```
