package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	//MaxWorker :
	MaxWorker = 100000
	//MaxQueue :
	MaxQueue = 100000
)

//JobQueue 一个带缓冲的全局channel，用来存放接收到的job
var JobQueue = make(chan Job, MaxQueue)

//Dispatcher :
type Dispatcher struct {
	WorkerPool chan chan Job //用来注册worker的池
	MaxWorkers int           //worker最大个数
}

// Job 代表一个任务，根据自己需求定义
type Job struct {
	//Id :
	ID string
	//Payload :
	Payload string
}

// Worker 用来处理job的实例
type Worker struct {
	JobChannel chan Job //用来接受任务的通道
	quit       chan bool
	id         int
}

//var Pool chan chan Job //需要注册到的worker池

//Done 一个简单的job处理函数
func (j *Job) Done() (string, error) {
	//time.Sleep(100 * time.Millisecond) //模拟任务处理时间
	time.Sleep(5 * time.Second)
	fmt.Println("job:", j.ID, j.Payload, "done...")
	return fmt.Sprintf("job %s done %s ....", j.ID, j.Payload), nil
}

//NewWorker :
func NewWorker(wid int) *Worker {
	return &Worker{
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		id:         wid,
	}
}

//Start :函数，启一个goroutine， 启动的时候将自己注册到worker池当中，然后就等待job被放到自己的jobChannel中
//一旦jobChannel中有job放入的时候，就开始处理这个job
//同时加入了一个quit channel可以用来控制销毁这个worker
func (w *Worker) Start(d *Dispatcher) {
	go func() {
		for {
			d.WorkerPool <- w.JobChannel //注册当前这个worker到worker池中，
			//fmt.Printf("Work_Id: %d\n", w.id)
			// 也就是将自己的jobChannel放入到池中，用来接收job
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Work_Id_job: %d\n", w.id)
				//channel中放入了一个job
				if _, err := job.Done(); err != nil {
					//处理这个job

				} else {
					//w.quit <- true
					w.Stop()
					fmt.Printf("Job done. Let it out. work_id: %d\n", w.id)
				}
			case <-w.quit:
				fmt.Printf("Work_Id_quit: %d\n", w.id)
				// 收到停止的信号，销毁这个worker
				return
			}
		}
	}()
}

// Stop 一个worker，相当于销毁一个worker，也就是从worker pool中去掉
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//NewDispatcher :
func NewDispatcher(maxWorkers int) *Dispatcher {
	Pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: Pool, //初始化worker池
		MaxWorkers: maxWorkers,
	}
}

//Run :
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		//对每一个worker进行初始化，也就是将worker注册到池中，更直接点就是将每个worker的jobChannel放入到池中
		worker := NewWorker(i)
		worker.Start(d)
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Printf("select job_id : %s.\n", job.ID)
			// 收到一个job
			go func(job Job) {
				// 从worker池中，选取一个worker的jobChannel，如果worker池中是空的，则会阻塞在这里
				jobChannel := <-d.WorkerPool
				// 将job放入到其中一个worker的jobChannel中，等待这个worker进行处理
				jobChannel <- job
				fmt.Printf("put job_id: %s in jobchannel\n", job.ID)
			}(job)
		}
	}
}

//JobCreator :
func JobCreator(w http.ResponseWriter, r *http.Request) {
	//id := uuid.NewV4()
	id := rand.Intn(1000000)
	oneJob := Job{
		ID:      strconv.Itoa(id),
		Payload: GetRandomString(10),
	}
	fmt.Printf("push job_id: %s in queue.\n", oneJob.ID)
	// 将job实例放入到job队列当中
	JobQueue <- oneJob
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Response OK")
}

//JobCreator2 :
func JobCreator2(w http.ResponseWriter, r *http.Request) {
	//id := uuid.NewV4()
	id := rand.Intn(1000000) + 1000000
	j := Job{ID: strconv.Itoa(id),
		Payload: GetRandomString(10)}

	// 将job实例放入到job队列当中
	time.Sleep(5 * time.Second) //模拟任务处理时间
	fmt.Println("job2:", j.ID, j.Payload, "done...")
	w.WriteHeader(http.StatusOK)
}

//GetRandomString :
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	//time.Sleep(5 * time.Second) //模拟任务处理时间
	return string(result)
}

func main() {
	//初始化一个dispatcher
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	http.HandleFunc("/job", JobCreator)      //设定访问的路径
	http.HandleFunc("/job2", JobCreator2)    //设定访问的路径
	err := http.ListenAndServe(":8887", nil) //设定端口和handler
	if err != nil {
		panic(err)
	}
}
