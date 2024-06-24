# schedulers: leak-goroutines


## how to run
``` shell
### terminal-01
GOMAXPROCS=1 GODEBUG=schedtrace=1 go run main.go

### terminal-02
for i in $(seq 10); do echo "## $i"; curl -s "http://localhost:8080/leak"; done
ab -n 100000 -c 20 http://localhost:8080/leak
curl -s http://localhost:8080/status

SCHED 9591ms: gomaxprocs=1 idleprocs=1 threads=6 spinningthreads=0 needspinning=0 idlethreads=1 runqueue=0 [0]
tempo de execução em ms
gomaxprocs: quantidade de cores em uso
idleprocs: quantidade de cores ociosos
threads: quantidade de goroutines geradas pelo runtime
spinningthreads: quantidade de threads de backup
needspinning: quantidade de spinningthreads sendo demandadas
idlethreads: threads ociosas
runqueue: tamanho global da fila [tamanho atual da fila]
```