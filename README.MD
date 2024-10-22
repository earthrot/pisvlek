## Fiber CLI
voor hot reloads

### installeren
```bash
go install github.com/gofiber/cli/fiber@latest
```

### draaien (in project map)
```bash
fiber dev
```


## sqlc
voor het compilen van sql queries naar Go structures

### installeren
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### draaien (in project map)
```bash
sqlc generate
```


## tailwind
voor het watchen/compilen van css

### installeren
het makkelijkst is om eerst scoop te installeren
```ps
# in powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
```

en daarna tailwindcss, hierdoor zit 'ie automatisch in je path
```bash
scoop install tailwindcss
```

### draaien (in project map)
```bash
tailwindcss -i src/pisvlek.css -o public/css/pisvlek.min.css --watch --minify
```