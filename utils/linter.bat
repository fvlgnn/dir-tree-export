@echo off

REM Trova la cartella src
if exist "%cd%\src\" (
    set "SOURCE_DIR=%cd%\src"
) else if exist "%cd%\..\src\" (
    set "SOURCE_DIR=%cd%\..\src"
) else (
    echo Errore: Nessuna cartella 'src' trovata nella directory corrente o superiore.
    exit /b 1
)

echo Eseguo GoLang-Lint in: %SOURCE_DIR%

docker run --rm -t -v "%SOURCE_DIR%":/app -w /app golangci/golangci-lint:v2.1.2 golangci-lint run

endlocal
exit /b 0
