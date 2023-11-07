// Santa is delivering presents to an infinite two-dimensional grid of houses.

// He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

// However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

// For example:

//     > delivers presents to 2 houses: one at the starting location, and one to the east.
//     ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
//     ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

use std::fs;
use std::format;
use std::collections::HashMap;

struct Coordinates {
    x: i32,
    y: i32
}

impl Coordinates {
    fn to_key(&self) -> String{
        format!("{},{}", self.x, self.y)
    }

    fn update_coordinates(&mut self, direction: &u8) {
        match direction {
            b'>' => self.x += 1,
            b'<' => self.x -= 1,
            b'^' => self.y += 1,
            b'v' => self.y -= 1,
            _ => self.y += 0
        };
    }
}

fn read_input() -> String {
    let input_file_path = "./day_3/present_delivery_input";
    match fs::read_to_string(input_file_path) {
        Ok(input) => input,
        Err(_) => "".to_string()
    }
}

fn unassisted_delivery(input: &String) {
    let mut delivery_map = HashMap::<String, u32>::new();

    let mut c = Coordinates {
        x: 0,
        y: 0
    };

    delivery_map.insert(c.to_key(), 1);

    for ch in input.as_bytes().iter() {
        c.update_coordinates(ch);
        if let Some(v) = delivery_map.get_mut(&c.to_key()) {
            *v += 1
        } else {
            delivery_map.insert(c.to_key(), 1);
        }
    }

    println!("Houses that recieve a present: {}", delivery_map.keys().len());
}

fn robo_assisted_delivery(input: &String) {
    let mut delivery_map = HashMap::<String, u32>::new();

    let mut sc = Coordinates {
        x: 0,
        y: 0
    };

    let mut rc = Coordinates {
        x: 0,
        y: 0
    };

    delivery_map.insert(sc.to_key(), 2);

    let mut is_robo = false;
    let mut key;

    for ch in input.as_bytes().iter() {
        if is_robo {
            rc.update_coordinates(ch);
            key = rc.to_key()
        } else {
            sc.update_coordinates(ch);
            key = sc.to_key()
        }
        if let Some(v) = delivery_map.get_mut(&key) {
            *v += 1
        } else {
            delivery_map.insert(key, 1);
        }
        is_robo = !is_robo
    }

    println!("Houses that recieve a present (Robo Assisted): {}", delivery_map.keys().len());
}

fn main() {
    let input_string = read_input();
    unassisted_delivery(&input_string);
    robo_assisted_delivery(&input_string);
}