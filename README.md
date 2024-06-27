# goexpert-gointernals


## Multitask - Timeline
- 1940-1960: pré-multitask; programação via cartões perfurados, tarefa por tarefa
- 1960-1970: sistemas de tempo compartilhado; multi-user, mainframes
- 1980: os-multitask; unix
- 1990-2000: hyper-threading; UI, processos e threads
- 2000: multi-core; multiplos cores permitindo tasks simultaneas, paralelismo real
- 2010: otimizações para nuvem, ia, etc: otimizações e especializações de tarefas



## Processos
- o que é um processo: instância de um programa em execução
- componentes de um processo:
    - endereçamento: região da memória dedicada a um processo
    - contextos:
        - conjunto de dados que o SO salva para gerenciar um processo
        - possui o endereço de memória da próxima instrução que o processador irá executar
        - auxilia no context-switch
    - registros do processador:
        - áreas que temporarioamente armazenam no GPU dados e endereços para realizar a execução
        - dados: exemplo=realiza operações aritméticas e lógicas
        - registro de endereços:
            - armazenamento em memória, incluindo stack pointers; exemplo=ao acessar uma variável o CPU possui um registro na memória para guardar seu valor
    - heap:
        - utilizado para alocação de memória dinâmica. Cresce e encolhe em tempo de execução conforme a necessidade de mais ou menos espaço
    - stack:
        - armazena informações de controle para chamadas de função, como endereços de retorno, e parâmetros de função
        - segue uma estrutura LIFO(Last In, First Out - Último a entrar, primeiro a sair)
    - registros de status/flags:
        - fornecem os status recentes das operações realizadas pelo CPU
        - trabalham atravéz de bits específicos(flags)
            - ex: flag-zero(z): resultado de uma operação o qual o resultado é zero. Decide o fluxo do programa baseado nesse valor
            - ex: flag-sigma(s) ou negative(n): indica o resultado de uma operação positiva ou negativa
            - ex: flag-overview: produz resultado além da capacidade



## Ciclo de vida de um processo
### Creation
- Um novo processo é criado quando um programa solicita a execução de um processo, por meio de chamadas de sistema como fork() no UNIX/Linux ou CreateProcess() no Windows
### Execution
- O processo está ativamente sendo executado pela CPU. Pode alterar entre os estados de "executando" e "pronto"(para ser executado)
- Waiting/Blocked:
    - o processo é suspenso e colocado em espera até que um evento externo ocorra. Comum em operações de I/O, onde o processo aguarda pelo término de uma leitura de disco ou recebimento de uma entrada de rede.
- Termination: 
    - o processo completa sua execução ou é forçadamente terminado
    - exit: conflusão bem-sucedida do processo após completar suas instruções
    - killed: interrupção por meio de um erro de execução ou por ser terminado por outro processo(por exemplo, através do comando kill)



## Criação de um novo processo no SO
- unix/linux
    - fork()
    - clona o processo atual
    - gerado um processo filho
    - fork() retorna um valor diferente para o processo pai(PID)
    - processo pai e filho são quase idênticos, porém os valores na memória são copiados para outro endereçamento separado e independente
    - processo pai recebe um PID(valor inteiro positivo) do filho quando o fork() é chamado
    - processo filho retornar o PID 0, indicando que é um processo filho



## Gerenciamento de processos
- scheduler
    - decide qual processo será executado
    - alterna entre processos
    - possui diversos algoritmos para atentar maximizar o uso da CPU
    - scheduper pode:
        - selecionar processos de uma fila que estão "ready queue"
        - alocar CPU: mudança de estado - ready to running
        - retirar CPU: I/O, etc
- tipos de schedulers
    - colaborativo / cooperativo: processos que estão sendo executados tem controle quando liberam a CPU para outros processos
        - contras: processos podem monipolizar a CPU
    - preemptivo: SO tem a capacidade de interromper um processo em execução e ceder o uso da CPU para outro processo. Trabalha de forma mais "justa"
        - contras: muitas mudanças de contexto



## Threads
- processos são instâncias de programs em execução
- threads são unidades básicas de utilização de CPU que fazem parte dos processos
- threads são sequências de execução dentro do mesmo processo, compartilhando o mesmo espaço de memória e recursos
- dentro de um único processo, várias threads podem existir, cada uma executando diferentes partes do programa
- paralelismo vs concorrência: com múltipos CPUs conseguimso atingir paralelismo. Com apenas um núcleo, trabalhamos de forma concorrente(simulando paralelismo)
- threads, obviamente, ocupam menos espaço na memória do que um processo, pois elas compartilham a mesma memória do processo
- cada thread possui sua stack indepentente e isolada
- cada thread ocupa ~= 2MB(linux)
- cada thread, no go, ocupa 2KB


---


## Runtime Architecture
- goroutines = threads do go, também conhecidas como light-threads, green-threads, vitual-threads, etc
- scheduler = scheduler do go; ajuda no agendamento da execução das goroutines
- channels = para trabalhar de forma concorrente e possibilitar comunicação e sincronização entre treads(goroutines)
- memory-allocation
- garbage-collector
- stack management
- network-poller
- reflection



## Padrão M:N
- Threads Virtuais(threads geradas pela linguagem de programação) vs Threads Reais(gerados pelo SO)
- Modelo de agendamento de tarefas
    - M: threads virtuais em "user land, green threds, light threads"
    - N: threads reais do sistema operacional
    - M tarefas para N threads



## Goroutines
- funções/métodos que são executadas de forma concorrente
- são "threads" gerenciadas pelo runtime do go
- muito mais "baratas" do que criar novas threads no sistema operacional(2KB)
- muito mais rápido de criar e destruir
- compartilham os mesmos endereços de memória do programa em go. Possuam stacks independentes



