package hashset

/*
hash set bentuknya seperti hash table tetapi saat terjadi collision,
value pada key yg sama dilakukan update sehingga pada hash set
tidak memiliki key double
*/

type HashSet interface{}

type hSet struct {
}
