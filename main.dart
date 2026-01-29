import 'dart:math';

mixin Shaper {
  double area();
  double perimeter();
  void info();
}

class Rectangle with Shaper {
  double width;
  double height;
  
  Rectangle(this.width, this.height);
  
  @override 
  double area(){
    return width * height;
  }
  
  @override
  double perimeter(){
    return 2 * (width + height);
  }
  
  @override
  void info(){
    print("___RectangleInfo___");
    print("Area=${area()}, Perimeter=${perimeter()}");
  }
}

class Circle with Shaper {
  double radius;
  
  Circle(this.radius);
  
  @override
  double area() {
    return pi * radius * radius;
  }
  
  @override 
  double perimeter() {
    return 2 * pi * radius;
  }
  
  @override
  void info(){
    print("___CircleInfo___");
    print("Area=${area()}, Perimeter=${perimeter()}");
  }
}

void main() {
  Rectangle r = Rectangle(4,5);
  r.info();
  Circle c = Circle(10);
  c.info();
}
