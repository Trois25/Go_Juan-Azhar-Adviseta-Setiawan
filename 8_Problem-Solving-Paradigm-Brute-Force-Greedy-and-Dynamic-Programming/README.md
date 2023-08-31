**Problem solving paradigm, sebuah paradigma dimana kita menyelesaikan masalah dengan beberapa prinsip

* Complete Search(Brute Force), melakukan scanning dari awal sampai akhir (O(n)). poblem solving ini digunakan apabila tidak ada penggunaan lain yang lebih cepat.

* Devide and conquer, pencarian digunakan dengan pemecahan data menjadi kelompok yang lebih kecil (binary tree). contoh penggunannya adalah binary search dan syarat penggunaannya adalah data harus di sorting dengan urut terlebih dahulu

* Greedy, menyelesaikan masalah dengan menggunakan lokal optimal. algoritma ini digunakan untuk data yang berbentuk graph, lokal optimal adalah mencari jalur yang terkecil ketika akan berpindah titik tanpa memikirkan jumlah total akhir nantinya.algoritma greedy lainnya : huffman coding, activity selection, dijkstra algorithm.

* Dynamic Programming, memecahkan problem yang besar menjadi problem yang lebih kecil dan lebih mudah untuk dikerjakan dengan mengutamakan fakta optimal setiap problem. dynamic programming memiliki karakteristik
*1 overlapping subproblems 
*2 optimal substructure property
*method top down with memoization adalah menyimpan value kedalam cache in case akan ada problem yang sama dan dapat mengambil data tanpa melakukan operasi. (menyelesaikan dari masalah utama ke masalah terkecil)
* method bottom up with tabulation akan menyelesaikan masalah dari yang terkecil terlebih dahulu (menyelesaikan masalah dari terkecil ke utama)

