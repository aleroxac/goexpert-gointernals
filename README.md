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
