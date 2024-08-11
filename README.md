# Moneyforward CLI
Chromeを起動して、[マネーフォワード ME](https://moneyforward.com/)を操作するためのツールです。


## ツールでできること
以下のことができます。

1. 金融機関からのデータを一括更新する (`./moneyforward collect`)
2. 家計簿から当月もしくは前月の収入・支出詳細をHTMLで取得する (`./moneyforward breaktable`)
3. 資産から資産の内訳をテキスト形式で取得する (`./moneyforward portfolio`)


## ユーザー名・パスワードの指定方法
ユーザー名・パスワードを指定し、コマンドを実行する必要があります。指定方法は複数あります。

1. コマンドのフラグとして指定する方法
2. 環境変数で指定する方法
3. 設定ファイルで指定する方法

### コマンドのフラグとして指定する方法
コマンドのフラグとして指定する場合は、以下のようにします。

```bash
$ ./moneyforward --user foo --password P@ssw0rd
```

### 環境変数で指定する方法
環境変数で指定する場合には、以下のようにします。

```bash
$ export MONEYFORWARD_USER="foo"
$ export MONEYFORWARD_PASSWORD="P@ssw0rd"
$ ./moneyforward
```

### 設定ファイルで指定する方法
設定ファイルの中で指定する場合には、以下のようにします。デフォルトでは、`${HOME}/.moneyforward.yaml`を参照します。

```bash
$ cat ~/.moneyforward.yaml
user: "foo"
password: "P@ssw0rd"

$ ./moneyforward
```

## tasks
Here are tasks to be executed by `xc` command.

### Initial Steps
Conduct initial steps.

```bash
rm go.*
go mod init gitea.kazu634.com/kazu634/$(basename $PWD)
```

### Use Cobra
Invoke `cobra-cli init`.

```bash
which cobra-cli > /dev/null
if [ $? -ne 0 ]; then
  echo Install cobra-cli first. Execute "go install github.com/spf13/cobra-cli@latest".
  exit 1
fi

cobra-cli init
```

### Install Go Modules
Install `go` modules to the local directory. In short,  execute `go mod vendor`.

```bash
go mod tidy
go mod vendor
```

### Generate .drone.yml
Generate example `.drone.yml` to the local directory.

```bash
cat .assets/.drone.yml | sed -e "s/__DIR__/$(basename ${PWD})/g" | tee .drone.yml
```

### Generate Dockerfile
Generate example `Dockerfile` to the local directory.

```bash
cat .assets/Dockerfile | sed -e "s/__DIR__/$(basename ${PWD})/g" | tee Dockerfile
```

### Generate .goreleaser.yml
Generate example `.goreleaser.yml` to the local directory.

```bash
cat .assets/.goreleaser.yml | sed -e "s/__DIR__/$(basename ${PWD})/g" | tee .goreleaser.yml
```
