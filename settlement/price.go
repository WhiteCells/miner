package settlement

var price = []float32{
	0.5, // 1
	1.0, // 2
	1.5, // 3
	2.0, // 4
	2.5, // 5
	3.0, // 6+
}

var discount = map[int]float32{
	50:   0.1, // 50+    => 10% discount
	100:  0.2, // 100+   => 20% discount
	250:  0.3, // 250+   => 30% discount
	500:  0.4, // 500+   => 40% discount
	1000: 0.5, // 1000+  => 50% discount
}
