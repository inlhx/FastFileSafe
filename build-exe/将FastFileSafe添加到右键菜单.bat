@ECHO OFF&(PUSHD "%~DP0")&(REG QUERY "HKU\S-1-5-19">NUL 2>&1)||(
powershell -Command "Start-Process '%~sdpnx0' -Verb RunAs"&&EXIT)
:: ע����Ϣ
:Menu
Echo.&Echo  ���FastFileSafe���Ҽ�,�ɷ�������ٶ��ļ����мӽ��ܣ�
Echo.&Echo  ���ܷ�ʽ����AES+ZIP,ÿ�μ�����������Ʊ��ܺ�����
Echo.&Echo  1������Ҽ�ʹ�� FastFileSafe ���ļ�����/����
Echo.&Echo  2��ɾ���Ҽ�ʹ�� FastFileSafe �Լ�����/����
Echo.&Echo  3���Ȳ����ֱ�����мӽ���

xcopy FastFileSafe.exe C:\Windows\ /y /e

ECHO:
set /p choice=���������ûس�����
if not "%choice%"=="" set choice=%choice:~0,1%
if /i "%choice%"=="1" Goto AddMenu
if /i "%choice%"=="2" Goto RemoveMenu
if /i "%choice%"=="3" Goto RunWork
ECHO ������Ч &PAUSE>NUL&CLS&GOTO MENU
:AddMenu
reg add "HKCR\*\shell\FastFileSafe" /f /ve /d "FastFileSafe-�ļ�����/����"  >NUL 2>NUL
reg add "HKCR\*\shell\FastFileSafe" /f /v "Icon" /t REG_EXPAND_SZ /d "%~dp0FastFileSafe.exe" >NUL 2>NUL
reg add "HKCR\*\shell\FastFileSafe\command" /f /ve /d "%~dp0FastFileSafe.exe \"%%1\"" >NUL 2>NUL


reg add "HKCR\Directory\shell\FastFileSafe" /f /ve /d "FastFileSafe-�ļ�[��]����/����"  >NUL 2>NUL
reg add "HKCR\Directory\shell\FastFileSafe" /f /v "Icon" /t REG_EXPAND_SZ /d "%~dp0FastFileSafe.exe" >NUL 2>NUL
reg add "HKCR\Directory\shell\FastFileSafe\command" /f /ve /d "%~dp0FastFileSafe.exe \"%%1\"" >NUL 2>NUL

ECHO:&ECHO ��� &PAUSE>NUL&CLS&GOTO MENU
:RemoveMenu
reg delete "HKCR\*\shell\FastFileSafe" /f >NUL 2>NUL
reg delete "HKCR\Directory\shell\FastFileSafe" /f >NUL 2>NUL
ECHO:&ECHO ��� &PAUSE>NUL&CLS&GOTO MENU

:RunWork
FastFileSafe.exe

ECHO:&ECHO ��� &PAUSE>NUL&CLS&GOTO MENU