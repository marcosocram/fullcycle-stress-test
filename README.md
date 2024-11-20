# fullcycle-stress-test

Este é um sistema CLI (Command-Line Interface) desenvolvido em Go para realizar testes de carga em serviços web. O objetivo é medir o desempenho de um serviço HTTP sob alta carga, simulando várias requisições simultâneas e gerando relatórios detalhados.

## Como funciona

O Stress Test funciona simulando múltiplos clientes acessando o serviço HTTP de forma simultânea.

1. Parâmetros de Entrada:
   * **URL** (`--url`): O endereço do serviço a ser testado.
   * **Requests** (`--requests`): O número total de requisições que serão enviadas.
   * **Concurrency** (`--concurrency`): O número de requisições enviadas simultaneamente.

2. Execução:
   * O sistema distribui as requisições entre várias goroutines.
   * Cada resposta HTTP é registrada, incluindo o código de status retornado.

3. Relatório:
   * Após a execução, um relatório é gerado com as seguintes informações:
     * URL testada.
     * Tempo total de execução.
     * Total de requisições realizadas.
     * Contagem de respostas com código HTTP 200.
     * Distribuição de outros códigos de status HTTP (ex.: 404, 500, etc.).
   * O relatório é salvo em um arquivo `.txt` no diretório mapeado do host.

## Requisitos
* **Docker** instalado na máquina.

## Como Executar

1. Clone o repositório:
   ```bash
    git clone https://github.com/marcosocram/fullcycle-stress-test.git
    cd fullcycle-stress-test
    ```

2. Construa a imagem Docker:
   ```bash
   docker build -t fullcycle-stress-test .
   ```
   
3. Execute o teste de stress com mapeamento de volume:
   ```bash
   docker run -v $(pwd)/reports:/app/reports fullcycle-stress-test --url=http://localhost:8080 --requests=1000 --concurrency=10
   ```
   * Substitua `http://localhost:8080` pela URL do serviço a ser testado.
   * Substitua `1000` pelo número total de requisições a serem enviadas.
   * Substitua `10` pelo número de requisições enviadas simultaneamente.
* O diretório `reports` no host será mapeado para `/app/reports` no container.
* O relatório gerado será salvo no diretório `reports` do host.

4. Verifique os relatórios gerados:
   ```bash
   ls -l reports/
   ```
   Exemplo de saída:
   ```
   reports/report-20241120-153045.txt
    ```
   Exemplo de conteúdo do relatório:
   ```txt
    Relatório de Teste:
    URL Testada: https://stackoverflow.com/
    Tempo Total: 4.199896654s
    Requests Totais: 500
    Status 200: 120
    Status 429: 380
   ```



