# 設計

## コマンド体系

```console
$ madao mv [options...] <source> [<destination>]
$ madao move [options...] <source> [<destination>]

$ madao mv SOURCE dest.md:9
```

```ebnf
source               = entire lines of file
                     | head lines of file
                     | tail lines of file
                     | middle lines of file ;

destination          = tail of file | a line of file ;

entire lines of file = filename ;
head lines of file   = filename , ":" , "~" , end line ;
tail lines of file   = filename , ":" , start line , "~" ;
middle lines of file = filename , ":" start line , "~" , end line ;

tail of file         = filename ;
a line of file       = filename , ":" , line ;

start line           = positive ;
end line             = line ;
line                 = positive | negative ;

digit sequence       = digit | digit sequence , digit ;
negative             = "-" , positive ;
positive             = digit excluding zero | digit excluding zero , digit sequence ;

digit excluding zero = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;
digit                = "0" | digit excluding zero ;

filename             = ? all characters ?
```

これだと曖昧になる（e.g.`xxx.md:1` という名前のファイルの全行、みたいな指定はできない）が、エッジケースなので無視する。

## 流れ

### sourceの切り出し

指定された場所の内容を切り出す。
ファイル全体の場合はファイルを削除する。

partial-sourceと呼ぶ。

### linkの変更を検出する

• 移動元から移動先は同じ階層か
• 移動元から移動先は同じファイルか

を別々に使うので

• source-dir
• dest-dir
• source-name
• dest-name

を別々に持ったほうが良さそう

### partial-sourceの中のIDを抽出

Pandoc Lua Filterにより、IDのリストが得られる。
得られたIDのリストを以降partial-source-idsと呼ぶ。

### partial-sourceの中のリンクを加工

- if sourceとdestが同じファイル: 特に何もしない
- elseif sourceとdestが同じ階層
    - if sourceの指定がファイル全体: 何もしない
    - else: 同一ファイルへのリンクで、partial-source-idsにないIDへのリンクだけ、source-nameをつける
- else: すべてのリンクを元の階層を意識した相対リンクに置き換える

### 他のドキュメントから、sourceへのリンクを置換する

- if sourceの指定がファイル全体: ひたすらリンク先ファイル名がsourceに一致するものを置き換えていく
- elseif リンク先がIDを含む場合: partial-source-idsにないものは置き換えない
- else: (parameterize?) 置き換えない

### destに書き込む

partial-sourceをdestに書き込む。
