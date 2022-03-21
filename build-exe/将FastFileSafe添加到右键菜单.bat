@ECHO OFF&(PUSHD "%~DP0")&(REG QUERY "HKU\S-1-5-19">NUL 2>&1)||(
powershell -Command "Start-Process '%~sdpnx0' -Verb RunAs"&&EXIT)
:: 注册信息
:Menu
Echo.&Echo  添加FastFileSafe到右键,可方便你快速对文件进行加解密！
Echo.&Echo  加密方式采用AES+ZIP,每次加密完成请妥善保管好密码
Echo.&Echo  1、添加右键使用 FastFileSafe 对文件加密/解密
Echo.&Echo  2、删除右键使用 FastFileSafe 对件加密/解密
Echo.&Echo  3、先不添加直接运行加解密

xcopy FastFileSafe.exe C:\Windows\ /y /e

ECHO:
set /p choice=输入数字敲回车键：
if not "%choice%"=="" set choice=%choice:~0,1%
if /i "%choice%"=="1" Goto AddMenu
if /i "%choice%"=="2" Goto RemoveMenu
if /i "%choice%"=="3" Goto RunWork
ECHO 输入无效 &PAUSE>NUL&CLS&GOTO MENU
:AddMenu
reg add "HKCR\*\shell\FastFileSafe" /f /ve /d "FastFileSafe-文件加密/解密"  >NUL 2>NUL
reg add "HKCR\*\shell\FastFileSafe" /f /v "Icon" /t REG_EXPAND_SZ /d "%~dp0FastFileSafe.exe" >NUL 2>NUL
reg add "HKCR\*\shell\FastFileSafe\command" /f /ve /d "%~dp0FastFileSafe.exe \"%%1\"" >NUL 2>NUL


reg add "HKCR\Directory\shell\FastFileSafe" /f /ve /d "FastFileSafe-文件[夹]加密/解密"  >NUL 2>NUL
reg add "HKCR\Directory\shell\FastFileSafe" /f /v "Icon" /t REG_EXPAND_SZ /d "%~dp0FastFileSafe.exe" >NUL 2>NUL
reg add "HKCR\Directory\shell\FastFileSafe\command" /f /ve /d "%~dp0FastFileSafe.exe \"%%1\"" >NUL 2>NUL

ECHO:&ECHO 完成 &PAUSE>NUL&CLS&GOTO MENU
:RemoveMenu
reg delete "HKCR\*\shell\FastFileSafe" /f >NUL 2>NUL
reg delete "HKCR\Directory\shell\FastFileSafe" /f >NUL 2>NUL
ECHO:&ECHO 完成 &PAUSE>NUL&CLS&GOTO MENU

:RunWork
FastFileSafe.exe

ECHO:&ECHO 完成 &PAUSE>NUL&CLS&GOTO MENU