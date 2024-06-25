# 環境構築

go mod init local.package/golang_todo

go mod tidy

### DB 接続方法

`sqlite3 wabapp.sql`実行

# go の基礎

## 暗黙的な変数の定義

Go 言語では、暗黙的な変数の定義は短変数宣言（short variable declaration）とも呼ばれ、`:=`を使って行われます。これは、特定のスコープ内で新しい変数を簡潔に宣言し、初期化するための方法です。以下に詳細と例を示します。

### 短変数宣言の基本

短変数宣言は、`var`キーワードを使わずに変数を宣言し、初期化する方法です。形式は以下のようになります。

```go
variableName := initialValue
```

### 例

1. **基本的な使用例**:

```go
package main

import "fmt"

func main() {
    // 短変数宣言で整数型の変数を定義
    x := 10
    fmt.Println(x) // 10

    // 文字列型の変数を定義
    name := "Alice"
    fmt.Println(name) // Alice

    // ブール型の変数を定義
    isActive := true
    fmt.Println(isActive) // true
}
```

2. **複数の変数を同時に定義**:

短変数宣言では、同じ行で複数の変数を同時に定義することもできます。

```go
package main

import "fmt"

func main() {
    // 複数の変数を同時に宣言および初期化
    a, b, c := 1, "hello", 3.14
    fmt.Println(a) // 1
    fmt.Println(b) // hello
    fmt.Println(c) // 3.14
}
```

3. **関数の戻り値を使った短変数宣言**:

関数の戻り値を使って変数を定義することもできます。

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // strconvパッケージのAtoi関数を使って整数に変換
    i, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted integer:", i) // 123
    }
}
```

### 注意点

1. **再宣言と再代入**:
   短変数宣言を使って変数を再度宣言する場合、少なくとも一つの新しい変数が含まれている必要があります。そうでないとコンパイルエラーになります。

   ```go
   func main() {
       x := 10
       x, y := 20, 30 // xは再代入され、yは新しい変数として宣言される
       fmt.Println(x, y) // 20 30
   }
   ```

2. **パッケージレベルでは使用できない**:
   短変数宣言は関数内でのみ使用可能で、パッケージレベルでは使用できません。

   ```go
   // これはコンパイルエラーになる
   // x := 10

   func main() {
       x := 10 // これは有効
       fmt.Println(x)
   }
   ```

3. **未使用変数の扱い**:
   Go 言語では未使用の変数があるとコンパイルエラーになります。短変数宣言で定義した変数も必ず使用する必要があります。

短変数宣言はコードを簡潔にし、特に一時的な変数を使う場合に便利です。しかし、スコープや変数の使用状況を適切に管理することが重要です。

## パブリック（公開）とプライベート（非公開）

Go 言語では、パブリック（公開）とプライベート（非公開）の概念は、識別子（変数、関数、型、構造体のフィールドなど）の名前の最初の文字の大文字/小文字によって決まります。これにより、パッケージ外部からアクセスできるかどうかが決まります。

### パブリックとプライベートのルール

1. **パブリック（公開）**:

   - 名前が大文字で始まる識別子はパブリックです。
   - パッケージ外からアクセス可能です。

2. **プライベート（非公開）**:
   - 名前が小文字で始まる識別子はプライベートです。
   - 同じパッケージ内からのみアクセス可能です。

### 例

#### パブリックとプライベートのフィールドとメソッド

```go
// ファイル: example.go
package example

// パブリックな構造体
type Person struct {
    // パブリックなフィールド
    Name string
    // プライベートなフィールド
    age int
}

// パブリックな関数
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, age: age}
}

// パブリックなメソッド
func (p *Person) GetName() string {
    return p.Name
}

// プライベートなメソッド
func (p *Person) getAge() int {
    return p.age
}

// パブリックなメソッド
func (p *Person) SetAge(age int) {
    p.age = age
}

// パブリックな関数
func (p *Person) GetAge() int {
    return p.age
}
```

#### 別のパッケージからのアクセス

```go
// ファイル: main.go
package main

import (
    "fmt"
    "example"
)

