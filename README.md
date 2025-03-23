# Consulta de IP e Hosts de Domínios em Go

Este projeto em **Go** permite consultar o **IP público** de um domínio e obter informações sobre onde ele está hospedado.

## 🚀 Funcionalidades
- Consultar o **IP** de um domínio.
- Recuperar informações sobre o domínio.

## 🛠 Tecnologias Utilizadas
- **Go (Golang)**
- **Pacote Linha de comando** ([urfave/cli](https://github.com/urfave/cli))

## 📦 Instalação
1. **Clone o repositório**:
   ```sh
   git clone https://github.com/gilberto-dev-silva/command-linegit
	 
   cd command-linegit
   ```
2. **Execute o seguinte comando**:
   ```sh
   go mod tidy
   ```

## 🔧 Como Usar
1. **Execute o programa**:
   ```sh
   go run main.go
   ```
2. **Saída esperada**:
   ```sh
   NAME:
   Consulta - IP Servidores
	USAGE:
		main.exe [global options] command [command options] [arguments...]

	COMMANDS:
		ip, i      Busca IPS de endereços na internet
		server, s  Busca domínios na internet
		help, h    Shows a list of commands or help for one command

	GLOBAL OPTIONS:
		--help, -h  show help
   ```
3. **Busca por IP**:
   ```sh
   go run main.go i --host amazon.com.br
   ```
4. **Saída esperada**:
	```sh
	52.94.225.243
	72.21.203.171
	54.239.26.87
	```
5. **Busca por Server**:
   ```sh
   go run main.go s --host amazon.com.br
   ```

6. **Saída esperada**:
	```sh
	ns1.amzndns.co.uk.
	ns1.amzndns.com.
	ns1.amzndns.net.
	ns1.amzndns.org.
	ns2.amzndns.co.uk.
	ns2.amzndns.com.
	ns2.amzndns.net.
	ns2.amzndns.org.
	```

## 🔧 Build do projeto
1. **Execute o comando**:
   ```sh
   go build
   ```

2. **Busca por servidor**:
   ```sh
   ./command-line.git.exe s --host amazon.com.br
   ```

3. **Busca por IP**:
   ```sh
   ./command-line.git.exe i --host amazon.com.br
   ```

## 📜 Licença
Este projeto está licenciado sob a **MIT License** - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 🤝 Contribuição
Sinta-se à vontade para contribuir! Abra uma **issue** ou envie um **pull request**. 😉

