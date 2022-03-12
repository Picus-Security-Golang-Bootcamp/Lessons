package main

/*
	Concurrency: Tek bir işlemi bağımsız bileşenlere bölmek ve bu bileşenlerin verileri nasıl güvenli bir şekilde
		paylaştığını belirtmek için kullanılan bilgisayar bilimi terimidir.
				Concurrency is not parallelism: https://www.youtube.com/watch?v=oV9rvDllKEg
				GopherCon 2016: Ivan Danyliuk - Visualizing Concurrency in Go: https://www.youtube.com/watch?v=KyuFeiG3Y60

	Goroutine: Programda bulunan diğer Goroutine'lerle bağlantılı olarak bağımsız ve aynı anda çalışan bir function
			veya methoddur.

	Channels: Goroutinelerin birbiriyle haberleşmesi için channel kullanılır. Slices, maps gibi chan de bir built-in
		typedır.
		ch := make(chan int)
		https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html


		https://www.youtube.com/watch?v=PAAkCSZUG1c
*/

func main() {
	// <- operatorü channeldan veriyi okumak veya yazmak için kullanılır.
	//ch := make(chan int)
	//a := <-ch // read a value from ch and assign it
	//
	//b := 8
	//ch <- b // write the value in b to ch

	/* Unbuffered channels (default channel)
	Gönderilen değeri almaya hazır ilgili bir alma (<-chan) varsa, yalnızca gönderimleri (chan <-) kabul edeceklerdir
	*/

	/* Buffered channales
	Bu değerler için karşılık gelen bir alıcı olmaksızın sınırlı sayıda değeri kabul eder.
	ch := make(chan int,10)
	*/
	//
	//
	// Bir channel loop'a girebilir
	//for v := range ch {
	//	fmt.Println(v)
	//Channel kapanana kadar bu loop çalışacaktır veya bir break ve return'e ulaşmadıysa.
	//}

	// Channel'ı kapattıktan sonra channel'a yazma işlemi olursa panic ile karşılaşırsınız.
	//close(ch)

	// Bu şekilde channel'ın açık mı veya kapalı mı olduğunu kontrol edebiliriz
	//v, ok := <-ch
	//}

}
