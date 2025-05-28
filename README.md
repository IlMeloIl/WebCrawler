# WebCrawler

Um web crawler simples em Go.

## Funcionalidades

*   **Limitado por Profundidade:** Permite especificar uma profundidade máxima para o rastreamento para controlar o escopo.
*   **Restrito ao Domínio:** Rastreia apenas páginas dentro do mesmo host da URL base inicial.
*   **Análise de HTML:** Analisa o conteúdo HTML para extrair links (atributo `href` dentro de tags `<a>`).
*   **Concorrência Configurável:** Permite definir o número máximo de operações de rastreamento concorrentes.

## Uso

Execute o crawler a partir da linha de comando, fornecendo a URL inicial, a concorrência máxima e a profundidade máxima como argumentos:

```bash
./crawler <baseURL> <maxConcurrency> <maxDepth>
```

Alternativamente, executar diretamente sem compilar

 ```bash
go run . <baseURL> <maxConcurrency> <maxDepth>
```
