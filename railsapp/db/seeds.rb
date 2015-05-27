p1 = Product.create(name: "Monitor", price: 10.0)
p2 = Product.create(name: "PlayStation", price: 220.0)
p3 = Product.create(name: "Macbook", price: 1110.0)

Attribute.create(key: "Size", value: "24", product: p1)
Attribute.create(key: "Disc", value: "500GB", product: p2)
Attribute.create(key: "Controller", value: "DualShock", product: p2)
Attribute.create(key: "Ram", value: "16GB", product: p3)
Attribute.create(key: "CPU", value: "i7", product: p3)
Attribute.create(key: "Display", value: "Retina", product: p3)

