# Disponibilidade

O arquivo k8s/hpa.yaml possui a configuração de "Horizontal Pod Autoscaler" da aplicação Trem Ligeiro.

Utilizamos a quantidade mínima de réplicas sendo 2 instâncias da aplicação, pois, em caso de alguma apresentar algum tipo de falha, haverá uma outra disponível para atender as requisições, enquanto o próprio k8s pode realizar o recycle da pod com problemas.

```bash
  minReplicas: 2
  maxReplicas: 6
```
O valor de 6 foi definido para uma margem de segurança considerando os momentos de pico de uso da aplicação. Conforme práticas de monitoramento, pode ser necessário ajustar este valor, como também, avaliar o tipo de instância dos componentes que não escalam horizontalmente (banco de dados relacional) e, se necessário, aumentar os recursos do mesmo.