Step01: ソースコードを用意しよう

## リポジトリのclone

```
$ git clone https://github.com/tenntenn/hosei24
```

## ブランチを作成しよう

```
$ git switch -c fix-section02-step01
```

## 修正をコミットしてみよう

```
$ cd skeleton/section02/step01
$ echo "#" >> README.md
$ git add README.md
$ git commit -m "fix README"
```

## mainブランチに戻して更新する

```
$ git switch main
$ git fetch -p
$ git pull --rebase origin main
```
