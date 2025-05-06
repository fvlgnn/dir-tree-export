@echo off

:: NOTE: Questo script crea e invia un tag Git su un repository.
:: tag-release.bat

set /p VERSION=Inserisci la versione (es. 1.2.3): 

:: Esegue i comandi git
echo Creazione del tag v%VERSION%...
git tag v%VERSION%

echo Push del tag su origin...
git push origin v%VERSION%

echo Tag v%VERSION% creato e inviato con successo!

exit /b 0