## Padrão M:P:G
- machine < processor > goroutine
- quando o go inicia, ele cria as threads reais e "gruda" 1 processor por machine
- em geral, será 1 processador lógico por machine
- ao longo da execução do programa, novas threads reais podem ser criadas
- o processador lógico é quem faz o link com a goroutine, e ele é quem fala com a thread real



## runtime.GOMAXPROCS()
- go cria um P(processor) por Núcleo Computacional
- go tente a criar um M(Machine - Threads) para atribuir para cada P
- o valor de uma Machine por processor não é fixo
    - go pode criar mais threads no SO se as atuais estiverem bloqueadas por I/O ou outro motivo de executar as Goroutines
    - o objetivo é sempre manter os Ps ocupados, sem tempo ocioso



## Scheduler: Pool de Ps
- gestão de como e quando as tarefas são executadas em threas do sistema operacional
- decide qual tarefa deve ser executada em qual thread e em que momento
- gerencia o balanceamento de carga entre diferentes threads ou processadores lógicos, garantindo que nenhuma thread fique sobrecarregada enquanto outras estão ociosas
- gerencia questões como sincronização, mutex, racing conditions, deadlocks, etc



## Scheduler no Go
- o trabalho do Scheduler é combinar em G(o código a ser executado), um M(onde executá-lo) e um P(os direitos e recursos para executá-lo).
- quando um M para de executar um código Go do usuário, por exemplo, ao entrrar em uma syscall, ele devolve seu P para o pool de P ociosos.
- para retomar a execução do código Go do usuário, por exemplo, ao retornar de uma syscall, ele deve adquirir um P do pool de ociosos.
- https://go.dev/src/runtime/HACKING
- detalhamento
    - scheduler faz parte do runtime. Trabalha de forma adaptativa
        - atribuição de tarefas
        - balanciamento de carga
        - gerenciamento de concorrência
    - trabalha de forma não cooperativa com preempção(versão >= 1.14)



## Scheduler no Go vs Goroutines
- scheduler determina o estado de cada goroutine
    - running
    - runnable(fila)
    - not runnable(bloqueada fazendo I/O, por exemplo)
- work stealing
    - se o P está ocioso(idle)
    - ele rouba goroutines de outro P ou mesmo da fila global de goroutines
        - verifica 1/61 do tempo, evitando overhead para evitar buscar na fila global o tempo todo



## Preempção no Go
- sinalização de preempção
    - nivel de sistema
        - go insere pontos de premepção usando recursos do SO(exemplo: signals)
    - verificação de pontos onde as goroutines podem ser seguramente interrompidas
        - localizados em funções que são chamadas de forma frequente / loops
    - funções longas
        - se uma função está sendo executada sem chamar outras funções ou fazendo I/O por muito tempo, ela está desafiando o scheduler. Logo, o go internamente vai realizar a preempção mesmo sem ter os pontos de sinalizadores



## Memória
### Conceitos básicos
- memória de acesso rápido; memória que fica no chip da CPU, utilizada como cache
    - l1: 64kb
    - l2: 0.5mb
    - l3: 8mb
- memória de acesso lento
    - DDR(double data rate); clock consegue ter acesso 2 vezes por ciclo
    - é ligada através de um barramento(canais de comunicação entre CPu e a memória)
- endereços de memória são referenciados em formato hexadecimal(de 0-9 e de A-F)
### Custo de memória
- threads, obviamente, ocupam menos espaço na memória do que um processo, pois elas compartilham a mesma memória do processo
- cada thread possui sua stack independente e isolada
- cada thread ocupa, em média, 2MB(linux)
- stack < heap < static data < literals < instructions
- function_c() > function_b() > function_a() > free memory space
### heap
- dynamic memory allocation
    - large memory pool; a heap pega um bloco de memória para si, que é muito maior do que a stack
    - flexibilidede; possibilitando organizar informações em diversos locais de memória
    - acessível globalmente
    - reutilizável
    - suporta estruturas complexas
    - gerenciamento completo
    - leaks
    - fragmentação; buracos nos blocos de memória causados durante processos de alocação e desalocação
    - mais lento que a stack
    - concorrência
### Fragmentação
- falsa sensação de que existe o espaço necessário na memória, mas quando temos blocos de dados maiores e que precisam ficar juntos, percebemos que não temos esse espaço
- arenas
    - blocos de espaço em memória, uma separação lógica para execução alguma tarefa; também possibilitando fazer subdivisões
    - subdivisão da heap em chunks
        - por velocidade
            - fast bin
            - small bin
            - large bin
        - por tamanho
            - 8KB
            - 64KB
            - >1MB
- alocadores
    - alocadores mais populares
        - malloc/free (C std library)
        - dlmalloc(Doug Lea's Malloc) - Não suporta de forma eficiente Multithreading
        - ptmalloc / ptmalloc2 (pthreads Malloc) - Utilização de arenas
        - jemalloc (Jason Evans) - otimizado por Facebook, Rust, Postgres
        - TCMalloc (Thread-Caching Malloc) - Google

## Memória no Go
- utiliza como baseo TCMalloc(desenvolvido pelo Google)
  - ao longo do tempo, o alocador tomou diferentes caminhos do TCMalloc
  - o próprio runtime do Go é responsável por trabalhar com a alocação de memória
- nome do alocador é 'mallocgc'
  - flow: G(goroutine) > P(processor) > mcache > mcentral > mheap > OS
  - tipos
    - tiny: objetos < 16 bytes
    - small: objetos entr 16 e 32KB
    - large: objetos > 32KB
  - gerenciamento de memória
    - separa os chunks em spans, que são blocos de páginas da memória heap
    - mheap[ N[spans] ] > N[mcentral(gerencia spans de N diversos tamanhos)] > N[mcache(cache local)]
