# dir-tree-export

Tool da linea di comando per esportare in un file di testo la struttura ad albero di directory e file a partire da una cartella radice.


## Descrizione

**`dir-tree-export`** è uno strumento da riga di comando che consente di esportare la struttura completa di una directory (inclusi file e sottocartelle) in formato testo. Include opzioni per:
- limitare la profondità
- filtrare directory nascoste
- includere dimensioni dei file
- personalizzare il nome del file `.txt` di output


## Argomenti

- `--depth`: profondità di scansione (default _3_)
- `--show-hidden`: mostra anche i file nascosti (default _false_)
- `--output`: nome del file di output (default _dir\_tree\_export\_output.txt_)
- `--show-size`: mostra la dimensione dei file (default _false_)
- `--only-dirs`: mostra solo directory (default _false_)


## Esempi di utilizzo

### Go

```sh
# Base
go run main.go

# Specifica un percorso
go run main.go ./my-folder

# Solo directory, massimo 4 livelli della cartella di esecuzione
go run main.go --only-dirs --depth=4 

# Mostra anche file nascosti specificando un percorso
go run main.go --show-hidden ~/projects

# Esporta con dimensioni e nome custom file output specificando un percorso
go run main.go --show-size --output=mytree.txt ../my-folder
```

### Windows

```powershell
# Base
.\dir-tree-export.exe

# Cartella specifica
.\dir-tree-export.exe C:\Users\foobar\the-folder
```

### Linux

```sh
# Base
sh ./dir-tree-export

# Cartella specifica
sh ./dir-tree-export ~/the-folder/subfolder
```
