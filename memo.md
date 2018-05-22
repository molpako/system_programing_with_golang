ファイルディスクリプタとは

- go で `syscall.Write()` と言うシステムコールが呼び出されてるのを確認した
- ファイルディスクリプタとは一種の識別子（数値）
- 以下の `syscall.Write()` の引数 `fd.Sysfd` がファイルディスクリプタに相当する

```
// Write implements io.Writer.
func (fd *FD) Write(p []byte) (int, error) {
	if err := fd.writeLock(); err != nil {
		return 0, err
	}
	defer fd.writeUnlock()
	if err := fd.pd.prepareWrite(fd.isFile); err != nil {
		return 0, err
	}
	var nn int
	for {
		max := len(p)
		if fd.IsStream && max-nn > maxRW {
			max = nn + maxRW
		}
		n, err := syscall.Write(fd.Sysfd, p[nn:max])
		if n > 0 {
			nn += n
		}
		if nn == len(p) {
			return nn, err
		}
		if err == syscall.EAGAIN && fd.pd.pollable() {
			if err = fd.pd.waitWrite(fd.isFile); err == nil {
				continue
			}
		}
		if err != nil {
			return nn, err
		}
		if n == 0 {
			return nn, io.ErrUnexpectedEOF
		}
	}
}
```

- os パッケージの File 型には Write() が定義されている
- `(f *File)` から、この定義が File 型の構造体へのポインタ f に対するメソッドであることを示す
- `Write(b []byte) (n int, err error)` は、[]byte 型のデータを引数とし int 型と error 型を返す
- つまりこの Write は、f が指し示すファイルに対してバイト列 b を書き込み、書き込んだバイト数 n と、エラーが起きた場合はそのエラー err を返す

```
func (f *File) Write(b []byte) (n int, err error) {
	if err := f.checkValid("write"); err != nil {
		return 0, err
	}
	n, e := f.write(b)
	if n < 0 {
		n = 0
	}
	if n != len(b) {
		err = io.ErrShortWrite
	}

	epipecheck(f, e)

	if e != nil {
		err = f.wrapErr("write", e)
	}

	return n, err
}
```

- 同じ仕様のメソッドを持つ方を統一的に扱えると便利
- Golang ではその場合に使える仕組みとして **インタフェース** という型が用意されている
- インタフェースはメソッド宣言の塊
- 上記のWriteメソッドが宣言されているインタフェースが `io.Writer`
- インタフェースで宣言されているすべてのメソッドがデータ型に対して定義されている場合、そのデータ型は「インタフェースを満たす」を表現する
- つまり、「*File は io.Writer インタフェースを満たす」という

