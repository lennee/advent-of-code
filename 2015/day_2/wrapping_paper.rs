// The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

// Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

// For example:

//     A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
//     A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.

// All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

// The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

// The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

// For example:

//     A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
//     A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.

// How many total feet of ribbon should they order?

use std::fs;

#[derive(Debug)]
struct Present {
    w: u32,
    h: u32,
    l: u32,
}

impl Present {
    fn required_wrapping_paper(&self) -> u32 {
        self.w * self.h * 2 + self.h * self.l * 2 + self.l * self.w * 2 + self.min()
    }

    fn required_ribbon(&self) -> u32 {
        let mut sides = vec![self.w, self.h, self.l];
        sides.sort();
        self.volume() + sides[0] * 2 + sides[1] * 2
    }

    fn volume(&self) -> u32 {
        self.w * self.h * self.l
    }
    
    fn min(&self) -> u32 {
        vec![self.w * self.h, self.h * self.l, self.l * self.w]
            .iter()
            .min()
            .unwrap()
            .clone()
    }
}

fn build_present_from_input(input_line: &str) -> Present {
    let input_arr = input_line.split("x").collect::<Vec<_>>();
    if input_arr.len() != 3 {
        return Present {
            w: 0,
            h: 0,
            l: 0
        }
    }
    Present {
        w: input_arr[0].parse::<u32>().unwrap(),
        h: input_arr[1].parse::<u32>().unwrap(),
        l: input_arr[2].parse::<u32>().unwrap()
    }
}


fn main() {
    let input_file_path = String::from("./day_2/paper_input");
    let input_string = match fs::read_to_string(input_file_path) {
        Ok(input) => input,
        Err(_) => "".to_string()
    };
    
    let mut total_paper = 0;
    let mut total_ribbon = 0;
    
    if input_string.len() > 0 {
        for input in input_string.split("\n") {
            let present = build_present_from_input(input);
            total_paper += present.required_wrapping_paper();
            total_ribbon += present.required_ribbon();
        }
    }
    
    println!("Total wrapping paper needed: {} sqft", total_paper);
    println!("Total ribbon needed: {} sqft", total_ribbon);
}