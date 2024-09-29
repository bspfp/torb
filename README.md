# torb (To Recycle Bin)

## 이게 뭔데? (What is this?)

Windows Command Prompt에서 입력된 경로를 휴지통으로 보내는 도구이다.
(It is a tool that sends the specified path to the Recycle Bin from the Windows Command Prompt.)

## 사용법 (How to use)

### Windows 스크립트와 파워쉘 스크립트 이용 (Using Windows Script and PowerShell Script)

1. torb.cmd와 torb-func.ps1 파일을 특정 경로에 복사한다. (Copy the torb.cmd and torb-func.ps1 files to a specific path.)
2. ```Command
   torb.cmd <경로> [...]
   ```

### go 프로그램 사용 (Using go program)

1. torb.go를 빌드한다. (Build torb.go.)
2. 생성된 실행 파일을 사용한다. (Use the generated executable file.)

#### 아이콘 적용 방법

```
go install github.com/tc-hib/go-winres@latest
go-winres simply --icon .\recyle-bin.ico
go build -o ./torb.exe .
```


## 리소스 & 툴 (Resources & Tools)

* VSCode
* Paint.NET
* Garbage icons created by Ridho Imam Prayogi - Flaticon
  * [garbage icons](https://www.flaticon.com/free-icons/garbage)