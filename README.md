# torb (To Recycle Bin)

## 이게 뭔데? (What is this?)

Windows Command Prompt에서 입력된 경로를 휴지통으로 보내는 도구이다.

(It is a tool that sends the specified path to the Recycle Bin from the Windows Command Prompt.)

go 프로그램을 작성했으나 뭔 이유로 Windows Defender가 바이러스라고 판단해서 삭제하였다.

(I wrote a go program but Windows Defender deleted it because it thought it was a virus.)

## 사용법 (How to use)

1. torb.cmd와 torb-func.ps1 파일을 특정 경로에 복사한다.<br />
   (Copy the torb.cmd and torb-func.ps1 files to a specific path.)<br />
   경로는 환경변수 PATH에 있는 경로인 것이 편하다.<br />
   (It is convenient to have a path in the environment variable PATH.)
2. 아래와 같이 사용한다<br />
   ```Windows
   torb.cmd <경로> [...]
   ```

## 리소스 & 툴 (Resources & Tools)

* VSCode