func main() {
    // NewPerson関数を使って新しいPersonを作成
    p := example.NewPerson("Alice", 30)

    // パブリックなフィールドとメソッドにアクセス
    fmt.Println(p.GetName()) // "Alice"

    // プライベートなフィールドやメソッドにアクセスしようとするとエラーになる
    // fmt.Println(p.age) // コンパイルエラー
    // fmt.Println(p.getAge()) // コンパイルエラー

    // パブリックなメソッドを使ってプライベートフィールドにアクセス
    fmt.Println(p.GetAge()) // 30

    // パブリックなメソッドを使ってプライベートフィールドを変更
    p.SetAge(31)
    fmt.Println(p.GetAge()) // 31
}
```

### ポイント

1. **名前の規則**:

   - 大文字で始まる識別子はパッケージ外部からアクセス可能（パブリック）。
   - 小文字で始まる識別子はパッケージ内部でのみアクセス可能（プライベート）。

2. **パッケージ設計**:

   - プライベートな識別子はパッケージの内部実装を隠蔽するために使われます。
   - パブリックな識別子はパッケージの外部インターフェースを定義します。

3. **カプセル化**:
   - プライベートフィールドを使ってデータのカプセル化を行い、パブリックメソッドを通じてアクセスを制御することで、オブジェクトの整合性を保ちます。

Go 言語のこのシンプルなパブリックとプライベートの概念は、コードの可読性を高め、適切なカプセル化とモジュール性を維持するのに役立ちます。

## ポインタ

Go 言語のポインタは、メモリ内の特定の変数のアドレスを保持する変数です。ポインタを使うと、変数の値を直接操作することができ、効率的なメモリ管理やデータの共有が可能になります。以下に基本的なポインタの使用方法を示します。

### ポインタの宣言と使用

1. **ポインタの宣言**:
   ポインタはアスタリスク（`*`）を使って宣言します。例えば、`*int`は整数型のポインタです。

   ```go
   var p *int
   ```

2. **ポインタへの値の割り当て**:
   `&`演算子を使って変数のアドレスを取得し、それをポインタに割り当てます。

   ```go
   var x int = 10
   var p *int
   p = &x
   ```

3. **ポインタを通じて値を操作する**:
   ポインタを使って変数の値を変更するには、`*`演算子を使います。
   ```go
   fmt.Println(*p) // xの値を出力する（10）
   *p = 20         // xの値を変更する
   fmt.Println(x)  // xの値を出力する（20）
   ```

### ポインタの例

以下に、ポインタを使って関数内で変数の値を変更する例を示します。

```go
package main

import "fmt"

// ポインタを受け取る関数
func updateValue(p *int) {
    *p = 200
}

func main() {
    var x int = 100
    fmt.Println("Before:", x) // 変更前の値を出力（100）

    updateValue(&x) // xのアドレスを渡す
    fmt.Println("After:", x)  // 変更後の値を出力（200）
}
```

### ポインタの利点

- **効率的なメモリ使用**: 大きなデータ構造を関数間で渡すときにコピーする代わりに、ポインタを使うことで効率的にメモリを使用できます。
- **データの共有**: 複数の関数で同じデータを操作することができます。

ポインタは適切に使用すると強力なツールですが、誤用するとプログラムのバグや予期しない動作を引き起こす可能性があるため、慎重に扱う必要があります。

## 構造体（struct）

Go 言語の構造体（struct）は、異なる型のフィールドをまとめてひとつのデータ型として扱うことができる、カスタムデータタイプです。構造体を使うと、複雑なデータモデルを定義して操作することができます。

### 構造体の宣言と使用方法

1. **構造体の宣言**:
   構造体は`type`キーワードを使って宣言します。以下は基本的な構造体の宣言例です。

   ```go
   type Person struct {
       Name string
       Age  int
       City string
   }
   ```

2. **構造体のインスタンス化**:
   構造体のインスタンスを作成し、フィールドに値を割り当てます。

   ```go
   var person1 Person
   person1.Name = "Alice"
   person1.Age = 30
   person1.City = "New York"

   // またはリテラルを使って一度に初期化することもできます。
   person2 := Person{Name: "Bob", Age: 25, City: "San Francisco"}
   ```

3. **構造体のフィールドアクセス**:
   構造体のフィールドにはドット（`.`）を使ってアクセスします。
   ```go
   fmt.Println(person1.Name) // "Alice"
   fmt.Println(person2.Age)  // 25
   ```

### 構造体の例

以下は、構造体を使って複数の人のデータを管理する例です。

```go
package main

import "fmt"

// Person構造体の定義
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Person構造体のインスタンスを作成
    person1 := Person{Name: "Alice", Age: 30, City: "New York"}
    person2 := Person{Name: "Bob", Age: 25, City: "San Francisco"}

    // 構造体のフィールドにアクセス
    fmt.Println(person1.Name) // "Alice"
    fmt.Println(person2.Age)  // 25

    // 構造体のフィールドを更新
    person1.Age = 31
    fmt.Println(person1.Age) // 31
}
```

### 構造体のメソッド

構造体にはメソッドを定義することができます。メソッドは、特定の型に関連付けられた関数です。

```go
package main

import "fmt"

// Person構造体の定義
type Person struct {
    Name string
    Age  int
    City string
}

// Person構造体に関連付けられたメソッド
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Charlie", Age: 28, City: "Los Angeles"}
    person.Greet() // "Hello, my name is Charlie and I am 28 years old."
}
```

### 構造体の利点

- **カスタムデータタイプ**: 構造体を使うことで、複雑なデータモデルを簡潔に表現できます。
- **データの整理**: 関連するデータを一つの単位としてまとめることができ、コードの可読性が向上します。
- **メソッドの追加**: 構造体に関連するメソッドを定義することで、オブジェクト指向的な設計が可能になります。

Go 言語の構造体は、データを整理し、関連するメソッドを追加するための強力なツールです。適切に使用することで、コードの保守性と可読性が向上します。
