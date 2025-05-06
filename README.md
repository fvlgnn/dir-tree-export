# dir-tree-export

🗂️ **dir-tree-export** è uno strumento da linea di comando per esportare in un file di testo la struttura ad albero di directory e file a partire da una cartella radice.

---

## 🚀 Caratteristiche principali

- Esporta la struttura completa di una directory in formato testo
- Opzioni per:
    - Limitare la profondità di scansione
    - Includere o escludere file e directory nascosti
    - Visualizzare la dimensione dei file
    - Mostrare solo directory
    - Personalizzare il nome del file di output

---

## 📥 Installazione

1. Scarica l'[ultima versione](https://github.com/fvlgnn/dir-tree-export/releases/latest) dell'eseguibile compatibile con il tuo sistema operativo e architettura, come da esempio riportato nella tabella seguente, oppure visita la sezione [Releases](https://github.com/fvlgnn/dir-tree-export/releases) per scegliere una versione specifica.

> 💡 **Nota**:
> 
> * Per **Raspberry Pi 2/3**, usa `linux-armv7`
> * Per **Raspberry Pi 4/5**, usa `linux-arm64`


| Sistema         | Architettura | Denominazione File                           |
|-----------------|--------------|----------------------------------------------|
| Linux           | amd64        | `dir-tree-export-versione-linux-amd64`       |
| Linux           | 386          | `dir-tree-export-versione-linux-386`         |
| Linux (RPi 2/3) | arm (v7)     | `dir-tree-export-versione-linux-armv7`       |
| Linux (RPi 4/5) | arm64        | `dir-tree-export-versione-linux-arm64`       |
| Windows         | amd64        | `dir-tree-export-versione-windows-amd64.exe` |
| Windows         | 386          | `dir-tree-export-versione-windows-386.exe`   |
| macOS (Darwin)  | amd64        | `dir-tree-export-versione-darwin-amd64`      |

2. Copia il file scaricato in una directory del **PATH** di sistema:
   
   * Su **Windows**: ad esempio `%USERPROFILE%\bin\` o aggiungilo al PATH
   * Su **Linux/macOS**: ad esempio `/usr/local/bin/`
 
3. Rinomina il file, se necessario, in `dir_tree_export` (o `dir_tree_export.exe` su Windows) per semplicità.

---

## ⚙️ Utilizzo di base

Puoi eseguire il programma con **doppio clic** (modalità interattiva desktop) o da **linea di comando** per usare gli argomenti.

Esempio base da terminale:

```bash
./dir_tree_export --depth 2 --show-hidden --output "albero.txt"
```

Genererà un file `albero.txt` con la struttura delle cartelle e file fino a una profondità di 2 sottocartelle, includendo anche i file nascosti.

> ⚠️ **Nota per gli utenti Windows**
> 
> All’avvio, Windows potrebbe mostrare un avviso che segnala l'applicazione come "autore non riconosciuto" o "potenzialmente non sicura" perché non è firmata digitalmente. Puoi ignorare l'avviso e procedere con l'esecuzione: il programma è sicuro, open source e trasparente.
>
> Il codice sorgente è disponibile [qui](https://github.com/fvlgnn/dir-tree-export/blob/main/src/main.go) e l'applicazione **non trasmette dati né esegue operazioni dannose.**

---

## 📝 Argomenti disponibili

| Argomento       | Descrizione                       | Default                      |
| --------------- | --------------------------------- | ---------------------------- |
| `--depth`       | Profondità massima di scansione   | `3`                          |
| `--show-hidden` | Includi file e directory nascosti | `false`                      |
| `--output`      | Nome del file di output           | `dir_tree_export_output.txt` |
| `--show-size`   | Mostra la dimensione dei file     | `false`                      |
| `--only-dirs`   | Mostra solo le directory          | `false`                      |

---

## 💡 Esempi di utilizzo

- Mostra solo le directory senza limiti di profondità:
```bash
./dir-tree-export --only-dirs --depth 0
```

- Esporta inclusi file nascosti e dimensioni dei file:
```bash
./dir-tree-export --show-hidden --show-size
```

- Esempio completo: tutti gli argomenti combinati
(esporta solo directory, include file nascosti, mostra dimensioni, profondità illimitata, output personalizzato su percorso specifico):
```bash
./dir-tree-export --only-dirs --show-hidden --show-size --depth 0 --output=/tmp/full_tree.txt /var/www
```

---

### 🐹 **Go (versione sviluppo)**

```bash
# Esecuzione base nella cartella corrente
go run main.go

# Specifica una cartella di partenza
go run main.go ./my-folder

# Solo directory, profondità massima 4 livelli
go run main.go --only-dirs --depth=4

# Includi anche file e cartelle nascosti partendo da ~/projects
go run main.go --show-hidden ~/projects

# Esporta con dimensioni dei file e nome personalizzato dell’output
go run main.go --show-size --output=mytree.txt ../my-folder

# Esempio completo: tutti gli argomenti combinati
go run main.go --only-dirs --show-hidden --show-size --depth=0 --output=full_tree.txt ~/projects
```

> 🚀 Se preferisci compilare:
> 
> ```bash
> go build -o dir-tree-export main.go
> ./dir-tree-export --show-hidden
> ```

---

### 💻 **Windows**

```powershell
# Esecuzione base
.\dir-tree-export.exe

# Specifica una cartella
.\dir-tree-export.exe C:\Users\foobar\the-folder

# Esporta con dimensioni dei file e nome personalizzato dell’output
.\dir-tree-export.exe --show-size --output="C:\Users\gianni\Desktop\foobar.txt" C:\Users\foobar\the-folder

# Esempio completo: tutti gli argomenti combinati
.\dir-tree-export.exe --only-dirs --show-hidden --show-size --depth=0 --output="C:\Users\gianni\Desktop\full_tree.txt" C:\inetpub\wwwroot
```

---

### 🐧 **Linux / macOS**

```bash
# Esecuzione base nella cartella corrente
./dir-tree-export

# Specifica una cartella di partenza
./dir-tree-export ~/the-folder/subfolder

# Esempio completo: tutti gli argomenti combinati
./dir-tree-export --only-dirs --show-hidden --show-size --depth=0 --output=/tmp/full_tree.txt /var/www
```

---

## 📄 Licenza

[MIT License](./LICENSE)

---

## 🤝 Contribuire

Hai idee, bug da segnalare o migliorie? Sentiti libero di aprire una **issue** o una **pull request**!

