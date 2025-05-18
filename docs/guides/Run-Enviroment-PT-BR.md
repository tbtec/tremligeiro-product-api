# Executando o ambiente

Utilizamos o Make para agregar todos os passos de uma execução e facilitar os passos necessários para buildar a imagem docker da aplicação e posteriormente subi-la junto aos demais componentes da arquitetura num cluster de Kubernetes.

Para isto, seguem os passos abaixo.

1- Realizar o build da imagem docker da aplicação através do comando:

```bash
make build-docker
```

2- A partir deste ponto considerando que o kubectl já está apontando para um cluster de kubernetes, basta seguir estes comandos do make que será realizar os passos de configuração e deploy

Este comando faz a criação de uma namespace e deploy de uma imagem postgresql.

```bash
make kube-config
```

Este comando faz o deploy dos manifestos necessários para a imagem da aplicação "tremligeiro".

É feito a criação de configmap, secrets, deployment, service e hpa.

A título didático, as secrets estão neste mesmo repositório. Para fins de produção, considerar separar o manifesto da secret em outro local, como também, utilização de cofres de senha.

```bash
make kube-deploy
```