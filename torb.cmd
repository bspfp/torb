@echo off

setlocal enableextensions
setlocal enabledelayedexpansion

rem 파워쉘 명령을 뭘 쓸까?
set pwsh_cmd=
call :get_pwsh_cmd pwsh || ^
call :get_pwsh_cmd powershell || ^
goto not_found

rem 파워쉘로 넘길 값을 준비
set pwsh_args=
for %%i in (%*) do (
    if "!pwsh_args!" equ "" (
        set "pwsh_args='%%~i'"
    ) else (
        set "pwsh_args=!pwsh_args!, '%%~i'"
    )
)

rem 휴지통으로 보내자
%pwsh_cmd% -NoProfile -ExecutionPolicy Bypass -Command "& { . '%~dp0\torb-func.ps1'; MoveTo-RecycleBin @(%pwsh_args%) }" && exit /b 0
exit /b 1

:get_pwsh_cmd
where %1 >nul 2>&1 || exit /b 1
set pwsh_cmd=%1
exit /b 0

:not_found
echo "powershell not found"
exit /b 1
