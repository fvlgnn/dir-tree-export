@echo off
setlocal enabledelayedexpansion

:: Recupera il numero di commit e il timestamp
for /f %%i in ('git rev-list --count HEAD') do set COUNT=%%i
for /f %%i in ('powershell -Command "[int](Get-Date -UFormat \"%%y%%m%%d\")"') do set TIMESTAMP=%%i

:: Definisco il nome del file
set FILE_NAME=dir_tree_export.exe

:: Costruisci la versione
set VERSION=dev.%TIMESTAMP%.%COUNT%

:: Esegui la build
echo Eseguo la build...
cd src
go build -ldflags="-X 'main.version=%VERSION%'" -o ..\%FILE_NAME% main.go
cd ..

echo Build %FILE_NAME% completata con versione %VERSION%

endlocal
exit /b 0




