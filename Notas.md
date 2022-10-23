
Foco do curso: Trabalhar com microserviços, apresentar esta forma de tecnologia e demonstrar um processo completo de desenvolvimento de arquitetura, ambiente, aplicação e posteriormente monitoramento.

Ferramentas a utilizar: Go Lang, NestJS, React, Kafta, Kafka Connect, Docker, Kubernets e Elastic Stack.



# Aula 1 - Tenha um dos perfis mais desejados do mercado de TI

**Microserviços**: Grandes sistemas importantes (que dependiam de funcionamento 100% do tempo) eram gerenciado por vários profissionais, cada um na sua especialização, porém está havendo uma mudança de cultura, cada grande serviço tem sido dividivo em serviços menores, os chamados microserviços. Para facilitar o processo de desenvolvimento dos microserviços, as empresas dividiram as equipes para ma cuidar dm microserviço, desta forma os desenvolvedores precisam entregar todo o ciclo de desenvolvimento do que foi designado a equipe, entendendo o processo como um todo.

**Full Cycle VS Full Stack**

**Full Cycle Developer:** Termo criado em 2018 através da Netflix, que chegou exatamente no modelo atual de full cycle após diversas tentativas fracassadas de melhoria de processos em suas tecnologias. 
Existem dois pilares que apoiam essa metodologia:
- Operate what you build (Operar o que você desenvolve)
- Ferramentas para escalar (Dominar ferramentas para entregar seu projeto a todo momento, como pipelines, docker, kubernets, ferramentas de monitoramento)

Quando falamos em Full Stack falamos de habilidades de desenvolvimento seja ele em front-end ou back-end, habilidade para desenvolver realmente todas camadas do projeto.

Porém, falando de Full Cycle, vamos além, e falamos de habilidades que englobam o ciclo todo de produção, desde o desenvolvimento da aplicação, ao desenvolvimento da arquitetura, entregando a aplicação já testada, "deployada" e ainda sendo já monitorada.

"Nem só de CRUD vive o desenvolvedor"



# Estudo de Caso Prático
# Code Delivery

- ## O que iremos desenvolver?
	- Sistema de entregas que permite visualizar em temp oreal o veículo do entregador
	- Há possibilidade de múltiplos entregadores simultâneos
	- Serviço simulador que enviará a posição em tempo real de cada entregador
	- Os dados de cada entrega, bem como as posições, serão armazenadas no Elasticsearch para futuras análises

- ## Alguns Desafios
	- Para evitar perda de informações caso o serviço backend fique indisponível por alguns momentos, NÃO trabalharemos com REST
		- Solução: Trabalharemos com o Apache Kafka para o envio e recebimento de dados entre os sistemas
	- Não é responsabilidade do serviço backend persistir os dados no Elasticsearch. Logo, como armazenar as informações no ELasticsearch?
		- Solução: Utilizaremos o Kafka Connect que também consumirá os dados do simulador e fará a inserção no Elasticsearch
	- Precisaremos exibir em tempo real a localização de cada entregador
		- Solução: Trabalharemos com websockets. O backend receberá os dados do simulador, e enviará as posições para o frontend via websocket.

- ## DInâmica do Sistema
	- ![[Pasted image 20221018213756.png]]

- ## Tecnologias a serem utilizadas
	- Simulador: Golang
	- Backend: Nest.js & Mongo
	- Frontend: React
	- Kafka & Kafka Connect
	- Elasticsearch & Kibana
	- Docker e Kubernets
	- Istio, Kiali, Prometheus & Grafana



# Apache Kafka (Conceitos Básicos)

Por quê utilizar?
Palavra-chave: Event-Driven

Event-driven
Tudo movido à eventos! Ex:
- Carros
- E-commerce
- Alarmes
- Monitoramento
- Microserviços

Atualmente, tudo se baseia em eventos gerados, o evento  gera todo um processamento por trás do evento, tudo simultâneo, e o Kafka vem para ajudar em tudo que possui eventos. O Kafka é um sistema em tempo real, ele tem uma latência baixa, e ao receber o evento dispara as ações muito rapidamente. Outro ponto é o histórico de dados, ele se sobresai nisto, pois guarda todo um histórico de eventos registrados, podemos escolher o tempo de retenção, e desta forma até mesmo para se adequar a LGPD, temos um rastro de todos dados.

Características do Kafka:
- Funcionar como plataforma
- Trabalha de forma distribuída
- Banco de dados integrado
- Extremamente rápido e com baixa latência
- Utiliza o disco ao invés de memória para processar os dados

**Não é apenas um sistema tradicional de filas como RabbitMQ.**

Conceito básico:
- Topic
	- Stream de dados que atua como um banco de dados
	- Todos os dados ficam armazenados, ou seja, cada "Topic" tem seu "local" para armazenar seus dados
	- Tópico possui diversas particões
		- Cada partição é definida por um número. Ex: 0, 1, 2
		- Você é obrigado a definir a quantidade de partições quando for criar um topic
		- ![[Pasted image 20221020143418.png]]
		- ![[Pasted image 20221020143616.png]]
- Kafka Cluster
	- Conjunto de Brokers (Máquina com Kafka instalado)
	- Cada Broker é um servidor
	- Cada Broker é responsável por armazenar os dados de uma partição
	- Cada partição de Topic está distribuído em diferentes Brokers
	- ![[Pasted image 20221020143949.png]]
	- ![[Pasted image 20221020144049.png]]
	- Replication Factor faz o Kafka ser muito poderoso, pois ele replica os dados de partições em vários brokers, assim os dados tem constante disponibilidade, com redundância de dados.
	- ![[Pasted image 20221020144239.png]]
	- O Kafka utiliza atualmente o ZooKeeper, que monitora os Brokers, e faz todo balanceamento na responsabilidade de cada Broker
	- O producer manda mensagens ao Kafka, e o Kafka fica responsável pelo encaminhamento da mensagem, após enviar ao broker ele vai controlando o Offset (log da mensagem)
	- ![[Pasted image 20221020161129.png]]
	- Após isto, temos os consumers, sistemas responsáveis por consumir as mensagens que chegam ao Kafta
- Ecossistema
	- Kafka Connect
		- Connectors (Recebe dados de um sistema, e toda vez que esse sistema atualizar ele envia a alteração ao Kafka, e do Kafka os dados são enviados a um destino, transporte de dados de um lado pro outro inteligente)
	- Confluent Schema Registry
		- Você cria o Schema da mensagem a ser enviada, garantindo o padrão das mensagens
	- Rest Proxy
	- ksqDB (Uma espécie de SQL do Kafka, permite diretamente no Kafka rodar "comandos de banco")
	- Streams (Manipulador de Informações do Kafka)










# Aula 2 - NestJS e o Backend orientado a microsserviços