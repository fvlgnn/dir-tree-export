@echo off

REM NOTE: Questo script semplifica la cancellazione di un tag.
REM tag-delete.bat

set /p VERSION=Dichiara la versione (es. 1.2.3): 

REM Esegue i comandi git
echo Cancellazione locale del tag v%VERSION%...
git tag -d v%VERSION%

echo Cancellazione remota del tag origin...
git push origin --delete v%VERSION%

echo Tag v%VERSION% cancellato. Se collegata una release ricorda di cancellarla sul repository GitHub!

@REM pause
exit /b 0