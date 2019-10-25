package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
	"runtime"
	"strconv"
)



//params for db
var (
	dbhostip="127.0.0.1"
	dbusername="boy_test"
	dbpassword="password"
	dbname="lm970"
	dbCon *sql.DB
	err error
)

//checker for db
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}

func connectDB(){
	dbCon,err = sql.Open("mysql",dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
	checkErr(err)


	// stmt,err:=db.Prepare("SELECT l_id, l_code, l_title FROM app_location where l_id=10")
	// checkErr(err)

	// row :=stmt.QueryRow()
}
func closeDB(){
	defer dbCon.Close()
	// defer stmt.Close()
	// defer row.Close()
}

//app params
var (
	sum int
	runCnt int
	oneRoundCnt int = 1000 //one round loop count
	roundCnt int = 5 //how many round to run
)


//one round function for once round
func oneRound(roundId int){
	runtime.Gosched()


	rand.Seed(time.Now().UnixNano())

	for i := 0; i < oneRoundCnt; i++ {


		tmpId := rand.Intn(30)
		// fmt.Println(">>> the l_id for this round is", tmpId)
		if (tmpId == 0) {
			tmpId = rand.Intn(30) + 2
			// tmpId = 2
			// fmt.Println(">>> l_id = 0 found, new id is", tmpId)
		}

		row := dbCon.QueryRow("SELECT l_id, l_code, l_title FROM app_location where l_id=?", tmpId)
		// checkErr(err)

		var l_id string
		var l_code []byte
		var l_title string
		err = row.Scan(&l_id,&l_code,&l_title)
	    if err != nil {
	        fmt.Println("scan报错：", err)
	    }
		// checkErr(err)
		// fmt.Println(l_id,l_code,l_title)
		// fmt.Println(string(l_code))


		tmpStringCode := string(l_code)
		tmpInt := 0

		if(tmpStringCode != ""){

			tmpInt, err = strconv.Atoi(tmpStringCode)
		    if err != nil {
		        fmt.Println("string to int wrong : ", err)
		    }
		}


		//force fake loop
		for i := 0; i < 10; i++ {
			if(i % 2 == 0){
				sum += tmpId
			}else{
				sum += tmpInt
			}
		}

		runCnt++

	}



	completeFlag <- roundId
}


var completeFlag chan int = make(chan int, 5)

/**
	connect mysql db test
	table lm970.app_location

	sum calculate with l_code + mt_rand(0,99)
**/
func main(){

	// fmt.Println(runtime.NumCPU()) //show CPU core number
	// runtime.GOMAXPROCS(5) //not working

	startTime := time.Now().UnixNano()

	connectDB()

	for i := 0; i < roundCnt; i++ {
		go oneRound(i)
	}

	for i := 0; i < roundCnt; i++ {
		<- completeFlag
	}


	closeDB()


	

	fmt.Println("end sum is ", sum)
	fmt.Println(">>TT run time :", runCnt)



	endTime := time.Now().UnixNano()
	
	diffTime := float64(endTime - startTime)/1000000000
	fmt.Printf(">>>> load db %v times, used  %.9f s \n", runCnt, diffTime)



	fmt.Println(startTime)
	fmt.Println(endTime)

}
