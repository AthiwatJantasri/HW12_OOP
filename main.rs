use std::f64::consts::PI;

trait Shaper {
  fn area(&self) -> f64;
  fn perimeter(&self) -> f64;
  fn info(&self);
}

struct Rectangle {
  width: f64,
  height: f64,
}

impl Shaper for Rectangle {
  fn area(&self) -> f64 {
    self.width * self.height
  }
  
  fn perimeter(&self) -> f64 {
    2.0 * (self.width + self.height)
  }
  
  fn info(&self){
    println!("__RectangleInfo__");
    println!("Area{}, Perimeter{}",self.area(), self.perimeter());
  }
}


struct Circle {
  radius: f64,
}

impl Shaper for Circle {
  fn area(&self) -> f64 {
    PI * self.radius * self.radius
  }
  
  fn perimeter(&self) -> f64 {
    2.0 * PI * self.radius
  }
  
  fn info(&self){
    println!("__CircleInfo__");
    println!("Area{}, Perimeter{}", self.area(), self.perimeter());
  }
}

fn main() {
    let r = Rectangle {
      width: 5.0,
      height: 5.0,
    };
    
    let c = Circle {
      radius: 10.0
    };
    
    r.info();
    c.info();
}
