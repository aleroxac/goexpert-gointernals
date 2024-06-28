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
- garbage collector
    - o garbage collector(GC) é um mecanismo automático de gerenciamento de memória que busca, identifica e libera memória que não está mais sendo utilizada pelo grograma. Isso é crucial para previnir vazamentos de memória e garantir a eficiência do uso de memória
    - caracteristicas do GC do go:
        - não-geracional: trata todos os objetos igualmente, sem distinção entre objetos novos e antigos
        - concorrente: executa a maior parte do trabalho de coleta de lixo concorrentemente com a aplicação, minimizando as pausas
        - baseado na técnica de "Tri-color Mark and Sweep": utiliza o algoritmo de marcação e varredura com tês cores(branco, cinza e preto) para gerenciar os objetos
    - objetos alcançáveis
        -  podem ser acessados direta ou indiretamente por referência em um ponto de entrada no programa
            - roots: são os pontos de entrada iniciais para a busca de objetos alcançáveis. Incluem variáveis globais, variáveis locais atualmente ativas nas stacks de execução, e registros de CPU
            - objetos referenciados: qualquer objeto que é referenciado direta ou indiretamente a partir de um objeto root
        - exemplos
            - 01
                - se uma variável global referencia um objeto A, e o objeto A referencia um objeto B, então ambos os objetos A e B são alcançáveis
            - 02
                - uma variável global globalVar refenrecia o objeto A
                - o objeto A referencia um objeto B
                - o objeto B referencia um objeto C
                - o objeto D não é referenciado por nenhum outro objeto
                - A, B e C são alcançáveis porque podem ser acessados a partir da variável global globalVar. O objeto D é inalcançável porque não há nenhuma referência a ele a partir das raízes ou de outros objetos alcançáveis
    - dinamica
        - aplicação > GC > write barrier> 
        - GC:
            - 1. SWT(Stop the World): para a execução do programa para rodar o GC
                - Write Barrier: intercepta e pausa todas as chamadas do programa
                - Mark Setup
            - 2. Marking Work(concurrent) - nesse momento o programa volta a ser executado, o resto do trabalho vai ser feito concorrentemente
                - 25% do CPU é alocado para fazer o processo de marcação de objetos alcançáveis
                - Mark Assist: pega outras goroutines para trabalhar junto com o GC para ajudar a fazer as marcações
                    - branco: objeto ainda não explorado
                    - cinza: objetos alcançáveis
                        - pendentes de processamento
                        - precisa buscar por referência
                        - trabalha de forma recursiva
                    - preto: objeto já explorado
            - 3. Mark Termination: novamente interrompe a execução do programa e faz a varredura em busco de novos objetos alcançáveis
                - finaliza a marcação
                - desliga o Write Barrier
            - Sweeping(Concorrente)
                - identifica e libera a memória de objetos não alcançáveis
                - varredura on-demand
    - GOGC
        - define o tamanho da heap quando o GC deve ser acionado
        - por padrão é 100%
        - exemplo
            - se a heap após a última coleta de lixo for de 4MB e o GC Percentage estiver definido como 100%, o próximo GC será acionado quando o tamanho total do heap atingir 8MB(4MB + 100% disso, ou seja, mais 4MB, totalizando 8MB)
        - quanto mais baixo o número, mais frequente sera ativado o GC
        - GOGC=100 (variável de ambiente)
    - GC Trace
        ```
        gc 1 @0.019s 0%:0.014+0.56+0.010 ms clock, 0.029+0/0.55/0+0.021 ms cpu, 4->4->1 MB, 5MB goal, 8 P
        ```
        - gc 1: numero do clico de GC< começando em 1 para o primeiro GC que ocorre após a inicialização do programa
        - @0.019s: o tempo desde o início do programa até o início deste ciclo de GC
        - 0%: a porcentagem do tempo total do programa gasto em GC até este ponto
        - 0.014+0.56+0.010 ms clock
            - 0.014: antes da coleta - este valor indica o tempo gasto antes de iniciar efetivamente a fase concorrente de marcação(marking). Pode incluir preparações iniciais e o tempo para iniciar a coleta de lixo. O "0.014ms" sugere que foram gastos 14 microssegundos em atividades preliminares antes de iniciar a marcação propriamente dita.
            - 0.56: fase concorrente - o valor "0.56ms" representa o tempo gasto na fase concorrente do GC< que geralmente envolve a marcação(marking) de objetos alcançáveis no heap. Durante esta fase, o programa continua em execução normalmente. enquanto o GC trabalha para identificar quais objetos ainda estão sendo usados. Os 560 microssegundos indicam o tempo total despendido nessa atividade concorrente.
            - 0.010: finalização da coleta - o "0.010ms" é o tempo gasto aós a fase concorrente, incluindo a finalização da marcação e a preparação para a fase de varredura(sweep). Esses 10 microssegundos podem cobrir a conclusão do trabalho de marcação e as atividades de limpesa necessárias antes de o GC prosseguir para a próxima etapa.
        - 0.029+0/0.55/0+0.021 ms cpu
            - 0.029+0: indica o tempo gasto na fase de STW_SWEEP_TERMINATION. "0.029" é o tempo de STW para esta fase, e o "+0" indica que não houve tempo adicional significativo gasto após a pausa inicial
            - 0.55: tempo gasto na fase de marcação e varredura(MARK_AND_SWEEP) que é feita de forma concorrente, sem STW. "0.55" indica o tempo gasto nessa fase
            - 0+0.021: tempo gasto na fase de STW_MARK_TERMINATION. "0" indica que não houve tempo inicial de STW sifnificativo antes dessa fase, e "+0.021" é o tempo de pausa STW para finalizar a marcação.
        - 4->4->1 MB: tamanho do heap antes do GC, o tamanhho do heap ao determinar iniciar o GC, e o tamanho do heap após o GC, respectivamente
        - 5MB goal: o tamanho alvo do heap para o próximo GC, baseado na heurística do GC para tentar manter o tempo de pausa ou a frequência do GC dentro de limites desejáveis
        - 8P: número de processadores lógicos(P's) usados pelo Go scheduler

        - clock-time: tempo total que o GC está tomando do ponto de vista de um observador externo. Isso inclui todos os aspectos de execução e espera. Tempo, do início ao fim
        - cpu-time: tempo que a CPU está ativamente trabalhando no GC, excluindo tempos de espera ou quando outras goroutines estão sendo executadas

        - histórico de melhorias do GC
            - go 1.5: introdução do GC concorrente
                - substitui o modelo de GC Stop-The-World(STW) por um modelo concorrente, reduzindo significativamente as pausas do GC e melhorando a performance geral das aplicações Go
            - go 1.8: otimizações no GC
                - redução das pausas do GC ao aprimorar a eficiência da fase de varredura(sweep) e da assistência de marcação(mark assist)
            - go 1.14: implementação do Non-Cooperative Preemption
                - embora não seja uma mudança no GC em si, essa otimização na preempção de goroutines teve impactos positivos na latência do GC, permitindo que o runtime interrompesse goroutines mais eficientemente para garantir que o GC pudesse rodar a tempo
            - go 1.15 e 1.16: redução de alocações
                - otimizações nessas versões reduziram as alocações descenessárias, diminuindo a pressão sobre o GC
            - go 1.19: Soft Memory Limit
                - introdução de um limite de memória "soft" para o GC, permitindo que os desenvolvedores definam um alvo de uso de memória que o GC tentará respeitar, melhorando a gestão da memória em sistemas com restrições



## Channels no Go
- channels são um mecanismo de comunicação e sincronização entre goroutines no Go. Eles permitem que goroutines troquem dados de maneira segura e eficiente, suportando a construção de programas concorrentes

- o problema da não utilização de channels
    - problemas de sincronização
        - quando a G1 quer passar um valor para a G2, ela altera um valor na memória para a G2 ter acesso
        - o grande problema é que outras goroutines e partes do programa também podem ter acesso aquele endereço de memória
        - ou a G1 não terminou completamente de alocar o valor em memória e a G2 já fez a leitura e eventualmente uma gravação no mesmo local
    - dificuldade de trabalhar com concorrência
        - data race(race condition)
        - para remediar o problema utilizamos Mutex(Mutual Exclusion)
            - fazemos um lock no valor na memória e durante esse momento, somente uma goroutine pode fazer alteração. Após isso, esse valor é liberado(unlock)
            - mutex e similares abrem muita marge para erro, pois tudo isso é feito de forma manual
    - deadlocks
        - quando uma goroutine-2 quer acessar algum dado bloqueado na goroutine-1, e a goroutine-1 quer acessar algum dado bloqueado pela goroutine-2, onde acontece o cenário do programa ficar indefinidamente bloqueado nessas 2 goroutines

- a frase que define com mais clareza a utilização de channels
    - "Do not communicate by sharing memory; instead, share memory by communicating"
        - essa frase encapsula um dos princípios fundamentais do design de sistemas concorrentes no Go. A ideia é que, ao usar channels para comunicação entre goroutines, você evita muitos dos problemas associados à concorrência e ao compartilhamento de memória direta, como condições de corrida e deadlocks
        - em vez de várias goroutines acessarem diretamente variáveis compartilhadas(o que requer mecanismos de sincronização como locks), elas se comunicam enviando dados através de channels, o que proporciona uma maneira segura e clara de coordenar o acesso aos dados

- channels são um mecanismo fundamental de comunicação entre goroutines, permitindo a troca segura e sincronizada de dados
- eles são utilizados para passar informações entre goroutines de forma eficiente, evitando a necessidade de locks explícitos e reduzindo a complexidade da sincronização

- tipos de channels
    - não-bufferizados: requerem que a operação de envio e recebimento ocorra simultaneamente. Ideal para sincronização direta entre goroutines
    - bufferizados: permitem que dados sejam armazenados temporariamente no buffer, permitindo que a goroutine de envio ea de recebimento sejam executadas em tempos diferentes
