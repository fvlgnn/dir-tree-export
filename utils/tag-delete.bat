@echo off

:: NOTE: Questo script semplifica la cancellazione di un tag.
:: tag-delete.bat

set /p VERSION=Dichiara la versione (es. 1.2.3): 

:: Esegue i comandi git
echo Cancellazione locale del tag v%VERSION%...
git tag -d v%VERSION%

echo Cancellazione remota del tag origin...
git push origin --delete v%VERSION%

echo Tag v%VERSION% cancellato. Se collegata una release ricorda di cancellarla sul repository GitHub!

exit /b 0