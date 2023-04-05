# Ping-Pong 

```Este código em Go implementa uma aplicação que usa goroutines e canais para sincronizar a execução de duas rotinas que enviam mensagens "ping" e "pong" e uma rotina que as imprime na tela. A função ping() envia a mensagem "ping" para o canal c e aguarda um segundo antes de enviar outra mensagem. A função pong() faz o mesmo, mas envia a mensagem "pong". A função imprimir() recebe mensagens do canal c e as imprime na tela. ```

```O uso de canais garante que as rotinas ping() e pong() esperem a outra rotina enviar uma mensagem antes de enviar outra mensagem. A rotina imprimir() está em um loop infinito, lendo do canal c e imprimindo cada mensagem recebida. O programa termina quando todas as três rotinas são concluídas.```

```Este código é um exemplo simples de como usar goroutines e canais em Go para implementar um modelo de produtor/consumidor. Ele demonstra como a sincronização pode ser alcançada usando canais para garantir que as rotinas enviem e recebam mensagens em uma ordem específica e como as rotinas podem ser coordenadas usando a API WaitGroup da biblioteca padrão do Go.```