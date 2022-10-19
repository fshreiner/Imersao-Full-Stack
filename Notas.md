
Foco do curso: Trabalhar com microserviços, apresentar esta forma de tecnologia e demonstrar um processo completo de desenvolvimento de arquitetura, ambiente, aplicação e posteriormente monitoramento.

Ferramentas a utilizar: Go Lang, Nactive JS, React, Kafta, Kafka Connect, Docker, Kubernets e Elastic Stack.



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


44:00

