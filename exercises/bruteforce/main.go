package main

import (
	"fmt"

	"github.com/pallat/force"
)

var hashed = []string{
	"4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce",
	"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b",
	"5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
	"d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
	"4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"ef2d127de37b942baad06145e54b0c619a1f22327b2ebbcfbec78f5564afe39d",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35",
	"5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"ef2d127de37b942baad06145e54b0c619a1f22327b2ebbcfbec78f5564afe39d",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"7902699be42c8a8e46fbbb4501726517e86b22c56a189f7625a6da49081b2451",
	"d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35",
	"4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce",
	"2c624232cdd221771294dfbb310aca000a0df6ac8b66b696d90ef06fdefb64a3",
	"19581e27de7ced00ff1ce50b2047e7a567c76b1cbaebabe5ef03f7c3017bb5b7",
	"ef2d127de37b942baad06145e54b0c619a1f22327b2ebbcfbec78f5564afe39d",
}

func main() {
	fmt.Print(force.Decrypt(hashed[0]))
}