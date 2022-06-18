// DATA STRUCTURES
// SINGLE LINKED LIST
// LINKED LIST DENGAN MEMBUNGKUS HEAD DENGAN STRUCT
package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// kita masukkan Head ke dalam struct lagi agar saat Head bisa nil saat pertama kali kita buat linked list,
// bisa saja kita buat linked list tanpa struct tambahan tetapi Head tidak bisa nil atau harus langsung kita isi Data nya si Head saat pertama kali kita buat
type List struct {
	Head *Node
}

func (this *List) Append(Data int) {
	newNode := &Node{Data: Data}

	// cek apakah belum ada node headnya, jika belum maka data yang masuk diberikan ke head
	if this.Head == nil {
		this.Head = newNode
		return
	}

	// jika sudah ada node headnya, maka lanjut ke next dari headnya apakah sudah ada
	thisNode := this.Head
	for {
		// jika next dari head belum ada, maka tambahkan data yang masuk di atas ke node baru
		if thisNode.Next == nil {
			thisNode.Next = newNode
			return
		}
		// jika next dari head sudah ada, maka lanjut cek ke next-nya lagi, dst
		thisNode = thisNode.Next
	}
}

func (this *List) PrintList() {
	last := this.Head // copy this.Head dan gunakan copian tsb untuk mencari, jangan gunakan langsung this.Head karena ini bisa berakibat mengubah this.Head tsb di seluruh tempat (artinya Head di myList akan berubah)
	for {
		if last == nil {
			break
		}
		fmt.Println(last.Data)
		last = last.Next
	}

	// for {
	// 	if this.Head == nil {	// Jika langsung this.Head (tidak menggunakan copiannya) maka this.Head akan berubah menjadi this.Head.Next dan di dalam method-method lainnya akan ikut terubah juga (gampangnya Head di myList akan berubah)
	// 		break
	// 	}
	// 	fmt.Println(this.Head.Data)
	// 	this.Head = this.Head.Next
	// }
}

// FUNGSI DELETE YANG INI MENGGUNAKAN VARIABEL *NODE
func (this *List) Delete(Data int) {
	// jika data yg ingin dihapus ternyata data si head, maka hapus head denga cara mengubah head dengan next nya si head
	if this.Head.Data == Data {
		this.Head = this.Head.Next // perubahan ini menggunakan metode mengubah alamat, bukan menggunakan metode
		return
	}

	// ambil depan, sekarang dan belakang
	prev := this.Head
	current := prev.Next
	next := current.Next

	for {
		// jika sudah di tail (next == nil) dan data si tail masih
		// juga tidak sama dengan data yg diminta (*current.Data != Data)
		// maka data tidak ditemukan
		if next == nil && current.Data != Data {
			fmt.Println("Data not-found")
			return
		}

		// jika sudah di tail (next == nil) dan data si tail ternyata
		// sama dengan data yg diminta (*current.Data == Data)
		// maka si tail kita hapus dgn cara mengisi nil ke alamat ke nextnya dari si sebelum tail
		if next == nil && current.Data == Data {
			prev.Next = nil // ini akan mengubah langsung datanya karena .Next dimana kita langsung mengakses datanya, jika ditampung dulu ke variabel misalnya prevNext = prev.Next lalu variabel prevNext kita ubah, maka prev.Next tidak akan berubah
			return
		}

		// jika data sudah sama
		// maka hapus node yg sekarang dengan cara
		// mengisi next dari node sebelumnya dengan alamat node setelah node ini
		if current.Data == Data {
			// metode perubahan ini menggunakan metode perubahan alamat, bukan menggunakan metode perubahan isi
			prev.Next = next // ini akan mengubah langsung datanya karena .Next dimana kita langsung mengakses datanya, jika ditampung dulu ke variabel misalnya prevNext = prev.Next lalu variabel prevNext kita ubah, maka prev.Next tidak akan berubah
			return
		}

		// ubah depan dan sekarang
		prev = current
		current = next

		// jika sudah di tail pasti next dari tail adalah nil maka variabel next jgn diisi alamat dari current Next, tetapi isi dengan nil
		if current.Next == nil {
			// ubah belakang
			next = nil
			continue
		}

		// ubah belakang
		// jika belum sampai tail maka lanjutkan iterasi dengan mengubah variabel next terlebih dahulu dengan alamat current Next
		next = current.Next
	}
}

