// DATA STRUCTURES
// SINGLE LINKED LIST
// LINKED LIST TANPA STRUCT TAMBAHAN
package main

import "fmt"

type Node struct {
	Count int
	Data int
	Next *Node
}

func (this *Node) Append(Data int) {
	newNode := &Node{Data: Data}

	// cek apakah belum ada node headnya, jika belum maka data yang masuk diberikan ke head
	if this.Data == 0 && this.Count == 0 {	// perlu menggunakan count agar Data dari head bisa 0, karena metode pengecekan head sudah ada atau belum menggunakan Data == 0, maka dari itu kita menggunakan Count, inilah kelemahannya sedikit ribet karena pada structure linked list kita kali ini tidak menggunakan struct tambahan yg mana jika menggunakan struct tambahan maka head bisa nil sehingga pengecekan seperti ini tidak perlu lagi termasuk menambahkan properti Count
		this.Data = Data	// jika belum ada head maka isi head dengan data yang masuk
		this.Next = nil
		this.Count++
		return
	}

	// jika sudah ada node headnya, maka lanjut ke next dari headnya apakah sudah ada
	thisNode := this
	for {
		// jika next dari head belum ada, maka tambahkan data yang masuk di atas ke node baru
		if thisNode.Next == nil {
			thisNode.Next = newNode
			break
		}
		// jika next dari head sudah ada, maka lanjut cek ke next-nya lagi, dst
		thisNode = thisNode.Next
	}
}

func (this *Node) PrintList() {
	last := this
	for {
		if last == nil {
			break
		}
		fmt.Println(last.Data)
		last = last.Next
	}
}

func main() {
	Head := Node{}
	Head.Append(0)
	Head.Append(2)
	Head.Append(0)
	Head.PrintList()
}