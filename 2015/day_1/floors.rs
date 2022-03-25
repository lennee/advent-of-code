use std::fs;

const FILENAME: &str = "./floors_input";

fn read_input_file(input_file_name: &str) -> String {
  return fs::read_to_string(input_file_name)
    .expect("Something went wrong reading the file");
}

fn find_final_floor_number(floors: &String, find_basement: bool) -> i32 {
  let mut current_floor: i32 = 0;
  for  (i, c) in floors.chars().enumerate() {
    if c == '(' {
      current_floor = &current_floor + 1
    } else if c == ')' {
      current_floor = &current_floor - 1
    }

    if current_floor < 0 && find_basement == true {
      return i as i32 + 1;
    }
  }
  return current_floor
}


fn main() {
  let floors: String = read_input_file(FILENAME);

  let final_floor: i32 = find_final_floor_number(&floors, false);

  let first_basement: i32 = find_final_floor_number(&floors, true);

  println!("Final Floor: {}\nFirst Basement Visit: {}", final_floor, first_basement);
}