func main() {
	myList := &List{}
	
	myList.Append(3)
	myList.Append(1)
	myList.Append(2)
	myList.Append(6)
	myList.Append(5)
	myList.Append(4)

	myList.PrintList()
}

func (this *List) Sort() {
	length := this.Length()
	for index := 1; index < length; index++ {
		this.Sorting(index)
	}
}

func (this *List) Length() int {
	last := this.Head
	var count int
	for {
		count++
		if last.Next == nil {
			return count
		}
		last = last.Next
	}
}

func (this *List) Sorting(headIndex int) {
	var count int = 1

	// ambil sebelum dan sekarang dari head
	var currentHeadPrev *Node = nil
	currentHead := this.Head
	for {
		if count == headIndex {
			break
		}
		currentHeadPrev = currentHead
		currentHead = currentHead.Next //// ini nilai awal, selama for loop di bawah berjalan ini tidak akan berubah, akan berubah jika for di bawah sudah mencapai akhir
		count++
	}
	// fungsi di atas untuk menentukan head nya

	// ambil depan dan sekarang dari smallest
	smallestPrev := currentHeadPrev // ini nilai awal, smallestPrev akan berubah di dalam for loop di bawah
	smallest := currentHead         // ini nilai awal, smallest akan berubah di dalam for loop di bawah

	// ambil depan dan sekarang dari Node yang sedang dicari untuk pembanding
	currentNodePrev := currentHead
	currentNode := currentHead.Next

	for {
		// ubah smallest jika ternyata lebih besar
		if currentNode == nil {
			currentHead = smallest
			smallestPrev.Next = smallest.Next // 1
			smallest.Next = currentHead       // 2
			return
		}

		if smallest.Data > currentNode.Data {
			smallestPrev = currentNodePrev
			smallest = currentNode
			smallest.Next = currentNode.Next
		}

		// jika sudah di akhir
		if currentNode.Next == nil {
			// jika ternyata data terkecil sama dengan data dari head yang pertama maka jangan ada perubahan
			if smallestPrev == nil {
				return
			}

			// jika ternyata data terkecil sama dengan data dari head sekarang maka jangan ada perubahan
			if smallest.Data == currentHead.Data {
				return
			}

			// jika ternyata current head adalah head yang paling pertama maka ubah
			if currentHeadPrev == nil {
				this.Head = smallest
				smallestPrev.Next = smallest.Next // 1
				smallest.Next = currentHead
				return
			}

			// jika ternyata current head adalah bukan head yang paling pertama maka ubah current head dengan smallest
			currentHeadPrev.Next = smallest   // 3
			smallestPrev.Next = smallest.Next // 1
			smallest.Next = currentHead       // 2
			return
		}
		currentNodePrev = currentNode
		currentNode = currentNode.Next
	}
}

func (this *List) InsertAfter(Data int, DataAfter int) {
	newNode := &Node{Data: Data}

	currentNode := this.Head

	for {
		if currentNode == nil {
			return
		}

		if currentNode.Data == DataAfter {
			newNode.Next = currentNode.Next
			currentNode.Next = newNode
			return
		}
		currentNode = currentNode.Next
	}
}

func (this *List) InsertBefore(Data int, DataBefore int) {
	newNode := &Node{Data: Data}

	var currentNodePrev *Node
	currentNodePrev = nil
	currentNode := this.Head

	for {
		if currentNode == nil {
			return
		}

		if currentNode.Data == DataBefore {
			if currentNodePrev == nil {
				this.Head = newNode
				newNode.Next = currentNode
				return
			}
			newNode.Next = currentNode
			currentNodePrev.Next = newNode
			return
		}
		currentNodePrev = currentNode
		currentNode = currentNode.Next
	}
}